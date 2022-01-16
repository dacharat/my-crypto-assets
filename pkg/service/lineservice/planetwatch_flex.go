package lineservice

import (
	"fmt"

	"github.com/dacharat/my-crypto-assets/pkg/service/platnetwatchservice"
	"github.com/dacharat/my-crypto-assets/pkg/util/pointer"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func createPlanetwatchComponent(incomes []*platnetwatchservice.Income) *linebot.BubbleContainer {
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
			Contents: []linebot.FlexComponent{
				&linebot.BoxComponent{
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
									Text:   "Date",
									Flex:   pointer.NewInt(6),
									Color:  redColor,
									Size:   linebot.FlexTextSizeTypeMd,
									Weight: linebot.FlexTextWeightTypeBold,
									Align:  linebot.FlexComponentAlignTypeStart,
								},
								&linebot.TextComponent{
									Type:   linebot.FlexComponentTypeText,
									Text:   "Income(PLANET)",
									Flex:   pointer.NewInt(6),
									Color:  purpleColor,
									Size:   linebot.FlexTextSizeTypeSm,
									Weight: linebot.FlexTextWeightTypeBold,
									Align:  linebot.FlexComponentAlignTypeEnd,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, income := range incomes {
		container.Body.Contents = append(container.Body.Contents, createIncomeComponent(income))
	}

	return container
}

func createIncomeComponent(account *platnetwatchservice.Income) *linebot.BoxComponent {
	box := &linebot.BoxComponent{
		Type:          linebot.FlexComponentTypeText,
		Layout:        linebot.FlexBoxLayoutTypeVertical,
		PaddingBottom: "3px",
		Contents: []linebot.FlexComponent{
			&linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeText,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:  linebot.FlexComponentTypeText,
						Text:  account.Date.Format("02 Jan 2006 15:04:05"),
						Flex:  pointer.NewInt(8),
						Color: grayColor,
						Size:  linebot.FlexTextSizeTypeXs,
						Align: linebot.FlexComponentAlignTypeStart,
					},
					&linebot.TextComponent{
						Type:  linebot.FlexComponentTypeText,
						Text:  fmt.Sprintf("%.3f", account.Amount),
						Flex:  pointer.NewInt(4),
						Color: blueGreenColor,
						Size:  linebot.FlexTextSizeTypeXs,
						Align: linebot.FlexComponentAlignTypeEnd,
					},
				},
			},
		},
	}

	return box
}
