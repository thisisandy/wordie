package middleware

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
	"wordie/core/server/service/user"
)

type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type userClaims struct {
	Email string
}

func GetAuthMiddleWare() *jwt.GinJWTMiddleware {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("wordie is the best"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "identity",

		SigningAlgorithm: "HS256",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*userClaims); ok {
				return jwt.MapClaims{
					"email": v.Email,
				}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(context *gin.Context) interface{} {
			claims := jwt.ExtractClaims(context)
			return claims
		},

		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVales login
			if err := c.Bind(&loginVales); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			email := loginVales.Email
			password := loginVales.Password
			userInfo, err := user.VerifyUser(email, password)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return &userClaims{
				Email: userInfo.Email,
			}, nil
		},

		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*userClaims); ok {
				return false
			}
			return true
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(code, gin.H{
				"code":    code,
				"token":   token,
				"expire":  expire.Format(time.RFC3339),
				"message": "success",
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
	return authMiddleware
}
