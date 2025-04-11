package main

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"os"

	shortenerv1 "github.com/palashb01/Go-url-shortener/gen/go/proto/shortener/v1"
	"github.com/palashb01/Go-url-shortener/internal/service"
	"github.com/redis/go-redis/v9"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	gw "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Pretty logs for dev
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Starting URL shortener server")

	ctx := context.Background()

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		TLSConfig: &tls.Config{},
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to Redis")
	}

	// gRPC server
	grpcServer := grpc.NewServer()
	shortenerService := service.NewShortenerServer(rdb)
	shortenerv1.RegisterShortenerServiceServer(grpcServer, shortenerService)

	// Start gRPC listener
	go func() {
		lis, err := net.Listen("tcp", ":9090")
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to start gRPC listener")
		}
		log.Info().Msg("gRPC server listening on :9090")
		reflection.Register(grpcServer)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal().Err(err).Msg("gRPC server failed")
		}
	}()

	// gRPC-Gateway (REST)
	mux := gw.NewServeMux()
	err := shortenerv1.RegisterShortenerServiceHandlerFromEndpoint(ctx, mux, ":9090", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start gRPC-Gateway")
	}

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Accept"},
		AllowCredentials: true,
	}).Handler(mux)

	log.Info().Msg("REST Gateway listening on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal().Err(err).Msg("REST Gateway failed")
	}
}
