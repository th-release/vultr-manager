package os

import "th-release/vultr-manager/utils"

type OsListResponse struct {
	OsList []OS       `json:"os"`
	Meta   utils.Meta `json:"meta"`
}

type OS struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Arch   string `json:"arch"`
	Family string `json:"family"`
}
