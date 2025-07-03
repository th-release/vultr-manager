package script

import "th-release/vultr-manager/utils"

type ListResponse struct {
	StartScripts []Script   `json:"startup_scripts"`
	Meta         utils.Meta `json:"meta"`
}

type DetailResponse struct {
	StartUpScript StartUpScript `json:"startup_script"`
}

type Script struct {
	DateCreated  string `json:"date_created"`
	DateModified string `json:"date_modified"`
	ID           string `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

type StartUpScript struct {
	DateCreated  string     `json:"date_created"`
	DateModified string     `json:"date_modified"`
	Id           string     `json:"id"`
	Name         string     `json:"name"`
	Script       string     `json:"script"`
	Type         ScriptType `json:"type"`
}
