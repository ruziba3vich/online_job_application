package service

import (
	"context"

	"github.com/ruziba3vich/countries/genprotos"
)

type (
	CountryServiceSt struct {
		genprotos.UnimplementedCountryServiceServer
		Service genprotos.CountryServiceServer
	}
)

func (s *CountryServiceSt) CreateCountry(ctx context.Context, req *genprotos.RawCountry) (*genprotos.Country, error) {
	return s.Service.CreateCountry(ctx, req)
}

func (s *CountryServiceSt) GetClosestCountry(ctx context.Context, req *genprotos.GetCountryRequest) (*genprotos.GetClosestCountryResponse, error) {
	return s.Service.GetClosestCountry(ctx, req)
}

func (s *CountryServiceSt) GetCountryById(ctx context.Context, req *genprotos.GetCountryRequest) (*genprotos.Country, error) {
	return s.Service.GetCountryById(ctx, req)
}
