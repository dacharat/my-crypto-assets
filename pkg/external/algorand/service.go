package algorand

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/util/httpclient"
)

//go:generate mockgen -source=./service.go -destination=./mock_algorand/mock_service.go -package=mock_algorand
type IAlgoland interface {
	GetAlgodAccountByID(ctx context.Context, account string) (Account, error)
	GetAssetByID(ctx context.Context, asset int) (AssetResponse, error)
}

type service struct {
	client httpclient.IClient
}

func NewAlgolandService(client httpclient.IClient) IAlgoland {
	return &service{
		client: client,
	}
}

func (s *service) GetAlgodAccountByID(ctx context.Context, account string) (Account, error) {
	// fallback to free api
	if config.Cfg.AlgorandClient.UseFreeApi {
		res, err := s.getAccountByID(ctx, account)
		return res.Account, err
	}

	path := fmt.Sprintf(config.Cfg.AlgorandClient.GetAccountPath, account)
	url := fmt.Sprintf("%s%s", config.Cfg.AlgorandClient.AlgodHost, path)

	header := http.Header{}
	if config.Cfg.AlgorandClient.ApiKey != "" {
		header.Set("x-api-key", config.Cfg.AlgorandClient.ApiKey)
	}

	resp, err := s.client.Get(ctx, url, header)
	if err != nil {
		return Account{}, err
	}

	var response Account
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

func (s *service) getAccountByID(ctx context.Context, account string) (AccountResponse, error) {
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
