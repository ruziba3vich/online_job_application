package client

// import (
// 	"context"

// 	"github.com/ruziba3vich/ishtopuz/countries/config"
// 	genprotos "github.com/ruziba3vich/ishtopuz/countries/genprotos/protos"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// type (
// 	Client struct {
// 		client genprotos.CountryServiceClient
// 	}
// )

// func New(cnfg *config.Config) (*Client, error) {
// 	connection, err := grpc.NewClient(cnfg.ServerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Client{
// 		client: genprotos.NewCountryServiceClient(connection),
// 	}, nil
// }

// func (c *Client) CreateCountry(ctx context.Context, req *genprotos.RawCountry) (*genprotos.Country, error) {
// 	response, err := c.client.CreateCountry(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (c *Client) GetClosestCountry(ctx context.Context, req *genprotos.GetCountryRequest) (*genprotos.GetClosestCountryResponse, error) {
// 	response, err := c.client.GetClosestCountry(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }

// func (c *Client) GetCountryById(ctx context.Context, req *genprotos.GetCountryRequest) (*genprotos.Country, error) {
// 	response, err := c.client.GetCountryById(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return response, nil
// }
