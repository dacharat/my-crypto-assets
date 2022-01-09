package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/util/httpclient"
)

//go:generate mockgen -source=./service.go -destination=./mock_coingecko/mock_service.go -package=mock_coingecko
type ICoingecko interface {
	GetPrice(ctx context.Context, c Chain) (GetPriceResponse, error)
	GetAllPrice(ctx context.Context) (GetPriceResponse, error)
}

type service struct {
	client httpclient.IClient
	cfg    *config.Coingecko
}

func NewCoingeckoService(client httpclient.IClient, cfg *config.Coingecko) ICoingecko {
	return &service{
		client: client,
		cfg:    cfg,
	}
}

func (s *service) GetPrice(ctx context.Context, c Chain) (GetPriceResponse, error) {
	ids := chain[c].IDs()
	url := fmt.Sprintf("%s%s?ids=%s&vs_currencies=usd", s.cfg.Host, s.cfg.GetSimplePrice, url.QueryEscape(strings.Join(ids, ",")))

	resp, err := s.client.Get(ctx, url, nil)
	if err != nil {
		return GetPriceResponse{}, err
	}

	var response GetPriceResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (s *service) GetAllPrice(ctx context.Context) (GetPriceResponse, error) {
	var ids []string
	for k := range chain {
		ids = append(ids, chain[k].IDs()...)
	}

	url := fmt.Sprintf("%s%s?ids=%s&vs_currencies=usd", s.cfg.Host, s.cfg.GetSimplePrice, url.QueryEscape(strings.Join(ids, ",")))

	resp, err := s.client.Get(ctx, url, nil)
	if err != nil {
		return GetPriceResponse{}, err
	}

	var response GetPriceResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}
