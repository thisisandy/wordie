package server

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"net/http"
	"regexp"
	"time"
	"wordie/core/server/api/middleware"
	"wordie/core/server/service/word"

	"github.com/gin-gonic/gin"
	UserService "wordie/core/server/service/user"
)
import "github.com/gin-contrib/cors"

type ArticleReq struct {
	Article string `json:"article" binding:"required"`
}

func StartServer() {
	r, authMiddleWare := gin.Default(), middleware.GetAuthMiddleWare()
	r.Use(cors.New(cors.Config{
		AllowOrigins:           []string{"https://foo.com", "https://localhost:3000"},
		AllowMethods:           []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:          []string{"Content-Length", "Authorization"},
		AllowCredentials:       true,
		MaxAge:                 12 * time.Hour,
		AllowBrowserExtensions: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "chrome-extension://gjcbigdhglhcbaollpobpleckhlbahfa"
		},
	}))
	r.POST(`/login`, authMiddleWare.LoginHandler)
	r.NoRoute(authMiddleWare.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		c.JSON(404, gin.H{
			"code":    http.StatusNotFound,
			"message": "Page not found",
			"data":    claims,
		})
	})
	auth := r.Group("/")
	auth.GET("/refresh_token", authMiddleWare.RefreshHandler)
	auth.Use(authMiddleWare.MiddlewareFunc())
	{
		auth.POST("/article", func(c *gin.Context) {
			// get the ArticleReq from the request in stream
			var article ArticleReq
			err := c.Bind(&article)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "invalid article",
				})
				return
			}
			// Split the ArticleReq into words using a regular expression
			re := regexp.MustCompile(`\b\w+\b`)
			words := re.FindAllString(article.Article, -1)
			email := jwt.ExtractClaims(c)["email"].(string)
			frequency, err := UserService.GetUserFrequency(email)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}
			strangeWords := word.FilterWordsWithSmallerFrequency(words, frequency)
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
				"data":    strangeWords,
			})
			return
		})
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", func(c *gin.Context) {
		userInfo := UserService.Info{}
		if c.Bind(&userInfo) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid user info",
			})
			return
		}
		err := UserService.Register(userInfo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	})
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
