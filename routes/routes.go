package routes

import (
	"net/http"
	"practice/controllers"
	"practice/logger"
	"practice/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.Use(logger.GinLogger(), logger.GinRecovery(true), Cors())
	r.POST("/zhuc", controllers.SignUpHandler)
	r.POST("/login", controllers.LoginHandler)
	r.Use(middleware.JWTAuthorization())
	{
		r.GET("/sort", controllers.CommunityHome)

		r.POST("/Create_article", controllers.CreateArticle)
		r.POST("/CommunityList", controllers.CommunityList)
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

// Cors 中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") //
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
