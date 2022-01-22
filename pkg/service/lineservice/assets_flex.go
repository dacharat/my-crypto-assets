package lineservice

import (
	"fmt"

	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/dacharat/my-crypto-assets/pkg/util/pointer"
	"github.com/dacharat/my-crypto-assets/pkg/util/price"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func createComponent(accounts []shared.Account, maxAsset int) *linebot.BubbleContainer {
	container := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Background: &linebot.BoxBackground{
				Type:       linebot.FlexBoxBackgroundTypeLinearGradient,
				Angle:      "90deg",
				StartColor: startBgColor,
				EndColor:   endBgColor,
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
	title := string(account.Platform)
	// if account.Address != "" && len(account.Address) > 10 {
	// 	title = fmt.Sprintf("%s(%s...%s)", string(account.Platform), account.Address[:5], account.Address[len(account.Address)-4:])
	// }
	box := &linebot.BoxComponent{
		Type:          linebot.FlexComponentTypeText,
		Layout:        linebot.FlexBoxLayoutTypeVertical,
		PaddingBottom: "5px",
		Contents: []linebot.FlexComponent{
			&linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeText,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:   linebot.FlexComponentTypeText,
						Text:   title,
						Flex:   pointer.NewInt(8),
						Color:  redColor,
						Size:   linebot.FlexTextSizeTypeMd,
						Weight: linebot.FlexTextWeightTypeBold,
						Align:  linebot.FlexComponentAlignTypeStart,
					},
					&linebot.TextComponent{
						Type:   linebot.FlexComponentTypeText,
						Text:   price.Dollar(account.TotalPrice),
						Flex:   pointer.NewInt(4),
						Color:  blueGreenColor,
						Size:   linebot.FlexTextSizeTypeSm,
						Weight: linebot.FlexTextWeightTypeBold,
						Align:  linebot.FlexComponentAlignTypeEnd,
					},
				},
			},
		},
	}

	line := maxAsset
	var allAssets bool
	if len(account.Assets) <= line {
		allAssets = true
		line = len(account.Assets)
	}

	for i := 0; i < line; i++ {
		// hide price less than $0.01
		if account.Assets[i].Price < 0.01 {
			continue
		}

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
				PaddingTop: "3px",
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:        linebot.FlexComponentTypeText,
						Text:        fmt.Sprintf("%.3f %s", asset.Amount, asset.Name),
						Flex:        pointer.NewInt(8),
						Color:       blueColor,
						Size:        linebot.FlexTextSizeTypeXxs,
						OffsetStart: linebot.FlexComponentOffsetTypeMd,
						Align:       linebot.FlexComponentAlignTypeStart,
					},
					&linebot.TextComponent{
						Type:  linebot.FlexComponentTypeText,
						Text:  price.Dollar(asset.Price),
						Flex:  pointer.NewInt(4),
						Color: grayColor,
						Size:  linebot.FlexTextSizeTypeXxs,
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
				PaddingTop: "3px",
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:   linebot.FlexComponentTypeText,
						Text:   "Total",
						Flex:   pointer.NewInt(8),
						Color:  yelloColor,
						Weight: linebot.FlexTextWeightTypeBold,
						Align:  linebot.FlexComponentAlignTypeStart,
					},
					&linebot.TextComponent{
						Type:   linebot.FlexComponentTypeText,
						Text:   price.Dollar(totalPrice),
						Flex:   pointer.NewInt(4),
						Color:  greenColor,
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
		Layout: linebot.FlexBoxLayoutTypeHorizontal,
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Type:        linebot.FlexComponentTypeText,
				Text:        ".....",
				Flex:        pointer.NewInt(8),
				Color:       blueColor,
				Size:        linebot.FlexTextSizeTypeXxs,
				OffsetStart: linebot.FlexComponentOffsetTypeMd,
				Align:       linebot.FlexComponentAlignTypeStart,
			},
			&linebot.TextComponent{
				Type:      linebot.FlexComponentTypeText,
				Text:      ".....",
				Flex:      pointer.NewInt(8),
				Color:     grayColor,
				Size:      linebot.FlexTextSizeTypeXxs,
				OffsetEnd: linebot.FlexComponentOffsetTypeMd,
				Align:     linebot.FlexComponentAlignTypeEnd,
			},
		},
	}
}
