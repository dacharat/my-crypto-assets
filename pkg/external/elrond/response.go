package elrond

type GetAccountResponse struct {
	Address                  string `json:"address"`
	Balance                  string `json:"balance"`
	Nonce                    int    `json:"nonce"`
	Shard                    int    `json:"shard"`
	Code                     string `json:"code"`
	CodeHash                 string `json:"codeHash"`
	RootHash                 string `json:"rootHash"`
	TxCount                  int    `json:"txCount"`
	ScrCount                 int    `json:"scrCount"`
	Username                 string `json:"username"`
	DeveloperReward          string `json:"developerReward"`
	OwnerAddress             string `json:"ownerAddress"`
	DeployedAt               int    `json:"deployedAt"`
	IsUpgradeable            bool   `json:"isUpgradeable"`
	IsReadable               bool   `json:"isReadable"`
	IsPayable                bool   `json:"isPayable"`
	IsPayableBySmartContract bool   `json:"isPayableBySmartContract"`
	ScamInfo                 struct {
	} `json:"scamInfo"`
}

type GetAccountTokenResponse struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Ticker     string `json:"ticker"`
	Owner      string `json:"owner"`
	Minted     string `json:"minted"`
	Burnt      string `json:"burnt"`
	Decimals   int    `json:"decimals"`
	IsPaused   bool   `json:"isPaused"`
	Assets     struct {
		Website         string `json:"website"`
		Description     string `json:"description"`
		LedgerSignature string `json:"ledgerSignature"`
		Social          struct {
			Twitter string `json:"twitter"`
		} `json:"social"`
		Status string `json:"status"`
		PngURL string `json:"pngUrl"`
		SvgURL string `json:"svgUrl"`
	} `json:"assets"`
	CanUpgrade     bool   `json:"canUpgrade"`
	CanMint        bool   `json:"canMint"`
	CanBurn        bool   `json:"canBurn"`
	CanChangeOwner bool   `json:"canChangeOwner"`
	CanPause       bool   `json:"canPause"`
	CanFreeze      bool   `json:"canFreeze"`
	CanWipe        bool   `json:"canWipe"`
	Balance        string `json:"balance"`
}

type GetAccountDelegationResponse struct {
	Address             string        `json:"address"`
	Contract            string        `json:"contract"`
	UserUnBondable      string        `json:"userUnBondable"`
	UserActiveStake     string        `json:"userActiveStake"`
	ClaimableRewards    string        `json:"claimableRewards"`
	UserUndelegatedList []interface{} `json:"userUndelegatedList"`
}

type GetAccountNftResponse struct {
	Identifier string `json:"identifier"`
	Collection string `json:"collection"`
	Timestamp  int    `json:"timestamp"`
	Attributes string `json:"attributes"`
	Nonce      int    `json:"nonce"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	Creator    string `json:"creator"`
	Royalties  struct {
	} `json:"royalties"`
	Uris  []string `json:"uris"`
	URL   string   `json:"url"`
	Media struct {
	} `json:"media"`
	IsWhitelistedStorage bool     `json:"isWhitelistedStorage"`
	ThumbnailURL         string   `json:"thumbnailUrl"`
	Tags                 []string `json:"tags"`
	Metadata             struct {
	} `json:"metadata"`
	Owner    string `json:"owner"`
	Balance  string `json:"balance"`
	Supply   string `json:"supply"`
	Decimals struct {
	} `json:"decimals"`
	Assets struct {
		Website     string `json:"website"`
		Description string `json:"description"`
		Status      string `json:"status"`
		PngURL      string `json:"pngUrl"`
		SvgURL      string `json:"svgUrl"`
	} `json:"assets"`
	Ticker   string `json:"ticker"`
	ScamInfo struct {
	} `json:"scamInfo"`
}
