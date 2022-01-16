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

type AccountTransactionResponse struct {
	CurrentRound int    `json:"current-round"`
	NextToken    string `json:"next-token"`
	Transactions []struct {
		ApplicationTransaction struct {
			ApplicationID    int      `json:"application-id"`
			OnCompletion     string   `json:"on-completion"`
			ApplicationArgs  []string `json:"application-args"`
			Accounts         []string `json:"accounts"`
			ForeignApps      []int    `json:"foreign-apps"`
			ForeignAssets    []int    `json:"foreign-assets"`
			LocalStateSchema struct {
				NumUint      int `json:"num-uint"`
				NumByteSlice int `json:"num-byte-slice"`
			} `json:"local-state-schema"`
			GlobalStateSchema struct {
				NumUint      int `json:"num-uint"`
				NumByteSlice int `json:"num-byte-slice"`
			} `json:"global-state-schema"`
			ApprovalProgram   string `json:"approval-program"`
			ClearStateProgram string `json:"clear-state-program"`
			ExtraProgramPages int    `json:"extra-program-pages"`
		} `json:"application-transaction"`
		AssetConfigTransaction struct {
			AssetID int `json:"asset-id"`
			Params  struct {
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
		} `json:"asset-config-transaction"`
		AssetFreezeTransaction struct {
			Address         string `json:"address"`
			AssetID         int    `json:"asset-id"`
			NewFreezeStatus bool   `json:"new-freeze-status"`
		} `json:"asset-freeze-transaction"`
		AssetTransferTransaction struct {
			Amount      int    `json:"amount"`
			AssetID     int    `json:"asset-id"`
			CloseAmount int    `json:"close-amount"`
			CloseTo     string `json:"close-to"`
			Receiver    string `json:"receiver"`
			Sender      string `json:"sender"`
		} `json:"asset-transfer-transaction"`
		AuthAddr                string `json:"auth-addr"`
		CloseRewards            int    `json:"close-rewards"`
		ClosingAmount           int    `json:"closing-amount"`
		ConfirmedRound          int    `json:"confirmed-round"`
		CreatedApplicationIndex int    `json:"created-application-index"`
		CreatedAssetIndex       int    `json:"created-asset-index"`
		Fee                     int    `json:"fee"`
		FirstValid              int    `json:"first-valid"`
		GenesisHash             string `json:"genesis-hash"`
		GenesisID               string `json:"genesis-id"`
		Group                   string `json:"group"`
		ID                      string `json:"id"`
		IntraRoundOffset        int    `json:"intra-round-offset"`
		KeyregTransaction       struct {
			NonParticipation          bool   `json:"non-participation"`
			SelectionParticipationKey string `json:"selection-participation-key"`
			VoteFirstValid            int    `json:"vote-first-valid"`
			VoteKeyDilution           int    `json:"vote-key-dilution"`
			VoteLastValid             int    `json:"vote-last-valid"`
			VoteParticipationKey      string `json:"vote-participation-key"`
		} `json:"keyreg-transaction"`
		LastValid          int    `json:"last-valid"`
		Lease              string `json:"lease"`
		Note               string `json:"note"`
		PaymentTransaction struct {
			Amount           int    `json:"amount"`
			CloseAmount      int    `json:"close-amount"`
			CloseRemainderTo string `json:"close-remainder-to"`
			Receiver         string `json:"receiver"`
		} `json:"payment-transaction"`
		ReceiverRewards int    `json:"receiver-rewards"`
		RekeyTo         string `json:"rekey-to"`
		RoundTime       int    `json:"round-time"`
		Sender          string `json:"sender"`
		SenderRewards   int    `json:"sender-rewards"`
		Signature       struct {
			Logicsig struct {
				Args              []string `json:"args"`
				Logic             string   `json:"logic"`
				MultisigSignature struct {
					Subsignature []struct {
						PublicKey string `json:"public-key"`
						Signature string `json:"signature"`
					} `json:"subsignature"`
					Threshold int `json:"threshold"`
					Version   int `json:"version"`
				} `json:"multisig-signature"`
				Signature string `json:"signature"`
			} `json:"logicsig"`
			Multisig struct {
				Subsignature []struct {
					PublicKey string `json:"public-key"`
					Signature string `json:"signature"`
				} `json:"subsignature"`
				Threshold int `json:"threshold"`
				Version   int `json:"version"`
			} `json:"multisig"`
			Sig string `json:"sig"`
		} `json:"signature"`
		TxType          string `json:"tx-type"`
		LocalStateDelta []struct {
			Address string `json:"address"`
			Delta   []struct {
				Key   string `json:"key"`
				Value struct {
					Action int    `json:"action"`
					Bytes  string `json:"bytes"`
					Uint   int    `json:"uint"`
				} `json:"value"`
			} `json:"delta"`
		} `json:"local-state-delta"`
		GlobalStateDelta []struct {
			Key   string `json:"key"`
			Value struct {
				Action int    `json:"action"`
				Bytes  string `json:"bytes"`
				Uint   int    `json:"uint"`
			} `json:"value"`
		} `json:"global-state-delta"`
		Logs      []string `json:"logs"`
		InnerTxns []string `json:"inner-txns"`
	} `json:"transactions"`
}
