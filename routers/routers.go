package routers

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/likunyan/learn-go/api/controllers"
	"github.com/likunyan/learn-go/api/middleware"
	"github.com/likunyan/learn-go/api/models"
	"log"
	"net/http"
	"os"
)

var identityKey = os.Getenv("IDENTITY_KEY")

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID": claims[identityKey],
		"name":   user.(*models.User).Name,
		"text":   "Hello World.",
	})
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.LoadHTMLGlob("templates/*.html")

	router.StaticFile("/favicon.ico", "./favicon.ico")

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	// QR-Code
	router.GET("/qrcode", controllers.RenderQRExport)

	// Query 和 post form
	router.POST("/query", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	authMiddleware := middleware.AuthMiddlewareInit()

	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	router.GET("/users/me/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": name,
		})
	})

	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/logout", authMiddleware.LogoutHandler)

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := router.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
	}

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.GET("/jokes", controllers.JokeHandler)
		api.POST("/jokes/like/:jokeID", controllers.LikeJoke)
	}

	return router
}
