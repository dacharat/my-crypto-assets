package line

import (
	"context"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type ILine interface {
	SendMessage(ctx context.Context) error
}

type service struct {
	client *linebot.Client
}

func NewLineService(client *linebot.Client) ILine {
	return &service{
		client: client,
	}
}

func (s *service) SendMessage(ctx context.Context) error {
	container := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Background: &linebot.BoxBackground{
				Type:       linebot.FlexBoxBackgroundTypeLinearGradient,
				Angle:      "90deg",
				StartColor: "#29323c",
				EndColor:   "#37434f",
			},
			Contents: []linebot.FlexComponent{
				&linebot.BoxComponent{
					Type:     linebot.FlexComponentTypeText,
					Layout:   linebot.FlexBoxLayoutTypeVertical,
					Contents: []linebot.FlexComponent{},
				},
				&linebot.SeparatorComponent{
					Margin: "8px",
				},
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "World!",
				},
			},
		},
	}

	_, err := s.client.PushMessage(config.Cfg.Line.UserID, linebot.NewFlexMessage("test", container)).WithContext(ctx).Do()
	if err != nil {
		return err
	}

	return nil
}
