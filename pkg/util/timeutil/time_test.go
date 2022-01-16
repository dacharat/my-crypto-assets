package timeutil_test

import (
	"testing"
	"time"

	"github.com/dacharat/my-crypto-assets/pkg/util/timeutil"
	"github.com/stretchr/testify/require"
)

func TestTimeutil(t *testing.T) {
	t.Run("BkkLoc", func(tt *testing.T) {
		tt.Run("is bangkok timezone", func(ttt *testing.T) {
			loc, err := time.LoadLocation("Asia/Bangkok")

			require.NoError(ttt, err)
			require.Equal(ttt, loc, timeutil.BkkLoc)
		})
	})

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

	t.Run("Unfreeze", func(tt *testing.T) {
		tt.Run("time change when unfreeze", func(ttt *testing.T) {
			timeutil.Now = func() time.Time {
				return time.Date(2011, 11, 11, 11, 11, 11, 11, timeutil.BkkLoc)
			}

			now := timeutil.Now()

			require.Equal(ttt, now.Day(), 11)
			require.Equal(ttt, now.Month(), time.Month(11))
			require.Equal(ttt, now.Year(), 2011)

			timeutil.Unfreeze()
			now = timeutil.Now()

			require.NotEqual(ttt, now.Day(), 11)
			require.NotEqual(ttt, now.Month(), time.Month(11))
			require.NotEqual(ttt, now.Year(), 2011)
		})
	})
}
