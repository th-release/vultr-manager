package scheduler

import (
	"context"
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

	var result orm.Result
	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// DB에서 기존 인스턴스 조회
		var existingInstances []utils.Instance
		err := tx.Model(&existingInstances).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// DB에 있지만 입력 인스턴스 ID와 다른 경우 삭제
		for _, existing := range existingInstances {
			if existing.Id != instance.ID {
				_, err := tx.Model(&utils.Instance{}).Where("id = ?", existing.Id).Delete()
				if err != nil {
					return err
				}
			}
		}

		// ID로 기존 레코드 확인
		var existing utils.Instance
		err = tx.Model(&existing).Where("id = ?", instance.ID).Select()
		if err == nil {
			// 존재하면 업데이트
			result, err = tx.Model(newInstance).Where("id = ?", instance.ID).Update()
			return err
		} else if err != pg.ErrNoRows {
			// 다른 에러 반환
			return err
		}

		// 존재하지 않으면 인서트
		result, err = tx.Model(newInstance).Insert()
		return err
	})

	if err != nil {
		return nil, nil, err
	}

	return newInstance, result, nil
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

	var result orm.Result
	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// DB에서 기존 애플리케이션 조회
		var existingApplications []utils.Application
		err := tx.Model(&existingApplications).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// DB에 있지만 입력 애플리케이션 ID와 다른 경우 삭제
		for _, existing := range existingApplications {
			if existing.Id != application.Id {
				_, err := tx.Model(&utils.Application{}).Where("id = ?", existing.Id).Delete()
				if err != nil {
					return err
				}
			}
		}

		// ID로 기존 레코드 확인
		var existing utils.Application
		err = tx.Model(&existing).Where("id = ?", application.Id).Select()
		if err == nil {
			// 존재하면 업데이트
			result, err = tx.Model(newApplication).Where("id = ?", application.Id).Update()
			return err
		} else if err != pg.ErrNoRows {
			// 다른 에러 반환
			return err
		}

		// 존재하지 않으면 인서트
		result, err = tx.Model(newApplication).Insert()
		return err
	})

	if err != nil {
		return nil, nil, err
	}

	return newApplication, result, nil
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

	var result orm.Result
	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// DB에서 기존 그룹 조회
		var existingGroups []utils.FirewallGroup
		err := tx.Model(&existingGroups).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// DB에 있지만 입력 그룹 ID와 다른 경우 삭제
		for _, existing := range existingGroups {
			if existing.ID != group.ID {
				_, err := tx.Model(&utils.FirewallGroup{}).Where("id = ?", existing.ID).Delete()
				if err != nil {
					return err
				}
			}
		}

		// ID로 기존 레코드 확인
		var existing utils.FirewallGroup
		err = tx.Model(&existing).Where("id = ?", group.ID).Select()
		if err == nil {
			// 존재하면 업데이트
			result, err = tx.Model(newGroup).Where("id = ?", group.ID).Update()
			return err
		} else if err != pg.ErrNoRows {
			// 다른 에러 반환
			return err
		}

		// 존재하지 않으면 인서트
		result, err = tx.Model(newGroup).Insert()
		return err
	})

	if err != nil {
		return nil, nil, err
	}

	return newGroup, result, nil
}

func InsertFirewallRule(db *pg.DB, rule firewall.FirewallRule) (*utils.FirewallRule, orm.Result, error) {
	newRule := &utils.FirewallRule{
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

	var result orm.Result
	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// DB에서 기존 규칙 조회
		var existingRules []utils.FirewallRule
		err := tx.Model(&existingRules).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// DB에 있지만 입력 규칙 ID와 다른 경우 삭제
		for _, existing := range existingRules {
			if existing.Id != rule.Id {
				_, err := tx.Model(&utils.FirewallRule{}).Where("id = ?", existing.Id).Delete()
				if err != nil {
					return err
				}
			}
		}

		// ID로 기존 레코드 확인
		var existing utils.FirewallRule
		err = tx.Model(&existing).Where("id = ?", rule.Id).Select()
		if err == nil {
			// 존재하면 업데이트
			result, err = tx.Model(newRule).Where("id = ?", rule.Id).Update()
			return err
		} else if err != pg.ErrNoRows {
			// 다른 에러 반환
			return err
		}

		// 존재하지 않으면 인서트
		result, err = tx.Model(newRule).Insert()
		return err
	})

	if err != nil {
		return nil, nil, err
	}

	return newRule, result, nil
}

