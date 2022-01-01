package algorand

type AccountResponse struct {
	Account      Account `json:"account"`
	CurrentRound int     `json:"current-round"`
}

type AssetResponse struct {
	Asset        AssetDetail `json:"asset"`
	CurrentRound int         `json:"current-round"`
}

type Account struct {
	Address                     string           `json:"address"`
	Amount                      int              `json:"amount"`
	AmountWithoutPendingRewards int              `json:"amount-without-pending-rewards"`
	AppsLocalState              []AppsLocalState `json:"apps-local-state"`
	AppsTotalSchema             AppsTotalSchema  `json:"apps-total-schema"`
	AppsTotalExtraPages         int              `json:"apps-total-extra-pages"`
	Assets                      []Asset          `json:"assets"`
	CreatedApps                 []CreatedApp     `json:"created-apps"`
	CreatedAssets               []CreatedAsset   `json:"created-assets"`
	Participation               Participation    `json:"participation"`
	PendingRewards              int              `json:"pending-rewards"`
	RewardBase                  int              `json:"reward-base"`
	Rewards                     int              `json:"rewards"`
	Round                       int              `json:"round"`
	Status                      string           `json:"status"`
	SigType                     string           `json:"sig-type"`
	AuthAddr                    string           `json:"auth-addr"`
	Deleted                     bool             `json:"deleted"`
	CreatedAtRound              int              `json:"created-at-round"`
	ClosedAtRound               int              `json:"closed-at-round"`
}

type AppsLocalState struct {
	ID               int  `json:"id"`
	Deleted          bool `json:"deleted"`
	OptedInAtRound   int  `json:"opted-in-at-round"`
	ClosedOutAtRound int  `json:"closed-out-at-round"`
	Schema           struct {
		NumUint      int `json:"num-uint"`
		NumByteSlice int `json:"num-byte-slice"`
	} `json:"schema"`
	KeyValue []struct {
		Key   string `json:"key"`
		Value struct {
			Type  int    `json:"type"`
			Bytes string `json:"bytes"`
			Uint  int    `json:"uint"`
		} `json:"value"`
	} `json:"key-value"`
}

type AppsTotalSchema struct {
	NumUint      int `json:"num-uint"`
	NumByteSlice int `json:"num-byte-slice"`
}

type Asset struct {
	Amount          int    `json:"amount"`
	AssetID         int    `json:"asset-id"`
	Creator         string `json:"creator"`
	IsFrozen        bool   `json:"is-frozen"`
	Deleted         bool   `json:"deleted"`
	OptedInAtRound  int    `json:"opted-in-at-round"`
	OptedOutAtRound int    `json:"opted-out-at-round"`
}

type CreatedApp struct {
	ID             int  `json:"id"`
	Deleted        bool `json:"deleted"`
	CreatedAtRound int  `json:"created-at-round"`
	DeletedAtRound int  `json:"deleted-at-round"`
	Params         struct {
		Creator           string `json:"creator"`
		ApprovalProgram   string `json:"approval-program"`
		ClearStateProgram string `json:"clear-state-program"`
		LocalStateSchema  struct {
			NumUint      int `json:"num-uint"`
			NumByteSlice int `json:"num-byte-slice"`
		} `json:"local-state-schema"`
		GlobalStateSchema struct {
			NumUint      int `json:"num-uint"`
			NumByteSlice int `json:"num-byte-slice"`
		} `json:"global-state-schema"`
		GlobalState []struct {
			Key   string `json:"key"`
			Value struct {
				Type  int    `json:"type"`
				Bytes string `json:"bytes"`
				Uint  int    `json:"uint"`
			} `json:"value"`
		} `json:"global-state"`
		ExtraProgramPages int `json:"extra-program-pages"`
	} `json:"params"`
}

type CreatedAsset struct {
	Index            int  `json:"index"`
	Deleted          bool `json:"deleted"`
	CreatedAtRound   int  `json:"created-at-round"`
	DestroyedAtRound int  `json:"destroyed-at-round"`
	Params           struct {
		Clawback      string `json:"clawback"`
		Creator       string `json:"creator"`
		Decimals      int    `json:"decimals"`
		DefaultFrozen bool   `json:"default-frozen"`
		Freeze        string `json:"freeze"`
		Manager       string `json:"manager"`
		MetadataHash  string `json:"metadata-hash"`
		Name          string `json:"name"`
		NameB64       string `json:"name-b64"`
		Reserve       string `json:"reserve"`
		Total         int    `json:"total"`
		UnitName      string `json:"unit-name"`
		UnitNameB64   string `json:"unit-name-b64"`
		URL           string `json:"url"`
		URLB64        string `json:"url-b64"`
	} `json:"params"`
}

type Participation struct {
	SelectionParticipationKey string `json:"selection-participation-key"`
	VoteFirstValid            int    `json:"vote-first-valid"`
	VoteKeyDilution           int    `json:"vote-key-dilution"`
	VoteLastValid             int    `json:"vote-last-valid"`
	VoteParticipationKey      string `json:"vote-participation-key"`
}

type AssetDetail struct {
	Index            int  `json:"index"`
	Deleted          bool `json:"deleted"`
	CreatedAtRound   int  `json:"created-at-round"`
	DestroyedAtRound int  `json:"destroyed-at-round"`
	Params           struct {
		Clawback      string `json:"clawback"`
		Creator       string `json:"creator"`
		Decimals      int    `json:"decimals"`
		DefaultFrozen bool   `json:"default-frozen"`
		Freeze        string `json:"freeze"`
		Manager       string `json:"manager"`
		MetadataHash  string `json:"metadata-hash"`
		Name          string `json:"name"`
		NameB64       string `json:"name-b64"`
		Reserve       string `json:"reserve"`
		Total         int    `json:"total"`
		UnitName      string `json:"unit-name"`
		UnitNameB64   string `json:"unit-name-b64"`
		URL           string `json:"url"`
		URLB64        string `json:"url-b64"`
	} `json:"params"`
}
