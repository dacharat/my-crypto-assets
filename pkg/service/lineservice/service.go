package lineservice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/line"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/dacharat/my-crypto-assets/pkg/util/pointer"
	"github.com/dacharat/my-crypto-assets/pkg/util/price"
	"github.com/line/line-bot-sdk-go/v7/linebot"
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

func createComponent(accounts []shared.Account, maxAsset int) *linebot.BubbleContainer {
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
		container.Body.Contents = append(container.Body.Contents, createAccountComponent(account, maxAsset))
		totalPrice += account.TotalPrice
	}

	container.Body.Contents = append(container.Body.Contents, createTotalAccountAssetsComponent(totalPrice))

	return container
}

func createAccountComponent(account shared.Account, maxAsset int) *linebot.BoxComponent {
	box := &linebot.BoxComponent{
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

	line := maxAsset
	var allAssets bool
	if len(account.Assets) < line {
		allAssets = true
		line = len(account.Assets)
	}

	for i := 0; i < line; i++ {
		box.Contents = append(box.Contents, createAssetComponent(account.Assets[i]))
	}
	if !allAssets {
		box.Contents = append(box.Contents, createHasMoreComponent())
	}

	return box
}

func createAssetComponent(asset *shared.Asset) *linebot.BoxComponent {
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
						Type:        linebot.FlexComponentTypeText,
						Text:        fmt.Sprintf("%.3f %s", asset.Amount, asset.Name),
						Flex:        pointer.NewInt(8),
						Color:       "#f5f7f8",
						Size:        linebot.FlexTextSizeTypeXs,
						OffsetStart: linebot.FlexComponentOffsetTypeMd,
						Align:       linebot.FlexComponentAlignTypeStart,
					},
					&linebot.TextComponent{
						Type:  linebot.FlexComponentTypeText,
						Text:  price.Dollar(asset.Price),
						Flex:  pointer.NewInt(4),
						Color: "#f5f7f8",
						Size:  linebot.FlexTextSizeTypeXs,
						Align: linebot.FlexComponentAlignTypeEnd,
					},
				},
			},
		},
	}
}

func createTotalAccountAssetsComponent(totalPrice float64) *linebot.BoxComponent {
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

func createHasMoreComponent() *linebot.BoxComponent {
	return &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeText,
		Layout: linebot.FlexBoxLayoutTypeVertical,
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Type:        linebot.FlexComponentTypeText,
				Text:        ".....",
				Flex:        pointer.NewInt(8),
				Color:       "#f5f7f8",
				Size:        linebot.FlexTextSizeTypeXs,
				OffsetStart: linebot.FlexComponentOffsetTypeMd,
				Align:       linebot.FlexComponentAlignTypeStart,
			},
		},
	}
}
