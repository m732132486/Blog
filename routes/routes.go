package routes

import (
	"practice/controllers"
	"practice/logger"
	"practice/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/zhuc", controllers.SignUpHandler)
	r.POST("/login", controllers.LoginHandler)
	r.Use(middleware.JWTAuthorization())
	{
		r.GET("/sort", controllers.CommunityHome)
		r.POST("/Create_article", controllers.CreateArticle)
		r.GET("/CommunityList", controllers.CommunityList)
		r.GET("/articles_arch", controllers.ArticleSearch)
		r.GET("/users_article", controllers.UserSArticle)
		r.GET("/title_id/:id", controllers.TitleId)
		r.POST("/user_favorites", controllers.UserFavorites)
		r.GET("/user_favorites_list", controllers.UserFavoritesList)
		r.GET("/delete/:id", controllers.DeleteArticle)
		r.GET("/update/:id", controllers.Like)
		r.GET("/de/:id", controllers.CancelLike)
		r.GET("/", func(c *gin.Context) {
			c.String(200, "Hello World")

		})
	}

	r.Run(":8081")
	return r

}
