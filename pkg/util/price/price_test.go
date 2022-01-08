package price_test

import (
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/util/price"
	"github.com/stretchr/testify/require"
)

func TestPrice(t *testing.T) {
	t.Run("Dollar", func(tt *testing.T) {
		tt.Run("should return $12,345.00", func(ttt *testing.T) {
			dollar := price.Dollar(12345)

			require.Equal(ttt, dollar, "$12,345.00")
		})

		tt.Run("should return $12,345.67", func(ttt *testing.T) {
			dollar := price.Dollar(12345.67)

			require.Equal(ttt, dollar, "$12,345.67")
		})

		tt.Run("should return $12,345.68", func(ttt *testing.T) {
			dollar := price.Dollar(12345.678)

			require.Equal(ttt, dollar, "$12,345.68")
		})

		tt.Run("should return $12,345.60", func(ttt *testing.T) {
			dollar := price.Dollar(12345.601)

			require.Equal(ttt, dollar, "$12,345.60")
		})

		tt.Run("should return $1,111,111,111.11", func(ttt *testing.T) {
			dollar := price.Dollar(1111111111.111)

			require.Equal(ttt, dollar, "$1,111,111,111.11")
		})
	})
}
