package bitkub

type orderBody struct {
	Symbol    string  `json:"sym,omitempty"`
	Amount    float64 `json:"amt,omitempty"` // for buy is amount of THB spend, for sell is amount of btc
	Rate      float64 `json:"rat,omitempty"`
	Type      string  `json:"typ,omitempty"`
	Ts        int64   `json:"ts,omitempty"`
	Signature string  `json:"sig,omitempty"`
}
