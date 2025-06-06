package main

import (
	"net"

	proto "github.com/GP-Hacks/proto/pkg/api/user"
	"github.com/GP-Hacks/users/internal/config"
	"github.com/GP-Hacks/users/internal/service_provider"
	"github.com/GP-Hacks/users/internal/utils/logger"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.LoadConfig("./config")
	logger.SetupLogger()
	serviceProvider := service_provider.NewServiceProvider()

	log.Info().Msg("Init app")

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(grpcServer)

	proto.RegisterUserServiceServer(grpcServer, serviceProvider.UserController())

	list, err := net.Listen("tcp", ":"+config.Cfg.Grpc.Port)
	if err != nil {
		log.Fatal().Msg("Failed start listen port")
	}

	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatal().Msg("Failed serve grpc")
	}

}
