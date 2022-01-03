package lineservice

import (
	"context"

	"github.com/dacharat/my-crypto-assets/pkg/external/line"
)

type ILineService interface {
	SendFlex(ctx context.Context)
}

type service struct {
	lineApi line.ILine
}

func NewLineService(lineApi line.ILine) ILineService {
	return &service{
		lineApi: lineApi,
	}
}

func (s *service) SendFlex(ctx context.Context) {
	s.lineApi.SendMessage(ctx)
}
