package lineservice

import (
	"math"

	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type button struct {
	label string
	text  string
}

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
				createButonBoxComponent(button{label: "Assets", text: "Assets"}, button{label: "Planetwatch", text: "Planetwatch"}),
			},
		},
	}

	length := math.Ceil(float64(len(shared.AvailablePlatform)) / 2.0)
	buttons := make([][2]button, int(length))
	index := 0
	index2 := 0
	for i, platform := range shared.AvailablePlatform {
		if i != 0 && i%2 == 0 {
			index++
			index2 = 0
		}
		buttons[index][index2] = button{label: string(platform), text: string(platform)}
		index2++
	}

	for _, button := range buttons {
		container.Body.Contents = append(container.Body.Contents, createButonBoxComponent(button[0], button[1]))
	}

	return container
}

func createButonBoxComponent(b1, b2 button) *linebot.BoxComponent {
	box := &linebot.BoxComponent{
		Type:           linebot.FlexComponentTypeBox,
		Layout:         linebot.FlexBoxLayoutTypeHorizontal,
		JustifyContent: linebot.FlexComponentJustifyContentTypeSpaceBetween,
		PaddingTop:     linebot.FlexComponentPaddingTypeMd,
		PaddingBottom:  linebot.FlexComponentPaddingTypeMd,
		Contents: []linebot.FlexComponent{
			createButtonComponent(b1),
		},
	}

	if b2.label != "" {
		box.Contents = append(box.Contents, createButtonComponent(b2))
	}

	return box
}

func createButtonComponent(b button) *linebot.BoxComponent {
	return &linebot.BoxComponent{
		Type:            linebot.FlexComponentTypeBox,
		Layout:          linebot.FlexBoxLayoutTypeBaseline,
		JustifyContent:  linebot.FlexComponentJustifyContentTypeCenter,
		BackgroundColor: blueGreenColor,
		Action:          linebot.NewMessageAction(b.label, b.text),
		PaddingTop:      linebot.FlexComponentPaddingTypeLg,
		PaddingBottom:   linebot.FlexComponentPaddingTypeLg,
		Width:           "48%",
		CornerRadius:    linebot.FlexComponentCornerRadiusTypeMd,
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Type:  linebot.FlexComponentTypeText,
				Align: linebot.FlexComponentAlignTypeCenter,
				Color: startBgColor,
				Text:  b.label,
			},
		},
	}
}
