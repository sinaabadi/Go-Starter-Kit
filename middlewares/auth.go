package middlewares

import (
	"log"
	"net/http"
	"starter-kit/config"
	"starter-kit/models"
	"starter-kit/utils"
	"time"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetAuthMiddlewares() *jwt.GinJWTMiddleware {
	appConfig := config.GetConfig()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       appConfig.Get(`appName`).(string),
		Key:         []byte(appConfig.Get(`jwtSecret`).(string)),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: `id`,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					`username`:   v.Username,
					`first_name`: v.FirstName,
					`last_name`:  v.LastName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Username: claims[`username`].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &models.User{
					Username: userID,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			utils.FormatAndSend(c, http.StatusUnauthorized, `Invalid credentials provided`, nil)
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatalln("JWT Error:" + err.Error())
	}
	return authMiddleware
}
