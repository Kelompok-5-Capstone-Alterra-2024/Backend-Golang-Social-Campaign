package middleware

import (
	"capstone/helper"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		authorization := c.Request().Header.Get("Authorization")
// 		if authorization == "" {
// 			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
// 				"message": "Unauthorized",
// 			})
// 		}

// 		jwtToken := helper.GetToken(authorization)

// 		payload, err := helper.DecodePayload(jwtToken)
// 		if err != nil {
// 			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
// 				"message": "Unauthorized",
// 			})
// 		}

// 		role, ok := payload["role"].(string)
// 		if !ok || role != "admin" {
// 			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
// 				"message": "Unauthorized",
// 			})
// 		}

// 		return next(c)
// 	}
// }

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")
		if authorization == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Unauthorized",
			})
		}

		jwtToken := helper.GetToken(authorization)

		token, err := jwt.ParseWithClaims(jwtToken, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors == jwt.ValidationErrorExpired {
					// Token is expired, refresh the token
					refreshToken := c.Request().Header.Get("Refresh-Token")
					newAccessToken, newRefreshToken, err := RefreshToken(refreshToken)
					if err != nil {
						return c.JSON(http.StatusUnauthorized, map[string]interface{}{
							"message": "Unauthorized",
						})
					}

					c.Response().Header().Set("Access-Token", newAccessToken)
					c.Response().Header().Set("Refresh-Token", newRefreshToken)
				} else {
					return c.JSON(http.StatusUnauthorized, map[string]interface{}{
						"message": "Unauthorized",
					})
				}
			} else {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "Unauthorized",
				})
			}
		} else {
			claims, ok := token.Claims.(*jwtCustomClaims)
			if !ok || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "Unauthorized",
				})
			}

			if claims.Role != "admin" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "Unauthorized",
				})
			}

			c.Set("userID", claims.ID)
			c.Set("role", claims.Role)
		}

		return next(c)
	}
}
