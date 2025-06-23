// auth/auth.go

package middleware

import (
	config "github.com/abdanhafidz/ai-visual-multi-modal-backend/config"
	models "github.com/abdanhafidz/ai-visual-multi-modal-backend/models"
	"github.com/abdanhafidz/ai-visual-multi-modal-backend/services"
	utils "github.com/abdanhafidz/ai-visual-multi-modal-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var salt = config.Salt
var secretKey = []byte(salt)

// VerifyPassword verifies if the provided password matches the hashed password

func AuthenticationMiddleware(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.ResponseFAIL(c, 401, models.Exception{
				Unauthorized: true,
				Message:      "You Have To Login First!",
			})
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &models.JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			utils.ResponseFAIL(c, 401, models.Exception{
				Unauthorized: true,
				Message:      "Invalid Authorization Token!",
			})
			return
		}

		claims, ok := token.Claims.(*models.JWTCustomClaims)
		if !ok {
			utils.ResponseFAIL(c, 401, models.Exception{
				Unauthorized: true,
				Message:      "Invalid Authorization Token!",
			})
			return
		}

		c.Set("user_id", claims.IdUser)
		c.Next()
	}
}
