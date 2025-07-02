package script

type DetailResponse struct {
	StartUpScript StartUpScript `json:"startup_script"`
}

type StartUpScript struct {
	DateCreated  string     `json:"date_created"`
	DateModified string     `json:"date_modified"`
	Id           string     `json:"id"`
	Name         string     `json:"name"`
	Script       string     `json:"script"`
	Type         ScriptType `json:"type"`
}
