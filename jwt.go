package main

import (
	"crypto/rsa"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func appendClaims(defaultClaims, customClaims jwt.MapClaims) jwt.MapClaims {
	if defaultClaims == nil {
		return customClaims
	}

	if customClaims == nil {
		return defaultClaims
	}

	for k, v := range customClaims {
		defaultClaims[k] = v
	}

	return defaultClaims
}

func ForgeToken(uid, email, role string, level int, key *rsa.PrivateKey, customClaims jwt.MapClaims) (string, error) {
	claims := appendClaims(jwt.MapClaims{
		"iat":         time.Now().Unix(),
		"jti":         strconv.FormatInt(time.Now().Unix(), 10),
		"exp":         time.Now().UTC().Add(time.Hour).Unix(),
		"sub":         "session",
		"iss":         "barong",
		"aud":         [2]string{"peatio", "barong"},
		"uid":         uid,
		"email":       email,
		"role":        role,
		"level":       level,
		"state":       "active",
		"referral_id": nil,
	}, customClaims)

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return t.SignedString(key)
}
