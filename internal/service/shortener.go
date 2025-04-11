package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	shortenerv1 "github.com/palashb01/Go-url-shortener/gen/go/proto/shortener/v1"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShortenerServer struct {
	shortenerv1.UnimplementedShortenerServiceServer
	redis *redis.Client
}

func NewShortenerServer(rdb *redis.Client) *ShortenerServer {
	return &ShortenerServer{
		redis: rdb,
	}
}

func (s *ShortenerServer) Shorten(ctx context.Context, req *shortenerv1.ShortenRequest) (*shortenerv1.ShortenResponse, error) {
	shortCode := generateShortCode(6)
	err := s.redis.Set(ctx, shortCode, req.OriginalUrl, 48*time.Hour).Err()
	if err != nil {
		log.Error().Err(err).Msg("Failed to store short URL in Redis")
		return nil, err
	}
	apiHost := os.Getenv("API_HOST")
	if apiHost == "" {
		apiHost = "http://localhost:8080"
	}
	shortURL := fmt.Sprintf("%s/v1/resolve/%s", apiHost, shortCode)
	log.Info().Str("short_url", shortURL).Str("original_url", req.OriginalUrl).Msg("Shortened URL")
	return &shortenerv1.ShortenResponse{
		ShortUrl: shortURL,
	}, nil
}

func (s *ShortenerServer) Resolve(ctx context.Context, req *shortenerv1.ResolveRequest) (*shortenerv1.ResolveResponse, error) {
	originalURL, err := s.redis.Get(ctx, req.ShortCode).Result()
	if err == redis.Nil {
		log.Warn().Str("short_code", req.ShortCode).Msg("Short code not found")
		return nil, grpcErrorNotFound("Short code not found")
	} else if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve original URL from Redis")
		return nil, err
	}
	log.Info().Str("short_code", req.ShortCode).Str("original_url", originalURL).Msg("Resolved short URL")
	return &shortenerv1.ResolveResponse{OriginalUrl: originalURL}, nil
}

func generateShortCode(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	code := base64.URLEncoding.EncodeToString(b)
	return strings.TrimRight(code, "=")[:length]
}

// Optionally: wrap not-found error as gRPC error (for better HTTP response)
func grpcErrorNotFound(msg string) error {
	return status.Errorf(codes.NotFound, msg)
}
