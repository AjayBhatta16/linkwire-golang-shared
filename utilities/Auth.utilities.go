package utilities

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetJWTUsername(tokenString string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenUnverifiable
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		log.Println("GetJWTUsername - Error parsing token:", err)
		return "", err
	}

	if !parsedToken.Valid {
		log.Println("GetJWTUsername - Invalid token")
		return "", jwt.ErrTokenInvalidClaims
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		log.Println("GetJWTUsername - Error extracting claims")
		return "", jwt.ErrTokenInvalidClaims
	}

	username, ok := claims["username"].(string)

	if !ok || username == "" {
		log.Println("GetJWTUsername - Username claim missing or invalid")
		return "", jwt.ErrTokenInvalidClaims
	}

	return username, nil
}

func ValidateJWTNotExpired(tokenString string) (bool, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenUnverifiable
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		log.Println("GetJWTUsername - Error parsing token:", err)
		return "", err
	}

	if !parsedToken.Valid {
		log.Println("GetJWTUsername - Invalid token")
		return "", jwt.ErrTokenInvalidClaims
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		log.Println("GetJWTUsername - Error extracting claims")
		return "", jwt.ErrTokenInvalidClaims
	}

	exp, ok := claims["exp"].(float64)
	
	if !ok {
		log.Println("ValidateJWTNotExpired - Error extracting expiration claim")
		return false, jwt.ErrTokenInvalidClaims
	}

	if time.Now().Unix() > int64(exp) {
		log.Println("ValidateJWTNotExpired - Token is expired")
		return false, jwt.ErrTokenExpired
	}

	return true, nil
}

func GenerateJWT(username string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		log.Println("GenerateJWT - Error signing token:", err)
		return "", err
	}

	return tokenString, nil
}

func GetSetCookieHeaderValue(tokenString string) string {
	return "token=" + tokenString + "; HttpOnly; Secure; Path=/; Max-Age=3600"
}