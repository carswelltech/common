package common

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

const (
	IndexOutOfBoundsErrorMsg string = "index out of bounds (%v)"
)

type (
	IndexOutOfBoundsError[I constraints.Integer] struct {
		Index I
	}
)

func (e *IndexOutOfBoundsError[I]) Error() string {
	return fmt.Sprintf(IndexOutOfBoundsErrorMsg, e.Index)
}
