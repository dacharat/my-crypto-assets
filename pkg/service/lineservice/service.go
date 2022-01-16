package lineservice

import (
	"context"
	"net/http"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/line"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

const (
	startBgColor = "#272c34"
	endBgColor   = "#242c34"

	yelloColor     = "#E5C07B"
	redColor       = "#d8474e"
	greenColor     = "#4fb973"
	blueGreenColor = "#00bdc7"
	grayColor      = "#A6B5C5"
	blueColor      = "#38acf5"
	// purpleColor    = "#9860dd"
)

//go:generate mockgen -source=./service.go -destination=./mock_line_service/mock_service.go -package=mock_line_service
type ILineService interface {
	IsOwner(userId string) bool
	ParseRequest(r *http.Request) ([]*linebot.Event, error)
	SendFlexMessage(ctx context.Context, token string, accounts []shared.Account) error
	ReplyTextMessage(ctx context.Context, token string, message string) error
	PushMessage(ctx context.Context, accounts []shared.Account) error
}

type service struct {
	lineApi line.ILine
	cfg     *config.User
	ownerId string
}

func NewService(lineApi line.ILine, cfg *config.User, ownerId string) ILineService {
	return &service{
		lineApi: lineApi,
		cfg:     cfg,
		ownerId: ownerId,
	}
}

func (s *service) IsOwner(userId string) bool {
	return s.ownerId == userId
}

func (s *service) ParseRequest(r *http.Request) ([]*linebot.Event, error) {
	return s.lineApi.ParseRequest(r)
}

func (s *service) SendFlexMessage(ctx context.Context, token string, accounts []shared.Account) error {
	return s.lineApi.SendFlexMessage(ctx, token, createComponent(accounts, s.cfg.MaxAssetsDisplay))
}

func (s *service) ReplyTextMessage(ctx context.Context, token string, message string) error {
	return s.lineApi.ReplyTextMessage(ctx, token, message)
}

func (s *service) PushMessage(ctx context.Context, accounts []shared.Account) error {
	return s.lineApi.PushMessage(ctx, createComponent(accounts, s.cfg.MaxAssetsDisplay))
}
