package common

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type (
	IndexOutOfBoundsError[I constraints.Integer] struct {
		Index I
	}
)

func (e *IndexOutOfBoundsError[I]) Error() string {
	return fmt.Sprintf("index out of bounds (%v)", e.Index)
}
