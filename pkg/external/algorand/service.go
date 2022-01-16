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
	GetTransaction(ctx context.Context, account string) (AccountTransactionResponse, error)
}

type service struct {
	client httpclient.IClient
	cfg    *config.Algorand
}

func NewAlgolandService(client httpclient.IClient, cfg *config.Algorand) IAlgoland {
	return &service{
		client: client,
		cfg:    cfg,
	}
}

func (s *service) GetAlgodAccountByID(ctx context.Context, account string) (Account, error) {
	// fallback to free api
	if s.cfg.UseFreeApi {
		res, err := s.getAccountByID(ctx, account)
		return res.Account, err
	}

	path := fmt.Sprintf(s.cfg.GetAccountPath, account)
	url := fmt.Sprintf("%s%s", s.cfg.AlgodHost, path)

	header := http.Header{}
	if s.cfg.ApiKey != "" {
		header.Set("x-api-key", s.cfg.ApiKey)
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
	path := fmt.Sprintf(s.cfg.GetAssetPath, asset)
	url := fmt.Sprintf("%s%s", s.cfg.Host, path)

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
	path := fmt.Sprintf(s.cfg.GetAccountPath, account)
	url := fmt.Sprintf("%s%s", s.cfg.Host, path)

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

func (s *service) GetTransaction(ctx context.Context, account string) (AccountTransactionResponse, error) {
	path := fmt.Sprintf(s.cfg.GetAccountTransactionsPath, account)
	url := fmt.Sprintf("%s%s?limit=10&asset-id=27165954&currency-greater-than=0", s.cfg.Host, path)

	resp, err := s.client.Get(ctx, url, nil)
	if err != nil {
		return AccountTransactionResponse{}, err
	}

	var response AccountTransactionResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// curl -X 'GET' \
//   'https://algoindexer.algoexplorerapi.io/v2/accounts/CF7TTWTP7KWZSI7SR2FCBDAWMQMZ6OKS25X5LUZZIEEIW4R4VWYCSQYPII/transactions?limit=10&asset-id=27165954&currency-greater-than=0' \
//   -H 'accept: application/json'
