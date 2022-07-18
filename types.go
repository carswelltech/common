package common

import "golang.org/x/exp/constraints"

type (
	KVPair[K Key, V any] struct {
		Key   K
		Value V
	}
	ArrayElement[I constraints.Integer, V any] struct {
		Index I
		Value V
	}
)
