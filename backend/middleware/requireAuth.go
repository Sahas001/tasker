package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sahas001/some-project/db/controller"
	"github.com/Sahas001/some-project/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var config, _ = util.LoadConfig(".")

func RequireAuth(store *controller.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("Authorization")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.JWTKey), nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			if time.Now().Unix() > int64(claims["exp"].(float64)) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
			}
		}
		userID, ok := claims["sub"].(float64)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "error getting user id"})
			return
		}
		user, err := store.GetUser(ctx, int(userID))
		if user.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "get user failed"})
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
