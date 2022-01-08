package timeutil_test

import (
	"testing"
	"time"

	"github.com/dacharat/my-crypto-assets/pkg/util/timeutil"
	"github.com/stretchr/testify/require"
)

func TestTimeutil(t *testing.T) {
	t.Run("Now", func(tt *testing.T) {
		tt.Run("can mock time now", func(ttt *testing.T) {
			timeutil.Now = func() time.Time {
				return time.Date(2011, 11, 11, 11, 11, 11, 11, timeutil.BkkLoc)
			}

			now := timeutil.Now()

			require.Equal(ttt, now.Day(), 11)
			require.Equal(ttt, now.Month(), time.Month(11))
			require.Equal(ttt, now.Year(), 2011)
		})
	})
}
