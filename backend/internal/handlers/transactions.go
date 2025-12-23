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

type TransactionHandler struct {
	db *database.Database
}

func NewTransactionHandler(db *database.Database) *TransactionHandler {
	return &TransactionHandler{db: db}
}

type CreateTransactionRequest struct {
	PostID        string  `json:"postId" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
	PaymentMethod string  `json:"paymentMethod" binding:"required"`
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	buyerID := c.GetString("userID")

	var req CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get post to determine seller
	var post models.Post
	err := h.db.Posts().FindOne(ctx, bson.M{"_id": req.PostID}).Decode(&post)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Create transaction
	transaction := models.Transaction{
		ID:            primitive.NewObjectID().Hex(),
		BuyerID:       buyerID,
		SellerID:      post.UserID,
		PostID:        req.PostID,
		Amount:        req.Amount,
		Status:        "pending",
		PaymentMethod: req.PaymentMethod,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	_, err = h.db.Transactions().InsertOne(ctx, transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	userID := c.GetString("userID")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"$or": []bson.M{
			{"buyer_id": userID},
			{"seller_id": userID},
		},
	}

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := h.db.Transactions().Find(ctx, filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}
	defer cursor.Close(ctx)

	var transactions []models.Transaction
	if err = cursor.All(ctx, &transactions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) GetTransaction(c *gin.Context) {
	transactionID := c.Param("id")
	userID := c.GetString("userID")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var transaction models.Transaction
	err := h.db.Transactions().FindOne(ctx, bson.M{"_id": transactionID}).Decode(&transaction)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	// Only buyer or seller can view the transaction
	if transaction.BuyerID != userID && transaction.SellerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot view this transaction"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) UpdateTransactionStatus(c *gin.Context) {
	transactionID := c.Param("id")
	userID := c.GetString("userID")

	var req struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get transaction
	var transaction models.Transaction
	err := h.db.Transactions().FindOne(ctx, bson.M{"_id": transactionID}).Decode(&transaction)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	// Only seller can update transaction status
	if transaction.SellerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only seller can update transaction status"})
		return
	}

	// Update status
	update := bson.M{
		"$set": bson.M{
			"status":     req.Status,
			"updated_at": time.Now(),
		},
	}

	_, err = h.db.Transactions().UpdateOne(ctx, bson.M{"_id": transactionID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully"})
}
