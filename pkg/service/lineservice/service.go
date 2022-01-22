package lineservice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/line"
	"github.com/dacharat/my-crypto-assets/pkg/service/platnetwatchservice"
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
	purpleColor    = "#9860dd"
)

//go:generate mockgen -source=./service.go -destination=./mock_line_service/mock_service.go -package=mock_line_service
type ILineService interface {
	IsOwner(userId string) bool
	ParseRequest(r *http.Request) ([]*linebot.Event, error)
	SendFlexMessage(ctx context.Context, token string, accounts []shared.Account) error
	ReplyTextMessage(ctx context.Context, token string, message string) error
	SendPlanetwatchFlexMessage(ctx context.Context, token string, summary platnetwatchservice.Summary) error
	SendAssetFlexMessage(ctx context.Context, token string, account shared.Account) error
	SendMenuFlexMessage(ctx context.Context, token string) error

	// For Development
	PushMessage(ctx context.Context, accounts []shared.Account) error
	PushPlanetwatchMessage(ctx context.Context, summary platnetwatchservice.Summary) error
	PushAssetMessage(ctx context.Context, account shared.Account) error
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
	return s.lineApi.SendFlexMessage(ctx, token, linebot.NewFlexMessage("my crypto assets", createComponent(accounts, s.cfg.MaxAssetsDisplay)))
}

func (s *service) ReplyTextMessage(ctx context.Context, token string, message string) error {
	return s.lineApi.ReplyTextMessage(ctx, token, message)
}

func (s *service) PushMessage(ctx context.Context, accounts []shared.Account) error {
	return s.lineApi.PushMessage(ctx, createComponent(accounts, s.cfg.MaxAssetsDisplay))
}

func (s *service) PushPlanetwatchMessage(ctx context.Context, summary platnetwatchservice.Summary) error {
	return s.lineApi.PushMessage(ctx, createPlanetwatchComponent(summary))
}

func (s *service) SendPlanetwatchFlexMessage(ctx context.Context, token string, summary platnetwatchservice.Summary) error {
	return s.lineApi.SendFlexMessage(ctx, token, linebot.NewFlexMessage("planetwatch history", createPlanetwatchComponent(summary)))
}

func (s *service) SendAssetFlexMessage(ctx context.Context, token string, account shared.Account) error {
	return s.lineApi.SendFlexMessage(ctx, token, linebot.NewFlexMessage(fmt.Sprintf("%s assets", account.Platform), createAssetContainer(account)))
}

func (s *service) PushAssetMessage(ctx context.Context, account shared.Account) error {
	return s.lineApi.PushMessage(ctx, createAssetContainer(account))
}

func (s *service) SendMenuFlexMessage(ctx context.Context, token string) error {
	return s.lineApi.SendFlexMessage(ctx, token, linebot.NewFlexMessage("Menu", createMenuContainer()))
}
