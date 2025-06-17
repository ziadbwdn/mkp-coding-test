package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"pubtrans-eticketing/internal/models"
	customErrors "pubtrans-eticketing/pkg/errors"
	"pubtrans-eticketing/internal/utils"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse( 
				"Authentication required",
				customErrors.ErrUnauthorized,
			))
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse(
				"Invalid authorization format",
				customErrors.ErrInvalidToken,
			))
			c.Abort()
			return
		}

		token := tokenParts[1]

		claims, err := utils.ValidateJWT(token, jwtSecret) 
		if err != nil {
			var appErr *customErrors.AppError
			if !errors.As(err, &appErr) {
				appErr = customErrors.NewAppError(
					customErrors.ErrInvalidToken.Code,
					customErrors.ErrInvalidToken.Message,
					err,
				)
			}
			c.JSON(http.StatusUnauthorized, models.ErrorResponse(
				appErr.Message,
				appErr,
			))
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}