package application

type ListRequest struct {
	Type    ApplicationType `query:"type" json:"type"`
	PerPage int64           `query:"per_page" json:"per_page"`
	Cursor  string          `query:"cursor" json:"cursor"`
}
