package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reaviseapp/rv-backend/internal/database"
	"github.com/reaviseapp/rv-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	db *database.Database
}

func NewUserHandler(db *database.Database) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := h.db.Users().FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	currentUserID := c.GetString("userID")
	userID := c.Param("id")

	// Users can only update their own profile
	if currentUserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot update other user's profile"})
		return
	}

	var updateData struct {
		Username     string `json:"username"`
		Bio          string `json:"bio"`
		Website      string `json:"website"`
		Location     string `json:"location"`
		ProfilePhoto string `json:"profilePhoto"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"username":      updateData.Username,
			"bio":           updateData.Bio,
			"website":       updateData.Website,
			"location":      updateData.Location,
			"profile_photo": updateData.ProfilePhoto,
			"updated_at":    time.Now(),
		},
	}

	result, err := h.db.Users().UpdateOne(ctx, bson.M{"_id": userID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (h *UserHandler) FollowUser(c *gin.Context) {
	followerID := c.GetString("userID")
	followeeID := c.Param("id")

	if followerID == followeeID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot follow yourself"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if already following
	var existingFollow models.Follow
	err := h.db.Follows().FindOne(ctx, bson.M{
		"follower_id": followerID,
		"followee_id": followeeID,
	}).Decode(&existingFollow)

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Already following this user"})
		return
	}

	// Create follow relationship
	follow := models.Follow{
		ID:         primitive.NewObjectID().Hex(),
		FollowerID: followerID,
		FolloweeID: followeeID,
		CreatedAt:  time.Now(),
	}

	_, err = h.db.Follows().InsertOne(ctx, follow)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow user"})
		return
	}

	// Update follower and followee counts
	_, _ = h.db.Users().UpdateOne(
		ctx,
		bson.M{"_id": followerID},
		bson.M{"$inc": bson.M{"following_count": 1}},
	)

	_, _ = h.db.Users().UpdateOne(
		ctx,
		bson.M{"_id": followeeID},
		bson.M{"$inc": bson.M{"followers_count": 1}},
	)

	c.JSON(http.StatusOK, gin.H{"message": "User followed successfully"})
}

func (h *UserHandler) UnfollowUser(c *gin.Context) {
	followerID := c.GetString("userID")
	followeeID := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := h.db.Follows().DeleteOne(ctx, bson.M{
		"follower_id": followerID,
		"followee_id": followeeID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unfollow user"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not following this user"})
		return
	}

	// Update follower and followee counts
	_, _ = h.db.Users().UpdateOne(
		ctx,
		bson.M{"_id": followerID},
		bson.M{"$inc": bson.M{"following_count": -1}},
	)

	_, _ = h.db.Users().UpdateOne(
		ctx,
		bson.M{"_id": followeeID},
		bson.M{"$inc": bson.M{"followers_count": -1}},
	)

	c.JSON(http.StatusOK, gin.H{"message": "User unfollowed successfully"})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	currentUserID := c.GetString("userID")
	userID := c.Param("id")

	// Users can only delete their own account
	if currentUserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot delete other user's account"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Delete user
	result, err := h.db.Users().DeleteOne(ctx, bson.M{"_id": userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// TODO: Delete all user's posts, comments, messages, etc.
	// This should be done in a transaction or background job

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
