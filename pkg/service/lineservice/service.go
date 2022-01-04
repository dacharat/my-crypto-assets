package lineservice

import (
	"context"

	"github.com/dacharat/my-crypto-assets/pkg/external/line"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/dacharat/my-crypto-assets/pkg/util/pointer"
	"github.com/dacharat/my-crypto-assets/pkg/util/price"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type ILineService interface {
	SendFlexMessage(ctx context.Context, token string, accounts []shared.Account) error
	PushMessage(ctx context.Context, accounts []shared.Account) error
}

type service struct {
	lineApi line.ILine
}

func NewLineService(lineApi line.ILine) ILineService {
	return &service{
		lineApi: lineApi,
	}
}

func (s *service) SendFlexMessage(ctx context.Context, token string, accounts []shared.Account) error {
	return s.lineApi.SendFlexMessage(ctx, token, createComponent(accounts))
}

func (s *service) PushMessage(ctx context.Context, accounts []shared.Account) error {
	return s.lineApi.PushMessage(ctx, createComponent(accounts))
}

func createComponent(accounts []shared.Account) *linebot.BubbleContainer {
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
		},
	}

	var totalPrice float64
	for _, account := range accounts {
		container.Body.Contents = append(container.Body.Contents, createAccountComponent(account))
		totalPrice += account.TotalPrice
	}

	container.Body.Contents = append(container.Body.Contents, createTotalAssetsComponent(totalPrice))

	return container
}

func createAccountComponent(account shared.Account) *linebot.BoxComponent {
	return &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeText,
		Layout: linebot.FlexBoxLayoutTypeVertical,
		Contents: []linebot.FlexComponent{
			&linebot.BoxComponent{
				Type:       linebot.FlexComponentTypeText,
				Layout:     linebot.FlexBoxLayoutTypeHorizontal,
				PaddingTop: "8px",
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:   linebot.FlexComponentTypeText,
						Text:   string(account.Platform),
						Flex:   pointer.NewInt(8),
						Color:  "#436AA9",
						Weight: linebot.FlexTextWeightTypeBold,
						Align:  linebot.FlexComponentAlignTypeStart,
					},
					&linebot.TextComponent{
						Type:   linebot.FlexComponentTypeText,
						Text:   price.Dollar(account.TotalPrice),
						Flex:   pointer.NewInt(4),
						Color:  "#436AA9",
						Weight: linebot.FlexTextWeightTypeBold,
						Align:  linebot.FlexComponentAlignTypeEnd,
					},
				},
			},
		},
	}
}

func createTotalAssetsComponent(totalPrice float64) *linebot.BoxComponent {
	return &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeText,
		Layout: linebot.FlexBoxLayoutTypeVertical,
		Contents: []linebot.FlexComponent{
			&linebot.BoxComponent{
				Type:       linebot.FlexComponentTypeText,
				Layout:     linebot.FlexBoxLayoutTypeHorizontal,
				PaddingTop: "8px",
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:   linebot.FlexComponentTypeText,
						Text:   "Total",
						Flex:   pointer.NewInt(8),
						Color:  "#436AA9",
						Weight: linebot.FlexTextWeightTypeBold,
						Align:  linebot.FlexComponentAlignTypeStart,
					},
					&linebot.TextComponent{
						Type:   linebot.FlexComponentTypeText,
						Text:   price.Dollar(totalPrice),
						Flex:   pointer.NewInt(4),
						Color:  "#436AA9",
						Weight: linebot.FlexTextWeightTypeBold,
						Align:  linebot.FlexComponentAlignTypeEnd,
					},
				},
			},
		},
	}
}
