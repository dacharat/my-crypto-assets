package binance

type GetAccountResponse struct {
	MakerCommission  int       `json:"makerCommission"`
	TakerCommission  int       `json:"takerCommission"`
	BuyerCommission  int       `json:"buyerCommission"`
	SellerCommission int       `json:"sellerCommission"`
	CanTrade         bool      `json:"canTrade"`
	CanWithdraw      bool      `json:"canWithdraw"`
	CanDeposit       bool      `json:"canDeposit"`
	UpdateTime       int       `json:"updateTime"`
	AccountType      string    `json:"accountType"`
	Balances         []Balance `json:"balances"`
	Permissions      []string  `json:"permissions"`
}

type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type GetSavingBalanceResponse struct {
	TotalAmountInBTC       string `json:"totalAmountInBTC"`
	TotalAmountInUSDT      string `json:"totalAmountInUSDT"`
	TotalFixedAmountInBTC  string `json:"totalFixedAmountInBTC"`
	TotalFixedAmountInUSDT string `json:"totalFixedAmountInUSDT"`
	TotalFlexibleInBTC     string `json:"totalFlexibleInBTC"`
	TotalFlexibleInUSDT    string `json:"totalFlexibleInUSDT"`
	PositionAmountVos      []struct {
		Asset        string `json:"asset"`
		Amount       string `json:"amount"`
		AmountInBTC  string `json:"amountInBTC"`
		AmountInUSDT string `json:"amountInUSDT"`
	} `json:"positionAmountVos"`
}

type GetTrickerResponse []struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
