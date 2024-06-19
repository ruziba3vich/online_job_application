package storage

import (
	"context"
	"database/sql"
	"log"
	"math"
	"sort"

	"github.com/ruziba3vich/countries/genprotos"
	"github.com/ruziba3vich/countries/internal/config"

	sq "github.com/Masterminds/squirrel"
)

type (
	CountrySt struct {
		db *sql.DB
	}
)

func New(config *config.Config) (*CountrySt, error) {

	db, err := ConnectDB(*config)
	if err != nil {
		return nil, err
	}

	return &CountrySt{
		db: db,
	}, nil
}

func (s *CountrySt) CreateCountry(ctx context.Context, req *genprotos.RawCountry) (*genprotos.Country, error) {
	queryBuilder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar) // postgres identifier

	query, args, err := queryBuilder.Insert("countries").
		Columns("country_name", "latitude", "longitude").
		Values(
			req.CountryName,
			req.Latitude,
			req.Longitude).
		Suffix("RETURNING country_id, country_name, latitude, longitude").
		ToSql()
	if err != nil {
		log.Println("Error generating SQL:", err)
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, query, args...)

	var response genprotos.Country

	if err := row.Scan(
		&response.CountryId,
		&response.CountryName,
		&response.Latitude,
		&response.Longitude); err != nil {
		log.Println("Scan error:", err)
		return nil, err
	}
	if err := row.Err(); err != nil {
		log.Println("Row error:", err)
		return nil, err
	}
	return &response, nil
}

func (s *CountrySt) GetCountryById(ctx context.Context, req *genprotos.GetCountryRequest) (*genprotos.Country, error) {
	query, args, err := sq.Select("country_id", "country_name", "latitude", "longitude").
		From("countries").
		Where(sq.Eq{"country_id": req.CountryId}).
		ToSql()

	if err != nil {
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, query, args...)

	var (
		response genprotos.Country
	)

	if err := row.Scan(
		&response.CountryId,
		&response.CountryName,
		&response.Latitude,
		&response.Longitude); err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *CountrySt) getAllCountries(ctx context.Context) ([]*genprotos.Country, error) {
	query, _, err := sq.Select("country_id", "latitude", "longitude").
		From("countries").
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var response []*genprotos.Country

	for rows.Next() {
		var responsecha genprotos.Country
		if err := rows.Scan(
			&responsecha.CountryId,
			&responsecha.CountryName,
			&responsecha.Latitude,
			&responsecha.Longitude); err != nil {
			return nil, err
		}
		response = append(response, &responsecha)
	}
	return response, nil
}

func (s *CountrySt) GetClosestCountry(ctx context.Context, req *genprotos.GetCountryRequest) (*genprotos.GetClosestCountryResponse, error) {
	return s.getSortedCountries(ctx, req)
}

func (s *CountrySt) haversineDistance(lat1, lon1, lat2, lon2 float32) float32 {
	const R = 6371
	lat1Rad := s.degreesToRadians(lat1)
	lon1Rad := s.degreesToRadians(lon1)
	lat2Rad := s.degreesToRadians(lat2)
	lon2Rad := s.degreesToRadians(lon2)

	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	a := math.Sin(float64(dLat/2))*math.Sin(float64(dLat/2)) +
		math.Cos(float64(lat1Rad))*math.Cos(float64(lat2Rad))*math.Sin(float64(dLon/2))*math.Sin(float64(dLon/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return float32(R * c)
}

func (s *CountrySt) degreesToRadians(degrees float32) float32 {
	return degrees * math.Pi / 180
}

func (s *CountrySt) getSortedCountries(ctx context.Context, req *genprotos.GetCountryRequest) (*genprotos.GetClosestCountryResponse, error) {
	originCountry, err := s.GetCountryById(ctx, req)
	if err != nil {
		return nil, err
	}

	allCountries, err := s.getAllCountries(ctx)
	if err != nil {
		return nil, err
	}

	type CountryDistance struct {
		Country  *genprotos.Country
		Distance float64
	}

	var distances []CountryDistance
	for _, country := range allCountries {
		distance := s.haversineDistance(originCountry.Latitude, originCountry.Longitude, country.Latitude, country.Longitude)
		distances = append(distances, CountryDistance{Country: country, Distance: float64(distance)})
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	sortedCountries := make([]*genprotos.Country, len(distances))
	for i := range distances {
		sortedCountries[i] = distances[i].Country
	}

	return &genprotos.GetClosestCountryResponse{
		Countries: sortedCountries,
	}, nil
}
