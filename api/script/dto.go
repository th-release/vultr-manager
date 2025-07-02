package script

type ListRequest struct {
	PerPage int64  `query:"per_page"`
	Cursor  string `query:"cursor"`
}

type ScriptCreateRequest struct {
	Name   string     `json:"name"`
	Script string     `json:"script"`
	Type   ScriptType `json:"type"`
}

type ScriptUpdateRequest struct {
	Name   string     `json:"name"`
	Script string     `json:"script"`
	Type   ScriptType `json:"type"`
}
