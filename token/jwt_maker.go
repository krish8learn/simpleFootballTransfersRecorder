package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSizeSecretKey = 32

//JSON web token maker
type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSizeSecretKey {
		return nil, fmt.Errorf("invalid key size, cannot be less than %v", minSizeSecretKey)
	}

	return &JWTMaker{secretKey: secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	//parse the token
	parsedToken, err := jwt.ParseWithClaims(token, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		//check the signing method(algorithm HS256)
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("wrong signing method")
		}
		return []byte(maker.secretKey), nil
	})
	
	if err != nil {
		//invalid token can be two type, expired or invalid
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, fmt.Errorf("token has expired")) {
			//expired token
			return nil, fmt.Errorf("expired token")
		}
		return nil, fmt.Errorf("invalid token")
	}

	payload, ok := parsedToken.Claims.(*Payload)
	if !ok {
		return nil, fmt.Errorf("token is valid")
	}
	return payload, nil
}
