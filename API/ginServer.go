package api

import (
	"drouotBack/controllers"
	"drouotBack/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	// resolve cors issues for dev stage. Modify for production
	r.Use(middleware.CORSMiddleware())

	// Users endpoints
	r.GET("/users", controllers.GetAllUsers)
	r.POST("/users", controllers.AddUser)
	r.GET("/users/:id", controllers.GetUserById)
	r.POST("/login", controllers.SignInUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.GET("/signout", controllers.SignOutUser)
	r.GET("/user/auctions/:id", controllers.FindUserAuctions)

	// Articles endpoints
	r.GET("/articles", controllers.GetAllArticles)
	r.POST("/articles", controllers.AddArticle)
	r.GET("/articles/:id", controllers.GetArticleById)
	r.PUT("/articles/:id", controllers.UpdateArticle)
	r.DELETE("/articles/:id", controllers.DeleteArticle)
	r.GET("/articles/bids/:id", controllers.GetBidsArticle)
	r.GET("/articles/bid/:id", controllers.GetHighestBidArticle)

	// Auctions endpoints
	r.GET("/auctions", controllers.GetAllAuctions)
	r.POST("/auctions", controllers.AddAuction)
	r.GET("/auctions/:id", controllers.GetAuctionById)
	r.GET("/auctions/articles/:id", controllers.GetArticlesAuction)
	r.PUT("/auctions/:id", controllers.UpdateAuction)
	r.DELETE("/auctions/:id", controllers.DeleteAuction)

	// Bids endpoints
	r.GET("/bids", controllers.GetAllbids)
	r.POST("/bids", controllers.AddBid)
	r.GET("/bids/:id", controllers.GetBidById)
	r.DELETE("/bids/:id", controllers.DeleteBid)
}
