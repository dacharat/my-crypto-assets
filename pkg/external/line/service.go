package line

import (
	"context"
	"net/http"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

//go:generate mockgen -source=./service.go -destination=./mock_line/mock_service.go -package=mock_line
type ILine interface {
	ParseRequest(r *http.Request) ([]*linebot.Event, error)
	SendFlexMessage(ctx context.Context, token string, message linebot.SendingMessage) error
	ReplyTextMessage(ctx context.Context, token string, message string) error
	PushMessage(ctx context.Context, container *linebot.BubbleContainer) error
}

type service struct {
	client *linebot.Client
	cfg    *config.Line
}

func NewLineService(client *linebot.Client, cfg *config.Line) ILine {
	return &service{
		client: client,
		cfg:    cfg,
	}
}

func (s *service) ParseRequest(r *http.Request) ([]*linebot.Event, error) {
	return s.client.ParseRequest(r)
}

func (s *service) SendFlexMessage(ctx context.Context, token string, message linebot.SendingMessage) error {
	_, err := s.client.ReplyMessage(token, message).WithContext(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

func (s *service) ReplyTextMessage(ctx context.Context, token string, message string) error {
	_, err := s.client.ReplyMessage(token, linebot.NewTextMessage(message)).WithContext(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}

func (s *service) PushMessage(ctx context.Context, container *linebot.BubbleContainer) error {
	_, err := s.client.PushMessage(s.cfg.UserID, linebot.NewFlexMessage("test", container)).WithContext(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}
