package common

type (
	Yield[V any]                      func(V)
	Generator[V any]                  func(Yield[V])
	IteratorFunc[V any]               func() (V, error)
	Filter[V any]                     func(V) (bool, error)
	Transformer[V1 any, V2 any]       func(V1) (V2, error)
	Mapper[V any]                     Transformer[V, V]
	Walker[V any]                     func(V) error
	Reducer[V1 any, V2 any]           func(V1, ...V2) (V2, error)
	Folder[V any]                     Reducer[V, V]
	SimpleIteratorFunc[V any]         func() (V, bool)
	SimpleFilter[V any]               func(V) bool
	SimpleTransformer[V1 any, V2 any] func(V1) V2
	SimpleMapper[V any]               SimpleTransformer[V, V]
	SimpleWalker[V any]               func(V)
	SimpleReducer[V1 any, V2 any]     func(V1) V2
	SimpleFolder[V any]               SimpleReducer[V, V]
)
