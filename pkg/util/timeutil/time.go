package timeutil

import (
	"fmt"
	"time"
)

var (
	BkkLoc *time.Location
	Now    = func() time.Time {
		return time.Now()
	}
)

func init() {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	BkkLoc = loc
}

func Unfreeze() {
	Now = func() time.Time {
		return time.Now()
	}
}

func TimeAgo(current time.Time, from time.Time) string {
	duration := current.Sub(from)
	if duration < time.Minute {
		return fmt.Sprintf("%0.fs ago", duration.Seconds())
	}

	if duration < time.Hour {
		return fmt.Sprintf("%0.fm ago", duration.Minutes())
	}

	if duration < 24*time.Hour {
		return fmt.Sprintf("%0.fh ago", duration.Hours())
	}

	return fmt.Sprintf("%.0fd ago", duration.Hours()/24)
}
