package filter

type TotalAware interface {
	TotalCount() int
	SetTotalCount(total int)
	TotalPages() int
	SetTotalPages(total int)
}

type ResultPage[T any] interface {
	TotalAware
	SetItems(results []*T)
	Items() []*T
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewEmptyResultPage[T any]() *resultPage[T] {
	return NewResultPage[T](0, 0, nil)
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewResultPage[T any](totalCount, totalPages int, items []*T) *resultPage[T] {
	return &resultPage[T]{
		totalCount: totalCount,
		totalPages: totalPages,
		items:      items,
	}
}

type resultPage[T any] struct {
	totalCount int
	totalPages int
	items      []*T
}

func (p *resultPage[T]) TotalCount() int {
	return p.totalCount
}

func (p *resultPage[T]) SetTotalCount(total int) {
	p.totalCount = total
}

func (p *resultPage[T]) TotalPages() int {
	return p.totalPages
}

func (p *resultPage[T]) SetTotalPages(total int) {
	p.totalPages = total
}

func (p *resultPage[T]) Items() []*T {
	return p.items
}

func (p *resultPage[T]) SetItems(results []*T) {
	p.items = results
}
