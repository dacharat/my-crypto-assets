package algorand_test

import (
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/external/algorand"
	"github.com/stretchr/testify/require"
)

func TestQuery(t *testing.T) {
	t.Run("WithLimit", func(tt *testing.T) {
		queryOtp := algorand.WithLimit(10)

		q := algorand.Query{}
		queryOtp(&q)

		require.Len(tt, q, 1)
	})

	t.Run("WithAssetID", func(tt *testing.T) {
		queryOtp := algorand.WithAssetID(10)

		q := algorand.Query{}
		queryOtp(&q)

		require.Len(tt, q, 1)
	})

	t.Run("WithCurrencyGreaterThan", func(tt *testing.T) {
		queryOtp := algorand.WithCurrencyGreaterThan(10)

		q := algorand.Query{}
		queryOtp(&q)

		require.Len(tt, q, 1)
	})

	t.Run("String", func(tt *testing.T) {
		queryOtp := []algorand.QueryOption{
			algorand.WithLimit(1),
			algorand.WithAssetID(2),
			algorand.WithCurrencyGreaterThan(3),
		}

		q := algorand.Query{}
		for _, opt := range queryOtp {
			opt(&q)
		}

		require.Len(tt, q, 3)
		require.Equal(tt, q.String(), "limit=1&asset-id=2&currency-greater-than=3")
	})
}
