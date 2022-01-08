package shared_test

import (
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/stretchr/testify/require"
)

func TestModel(t *testing.T) {
	t.Run("Assets", func(tt *testing.T) {
		tt.Run("TotalPrice", func(ttt *testing.T) {
			assets := shared.Assets{
				{
					Price: 10,
				},
				{
					Price: 20,
				},
				{
					Price: 30,
				},
			}

			totalPrice := assets.TotalPrice()

			require.Equal(ttt, totalPrice, 60.0)
		})

		tt.Run("Sort", func(ttt *testing.T) {
			assets := shared.Assets{
				{
					Name:  "A",
					Price: 10,
				},
				{
					Name:  "B",
					Price: 20,
				},
				{
					Name:  "C",
					Price: 30,
				},
			}

			sortedAssets := assets.Sort()

			require.Equal(ttt, sortedAssets[0].Name, "C")
			require.Equal(ttt, sortedAssets[1].Name, "B")
			require.Equal(ttt, sortedAssets[2].Name, "A")
		})
	})
}
