package algorand

import (
	"fmt"
	"strings"
)

type Query []string
type QueryOption func(*Query)

func (q Query) String() string {
	return strings.Join(q, "&")
}

func WithLimit(limit int) QueryOption {
	return func(q *Query) {
		*q = append(*q, fmt.Sprintf("limit=%d", limit))
	}
}

func WithAssetID(id int) QueryOption {
	return func(q *Query) {
		*q = append(*q, fmt.Sprintf("asset-id=%d", id))
	}
}

func WithCurrencyGreaterThan(amount int) QueryOption {
	return func(q *Query) {
		*q = append(*q, fmt.Sprintf("currency-greater-than=%d", amount))
	}
}
