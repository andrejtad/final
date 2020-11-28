package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/andrejtad/final"
	"github.com/andrejtad/final/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt = "jfkl;sdfg ;kejrj;34tefvse56"
	singingKey = "gfsdfg*)jgkdsfDFFGDFGgefgerg"
	tokenTTL = 12 * time.Hour
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorisation
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user final.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
    user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	    user.Id,
	} )
	return token.SignedString([]byte(singingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error)  {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(singingKey), nil
	})
	if err != nil {
		return 0, err
	}
	climes, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("token climes are not of type *TokenClaims")
	}
	return climes.UserId, nil
}

func generatePasswordHash(password string) string  {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}