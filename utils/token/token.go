package token

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"waizly-tech-test/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var API_SECRET = utils.Getenv("API_SECRET", "initokenrafi")

func GenerateToken(user_id uint) (string, error) {
	token_lifespan, err := strconv.Atoi(utils.Getenv("TOKEN_HOUR_LIFESPAN", "1"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(API_SECRET))

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	// token := c.Query("token")
	// if token != "" {
	// 	return token
	// }

	// bearerToken := c.Request.Header.Get("Authorization")
	// splitToken := strings.Split(bearerToken, "Bearer ")
	// if len(splitToken) != 2 {
	// 	return ""
	// }
	// return splitToken[1]

	// authHeader := c.GetHeader("Authorization")
	// if authHeader == "" {
	// 	return ""
	// }

	// tokenParts := strings.Split(authHeader, "Bearer ")
	// if len(tokenParts) != 2 {
	// 	return ""
	// }

	// return strings.TrimSpace(tokenParts[1])

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ("Authorization header missing")
	}

	// Check if the header starts with "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return ("Invalid token format: Bearer token expected")
	}

	// Extract the token after "Bearer "
	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	if token == "" {
		return ("Token value is empty")
	}

	return token

}

func ExtractTokenID(c *gin.Context) (uint, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("ini dia errornya: ")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		customer_id, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(customer_id), nil
	}
	return 0, nil
}
