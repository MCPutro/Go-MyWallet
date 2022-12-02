package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtService interface {
	GenerateToken(UserId string) string
	ValidateToken(Token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	ID string
	jwt.RegisteredClaims
}

type jwtServiceImpl struct {
	secretKey string
	issuer    string
}

func NewJwtService(secretKey string, issuer string) JwtService {
	return &jwtServiceImpl{secretKey: secretKey, issuer: issuer}
}

func (j *jwtServiceImpl) GenerateToken(UserId string) string {
	claims := jwtCustomClaim{
		ID: UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(j.secretKey))
	//fmt.Printf("%v %v", ss, err)
	if err != nil {
		return ""
	}
	return ss
}

func (j *jwtServiceImpl) ValidateToken(Token string) (*jwt.Token, error) {
	t, err := jwt.Parse(Token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if t.Valid {
		//fmt.Println("You look nice today")
		return t, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		//fmt.Println("That's not even a token")
		return nil, errors.New("that's not even a token")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		//fmt.Println("Timing is everything")
		return nil, errors.New("token is expired")
	} else {
		//fmt.Println("Couldn't handle this token:", err)
		return nil, err
	}
}
