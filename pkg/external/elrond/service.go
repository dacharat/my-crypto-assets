package elrond

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/util/httpclient"
)

//go:generate mockgen -source=./service.go -destination=./mock_elrond/mock_service.go -package=mock_elrond
type IElrond interface {
	GetAccount(ctx context.Context, address string) (GetAccountResponse, error)
	GetAccountDelegation(ctx context.Context, address string) ([]GetAccountDelegationResponse, error)
	GetAccountNfts(ctx context.Context, address string) ([]GetAccountNftResponse, error)
}

type service struct {
	client httpclient.IClient
	cfg    *config.Elrond
}

func NewService(client httpclient.IClient, cfg *config.Elrond) IElrond {
	return &service{
		client: client,
		cfg:    cfg,
	}
}

func (s *service) GetAccount(ctx context.Context, address string) (GetAccountResponse, error) {
	path := fmt.Sprintf(s.cfg.GetAccount, address)
	url := fmt.Sprintf("%s%s", s.cfg.Host, path)

	resp, err := s.client.Get(ctx, url, nil, httpclient.WithoutResLog())
	if err != nil {
		return GetAccountResponse{}, err
	}

	var response GetAccountResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (s *service) GetAccountDelegation(ctx context.Context, address string) ([]GetAccountDelegationResponse, error) {
	path := fmt.Sprintf(s.cfg.GetAccountDelegations, address)
	url := fmt.Sprintf("%s%s", s.cfg.DelegationHost, path)

	resp, err := s.client.Get(ctx, url, nil, httpclient.WithoutResLog())
	if err != nil {
		return nil, err
	}

	var response []GetAccountDelegationResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (s *service) GetAccountNfts(ctx context.Context, address string) ([]GetAccountNftResponse, error) {
	path := fmt.Sprintf(s.cfg.GetAccountNfts, address)
	url := fmt.Sprintf("%s%s?type=MetaESDT", s.cfg.Host, path)

	resp, err := s.client.Get(ctx, url, nil, httpclient.WithoutResLog())
	if err != nil {
		return nil, err
	}

	var response []GetAccountNftResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}
