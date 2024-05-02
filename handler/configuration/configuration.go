package configuration

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ec "github.com/hokkung/configuration/entity/configuration"
	"github.com/hokkung/configuration/service/configuration"
)

type ConfigurationHandler struct {
	configurationService configuration.ConfigurationService
}

func (h *ConfigurationHandler) Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong from config",
		})
	}
}

func (h *ConfigurationHandler) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.Param("key")
		res, err := h.configurationService.Get(ctx, key)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if res == nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			res.Key: res.Val,
		})
	}
}

func (h *ConfigurationHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var config ec.Configuration
		if err := ctx.BindJSON(&config); err != nil {
			return
		}

		err := h.configurationService.Create(ctx, &config)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			config.Key: config.Val,
		})
	}
}

func NewConfigurationHandler(r configuration.ConfigurationService) *ConfigurationHandler {
	return &ConfigurationHandler{
		configurationService: r,
	}
}

func ProvideConfigurationHandler(r configuration.ConfigurationService) (*ConfigurationHandler, func(), error) {
	return NewConfigurationHandler(r), func() {}, nil
}
