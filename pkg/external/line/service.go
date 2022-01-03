package line

import (
	"context"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type ILine interface {
	SendFlexMessage(ctx context.Context, container *linebot.BubbleContainer) error
}

type service struct {
	client *linebot.Client
}

func NewLineService(client *linebot.Client) ILine {
	return &service{
		client: client,
	}
}

func (s *service) SendFlexMessage(ctx context.Context, container *linebot.BubbleContainer) error {
	_, err := s.client.PushMessage(config.Cfg.Line.UserID, linebot.NewFlexMessage("test", container)).WithContext(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}
