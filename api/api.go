package api

import (
	"github.com/gin-gonic/gin"
	"github.com/harisaginting/krdv/common/middleware"
	"github.com/harisaginting/krdv/common/wire"
	"gorm.io/gorm"
)

// Swagger Config
// @title GUIN
// @version 1.0
// @description GUIN
// @host localhost:4000
// @BasePath /
// @schemes http
// @query.collection.format multi
// @contact.name Harisa Ginting
// @contact.url ”
func V1(r *gin.RouterGroup, db *gorm.DB) {
	// Dependency injection
	apiAuth := wire.ApiAuth(db)
	apiUser := wire.ApiUser(db)
	apiWatchlist := wire.ApiWatchlist(db)
	apiMovie := wire.ApiMovie(db)

	member := middleware.Start(1)
	// group rest
	rest := r.Group("rest")
	{
		// group v1
		v1 := rest.Group("v1")
		{
			// auth
			apiAuthGroup := v1.Group("auth")
			{
				apiAuthGroup.POST("/register", apiAuth.Register)
				apiAuthGroup.POST("/login", apiAuth.Login)
				apiAuthGroup.GET("/me", member.MustMember(), apiAuth.Me)
			}
			// user
			apiUserGroup := v1.Group("user")
			{
				apiUserGroup.GET("/", apiUser.List)
			}

			// movie
			apiMovieGroup := v1.Group("movie")
			{
				apiMovieGroup.GET("/", apiMovie.List)
				apiMovieGroup.GET("/:id", apiMovie.Get)
			}

			// Watchlist
			ApiWatchlistGroup := v1.Group("watchlist")
			{
				ApiWatchlistGroup.GET("/", member.MustMember(), apiWatchlist.List)
				ApiWatchlistGroup.POST("/", member.MustMember(), apiWatchlist.Create)
				ApiWatchlistGroup.GET("/:id", member.MustMember(), apiWatchlist.Get)
				ApiWatchlistGroup.POST("/:id", member.MustMember(), apiWatchlist.Update)
				ApiWatchlistGroup.DELETE("/:id", member.MustMember(), apiWatchlist.Delete)
				ApiWatchlistGroup.POST("/:id/:movie_id", member.MustMember(), apiWatchlist.AddMovie)
				ApiWatchlistGroup.DELETE("/:id/:movie_id", member.MustMember(), apiWatchlist.DeleteMovie)
			}
		}
	}

}
