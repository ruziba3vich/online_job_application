package service

import (
	"context"

	"github.com/ruziba3vich/countries/genprotos"
	"github.com/ruziba3vich/countries/internal/storage"
)

type (
	CountryServiceSt struct {
		genprotos.UnimplementedCountryServiceServer
		service storage.CountrySt
	}
)

func New(service storage.CountrySt) *CountryServiceSt {
	return &CountryServiceSt{
		service: service,
	}
}

func (s *CountryServiceSt) CreateCountry(ctx context.Context, req *genprotos.RawCountry) (*genprotos.Country, error) {
	return s.service.CreateCountry(ctx, req)
}

func (s *CountryServiceSt) GetClosestCountry(ctx context.Context, req *genprotos.GetCountryRequest) (*genprotos.GetClosestCountryResponse, error) {
	return s.service.GetClosestCountry(ctx, req)
}

func (s *CountryServiceSt) GetCountryById(ctx context.Context, req *genprotos.GetCountryRequest) (*genprotos.Country, error) {
	return s.service.GetCountryById(ctx, req)
}
