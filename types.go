package common

type (
	KVPair[K Key, V any] struct {
		Key   K
		Value V
	}
)
