package di

import (
	"github.com/google/wire"
	"github.com/hokkung/configuration/config"
	hc "github.com/hokkung/configuration/handler/configuration"
	"github.com/hokkung/configuration/pkg/radis"
	"github.com/hokkung/configuration/repository/configuration"
	"github.com/hokkung/configuration/server"
	srv "github.com/hokkung/srv/server"
)

var APISet = wire.NewSet(
	ConfigSet,
	RepositorySet,
	HandlerSet,
	server.ProvideCustomizer,
	srv.ProvideServer,
	wire.Struct(new(ApplicationAPI), "*"),
)

var ConfigSet = wire.NewSet(
	config.ProvideConfiguration,
)

var RepositorySet = wire.NewSet(
	radis.ProvideRedis,
	configuration.ProvideConfigurationRepository,
)

var HandlerSet = wire.NewSet(
	hc.ProvideConfigurationHandler,
)

type ApplicationAPI struct {
	Server *srv.Server
}
