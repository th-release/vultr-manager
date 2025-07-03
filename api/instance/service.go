package instance

import (
	"th-release/vultr-manager/utils"

	"github.com/go-pg/pg/v10"
	"github.com/go-resty/resty/v2"
)

func InstanceList(queryParams map[string]string, token string) (*resty.Response, InstsanceListResponse, interface{}, error) {
	client := resty.New()
	resp, res, errResp, err := utils.GetRequest[InstsanceListResponse](client, "https://api.vultr.com/v2/instances", queryParams, token)
	if err != nil {
		return nil, res, errResp, err
	}

	return resp, res, errResp, nil
}

func DatabaseInstanceList(db *pg.DB, page, limit int) ([]utils.Instance, error) {
	var instances []utils.Instance

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := db.Model(&instances).
		Limit(limit).
		Offset(offset)

	err := query.Select()
	if err != nil {
		return nil, err
	}

	if instances == nil {
		instances = []utils.Instance{}
	}

	return instances, err
}

func DatabaseInstanceDetail(db *pg.DB, uuid string) (*utils.Instance, error) {
	instance := new(utils.Instance)
	err := db.Model(instance).Where("id = ?", uuid).Select()

	return instance, err
}
