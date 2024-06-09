package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userId uint, userName, userRole string) (string, string) {
	accessClaims := jwtCustomClaims{
		ID:       userId,
		Username: userName,
		Role:     userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 2)),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, _ := accessToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	refreshClaims := jwtCustomClaims{
		ID:       userId,
		Username: userName,
		Role:     userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), // 7 days
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, _ := refreshToken.SignedString([]byte(os.Getenv("REFRESH_SECRET_KEY")))

	return accessTokenString, refreshTokenString
}

func VerifyRefreshToken(refreshToken string) (*jwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(*jwtCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// func (c *jwtCustomClaims) Valid() error {
// 	if c.ExpiresAt == nil || !c.ExpiresAt.Time.After(time.Now()) {
// 		return fmt.Errorf("token is expired")
// 	}
// 	return nil
// }

// func GenerateToken(userId uint, userName, userRole string) (string, string, error) {
// 	secretKey := os.Getenv("SECRET_KEY")
// 	refreshSecretKey := os.Getenv("REFRESH_SECRET_KEY")

// 	// Create access token
// 	accessUUID := uuid.New().String()
// 	accessClaims := jwtCustomClaims{
// 		userId, userName, userRole, accessUUID,
// 		jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)), // Access token expires in 30 minutes
// 		},
// 	}
// 	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
// 	accessSignedToken, err := accessToken.SignedString([]byte(secretKey))
// 	if err != nil {
// 		return "", "", err
// 	}

// 	// Create refresh token
// 	refreshUUID := uuid.New().String()
// 	refreshClaims := jwtCustomClaims{
// 		userId, userName, userRole, refreshUUID,
// 		jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), // Refresh token expires in 7 days
// 		},
// 	}
// 	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
// 	refreshSignedToken, err := refreshToken.SignedString([]byte(refreshSecretKey))
// 	if err != nil {
// 		return "", "", err
// 	}

// 	return accessSignedToken, refreshSignedToken, nil
// }

// func GenerateToken(userId uint, userName, userRole string) string {
// 	var userClaims = jwtCustomClaims{
// 		userId, userName, userRole,
// 		jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

// 	resultJWT, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

// 	return resultJWT
// }
