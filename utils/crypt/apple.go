package crypt

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"math/big"

	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
	"pentag.kr/distimer/configs"
)

const AppleIssuer = "https://appleid.apple.com"

type ApplePublicKey struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}

type ApplePublicKeyList struct {
	Keys []ApplePublicKey `json:"keys"`
}

type AppleTokenClaims struct {
	ISS           string `json:"iss"`
	AUD           string `json:"aud"`
	EXP           int64  `json:"exp"`
	IAT           int64  `json:"iat"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

func getApplePublicKey() ([]ApplePublicKey, error) {
	var keys ApplePublicKeyList
	c := resty.New()
	resp, err := c.R().
		SetResult(&keys).
		Get("https://appleid.apple.com/auth/keys")
	if err != nil {
		return nil, err
	} else if resp.IsError() {
		return nil, fmt.Errorf("apple Public Key API Error: %s", resp.Status())
	}
	return keys.Keys, nil
}

func GetApplePublicKeyByKid(kid string) (*ApplePublicKey, error) {
	keys, err := getApplePublicKey()
	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		if key.Kid == kid {
			return &key, nil
		}
	}

	return nil, fmt.Errorf("apple Public Key Not Found")
}

func VerifyAppleToken(token string) (*AppleTokenClaims, error) {
	idTk, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		kid := token.Header["kid"].(string)
		publicKey, err := GetApplePublicKeyByKid(kid)
		if err != nil {
			return nil, err
		}
		n, _ := base64.RawURLEncoding.DecodeString(publicKey.N)
		eBytes, _ := base64.StdEncoding.DecodeString(publicKey.E)
		e := new(big.Int).SetBytes(eBytes)

		rsaKey := &rsa.PublicKey{
			N: new(big.Int).SetBytes(n),
			E: int(e.Int64()),
		}
		return rsaKey, nil
	}, jwt.WithAudience(configs.Env.AppleClientID), jwt.WithIssuer(AppleIssuer))
	if err != nil {
		return nil, err
	}
	if !idTk.Valid {
		return nil, fmt.Errorf("invalid Apple Token")
	}

	claims, ok := idTk.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid Apple Token")
	}

	appleClaims := AppleTokenClaims{
		ISS:           claims["iss"].(string),
		AUD:           claims["aud"].(string),
		EXP:           int64(claims["exp"].(float64)),
		IAT:           int64(claims["iat"].(float64)),
		Sub:           claims["sub"].(string),
		Email:         claims["email"].(string),
		EmailVerified: claims["email_verified"].(bool),
	}
	return &appleClaims, nil
}
