package region

type ListRequest struct {
	PerPage int64  `query:"per_page"`
	Cursor  string `query:"cursor"`
}
