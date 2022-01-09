package binance

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/util/httpclient"
	"github.com/dacharat/my-crypto-assets/pkg/util/timeutil"
)

//go:generate mockgen -source=./service.go -destination=./mock_binance/mock_service.go -package=mock_binance
// NOTE: binance doesn't support Staking account yet.
type IBinance interface {
	GetAccount(ctx context.Context) (GetAccountResponse, error)
	GetSavingBalance(ctx context.Context) (GetSavingBalanceResponse, error)
	GetTricker(ctx context.Context) (map[string]float64, error)
}

type service struct {
	client httpclient.IClient
	cfg    *config.Binance
}

func NewBinanceService(client httpclient.IClient, cfg *config.Binance) IBinance {
	return &service{
		client: client,
		cfg:    cfg,
	}
}

// GetAccount get account spot and saving balance
func (s *service) GetAccount(ctx context.Context) (GetAccountResponse, error) {
	nowMilli := timeutil.Now().UnixMilli()
	query := url.Values{
		"timestamp": []string{fmt.Sprintf("%d", nowMilli)},
	}.Encode()

	uri := fmt.Sprintf("%s%s?%s&signature=%s", s.cfg.Host, s.cfg.GetAccount, query, s.signRequest([]byte(query)))

	resp, err := s.client.Get(ctx, uri, s.generateHeader(), httpclient.WithoutResLog())
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

// GetSavingBalance get account saving balance
func (s *service) GetSavingBalance(ctx context.Context) (GetSavingBalanceResponse, error) {
	nowMilli := timeutil.Now().UnixMilli()
	query := url.Values{
		"timestamp": []string{fmt.Sprintf("%d", nowMilli)},
	}.Encode()

	uri := fmt.Sprintf("%s%s?%s&signature=%s", s.cfg.Host, s.cfg.GetSaving, query, s.signRequest([]byte(query)))

	resp, err := s.client.Get(ctx, uri, s.generateHeader())
	if err != nil {
		return GetSavingBalanceResponse{}, err
	}

	var response GetSavingBalanceResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}
	return response, nil
}

// GetTricker get price
func (s *service) GetTricker(ctx context.Context) (map[string]float64, error) {
	uri := fmt.Sprintf("%s%s", s.cfg.Host, s.cfg.GetTricker)

	resp, err := s.client.Get(ctx, uri, s.generateHeader(), httpclient.WithoutResLog())
	if err != nil {
		return nil, err
	}

	var response GetTrickerResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	tricker := make(map[string]float64, len(response))
	for _, r := range response {
		price, _ := strconv.ParseFloat(r.Price, 64)
		tricker[r.Symbol] = price
	}

	return tricker, nil
}

func (s *service) signRequest(body []byte) string {
	h := hmac.New(sha256.New, []byte(s.cfg.ApiSecret))
	if body != nil {
		_, _ = h.Write(body)
	}
	hmacSigned := h.Sum(nil)

	return hex.EncodeToString(hmacSigned)
}

func (s *service) generateHeader() http.Header {
	header := http.Header{}
	header.Set("X-MBX-APIKEY", s.cfg.ApiKey)
	header.Set("Content-Type", "application/json")
	header.Set("Accept", "application/json")

	return header
}
