package shared

import "context"

type IAssetsService interface {
	Type() string
	GetAccount(ctx context.Context) (Account, error)
}
