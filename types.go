package common

import "golang.org/x/exp/constraints"

type (

	// KVPair is a struct to represent a key/value pair for a map-like generic type.  Useful for iterating maps.
	KVPair[K Key, V any] struct {
		Key   K
		Value V
	}

	// ArrayElement is a struct to represent a key/value pair for an array-like generic type.  Useful for iterators.
	ArrayElement[I constraints.Integer, V any] struct {
		Index I
		Value V
	}
)
