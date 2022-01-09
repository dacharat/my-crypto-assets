package shared

import "context"

//go:generate mockgen -source=./interface.go -destination=./mock_assets_service/mock_interface.go -package=mock_assets_service
type IAssetsService interface {
	Type() string
	GetAccount(context.Context, GetAccountReq) (Account, error)
}
