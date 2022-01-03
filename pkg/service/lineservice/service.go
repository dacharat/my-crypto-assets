package lineservice

import (
	"context"
	"fmt"

	"github.com/dacharat/my-crypto-assets/pkg/external/line"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/dacharat/my-crypto-assets/pkg/util/pointer"
	"github.com/dacharat/my-crypto-assets/pkg/util/price"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type ILineService interface {
	SendFlex(ctx context.Context, accounts []shared.Account)
}

type service struct {
	lineApi line.ILine
}

func NewLineService(lineApi line.ILine) ILineService {
	return &service{
		lineApi: lineApi,
	}
}

func (s *service) SendFlex(ctx context.Context, accounts []shared.Account) {
	err := s.lineApi.SendFlexMessage(ctx, createComponent(accounts))
	fmt.Println("=================>", err)
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
	for _, account := range accounts {
		container.Body.Contents = append(container.Body.Contents, createAccountComponent(account))
	}
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
