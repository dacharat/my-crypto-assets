package algorand

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/util/httpclient"
)

type IAlgoland interface {
	GetAccountByID(ctx context.Context, account string) (AccountResponse, error)
	GetAssetByID(ctx context.Context, asset int) (AssetResponse, error)
}

type service struct {
	client httpclient.Client
}

func NewAlgolandService() IAlgoland {
	return &service{
		client: httpclient.NewClient(),
	}
}

func (s *service) GetAccountByID(ctx context.Context, account string) (AccountResponse, error) {
	path := fmt.Sprintf(config.Cfg.AlgorandClient.GetAccountPath, account)
	url := fmt.Sprintf("%s%s", config.Cfg.AlgorandClient.Host, path)

	resp, err := s.client.Get(ctx, url, nil)
	if err != nil {
		return AccountResponse{}, err
	}

	// if resp.StatusCode != http.StatusOK {
	// }

	var response AccountResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (s *service) GetAssetByID(ctx context.Context, asset int) (AssetResponse, error) {
	path := fmt.Sprintf(config.Cfg.AlgorandClient.GetAssetPath, asset)
	url := fmt.Sprintf("%s%s", config.Cfg.AlgorandClient.Host, path)

	resp, err := s.client.Get(ctx, url, nil)
	if err != nil {
		return AssetResponse{}, err
	}
	var response AssetResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return AssetResponse{}, err
	}

	return response, nil
}
