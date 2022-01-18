package platnetwatchservice

import "time"

type Income struct {
	Date   time.Time
	Amount float64
}

type StreamData struct {
	Date time.Time
}

type Summary struct {
	Incomes    []*Income
	StreamData []*StreamData
}
