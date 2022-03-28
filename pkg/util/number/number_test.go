package number_test

import (
	"math/big"
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/util/number"
	"github.com/stretchr/testify/require"
)

func TestNumber(t *testing.T) {
	t.Run("ToFloat", func(tt *testing.T) {
		tt.Run("should return 12.345", func(ttt *testing.T) {
			n := number.ToFloat(12345, 3)

			require.Equal(ttt, n, 12.345)
		})

		tt.Run("should return 12345", func(ttt *testing.T) {
			n := number.ToFloat(12345, 0)

			require.Equal(ttt, n, 12345.0)
		})

		tt.Run("should return 0", func(ttt *testing.T) {
			n := number.ToFloat(0, 3)

			require.Equal(ttt, n, 0.0)
		})
	})

	t.Run("BigIntToFloat", func(tt *testing.T) {
		tt.Run("should return 12.345", func(ttt *testing.T) {
			n := number.BigIntToFloat(big.NewInt(12345), 3)

			require.Equal(ttt, n, 12.345)
		})

		tt.Run("should return 12345", func(ttt *testing.T) {
			n := number.BigIntToFloat(big.NewInt(12345), 0)

			require.Equal(ttt, n, 12345.0)
		})

		tt.Run("should return 0", func(ttt *testing.T) {
			n := number.BigIntToFloat(big.NewInt(0), 3)

			require.Equal(ttt, n, 0.0)
		})
	})

	t.Run("LargeStringToFloat", func(tt *testing.T) {
		tt.Run("should return 5000", func(ttt *testing.T) {
			n := number.LargeStringToFloat("5000000000000000000000", 18)

			require.Equal(ttt, n, 5000.0)
		})

		tt.Run("should return 0", func(ttt *testing.T) {
			n := number.LargeStringToFloat("0", 18)

			require.Equal(ttt, n, 0.0)
		})

		tt.Run("should return 0 with string contain decimal", func(ttt *testing.T) {
			n := number.LargeStringToFloat("0.0", 18)

			require.Equal(ttt, n, 0.0)
		})
	})
}
