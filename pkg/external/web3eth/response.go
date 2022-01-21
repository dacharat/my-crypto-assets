package web3eth

import "math/big"

type Token struct {
	Balance  *big.Int
	Symbol   string
	Decimals uint8
}
