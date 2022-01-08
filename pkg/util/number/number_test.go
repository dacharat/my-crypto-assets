package number_test

import (
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
}
