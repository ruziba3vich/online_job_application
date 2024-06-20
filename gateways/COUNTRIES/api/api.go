package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/ishtopuz/countries/api/handler"
	"github.com/ruziba3vich/ishtopuz/countries/config"
)

type (
	API struct {
		logger *log.Logger
		config *config.Config
	}
)

func New(logger *log.Logger, config *config.Config) *API {
	return &API{
		logger: logger,
		config: config,
	}
}

func (a *API) RUN() error {
	router := gin.Default()

	handlerCfg := handler.HanderCfg{
		Logger: a.logger,
		Config: a.config,
	}

	handler := handler.New(&handlerCfg)

	router.POST("/create/country", handler.CreateCountry)
	router.GET("/country/:id", handler.GetCountryById)
	router.GET("/closest/country/to/:id", handler.GetClosestCountry)

	return router.Run(a.config.OwnHost)
}
