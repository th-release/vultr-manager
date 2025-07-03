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

	result, err := db.Model(newScript).Insert()

	return newScript, result, err
}
