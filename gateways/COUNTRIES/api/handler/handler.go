package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/ishtopuz/countries/config"
	genprotos "github.com/ruziba3vich/ishtopuz/countries/genprotos/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	handler struct {
		logger *log.Logger
		clnt   genprotos.CountryServiceClient
	}

	HanderCfg struct {
		Logger *log.Logger
		Config *config.Config
	}
)

func New(h *HanderCfg) *handler {
	conn, err := grpc.NewClient(h.Config.ServerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		h.Logger.Fatal(err)
	}
	return &handler{
		logger: h.Logger,
		clnt:   genprotos.NewCountryServiceClient(conn),
	}
}

func (h *handler) CreateCountry(c *gin.Context) {
	var req genprotos.RawCountry
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	response, err := h.clnt.CreateCountry(c, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, response)
}

func (h *handler) GetClosestCountry(c *gin.Context) {
	countryStrId := c.Param("id")
	log.Println(countryStrId)
	countryId, err := strconv.Atoi(countryStrId)
	if err != nil {
		h.logger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	req := genprotos.GetCountryRequest{
		CountryId: int32(countryId),
	}

	response, err := h.clnt.GetClosestCountry(c, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		h.logger.Println(err)
		return
	}
	c.JSON(http.StatusOK, response.Countries)
}

func (h *handler) GetCountryById(c *gin.Context) {
	countryStrId := c.Param("id")
	countryId, err := strconv.Atoi(countryStrId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	req := genprotos.GetCountryRequest{
		CountryId: int32(countryId),
	}

	response, err := h.clnt.GetCountryById(c, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, response)
}
