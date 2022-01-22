package lineservice

import (
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func createMenuContainer() *linebot.BubbleContainer {
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
				createButtonComponent("All assets", "Assets", "Assets"),
				createButtonComponent("Planetwatch", "Planetwatch", "Planetwatch"),
			},
		},
	}

	for _, platform := range shared.AvailablePlatform {
		container.Body.Contents = append(container.Body.Contents, createButtonComponent(string(platform), string(platform), string(platform)))
	}

	return container
}

func createButtonComponent(label, data, text string) *linebot.BoxComponent {
	return &linebot.BoxComponent{
		Type:            linebot.FlexComponentTypeBox,
		Layout:          linebot.FlexBoxLayoutTypeVertical,
		JustifyContent:  linebot.FlexComponentJustifyContentTypeCenter,
		BackgroundColor: blueColor,
		Action:          linebot.NewPostbackAction(label, data, text, ""),
		PaddingAll:      linebot.FlexComponentPaddingTypeSm,
		Margin:          linebot.FlexComponentMarginTypeMd,
		CornerRadius:    linebot.FlexComponentCornerRadiusTypeMd,
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Type:  linebot.FlexComponentTypeText,
				Align: linebot.FlexComponentAlignTypeCenter,
				Text:  label,
			},
		},
	}
}
