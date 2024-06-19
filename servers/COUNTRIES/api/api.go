package api

import (
	"log"
	"net"

	"github.com/ruziba3vich/countries/genprotos"
	"github.com/ruziba3vich/countries/internal/config"
	"google.golang.org/grpc"
)

type (
	API struct {
		service genprotos.CountryServiceServer
	}
)

func New(service genprotos.CountryServiceServer) *API {
	return &API{
		service: service,
	}
}

func (a *API) RUN(config *config.Config) error {
	listener, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		return err
	}

	serverRegisterer := grpc.NewServer()
	genprotos.RegisterCountryServiceServer(serverRegisterer, a.service)

	log.Println("server has started running on port", config.Server.Port)

	return serverRegisterer.Serve(listener)
}
