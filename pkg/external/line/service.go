package line

import (
	"context"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type ILine interface {
	SendFlexMessage(ctx context.Context, token string, container *linebot.BubbleContainer) error
	ReplyTextMessage(ctx context.Context, token string, message string) error
	PushMessage(ctx context.Context, container *linebot.BubbleContainer) error
}

type service struct {
	client *linebot.Client
}

func NewLineService(client *linebot.Client) ILine {
	return &service{
		client: client,
	}
}

func (s *service) SendFlexMessage(ctx context.Context, token string, container *linebot.BubbleContainer) error {
	_, err := s.client.ReplyMessage(token, linebot.NewFlexMessage("my crypto assets", container)).WithContext(ctx).Do()
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
	_, err := s.client.PushMessage(config.Cfg.Line.UserID, linebot.NewFlexMessage("test", container)).WithContext(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}
