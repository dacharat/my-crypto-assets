package lineservice

import (
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func createAssetContainer(account shared.Account) *linebot.BubbleContainer {
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
				createAccountComponent(account, 100),
			},
		},
	}

	return container
}
