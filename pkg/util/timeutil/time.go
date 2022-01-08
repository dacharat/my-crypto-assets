package timeutil

import "time"

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
