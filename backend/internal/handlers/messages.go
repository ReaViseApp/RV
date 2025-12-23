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

type MessageHandler struct {
	db *database.Database
}

func NewMessageHandler(db *database.Database) *MessageHandler {
	return &MessageHandler{db: db}
}

type SendMessageRequest struct {
	ReceiverID string `json:"receiverId" binding:"required"`
	Text       string `json:"text" binding:"required"`
}

func (h *MessageHandler) SendMessage(c *gin.Context) {
	senderID := c.GetString("userID")

	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message := models.Message{
		ID:         primitive.NewObjectID().Hex(),
		SenderID:   senderID,
		ReceiverID: req.ReceiverID,
		Text:       req.Text,
		IsRead:     false,
		CreatedAt:  time.Now(),
	}

	_, err := h.db.Messages().InsertOne(ctx, message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusCreated, message)
}

func (h *MessageHandler) GetConversations(c *gin.Context) {
	userID := c.GetString("userID")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get unique conversation partners with their last message
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"$or": []bson.M{
					{"sender_id": userID},
					{"receiver_id": userID},
				},
			},
		},
		{
			"$sort": bson.M{"created_at": -1},
		},
		{
			"$group": bson.M{
				"_id": bson.M{
					"$cond": []interface{}{
						bson.M{"$eq": []string{"$sender_id", userID}},
						"$receiver_id",
						"$sender_id",
					},
				},
				"lastMessage": bson.M{"$first": "$$ROOT"},
			},
		},
	}

	cursor, err := h.db.Messages().Aggregate(ctx, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch conversations"})
		return
	}
	defer cursor.Close(ctx)

	var conversations []map[string]interface{}
	if err = cursor.All(ctx, &conversations); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode conversations"})
		return
	}

	c.JSON(http.StatusOK, conversations)
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
	userID := c.GetString("userID")
	otherUserID := c.Param("userId")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"$or": []bson.M{
			{"sender_id": userID, "receiver_id": otherUserID},
			{"sender_id": otherUserID, "receiver_id": userID},
		},
	}

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: 1}})

	cursor, err := h.db.Messages().Find(ctx, filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
		return
	}
	defer cursor.Close(ctx)

	var messages []models.Message
	if err = cursor.All(ctx, &messages); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode messages"})
		return
	}

	// Mark messages as read
	_, _ = h.db.Messages().UpdateMany(
		ctx,
		bson.M{
			"sender_id":   otherUserID,
			"receiver_id": userID,
			"is_read":     false,
		},
		bson.M{"$set": bson.M{"is_read": true}},
	)

	c.JSON(http.StatusOK, messages)
}
