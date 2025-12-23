package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/reaviseapp/rv-backend/internal/database"
	"github.com/reaviseapp/rv-backend/internal/handlers"
	"github.com/reaviseapp/rv-backend/internal/middleware"
	"github.com/reaviseapp/rv-backend/internal/services"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Get configuration from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		dbName = "reavise"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET must be set")
	}

	// Initialize database
	db, err := database.NewDatabase(mongoURI, dbName)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize services
	authService := services.NewAuthService(jwtSecret)

	// Initialize services
	paymentService := services.NewPaymentService()
	recommendationService := services.NewRecommendationService(db)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(db, authService)
	postHandler := handlers.NewPostHandler(db)
	userHandler := handlers.NewUserHandler(db)
	messageHandler := handlers.NewMessageHandler(db)
	commentHandler := handlers.NewCommentHandler(db)
	transactionHandler := handlers.NewTransactionHandler(db)
	nftHandler := handlers.NewNFTHandler(db)

	// Setup Gin router
	router := gin.Default()

	// Middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.Use(middleware.Logger())
	router.Use(middleware.ErrorHandler())

	// Public routes
	api := router.Group("/api")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.GET("/me", middleware.AuthMiddleware(authService), authHandler.GetCurrentUser)
		}

		// Posts routes
		posts := api.Group("/posts")
		{
			posts.GET("", postHandler.GetPosts)
			posts.GET("/:id", postHandler.GetPost)
			
			// Protected routes
			posts.POST("", middleware.AuthMiddleware(authService), postHandler.CreatePost)
			posts.POST("/:id/like", middleware.AuthMiddleware(authService), postHandler.LikePost)
		}

		// User routes
		users := api.Group("/users")
		{
			users.GET("/:id", userHandler.GetUser)
			
			// Protected routes
			users.PUT("/:id", middleware.AuthMiddleware(authService), userHandler.UpdateUser)
			users.POST("/:id/follow", middleware.AuthMiddleware(authService), userHandler.FollowUser)
			users.DELETE("/:id/follow", middleware.AuthMiddleware(authService), userHandler.UnfollowUser)
			users.DELETE("/:id", middleware.AuthMiddleware(authService), userHandler.DeleteUser)
		}

		// Message routes (all protected)
		messages := api.Group("/messages", middleware.AuthMiddleware(authService))
		{
			messages.POST("", messageHandler.SendMessage)
			messages.GET("/conversations", messageHandler.GetConversations)
			messages.GET("/:userId", messageHandler.GetMessages)
		}

		// Comment routes
		comments := api.Group("/comments")
		{
			comments.GET("/post/:postId", commentHandler.GetComments)
			
			// Protected routes
			comments.POST("/post/:postId", middleware.AuthMiddleware(authService), commentHandler.CreateComment)
			comments.DELETE("/:id", middleware.AuthMiddleware(authService), commentHandler.DeleteComment)
		}

		// Transaction routes (all protected)
		transactions := api.Group("/transactions", middleware.AuthMiddleware(authService))
		{
			transactions.POST("", transactionHandler.CreateTransaction)
			transactions.GET("", transactionHandler.GetTransactions)
			transactions.GET("/:id", transactionHandler.GetTransaction)
			transactions.PUT("/:id/status", transactionHandler.UpdateTransactionStatus)
		}

		// NFT routes
		nfts := api.Group("/nft")
		{
			nfts.GET("", nftHandler.GetNFTListings)
			nfts.GET("/:id", nftHandler.GetNFTListing)
			
			// Protected routes
			nfts.POST("", middleware.AuthMiddleware(authService), nftHandler.CreateNFTListing)
			nfts.POST("/:id/bid", middleware.AuthMiddleware(authService), nftHandler.PlaceBid)
			nfts.POST("/:id/complete", middleware.AuthMiddleware(authService), nftHandler.CompleteAuction)
		}

		// Recommendation routes (protected)
		recommendations := api.Group("/recommendations", middleware.AuthMiddleware(authService))
		{
			recommendations.GET("/foryou", func(c *gin.Context) {
				userID := c.GetString("userID")
				posts, err := recommendationService.GetRecommendedPosts(userID, 50)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get recommendations"})
					return
				}
				c.JSON(http.StatusOK, posts)
			})

			recommendations.GET("/following", func(c *gin.Context) {
				userID := c.GetString("userID")
				posts, err := recommendationService.GetFollowingPosts(userID, 50)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get following posts"})
					return
				}
				c.JSON(http.StatusOK, posts)
			})
		}

		// Payment routes (protected)
		payment := api.Group("/payment", middleware.AuthMiddleware(authService))
		{
			payment.POST("/create-intent", func(c *gin.Context) {
				var req struct {
					Amount   float64 `json:"amount" binding:"required"`
					Currency string  `json:"currency" binding:"required"`
				}

				if err := c.ShouldBindJSON(&req); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				intent, err := paymentService.CreatePaymentIntent(req.Amount, req.Currency)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment intent"})
					return
				}

				c.JSON(http.StatusOK, intent)
			})
		}
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
