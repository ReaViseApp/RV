package models

import (
	"time"
)

type User struct {
	ID                string    `json:"id" bson:"_id,omitempty"`
	Username          string    `json:"username" bson:"username"`
	Email             string    `json:"email" bson:"email"`
	PasswordHash      string    `json:"-" bson:"password_hash"`
	ProfilePhoto      string    `json:"profilePhoto,omitempty" bson:"profile_photo,omitempty"`
	Bio               string    `json:"bio,omitempty" bson:"bio,omitempty"`
	Website           string    `json:"website,omitempty" bson:"website,omitempty"`
	Location          string    `json:"location,omitempty" bson:"location,omitempty"`
	FollowersCount    int       `json:"followersCount" bson:"followers_count"`
	FollowingCount    int       `json:"followingCount" bson:"following_count"`
	IsBusinessAccount bool      `json:"isBusinessAccount" bson:"is_business_account"`
	IsVerified        bool      `json:"isVerified" bson:"is_verified"`
	CreatedAt         time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt         time.Time `json:"updatedAt" bson:"updated_at"`
}

type Post struct {
	ID             string      `json:"id" bson:"_id,omitempty"`
	UserID         string      `json:"userId" bson:"user_id"`
	Username       string      `json:"username" bson:"username"`
	UserAvatar     string      `json:"userAvatar,omitempty" bson:"user_avatar,omitempty"`
	UserLocation   string      `json:"userLocation,omitempty" bson:"user_location,omitempty"`
	Media          []MediaItem `json:"media" bson:"media"`
	Description    string      `json:"description" bson:"description"`
	Category       string      `json:"category" bson:"category"` // lot, design, reavise
	Hashtags       []string    `json:"hashtags" bson:"hashtags"`
	LikesCount     int         `json:"likesCount" bson:"likes_count"`
	CommentsCount  int         `json:"commentsCount" bson:"comments_count"`
	CreatedAt      time.Time   `json:"createdAt" bson:"created_at"`
	UpdatedAt      time.Time   `json:"updatedAt" bson:"updated_at"`
}

type MediaItem struct {
	URL      string `json:"url" bson:"url"`
	Type     string `json:"type" bson:"type"` // image, video
	Category string `json:"category,omitempty" bson:"category,omitempty"`
}

type Comment struct {
	ID         string    `json:"id" bson:"_id,omitempty"`
	PostID     string    `json:"postId" bson:"post_id"`
	UserID     string    `json:"userId" bson:"user_id"`
	Username   string    `json:"username" bson:"username"`
	UserAvatar string    `json:"userAvatar,omitempty" bson:"user_avatar,omitempty"`
	Text       string    `json:"text" bson:"text"`
	CreatedAt  time.Time `json:"createdAt" bson:"created_at"`
}

type Message struct {
	ID         string    `json:"id" bson:"_id,omitempty"`
	SenderID   string    `json:"senderId" bson:"sender_id"`
	ReceiverID string    `json:"receiverId" bson:"receiver_id"`
	Text       string    `json:"text" bson:"text"`
	IsRead     bool      `json:"isRead" bson:"is_read"`
	CreatedAt  time.Time `json:"createdAt" bson:"created_at"`
}

type Transaction struct {
	ID            string    `json:"id" bson:"_id,omitempty"`
	BuyerID       string    `json:"buyerId" bson:"buyer_id"`
	SellerID      string    `json:"sellerId" bson:"seller_id"`
	PostID        string    `json:"postId" bson:"post_id"`
	Amount        float64   `json:"amount" bson:"amount"`
	Status        string    `json:"status" bson:"status"` // pending, completed, cancelled
	PaymentMethod string    `json:"paymentMethod" bson:"payment_method"` // stripe, paypal
	PaymentID     string    `json:"paymentId,omitempty" bson:"payment_id,omitempty"`
	CreatedAt     time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" bson:"updated_at"`
}

type NFTListing struct {
	ID             string    `json:"id" bson:"_id,omitempty"`
	PostID         string    `json:"postId" bson:"post_id"`
	OwnerID        string    `json:"ownerId" bson:"owner_id"`
	StartingBid    float64   `json:"startingBid" bson:"starting_bid"`
	CurrentBid     float64   `json:"currentBid" bson:"current_bid"`
	AuctionEndDate time.Time `json:"auctionEndDate" bson:"auction_end_date"`
	Status         string    `json:"status" bson:"status"` // active, sold, expired
	CreatedAt      time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" bson:"updated_at"`
}

type Like struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	UserID    string    `json:"userId" bson:"user_id"`
	PostID    string    `json:"postId" bson:"post_id"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
}

type Follow struct {
	ID         string    `json:"id" bson:"_id,omitempty"`
	FollowerID string    `json:"followerId" bson:"follower_id"`
	FolloweeID string    `json:"followeeId" bson:"followee_id"`
	CreatedAt  time.Time `json:"createdAt" bson:"created_at"`
}
