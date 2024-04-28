//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
)

func InitializeApplication() (*ApplicationAPI, func(), error) {
	wire.Build(APISet)
	return &ApplicationAPI{}, func() {}, nil
}
