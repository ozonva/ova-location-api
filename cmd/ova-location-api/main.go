package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ozonva/ova-location-api/internal/api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"net/http"

	desc "github.com/ozonva/ova-location-api/pkg/ova-location-api"
)

const (
	httpPort = ":8081"
	grpcPort = ":8082"
)

func listenHttp() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterApiHandlerFromEndpoint(ctx, mux, "localhost"+grpcPort, opts)
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось зарегистрировать прокси http->grpc")
	}

	err = http.ListenAndServe(httpPort, mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось запустить обработчик http")
	}
}

func listenGrpc() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось запустить слушатель grpc")
	}

	s := grpc.NewServer()
	desc.RegisterApiServer(s, api.New())

	if err := s.Serve(listen); err != nil {
		log.Fatal().Err(err).Msg("Не удалось запустить обработчик grpc")
	}
}

func main() {
	go listenHttp()

	listenGrpc()
}
