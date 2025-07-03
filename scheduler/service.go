package scheduler

import (
	"th-release/vultr-manager/api/application"
	"th-release/vultr-manager/api/firewall"
	"th-release/vultr-manager/api/instance"
	"th-release/vultr-manager/api/os"
	"th-release/vultr-manager/api/plan"
	"th-release/vultr-manager/api/region"
	"th-release/vultr-manager/api/script"
	"th-release/vultr-manager/utils"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func InsertInstance(db *pg.DB, instance instance.Instance) (*utils.Instance, orm.Result, error) {
	newInstance := &utils.Instance{
		Id:               instance.ID,
		AllowedBandwidth: instance.AllowedBandwidth,
		AppID:            instance.AppID,
		DateCreated:      instance.DateCreated,
		Disk:             instance.Disk,
		Features:         instance.Features,
		FirewallGroupID:  instance.FirewallGroupID,
		GatewayV4:        instance.GatewayV4,
		Hostname:         instance.Hostname,
		ImageID:          instance.ImageID,
		InternalIP:       instance.InternalIP,
		KVM:              instance.KVM,
		Label:            instance.Label,
		MainIP:           instance.MainIP,
		NetmaskV4:        instance.NetmaskV4,
		OS:               instance.OS,
		OsID:             instance.OsID,
		PendingCharges:   instance.PendingCharges,
		Plan:             instance.Plan,
		PowerStatus:      instance.PowerStatus,
		RAM:              instance.RAM,
		Region:           instance.Region,
		ServerStatus:     instance.ServerStatus,
		Status:           instance.Status,
		Tag:              instance.Tag,
		Tags:             instance.Tags,
		UserScheme:       instance.UserScheme,
		V6MainIP:         instance.V6MainIP,
		V6Network:        instance.V6Network,
		V6NetworkSize:    instance.V6NetworkSize,
		VcpuCount:        instance.VcpuCount,
		Vpcs:             instance.Vpcs,
	}

	var existing utils.Instance
	err := db.Model(&existing).Where("id = ?", newInstance.Id).Select()
	if err == nil {
		// 존재하면 업데이트
		result, err := db.Model(newInstance).
			Where("id = ?", instance.ID).
			Update()
		return newInstance, result, err
	} else if err != pg.ErrNoRows {
		// 다른 에러 발생 시 반환
		return nil, nil, err
	}

	result, err := db.Model(newInstance).Insert()

	return newInstance, result, err
}

func InsertApplication(db *pg.DB, application application.Application) (*utils.Application, orm.Result, error) {
	newApplication := &utils.Application{
		Id:         application.Id,
		Name:       application.Name,
		ShortName:  application.ShortName,
		DeployName: application.DeployName,
		Type:       application.Type,
		Vendor:     application.Vendor,
		ImageId:    application.ImageId,
	}

	var existing utils.Application
	err := db.Model(&existing).Where("id = ?", application.Id).Select()
	if err == nil {
		// 존재하면 업데이트
		result, err := db.Model(newApplication).
			Where("id = ?", application.Id).
			Update()
		return newApplication, result, err
	} else if err != pg.ErrNoRows {
		// 다른 에러 발생 시 반환
		return nil, nil, err
	}

	result, err := db.Model(newApplication).Insert()

	return newApplication, result, err
}

func InsertFirewallGroup(db *pg.DB, group firewall.FirewallGroup) (*utils.FirewallGroup, orm.Result, error) {
	newGroup := &utils.FirewallGroup{
		ID:            group.ID,
		Description:   group.Description,
		DateCreated:   group.DateCreated,
		DateModified:  group.DateModified,
		InstanceCount: group.InstanceCount,
		RuleCount:     group.RuleCount,
		MaxRuleCount:  group.MaxRuleCount,
	}

	// ID로 기존 레코드 확인
	var existing utils.FirewallGroup
	err := db.Model(&existing).Where("id = ?", group.ID).Select()
	if err == nil {
		// 존재하면 업데이트
		result, err := db.Model(newGroup).
			Where("id = ?", group.ID).
			Update()
		return newGroup, result, err
	} else if err != pg.ErrNoRows {
		// 다른 에러 발생 시 반환
		return nil, nil, err
	}

	// 존재하지 않으면 인서트
	result, err := db.Model(newGroup).Insert()
	return newGroup, result, err
}

func InsertFirewallRule(db *pg.DB, rule firewall.FirewallRule) (*utils.FirewallRule, orm.Result, error) {
	newRules := &utils.FirewallRule{
		Id:         rule.Id,
		Type:       rule.Type,
		IpType:     rule.IpType,
		Action:     rule.Action,
		Protocol:   rule.Protocol,
		Port:       rule.Port,
		Subnet:     rule.Subnet,
		SubnetSize: rule.SubnetSize,
		Source:     rule.Source,
		Notes:      rule.Notes,
	}

	var existing utils.FirewallRule
	err := db.Model(&existing).Where("id = ?", rule.Id).Select()
	if err == nil {
		// 존재하면 업데이트
		result, err := db.Model(newRules).
			Where("id = ?", rule.Id).
			Update()
		return newRules, result, err
	} else if err != pg.ErrNoRows {
		// 다른 에러 발생 시 반환
		return nil, nil, err
	}

	result, err := db.Model(newRules).Insert()

	return newRules, result, err
}

func InsertOs(db *pg.DB, os os.OS) (*utils.OS, orm.Result, error) {
	newOS := &utils.OS{
		Id:     os.Id,
		Name:   os.Name,
		Arch:   os.Arch,
		Family: os.Family,
	}

	var existing utils.OS
	err := db.Model(&existing).Where("id = ?", os.Id).Select()
	if err == nil {
		// 존재하면 업데이트
		result, err := db.Model(newOS).
			Where("id = ?", os.Id).
			Update()
		return newOS, result, err
	} else if err != pg.ErrNoRows {
		// 다른 에러 발생 시 반환
		return nil, nil, err
	}

	result, err := db.Model(newOS).Insert()

	return newOS, result, err
}

func InsertPlan(db *pg.DB, plan plan.Plan) (*utils.Plan, orm.Result, error) {
	newPlan := &utils.Plan{
		Bandwidth:   plan.Bandwidth,
		CpuVendor:   plan.CpuVendor,
		Disk:        plan.Disk,
		DiskCount:   plan.DiskCount,
		DiskType:    plan.DiskType,
		HourlyCost:  plan.HourlyCost,
		ID:          plan.ID,
		InvoiceType: plan.InvoiceType,
		Locations:   plan.Locations,
		MonthlyCost: plan.MonthlyCost,
		RAM:         plan.RAM,
		Type:        plan.Type,
		VcpuCount:   plan.VcpuCount,
	}

	var existing utils.Plan
	err := db.Model(&existing).Where("id = ?", plan.ID).Select()
	if err == nil {
		// 존재하면 업데이트
		result, err := db.Model(newPlan).
			Where("id = ?", plan.ID).
			Update()
		return newPlan, result, err
	} else if err != pg.ErrNoRows {
		// 다른 에러 발생 시 반환
		return nil, nil, err
	}

	result, err := db.Model(newPlan).Insert()

	return newPlan, result, err
}

func InsertRegion(db *pg.DB, region region.Region) (*utils.Region, orm.Result, error) {
	newRegion := &utils.Region{
		ID:        region.ID,
		City:      region.City,
		Country:   region.Country,
		Continent: region.Continent,
		Options:   region.Options,
	}

	var existing utils.Region
	err := db.Model(&existing).Where("id = ?", region.ID).Select()
	if err == nil {
		// 존재하면 업데이트
		result, err := db.Model(newRegion).
			Where("id = ?", region.ID).
			Update()
		return newRegion, result, err
	} else if err != pg.ErrNoRows {
		// 다른 에러 발생 시 반환
		return nil, nil, err
	}

	result, err := db.Model(newRegion).Insert()

	return newRegion, result, err
}

func InsertScript(db *pg.DB, script script.Script) (*utils.Script, orm.Result, error) {
	newScript := &utils.Script{
		DateCreated:  script.DateCreated,
		DateModified: script.DateModified,
		ID:           script.ID,
		Name:         script.Name,
		Type:         script.Type,
	}

	var existing utils.Script
	err := db.Model(&existing).Where("id = ?", script.ID).Select()
	if err == nil {
		// 존재하면 업데이트
		result, err := db.Model(newScript).
			Where("id = ?", script.ID).
			Update()
		return newScript, result, err
	} else if err != pg.ErrNoRows {
		// 다른 에러 발생 시 반환
		return nil, nil, err
	}

	result, err := db.Model(newScript).Insert()

	return newScript, result, err
}
