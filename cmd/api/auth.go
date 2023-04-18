package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Auth struct {
	Issuer        string
	Audience      string //who is the token intended for
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

type jwtUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	jwt.RegisteredClaims
}

func (j *Auth) GenerateTokenPair(user *jwtUser) (TokenPairs, error) {
	//create a token
	token := jwt.New(jwt.SigningMethodHS256)
	// claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	claims["sub"] = fmt.Sprint(user.ID)
	claims["aud"] = j.Audience
	claims["iss"] = j.Issuer
	claims["iat"] = time.Now().UTC().Unix() //when was it issued. Conversion is because this package wants it in this format
	claims["typ"] = "JWT"

	//expiry
	claims["exp"] = time.Now().UTC().Add(j.TokenExpiry).Unix()

	//sign token
	signedAccessToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	//refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	//claims
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = fmt.Sprint(user.ID)
	refreshTokenClaims["iat"] = time.Now().UTC().Unix() //when was it issued. Conversion is because this package wants it in this format
	refreshTokenClaims["typ"] = "JWT"

	refreshTokenClaims["exp"] = time.Now().UTC().Add(j.RefreshExpiry).Unix()

	signedRefreshToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	return TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}, nil
}
