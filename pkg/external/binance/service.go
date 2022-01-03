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
	"time"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/util/httpclient"
)

type IBinance interface {
	GetAccount(ctx context.Context) (GetAccountResponse, error)
	GetSavingBalance(ctx context.Context) (GetSavingBalanceResponse, error)
	GetTricker(ctx context.Context) (map[string]float64, error)
}

type service struct {
	client httpclient.Client
}

func NewBinanceService() IBinance {
	return &service{
		client: httpclient.NewClient(),
	}
}

// GetAccount get account spot and saving balance
func (s *service) GetAccount(ctx context.Context) (GetAccountResponse, error) {
	nowMilli := time.Now().UnixMilli()
	query := url.Values{
		"timestamp": []string{fmt.Sprintf("%d", nowMilli)},
	}.Encode()

	uri := fmt.Sprintf("%s%s?%s&signature=%s", config.Cfg.Binance.Host, config.Cfg.Binance.GetAccount, query, signRequest([]byte(query)))

	resp, err := s.client.Get(ctx, uri, generateHeader())
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
	nowMilli := time.Now().UnixMilli()
	query := url.Values{
		"timestamp": []string{fmt.Sprintf("%d", nowMilli)},
	}.Encode()

	uri := fmt.Sprintf("%s%s?%s&signature=%s", config.Cfg.Binance.Host, config.Cfg.Binance.GetSaving, query, signRequest([]byte(query)))

	resp, err := s.client.Get(ctx, uri, generateHeader())
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
	uri := fmt.Sprintf("%s%s", config.Cfg.Binance.Host, config.Cfg.Binance.GetTricker)

	resp, err := s.client.Get(ctx, uri, generateHeader())
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

func signRequest(body []byte) string {
	h := hmac.New(sha256.New, []byte(config.Cfg.Binance.ApiSecret))
	if body != nil {
		h.Write(body)
	}
	hmacSigned := h.Sum(nil)

	return hex.EncodeToString(hmacSigned)
}

func generateHeader() http.Header {
	header := http.Header{}
	header.Set("X-MBX-APIKEY", config.Cfg.Binance.ApiKey)
	header.Set("Content-Type", "application/json")
	header.Set("Accept", "application/json")

	return header
}
