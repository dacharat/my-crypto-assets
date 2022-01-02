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

type ICoingecko interface {
	GetPrice(ctx context.Context, c Chain) (GetPriceResponse, error)
}

type service struct {
	client httpclient.Client
}

func NewCoingeckoService() ICoingecko {
	return &service{
		client: httpclient.NewClient(),
	}
}

func (s *service) GetPrice(ctx context.Context, c Chain) (GetPriceResponse, error) {
	ids := chain[c].IDs()
	url := fmt.Sprintf("%s%s?ids=%s&vs_currencies=usd", config.Cfg.Coingecko.Host, config.Cfg.Coingecko.GetSimplePrice, url.QueryEscape(strings.Join(ids, ",")))

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
