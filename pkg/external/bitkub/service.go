package bitkub

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/util/httpclient"
	"github.com/dacharat/my-crypto-assets/pkg/util/timeutil"
)

//go:generate mockgen -source=./service.go -destination=./mock_bitkub/mock_service.go -package=mock_bitkub
type IBitkub interface {
	GetWallet(ctx context.Context) (GetWalletResponse, error)
	GetTricker(ctx context.Context) (GetTrickerResponse, error)
}

type service struct {
	client httpclient.IClient
}

func NewBitkubService(client httpclient.IClient) IBitkub {
	return &service{
		client: client,
	}
}

func (s *service) GetWallet(ctx context.Context) (GetWalletResponse, error) {
	uri := fmt.Sprintf("%s%s", config.Cfg.Bitkub.Host, config.Cfg.Bitkub.GetWallet)
	header := generateHeader()

	t := timeutil.Now().Unix()
	body := orderBody{
		Ts: t,
	}
	//create request body
	byteBody, err := json.Marshal(body)
	if err != nil {
		return GetWalletResponse{}, err
	}
	// create signature and add it to request
	sig := signRequest(byteBody)
	body.Signature = sig
	signedByteBody, err := json.Marshal(body)
	if err != nil {
		return GetWalletResponse{}, err
	}

	resp, err := s.client.Post(ctx, uri, header, bytes.NewBuffer(signedByteBody))
	if err != nil {
		return GetWalletResponse{}, err
	}

	var response GetWalletResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (s *service) GetTricker(ctx context.Context) (GetTrickerResponse, error) {
	uri := fmt.Sprintf("%s%s", config.Cfg.Bitkub.Host, config.Cfg.Bitkub.GetTricker)
	resp, err := s.client.Get(ctx, uri, nil, httpclient.WithoutResLog())
	if err != nil {
		return GetTrickerResponse{}, err
	}

	var response GetTrickerResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func signRequest(body []byte) string {
	h := hmac.New(sha256.New, []byte(config.Cfg.Bitkub.ApiSecret))
	h.Write(body)
	hmacSigned := h.Sum(nil)

	return hex.EncodeToString(hmacSigned)
}

func generateHeader() http.Header {
	header := http.Header{}
	header.Set("X-BTK-APIKEY", config.Cfg.Bitkub.ApiKey)
	header.Set("Content-Type", "application/json")
	header.Set("Accept", "application/json")

	return header
}
