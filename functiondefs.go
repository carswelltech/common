package common

type (
	Filter[V any]               func(V) (bool, error)
	Transformer[V1 any, V2 any] func(V1) (V2, error)
	Mapper[V any]               Transformer[V, V]
	Walker[V any]               func(V) error
	Reducer[V1 any, V2 any]     func(V1, ...V2) (V2, error)
	Folder[V any]               Reducer[V, V]
)
