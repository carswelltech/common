package common

import (
	"golang.org/x/exp/constraints"
)

type (

	// Key is a type constraint for associative arrays, and generic map types.  Defined here only for convenience and consistency across modules.
	Key interface{ comparable }

	// SortableKey is a type constraint for associative arrays, and generic types based on map, with an ordered quality. Defined here only for convenience and consistency across modules.
	SortableKey interface{ constraints.Ordered }

	// Index is a type constraint for indices for array-like generic types. Defined here only for convenience and consistency across modules.
	Index interface {
		constraints.Integer | constraints.Float
	}
)
