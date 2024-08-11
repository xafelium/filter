package filter

type Query struct {
	Filter Condition
	Page   Pagination
	Fields []string
}

func DefaultQuery() Query {
	return Query{
		Filter: Where(nil),
		Page:   DefaultPagination(),
		Fields: nil,
	}
}

func (o *Query) Options() []QueryOption {
	return []QueryOption{
		WithQueryFilter(o.Filter),
		WithQueryPage(o.Page),
		WithQueryFields(o.Fields),
	}
}

func QueryWithOptions(opts ...QueryOption) Query {
	o := DefaultQuery()
	for _, applyOption := range opts {
		applyOption(&o)
	}
	return o
}

type QueryOption func(o *Query)

func WithQueryFilter(f Condition) QueryOption {
	return func(o *Query) {
		o.Filter = f
	}
}

func WithQueryPage(p Pagination) QueryOption {
	return func(o *Query) {
		o.Page = p
	}
}

func WithQueryFields(fields []string) QueryOption {
	return func(o *Query) {
		o.Fields = fields
	}
}
