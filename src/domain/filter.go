package domain

type Filter struct {
	CreatedAt FilterValue

	Limit int
}

type FilterValue struct {
	Value  interface{}
	Sort   string
	Equal  bool
	Bigger bool
	Lower  bool
}
