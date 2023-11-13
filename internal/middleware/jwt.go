package middleware

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidMetadata = status.Error(codes.Internal, "missing metadata")
	ErrMissingAuthHeader = status.Error(codes.InvalidArgument, "missing authorization header")
	ErrInvalidAuthHeaderFormat = status.Error(codes.InvalidArgument, "invalid authorization header format")
	ErrUnsupportedAuthType = status.Error(codes.InvalidArgument, "unsupported authorization type")
	ErrInvalidAuthToken = status.Error(codes.InvalidArgument, "invalid authorization token")
	ErrExpiredAuthToken = status.Error(codes.PermissionDenied, "authorization token is expired or invalid")
)
const (
	autorhizationHeader = "Authorization"
	authorizationBearer = "bearer"
)

func AuthMiddleware(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ErrInvalidMetadata
	}
	values := md.Get(autorhizationHeader)
	if len(values) == 0 {
		return ErrMissingAuthHeader
	}

	header := values[0]
	fields := strings.Fields(header)
	if len(fields) < 2 {
		return ErrInvalidAuthHeaderFormat
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return ErrUnsupportedAuthType
	}

	tokenString := fields[1]
	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if token != nil && !token.Valid {
		return ErrExpiredAuthToken
	}

	if err != nil {
		return err
	}

	return nil
}