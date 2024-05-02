// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/hokkung/configuration/config"
	configuration3 "github.com/hokkung/configuration/handler/configuration"
	"github.com/hokkung/configuration/pkg/radis"
	"github.com/hokkung/configuration/repository/configuration"
	"github.com/hokkung/configuration/server"
	configuration2 "github.com/hokkung/configuration/service/configuration"
	server2 "github.com/hokkung/srv/server"
)

// Injectors from wire.go:

func InitializeApplication() (*ApplicationAPI, func(), error) {
	configConfiguration, err := config.ProvideConfiguration()
	if err != nil {
		return nil, nil, err
	}
	client, err := radis.ProvideRedis(configConfiguration)
	if err != nil {
		return nil, nil, err
	}
	configurationRepository, cleanup, err := configuration.ProvideConfigurationRepository(client)
	if err != nil {
		return nil, nil, err
	}
	configurationService := configuration2.ProvideConfigurationService(configurationRepository)
	configurationHandler, cleanup2, err := configuration3.ProvideConfigurationHandler(configurationService)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	serverCustomizer, cleanup3, err := server.ProvideCustomizer(configurationHandler)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serverServer, cleanup4, err := server2.ProvideServer(serverCustomizer)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	applicationAPI := &ApplicationAPI{
		Server: serverServer,
	}
	return applicationAPI, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
