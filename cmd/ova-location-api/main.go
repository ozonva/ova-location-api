package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/ozonva/ova-location-api/config"
	"github.com/ozonva/ova-location-api/internal/api"
	"github.com/ozonva/ova-location-api/internal/repo"
	desc "github.com/ozonva/ova-location-api/pkg/ova-location-api"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось загрузить конфигурацию приложения")
	}
}

func listenHttp() {
	cfg := config.Get()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterApiHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", cfg.Server.GrpcPort), opts)
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось зарегистрировать прокси http->grpc")
	}

	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.HttpPort), mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось запустить обработчик http")
	}
}

func listenGrpc() {
	cfg := config.Get()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.GrpcPort))
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось запустить слушатель grpc")
	}

	db, err := sqlx.Connect(cfg.Db.Driver, cfg.Db.GetDsn())
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось подключиться к БД")
	}

	defer db.Close()

	grpcServer := grpc.NewServer()
	locationRepo := repo.New(db)
	desc.RegisterApiServer(grpcServer, api.New(locationRepo))

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal().Err(err).Msg("Не удалось запустить обработчик grpc")
	}
}

func main() {
	go listenHttp()

	listenGrpc()
}
