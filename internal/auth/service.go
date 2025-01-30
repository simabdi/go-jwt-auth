package config

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/simabdi/go-jwt-auth/internal/helper"
	"log"
	"os"
	"strconv"
	"time"
)

type Service interface {
	GenerateToken(uuid string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	VerifyToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewJwtService() *jwtService {
	return &jwtService{}
}

func (js *jwtService) GenerateToken(uuid string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("[Error] : Error loading .env file")
	}

	lifeTime, _ := strconv.Atoi(os.Getenv("LIFETIME"))
	ttl := time.Duration(lifeTime) * time.Second
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": uuid,
		"exp":  time.Now().UTC().Add(ttl).Unix(),
	})

	signedToken, err := token.SignedString([]byte(helper.Std64Decode(os.Getenv("JWT_SECRET_KEY"))))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (js *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	resultToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte("secret_key"), nil
	})

	if err != nil {
		return resultToken, err
	}

	return resultToken, nil
}

func (js *jwtService) VerifyToken(tokenString string) (*jwt.Token, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("[Error] : Error loading .env file")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(helper.Std64Decode(os.Getenv("JWT_SECRET_KEY"))), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
