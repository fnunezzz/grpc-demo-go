package middleware

import (
	"context"
	"crypto/rand"
	rsaKeys "crypto/rsa"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/metadata"
)

const (
	MockMD = "MD"
)

func TestNoMetadata(t *testing.T) {
	ctx := context.Background()
	err := AuthMiddleware(ctx)

	if err != nil && errors.Is(err, ErrInvalidMetadata) {
		t.Log("Returned error as expected")
	} else {
		t.Error("Should have returned an error")
	}

}

func TestNoAuthHeader(t *testing.T) {
	ctx := context.Background()
	md := metadata.New(map[string]string{"key1": "val1", "key2": "val2"})
	ctx = metadata.NewIncomingContext(ctx, md)
	err := AuthMiddleware(ctx)

	if err != nil && errors.Is(err, ErrMissingAuthHeader) {
		t.Log("Returned error as expected")
	} else {
		t.Error("Should have returned an error")
	}
}


func TestInvalidAuthHeaderFormat(t *testing.T) {
	ctx := context.Background()
	md := metadata.New(map[string]string{"Authorization": "Bearer"})
	ctx = metadata.NewIncomingContext(ctx, md)
	err := AuthMiddleware(ctx)

	if err != nil && errors.Is(err, ErrInvalidAuthHeaderFormat) {
		t.Log("Returned error as expected")
	} else {
		t.Error("Should have returned an error")
	}
}

func TestUnsupportedAuthType(t *testing.T) {
	ctx := context.Background()
	md := metadata.New(map[string]string{"Authorization": "WrongAuth ABC"})
	ctx = metadata.NewIncomingContext(ctx, md)
	err := AuthMiddleware(ctx)

	if err != nil && errors.Is(err, ErrUnsupportedAuthType) {
		t.Log("Returned error as expected")
	} else {
		t.Error("Should have returned an error")
	}
}

func TestInvalidAuthToken(t *testing.T) {
	currentPrivateKey, _ := rsaKeys.GenerateKey(rand.Reader, 512)
	claims := jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(currentPrivateKey)
	if err != nil {
		t.Error("Error creating JWT token > ", err.Error())
	}
	ctx := context.Background()
	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", tokenString)})
	ctx = metadata.NewIncomingContext(ctx, md)
	err = AuthMiddleware(ctx)

	if err != nil && errors.Is(err, ErrExpiredAuthToken) {
		t.Log("Returned error as expected")
	} else {
		t.Error("Should have returned an error")
	}
}

func TestExpiredAuthToken(t *testing.T) {
	os.Setenv("JWT_KEY", "secret")
	claims := jwt.MapClaims{
		"exp":   time.Now().Add(time.Second * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var secretKey = []byte(os.Getenv("JWT_KEY"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		t.Error("Error creating JWT token > ", err.Error())
	}
	ctx := context.Background()
	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", tokenString)})
	ctx = metadata.NewIncomingContext(ctx, md)
	time.Sleep(time.Second * 1)
	err = AuthMiddleware(ctx)

	if err != nil && errors.Is(err, ErrExpiredAuthToken) {
		t.Log("Returned error as expected")
	} else {
		t.Error("Should have returned an error")
	}
}

func TestValidToken(t *testing.T) {
	os.Setenv("JWT_KEY", "secret")
	claims := jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var secretKey = []byte(os.Getenv("JWT_KEY"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		t.Error("Error creating JWT token > ", err.Error())
	}
	ctx := context.Background()
	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", tokenString)})
	ctx = metadata.NewIncomingContext(ctx, md)
	err = AuthMiddleware(ctx)

	if err == nil {
		t.Logf("Token is valid")
	} else {
		t.Error("Error token is not valid > ", err.Error())
	}
}