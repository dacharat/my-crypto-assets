package stringutil_test

import (
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/util/stringutil"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	t.Run("ContainStr", func(tt *testing.T) {
		tt.Run("should return false empty array", func(ttt *testing.T) {
			contain := stringutil.ContainStr("", []string{})

			require.False(ttt, contain)
		})
	})

	t.Run("ContainStr", func(tt *testing.T) {
		tt.Run("should return false not found", func(ttt *testing.T) {
			contain := stringutil.ContainStr("D", []string{"A", "B", "C"})

			require.False(ttt, contain)
		})
	})

	t.Run("ContainStr", func(tt *testing.T) {
		tt.Run("should return false case sensitive", func(ttt *testing.T) {
			contain := stringutil.ContainStr("a", []string{"A", "B", "C"})

			require.False(ttt, contain)
		})
	})

	t.Run("ContainStr", func(tt *testing.T) {
		tt.Run("should return true", func(ttt *testing.T) {
			contain := stringutil.ContainStr("A", []string{"A", "B", "C"})

			require.True(ttt, contain)
		})
	})
}
