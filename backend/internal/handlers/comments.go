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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommentHandler struct {
	db *database.Database
}

func NewCommentHandler(db *database.Database) *CommentHandler {
	return &CommentHandler{db: db}
}

type CreateCommentRequest struct {
	Text string `json:"text" binding:"required"`
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	userID := c.GetString("userID")
	postID := c.Param("postId")

	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get user info
	var user models.User
	err := h.db.Users().FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Create comment
	comment := models.Comment{
		ID:         primitive.NewObjectID().Hex(),
		PostID:     postID,
		UserID:     userID,
		Username:   user.Username,
		UserAvatar: user.ProfilePhoto,
		Text:       req.Text,
		CreatedAt:  time.Now(),
	}

	_, err = h.db.Comments().InsertOne(ctx, comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	// Increment comments count
	_, _ = h.db.Posts().UpdateOne(
		ctx,
		bson.M{"_id": postID},
		bson.M{"$inc": bson.M{"comments_count": 1}},
	)

	c.JSON(http.StatusCreated, comment)
}

func (h *CommentHandler) GetComments(c *gin.Context) {
	postID := c.Param("postId")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := h.db.Comments().Find(ctx, bson.M{"post_id": postID}, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}
	defer cursor.Close(ctx)

	var comments []models.Comment
	if err = cursor.All(ctx, &comments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	userID := c.GetString("userID")
	commentID := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if comment belongs to user
	var comment models.Comment
	err := h.db.Comments().FindOne(ctx, bson.M{"_id": commentID}).Decode(&comment)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if comment.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot delete other user's comment"})
		return
	}

	// Delete comment
	result, err := h.db.Comments().DeleteOne(ctx, bson.M{"_id": commentID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Decrement comments count
	_, _ = h.db.Posts().UpdateOne(
		ctx,
		bson.M{"_id": comment.PostID},
		bson.M{"$inc": bson.M{"comments_count": -1}},
	)

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
