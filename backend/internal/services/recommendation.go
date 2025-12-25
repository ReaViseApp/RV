package services

import (
	"context"
	"math/rand"
	"time"

	"github.com/reaviseapp/rv-backend/internal/database"
	"github.com/reaviseapp/rv-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RecommendationService struct {
	db *database.Database
}

func NewRecommendationService(db *database.Database) *RecommendationService {
	return &RecommendationService{db: db}
}

// GetRecommendedPosts returns recommended posts for a user
func (s *RecommendationService) GetRecommendedPosts(userID string, limit int) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get user's liked posts to understand preferences
	likesCursor, err := s.db.Likes().Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}

	var likes []models.Like
	if err = likesCursor.All(ctx, &likes); err != nil {
		return nil, err
	}

	// Get categories from liked posts
	var likedPostIDs []string
	for _, like := range likes {
		likedPostIDs = append(likedPostIDs, like.PostID)
	}

	// Simple recommendation: get posts from followed users and popular posts
	// Exclude already liked posts
	filter := bson.M{
		"_id": bson.M{"$nin": likedPostIDs},
	}

	opts := options.Find().
		SetSort(bson.D{{Key: "likes_count", Value: -1}, {Key: "created_at", Value: -1}}).
		SetLimit(int64(limit))

	cursor, err := s.db.Posts().Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []models.Post
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	// Shuffle posts for variety
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(posts), func(i, j int) {
		posts[i], posts[j] = posts[j], posts[i]
	})

	return posts, nil
}

// GetFollowingPosts returns posts from users that the given user follows
func (s *RecommendationService) GetFollowingPosts(userID string, limit int) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get followed users
	followsCursor, err := s.db.Follows().Find(ctx, bson.M{"follower_id": userID})
	if err != nil {
		return nil, err
	}

	var follows []models.Follow
	if err = followsCursor.All(ctx, &follows); err != nil {
		return nil, err
	}

	// Extract followed user IDs
	var followedUserIDs []string
	for _, follow := range follows {
		followedUserIDs = append(followedUserIDs, follow.FolloweeID)
	}

	if len(followedUserIDs) == 0 {
		// Return empty if not following anyone
		return []models.Post{}, nil
	}

	// Get posts from followed users
	filter := bson.M{
		"user_id": bson.M{"$in": followedUserIDs},
	}

	opts := options.Find().
		SetSort(bson.D{{Key: "created_at", Value: -1}}).
		SetLimit(int64(limit))

	cursor, err := s.db.Posts().Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []models.Post
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}
