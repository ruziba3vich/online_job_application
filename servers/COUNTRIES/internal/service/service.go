package service

import (
	"context"
	"log"
	"os"

	"github.com/ruziba3vich/countries/genprotos"
	"github.com/ruziba3vich/countries/internal/storage"
)

type (
	CountryServiceSt struct {
		genprotos.UnimplementedCountryServiceServer
		service storage.CountrySt
		logger  *log.Logger
	}
)

func New(service storage.CountrySt) *CountryServiceSt {
	return &CountryServiceSt{
		service: service,
		logger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (s *CountryServiceSt) CreateCountry(ctx context.Context, req *genprotos.RawCountry) (*genprotos.Country, error) {
	s.logger.Println("create country request")
	return s.service.CreateCountry(ctx, req)
}

func (s *CountryServiceSt) GetClosestCountry(ctx context.Context, req *genprotos.GetCountryRequest) (*genprotos.GetClosestCountryResponse, error) {
	s.logger.Println("get closest country request")
	return s.service.GetClosestCountry(ctx, req)
}

func (s *CountryServiceSt) GetCountryById(ctx context.Context, req *genprotos.GetCountryRequest) (*genprotos.Country, error) {
	s.logger.Println("get country by id request")
	return s.service.GetCountryById(ctx, req)
}
