package di

import (
	"github.com/google/wire"
	"github.com/hokkung/configuration/server"
	srv "github.com/hokkung/srv/server"
)

var APISet = wire.NewSet(
	server.ProvideCustomizer,
	srv.ProvideServer,
	wire.Struct(new(ApplicationAPI), "*"),
)

type ApplicationAPI struct {
	Server *srv.Server
}
