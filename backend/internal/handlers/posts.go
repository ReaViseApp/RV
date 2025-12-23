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

type PostHandler struct {
	db *database.Database
}

func NewPostHandler(db *database.Database) *PostHandler {
	return &PostHandler{db: db}
}

type CreatePostRequest struct {
	Media       []models.MediaItem `json:"media" binding:"required"`
	Description string             `json:"description" binding:"required"`
	Category    string             `json:"category" binding:"required"`
	Hashtags    []string           `json:"hashtags"`
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	userID := c.GetString("userID")

	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user info
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := h.db.Users().FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Create post
	post := models.Post{
		ID:            primitive.NewObjectID().Hex(),
		UserID:        userID,
		Username:      user.Username,
		UserAvatar:    user.ProfilePhoto,
		UserLocation:  user.Location,
		Media:         req.Media,
		Description:   req.Description,
		Category:      req.Category,
		Hashtags:      req.Hashtags,
		LikesCount:    0,
		CommentsCount: 0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	_, err = h.db.Posts().InsertOne(ctx, post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get query parameters
	category := c.Query("category")
	userID := c.Query("userId")

	filter := bson.M{}
	if category != "" {
		filter["category"] = category
	}
	if userID != "" {
		filter["user_id"] = userID
	}

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}).SetLimit(50)

	cursor, err := h.db.Posts().Find(ctx, filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	defer cursor.Close(ctx)

	var posts []models.Post
	if err = cursor.All(ctx, &posts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) GetPost(c *gin.Context) {
	postID := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var post models.Post
	err := h.db.Posts().FindOne(ctx, bson.M{"_id": postID}).Decode(&post)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) LikePost(c *gin.Context) {
	userID := c.GetString("userID")
	postID := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if already liked
	var existingLike models.Like
	err := h.db.Likes().FindOne(ctx, bson.M{"user_id": userID, "post_id": postID}).Decode(&existingLike)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Post already liked"})
		return
	}

	// Create like
	like := models.Like{
		ID:        primitive.NewObjectID().Hex(),
		UserID:    userID,
		PostID:    postID,
		CreatedAt: time.Now(),
	}

	_, err = h.db.Likes().InsertOne(ctx, like)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like post"})
		return
	}

	// Increment likes count
	_, err = h.db.Posts().UpdateOne(
		ctx,
		bson.M{"_id": postID},
		bson.M{"$inc": bson.M{"likes_count": 1}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update likes count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post liked successfully"})
}
