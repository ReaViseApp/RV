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

type NFTHandler struct {
	db *database.Database
}

func NewNFTHandler(db *database.Database) *NFTHandler {
	return &NFTHandler{db: db}
}

type CreateNFTListingRequest struct {
	PostID         string    `json:"postId" binding:"required"`
	StartingBid    float64   `json:"startingBid" binding:"required"`
	AuctionEndDate time.Time `json:"auctionEndDate" binding:"required"`
}

func (h *NFTHandler) CreateNFTListing(c *gin.Context) {
	ownerID := c.GetString("userID")

	var req CreateNFTListingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Verify post belongs to user
	var post models.Post
	err := h.db.Posts().FindOne(ctx, bson.M{"_id": req.PostID, "user_id": ownerID}).Decode(&post)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found or you don't own it"})
		return
	}

	// Create NFT listing
	nftListing := models.NFTListing{
		ID:             primitive.NewObjectID().Hex(),
		PostID:         req.PostID,
		OwnerID:        ownerID,
		StartingBid:    req.StartingBid,
		CurrentBid:     req.StartingBid,
		AuctionEndDate: req.AuctionEndDate,
		Status:         "active",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	_, err = h.db.NFTListings().InsertOne(ctx, nftListing)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create NFT listing"})
		return
	}

	c.JSON(http.StatusCreated, nftListing)
}

func (h *NFTHandler) GetNFTListings(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	status := c.Query("status")
	filter := bson.M{}
	if status != "" {
		filter["status"] = status
	}

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := h.db.NFTListings().Find(ctx, filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch NFT listings"})
		return
	}
	defer cursor.Close(ctx)

	var listings []models.NFTListing
	if err = cursor.All(ctx, &listings); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode NFT listings"})
		return
	}

	c.JSON(http.StatusOK, listings)
}

func (h *NFTHandler) GetNFTListing(c *gin.Context) {
	listingID := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var listing models.NFTListing
	err := h.db.NFTListings().FindOne(ctx, bson.M{"_id": listingID}).Decode(&listing)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "NFT listing not found"})
		return
	}

	c.JSON(http.StatusOK, listing)
}

type PlaceBidRequest struct {
	BidAmount float64 `json:"bidAmount" binding:"required"`
}

func (h *NFTHandler) PlaceBid(c *gin.Context) {
	listingID := c.Param("id")
	bidderID := c.GetString("userID")

	var req PlaceBidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get listing
	var listing models.NFTListing
	err := h.db.NFTListings().FindOne(ctx, bson.M{"_id": listingID}).Decode(&listing)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "NFT listing not found"})
		return
	}

	// Validate bid
	if listing.Status != "active" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auction is not active"})
		return
	}

	if time.Now().After(listing.AuctionEndDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auction has ended"})
		return
	}

	if req.BidAmount <= listing.CurrentBid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bid must be higher than current bid"})
		return
	}

	if bidderID == listing.OwnerID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot bid on your own listing"})
		return
	}

	// Update current bid
	update := bson.M{
		"$set": bson.M{
			"current_bid": req.BidAmount,
			"updated_at":  time.Now(),
		},
	}

	_, err = h.db.NFTListings().UpdateOne(ctx, bson.M{"_id": listingID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place bid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Bid placed successfully",
		"currentBid": req.BidAmount,
	})
}

func (h *NFTHandler) CompleteAuction(c *gin.Context) {
	listingID := c.Param("id")
	ownerID := c.GetString("userID")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get listing
	var listing models.NFTListing
	err := h.db.NFTListings().FindOne(ctx, bson.M{"_id": listingID}).Decode(&listing)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "NFT listing not found"})
		return
	}

	// Verify ownership
	if listing.OwnerID != ownerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only owner can complete auction"})
		return
	}

	// Check if auction has ended
	if time.Now().Before(listing.AuctionEndDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Auction has not ended yet"})
		return
	}

	// Update status
	update := bson.M{
		"$set": bson.M{
			"status":     "sold",
			"updated_at": time.Now(),
		},
	}

	_, err = h.db.NFTListings().UpdateOne(ctx, bson.M{"_id": listingID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to complete auction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Auction completed successfully"})
}
