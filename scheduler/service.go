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
)

func MappingInstances(db *pg.DB, instances []instance.Instance) ([]utils.Instance, error) {
	var updatedInstances []utils.Instance

	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// 1. 현재 DB에 있는 모든 인스턴스 조회
		var existingInstances []utils.Instance
		err := tx.Model(&existingInstances).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// 2. 기존 인스턴스 ID를 맵으로 저장 (빠른 검색을 위해)
		existingIDMap := make(map[string]bool)
		for _, existing := range existingInstances {
			existingIDMap[existing.Id] = true
		}

		// 3. 입력 배열의 ID를 맵으로 저장
		inputIDMap := make(map[string]bool)
		for _, inst := range instances {
			inputIDMap[inst.ID] = true
		}

		// 4. DB에는 있지만 입력 배열에는 없는 인스턴스 삭제
		for _, existing := range existingInstances {
			if !inputIDMap[existing.Id] {
				_, err := tx.Model(&utils.Instance{}).Where("id = ?", existing.Id).Delete()
				if err != nil {
					return err
				}
			}
		}

		// 5. 입력 배열의 각 인스턴스에 대해 업데이트 또는 삽입
		for _, inst := range instances {
			newInstance := &utils.Instance{
				Id:               inst.ID,
				AllowedBandwidth: inst.AllowedBandwidth,
				AppID:            inst.AppID,
				DateCreated:      inst.DateCreated,
				Disk:             inst.Disk,
				Features:         inst.Features,
				FirewallGroupID:  inst.FirewallGroupID,
				GatewayV4:        inst.GatewayV4,
				Hostname:         inst.Hostname,
				ImageID:          inst.ImageID,
				InternalIP:       inst.InternalIP,
				KVM:              inst.KVM,
				Label:            inst.Label,
				MainIP:           inst.MainIP,
				NetmaskV4:        inst.NetmaskV4,
				OS:               inst.OS,
				OsID:             inst.OsID,
				PendingCharges:   inst.PendingCharges,
				Plan:             inst.Plan,
				PowerStatus:      inst.PowerStatus,
				RAM:              inst.RAM,
				Region:           inst.Region,
				ServerStatus:     inst.ServerStatus,
				Status:           inst.Status,
				Tag:              inst.Tag,
				Tags:             inst.Tags,
				UserScheme:       inst.UserScheme,
				V6MainIP:         inst.V6MainIP,
				V6Network:        inst.V6Network,
				V6NetworkSize:    inst.V6NetworkSize,
				VcpuCount:        inst.VcpuCount,
				Vpcs:             inst.Vpcs,
			}

			if existingIDMap[inst.ID] {
				// 기존에 존재하면 업데이트
				_, err := tx.Model(newInstance).Where("id = ?", inst.ID).Update()
				if err != nil {
					return err
				}
			} else {
				// 존재하지 않으면 삽입
				_, err := tx.Model(newInstance).Insert()
				if err != nil {
					return err
				}
			}

			updatedInstances = append(updatedInstances, *newInstance)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedInstances, nil
}

// 방법 1: 기본 접근 방식 (개별 업데이트/삽입)
func MappingApplications(db *pg.DB, applications []application.Application) ([]utils.Application, error) {
	var updatedApplications []utils.Application

	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// 1. 현재 DB에 있는 모든 애플리케이션 조회
		var existingApplications []utils.Application
		err := tx.Model(&existingApplications).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// 2. 기존 애플리케이션 ID를 맵으로 저장 (빠른 검색을 위해)
		existingIDMap := make(map[int]bool)
		for _, existing := range existingApplications {
			existingIDMap[existing.Id] = true
		}

		// 3. 입력 배열의 ID를 맵으로 저장
		inputIDMap := make(map[int]bool)
		for _, app := range applications {
			inputIDMap[app.Id] = true
		}

		// 4. DB에는 있지만 입력 배열에는 없는 애플리케이션 삭제
		for _, existing := range existingApplications {
			if !inputIDMap[existing.Id] {
				_, err := tx.Model(&utils.Application{}).Where("id = ?", existing.Id).Delete()
				if err != nil {
					return err
				}
			}
		}

		// 5. 입력 배열의 각 애플리케이션에 대해 업데이트 또는 삽입
		for _, app := range applications {
			newApplication := &utils.Application{
				Id:         app.Id,
				Name:       app.Name,
				ShortName:  app.ShortName,
				DeployName: app.DeployName,
				Type:       app.Type,
				Vendor:     app.Vendor,
				ImageId:    app.ImageId,
			}

			if existingIDMap[app.Id] {
				// 기존에 존재하면 업데이트
				_, err := tx.Model(newApplication).Where("id = ?", app.Id).Update()
				if err != nil {
					return err
				}
			} else {
				// 존재하지 않으면 삽입
				_, err := tx.Model(newApplication).Insert()
				if err != nil {
					return err
				}
			}

			updatedApplications = append(updatedApplications, *newApplication)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedApplications, nil
}

func MappingFirewallGroups(db *pg.DB, groups []firewall.FirewallGroup) ([]utils.FirewallGroup, error) {
	var updatedGroups []utils.FirewallGroup

	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// 1. 현재 DB에 있는 모든 방화벽 그룹 조회
		var existingGroups []utils.FirewallGroup
		err := tx.Model(&existingGroups).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// 2. 기존 그룹 ID를 맵으로 저장 (빠른 검색을 위해)
		existingIDMap := make(map[string]bool)
		for _, existing := range existingGroups {
			existingIDMap[existing.ID] = true
		}

		// 3. 입력 배열의 ID를 맵으로 저장
		inputIDMap := make(map[string]bool)
		for _, group := range groups {
			inputIDMap[group.ID] = true
		}

		// 4. DB에는 있지만 입력 배열에는 없는 그룹 삭제
		for _, existing := range existingGroups {
			if !inputIDMap[existing.ID] {
				_, err := tx.Model(&utils.FirewallGroup{}).Where("id = ?", existing.ID).Delete()
				if err != nil {
					return err
				}
			}
		}

		// 5. 입력 배열의 각 그룹에 대해 업데이트 또는 삽입
		for _, group := range groups {
			newGroup := &utils.FirewallGroup{
				ID:            group.ID,
				Description:   group.Description,
				DateCreated:   group.DateCreated,
				DateModified:  group.DateModified,
				InstanceCount: group.InstanceCount,
				RuleCount:     group.RuleCount,
				MaxRuleCount:  group.MaxRuleCount,
			}

			if existingIDMap[group.ID] {
				// 기존에 존재하면 업데이트
				_, err := tx.Model(newGroup).Where("id = ?", group.ID).Update()
				if err != nil {
					return err
				}
			} else {
				// 존재하지 않으면 삽입
				_, err := tx.Model(newGroup).Insert()
				if err != nil {
					return err
				}
			}

			updatedGroups = append(updatedGroups, *newGroup)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedGroups, nil
}

func MappingFirewallRules(db *pg.DB, rules []firewall.FirewallRule) ([]utils.FirewallRule, error) {
	var updatedRules []utils.FirewallRule

	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// 1. 현재 DB에 있는 모든 방화벽 규칙 조회
		var existingRules []utils.FirewallRule
		err := tx.Model(&existingRules).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// 2. 기존 규칙 ID를 맵으로 저장 (빠른 검색을 위해)
		existingIDMap := make(map[int64]bool)
		for _, existing := range existingRules {
			existingIDMap[existing.Id] = true
		}

		// 3. 입력 배열의 ID를 맵으로 저장
		inputIDMap := make(map[int64]bool)
		for _, rule := range rules {
			inputIDMap[rule.Id] = true
		}

		// 4. DB에는 있지만 입력 배열에는 없는 규칙 삭제
		for _, existing := range existingRules {
			if !inputIDMap[existing.Id] {
				_, err := tx.Model(&utils.FirewallRule{}).Where("id = ?", existing.Id).Delete()
				if err != nil {
					return err
				}
			}
		}

		// 5. 입력 배열의 각 규칙에 대해 업데이트 또는 삽입
		for _, rule := range rules {
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

			if existingIDMap[rule.Id] {
				// 기존에 존재하면 업데이트
				_, err := tx.Model(newRule).Where("id = ?", rule.Id).Update()
				if err != nil {
					return err
				}
			} else {
				// 존재하지 않으면 삽입
				_, err := tx.Model(newRule).Insert()
				if err != nil {
					return err
				}
			}

			updatedRules = append(updatedRules, *newRule)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedRules, nil
}

func MappingOs(db *pg.DB, oses []os.OS) ([]utils.OS, error) {
	var updatedOSes []utils.OS

	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// 1. 현재 DB에 있는 모든 OS 조회
		var existingOSes []utils.OS
		err := tx.Model(&existingOSes).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// 2. 기존 OS ID를 맵으로 저장 (빠른 검색을 위해)
		existingIDMap := make(map[int]bool)
		for _, existing := range existingOSes {
			existingIDMap[existing.Id] = true
		}

		// 3. 입력 배열의 ID를 맵으로 저장
		inputIDMap := make(map[int]bool)
		for _, osItem := range oses {
			inputIDMap[osItem.Id] = true
		}

		// 4. DB에는 있지만 입력 배열에는 없는 OS 삭제
		for _, existing := range existingOSes {
			if !inputIDMap[existing.Id] {
				_, err := tx.Model(&utils.OS{}).Where("id = ?", existing.Id).Delete()
				if err != nil {
					return err
				}
			}
		}

		// 5. 입력 배열의 각 OS에 대해 업데이트 또는 삽입
		for _, osItem := range oses {
			newOS := &utils.OS{
				Id:     osItem.Id,
				Name:   osItem.Name,
				Arch:   osItem.Arch,
				Family: osItem.Family,
			}

			if existingIDMap[osItem.Id] {
				// 기존에 존재하면 업데이트
				_, err := tx.Model(newOS).Where("id = ?", osItem.Id).Update()
				if err != nil {
					return err
				}
			} else {
				// 존재하지 않으면 삽입
				_, err := tx.Model(newOS).Insert()
				if err != nil {
					return err
				}
			}

			updatedOSes = append(updatedOSes, *newOS)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedOSes, nil
}

func MappingPlans(db *pg.DB, plans []plan.Plan) ([]utils.Plan, error) {
	var updatedPlans []utils.Plan

	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// 1. 현재 DB에 있는 모든 플랜 조회
		var existingPlans []utils.Plan
		err := tx.Model(&existingPlans).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// 2. 기존 플랜 ID를 맵으로 저장 (빠른 검색을 위해)
		existingIDMap := make(map[string]bool)
		for _, existing := range existingPlans {
			existingIDMap[existing.ID] = true
		}

		// 3. 입력 배열의 ID를 맵으로 저장
		inputIDMap := make(map[string]bool)
		for _, planItem := range plans {
			inputIDMap[planItem.ID] = true
		}

		// 4. DB에는 있지만 입력 배열에는 없는 플랜 삭제
		for _, existing := range existingPlans {
			if !inputIDMap[existing.ID] {
				_, err := tx.Model(&utils.Plan{}).Where("id = ?", existing.ID).Delete()
				if err != nil {
					return err
				}
			}
		}

		// 5. 입력 배열의 각 플랜에 대해 업데이트 또는 삽입
		for _, planItem := range plans {
			newPlan := &utils.Plan{
				Bandwidth:   planItem.Bandwidth,
				CpuVendor:   planItem.CpuVendor,
				Disk:        planItem.Disk,
				DiskCount:   planItem.DiskCount,
				DiskType:    planItem.DiskType,
				HourlyCost:  planItem.HourlyCost,
				ID:          planItem.ID,
				InvoiceType: planItem.InvoiceType,
				Locations:   planItem.Locations,
				MonthlyCost: planItem.MonthlyCost,
				RAM:         planItem.RAM,
				Type:        planItem.Type,
				VcpuCount:   planItem.VcpuCount,
			}

			if existingIDMap[planItem.ID] {
				// 기존에 존재하면 업데이트
				_, err := tx.Model(newPlan).Where("id = ?", planItem.ID).Update()
				if err != nil {
					return err
				}
			} else {
				// 존재하지 않으면 삽입
				_, err := tx.Model(newPlan).Insert()
				if err != nil {
					return err
				}
			}

			updatedPlans = append(updatedPlans, *newPlan)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedPlans, nil
}

func MappingRegions(db *pg.DB, regions []region.Region) ([]utils.Region, error) {
	var updatedRegions []utils.Region

	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// 1. 현재 DB에 있는 모든 리전 조회
		var existingRegions []utils.Region
		err := tx.Model(&existingRegions).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// 2. 기존 리전 ID를 맵으로 저장 (빠른 검색을 위해)
		existingIDMap := make(map[string]bool)
		for _, existing := range existingRegions {
			existingIDMap[existing.ID] = true
		}

		// 3. 입력 배열의 ID를 맵으로 저장
		inputIDMap := make(map[string]bool)
		for _, regionItem := range regions {
			inputIDMap[regionItem.ID] = true
		}

		// 4. DB에는 있지만 입력 배열에는 없는 리전 삭제
		for _, existing := range existingRegions {
			if !inputIDMap[existing.ID] {
				_, err := tx.Model(&utils.Region{}).Where("id = ?", existing.ID).Delete()
				if err != nil {
					return err
				}
			}
		}

		// 5. 입력 배열의 각 리전에 대해 업데이트 또는 삽입
		for _, regionItem := range regions {
			newRegion := &utils.Region{
				ID:        regionItem.ID,
				City:      regionItem.City,
				Country:   regionItem.Country,
				Continent: regionItem.Continent,
				Options:   regionItem.Options,
			}

			if existingIDMap[regionItem.ID] {
				// 기존에 존재하면 업데이트
				_, err := tx.Model(newRegion).Where("id = ?", regionItem.ID).Update()
				if err != nil {
					return err
				}
			} else {
				// 존재하지 않으면 삽입
				_, err := tx.Model(newRegion).Insert()
				if err != nil {
					return err
				}
			}

			updatedRegions = append(updatedRegions, *newRegion)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedRegions, nil
}

func MappingScripts(db *pg.DB, scripts []script.Script) ([]utils.Script, error) {
	var updatedScripts []utils.Script

	err := db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		// 1. 현재 DB에 있는 모든 스크립트 조회
		var existingScripts []utils.Script
		err := tx.Model(&existingScripts).Select()
		if err != nil && err != pg.ErrNoRows {
			return err
		}

		// 2. 기존 스크립트 ID를 맵으로 저장 (빠른 검색을 위해)
		existingIDMap := make(map[string]bool)
		for _, existing := range existingScripts {
			existingIDMap[existing.ID] = true
		}

		// 3. 입력 배열의 ID를 맵으로 저장
		inputIDMap := make(map[string]bool)
		for _, scriptItem := range scripts {
			inputIDMap[scriptItem.ID] = true
		}

		// 4. DB에는 있지만 입력 배열에는 없는 스크립트 삭제
		for _, existing := range existingScripts {
			if !inputIDMap[existing.ID] {
				_, err := tx.Model(&utils.Script{}).Where("id = ?", existing.ID).Delete()
				if err != nil {
					return err
				}
			}
		}

		// 5. 입력 배열의 각 스크립트에 대해 업데이트 또는 삽입
		for _, scriptItem := range scripts {
			newScript := &utils.Script{
				DateCreated:  scriptItem.DateCreated,
				DateModified: scriptItem.DateModified,
				ID:           scriptItem.ID,
				Name:         scriptItem.Name,
				Type:         scriptItem.Type,
			}

			if existingIDMap[scriptItem.ID] {
				// 기존에 존재하면 업데이트
				_, err := tx.Model(newScript).Where("id = ?", scriptItem.ID).Update()
				if err != nil {
					return err
				}
			} else {
				// 존재하지 않으면 삽입
				_, err := tx.Model(newScript).Insert()
				if err != nil {
					return err
				}
			}

			updatedScripts = append(updatedScripts, *newScript)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedScripts, nil
}
