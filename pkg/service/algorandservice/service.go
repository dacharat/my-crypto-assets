package algorandservice

import (
	"context"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/dacharat/my-crypto-assets/pkg/util/number"
)

type service struct {
	api algorand.IAlgoland
	cfg *config.Algorand
}

func NewService(api algorand.IAlgoland, cfg *config.Algorand) shared.IAssetsService {
	return &service{
		api: api,
		cfg: cfg,
	}
}

func (s *service) Platform() shared.Platform {
	return shared.Algorand
}

func (s *service) GetAccount(ctx context.Context, req shared.GetAccountReq) (shared.Account, error) {
	res, err := s.api.GetAlgodAccountByID(ctx, req.WalletAddress)
	if err != nil {
		return shared.Account{}, err
	}

	return s.mapToAccount(ctx, res), nil
}

func (s *service) mapToAccount(ctx context.Context, resAcount algorand.Account) shared.Account {
	algoAmount := s.toAmount(resAcount.Amount)
	algo := &shared.Asset{
		Amount: algoAmount,
		Name:   "ALGO",
	}

	assets := make(shared.Assets, 0, len(resAcount.Assets))
	for _, asset := range resAcount.Assets {
		if asset.IsFrozen {
			continue
		}

		detail, _ := s.api.GetAssetByID(ctx, asset.AssetID)

		assetAmount := s.toAmount(asset.Amount)
		assets = append(assets, &shared.Asset{
			Amount: assetAmount,
			ID:     asset.AssetID,
			Name:   detail.Asset.Params.Name,
		})
	}

	assets = append(assets, algo)

	return shared.Account{
		Platform:     shared.Algorand,
		Address:      resAcount.Address,
		Assets:       assets,
		TotalPrice:   assets.TotalPrice(),
		NeedCgkPrice: true,
	}
}

func (s *service) toAmount(amount int) float64 {
	return number.ToFloat(amount, s.cfg.DefaultDecimal)
}
