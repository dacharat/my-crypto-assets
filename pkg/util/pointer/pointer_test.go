package pointer_test

import (
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/util/pointer"
	"github.com/stretchr/testify/require"
)

func TestPointer(t *testing.T) {
	t.Run("NewInt", func(tt *testing.T) {
		tt.Run("should return 1", func(ttt *testing.T) {
			n := pointer.NewInt(1)

			require.EqualValues(ttt, *n, 1)
		})
	})
}