func InsertOs(db *pg.DB, os os.OS) (*utils.OS, orm.Result, error) {
	newOS := &utils.OS{
		Id:     os.Id,
		Name:   os.Name,
		Arch:   os.Arch,
		Family: os.Family,
	}

	var result orm.Result
	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// DB에서 기존 OS 조회
		var existingOSes []utils.OS
		err := tx.Model(&existingOSes).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// DB에 있지만 입력 OS ID와 다른 경우 삭제
		for _, existing := range existingOSes {
			if existing.Id != os.Id {
				_, err := tx.Model(&utils.OS{}).Where("id = ?", existing.Id).Delete()
				if err != nil {
					return err
				}
			}
		}

		// ID로 기존 레코드 확인
		var existing utils.OS
		err = tx.Model(&existing).Where("id = ?", os.Id).Select()
		if err == nil {
			// 존재하면 업데이트
			result, err = tx.Model(newOS).Where("id = ?", os.Id).Update()
			return err
		} else if err != pg.ErrNoRows {
			// 다른 에러 반환
			return err
		}

		// 존재하지 않으면 인서트
		result, err = tx.Model(newOS).Insert()
		return err
	})

	if err != nil {
		return nil, nil, err
	}

	return newOS, result, nil
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

	var result orm.Result
	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// DB에서 기존 플랜 조회
		var existingPlans []utils.Plan
		err := tx.Model(&existingPlans).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// DB에 있지만 입력 플랜 ID와 다른 경우 삭제
		for _, existing := range existingPlans {
			if existing.ID != plan.ID {
				_, err := tx.Model(&utils.Plan{}).Where("id = ?", existing.ID).Delete()
				if err != nil {
					return err
				}
			}
		}

		// ID로 기존 레코드 확인
		var existing utils.Plan
		err = tx.Model(&existing).Where("id = ?", plan.ID).Select()
		if err == nil {
			// 존재하면 업데이트
			result, err = tx.Model(newPlan).Where("id = ?", plan.ID).Update()
			return err
		} else if err != pg.ErrNoRows {
			// 다른 에러 반환
			return err
		}

		// 존재하지 않으면 인서트
		result, err = tx.Model(newPlan).Insert()
		return err
	})

	if err != nil {
		return nil, nil, err
	}

	return newPlan, result, nil
}

func InsertRegion(db *pg.DB, region region.Region) (*utils.Region, orm.Result, error) {
	newRegion := &utils.Region{
		ID:        region.ID,
		City:      region.City,
		Country:   region.Country,
		Continent: region.Continent,
		Options:   region.Options,
	}

	var result orm.Result
	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// DB에서 기존 리전 조회
		var existingRegions []utils.Region
		err := tx.Model(&existingRegions).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// DB에 있지만 입력 리전 ID와 다른 경우 삭제
		for _, existing := range existingRegions {
			if existing.ID != region.ID {
				_, err := tx.Model(&utils.Region{}).Where("id = ?", existing.ID).Delete()
				if err != nil {
					return err
				}
			}
		}

		// ID로 기존 레코드 확인
		var existing utils.Region
		err = tx.Model(&existing).Where("id = ?", region.ID).Select()
		if err == nil {
			// 존재하면 업데이트
			result, err = tx.Model(newRegion).Where("id = ?", region.ID).Update()
			return err
		} else if err != pg.ErrNoRows {
			// 다른 에러 반환
			return err
		}

		// 존재하지 않으면 인서트
		result, err = tx.Model(newRegion).Insert()
		return err
	})

	if err != nil {
		return nil, nil, err
	}

	return newRegion, result, nil
}

func InsertScript(db *pg.DB, script script.Script) (*utils.Script, orm.Result, error) {
	newScript := &utils.Script{
		DateCreated:  script.DateCreated,
		DateModified: script.DateModified,
		ID:           script.ID,
		Name:         script.Name,
		Type:         script.Type,
	}

	var result orm.Result
	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// DB에서 기존 스크립트 조회
		var existingScripts []utils.Script
		err := tx.Model(&existingScripts).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// DB에 있지만 입력 스크립트 ID와 다른 경우 삭제
		for _, existing := range existingScripts {
			if existing.ID != script.ID {
				_, err := tx.Model(&utils.Script{}).Where("id = ?", existing.ID).Delete()
				if err != nil {
					return err
				}
			}
		}

		// ID로 기존 레코드 확인
		var existing utils.Script
		err = tx.Model(&existing).Where("id = ?", script.ID).Select()
		if err == nil {
			// 존재하면 업데이트
			result, err = tx.Model(newScript).Where("id = ?", script.ID).Update()
			return err
		} else if err != pg.ErrNoRows {
			// 다른 에러 반환
			return err
		}

		// 존재하지 않으면 인서트
		result, err = tx.Model(newScript).Insert()
		return err
	})

	if err != nil {
		return nil, nil, err
	}

	return newScript, result, nil
}
