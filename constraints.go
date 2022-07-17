package common

import (
	"golang.org/x/exp/constraints"
)

type (
	Key         interface{ comparable }
	SortableKey interface{ constraints.Ordered }
	Index       interface {
		constraints.Integer | constraints.Float
	}
)
