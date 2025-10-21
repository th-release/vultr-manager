package scheduler

import (
	"log"
	"th-release/vultr-manager/api/application"
	"th-release/vultr-manager/api/firewall"
	"th-release/vultr-manager/api/instance"
	"th-release/vultr-manager/api/os"
	"th-release/vultr-manager/api/plan"
	"th-release/vultr-manager/api/region"
	"th-release/vultr-manager/api/script"
	"th-release/vultr-manager/utils"
	"time"

	"github.com/robfig/cron/v3"
)

type LoadVultr struct {
	Config utils.Config
}

func (j LoadVultr) Run() {
	config := utils.GetConfig()

	if !config.SyncDatabase {
		return
	}

	db := utils.NewDB(*config)
	defer db.Close()

	applicationQuery := map[string]string{
		"type":     string(application.TypeAll),
		"per_page": "500",
	}

	_, applicationRes, applicationErrResp, applicationErr := application.ApplicationList(applicationQuery, config.ApiKey)

	if applicationErr == nil && applicationErrResp == nil {
		MappingApplications(db, applicationRes.Applications)
	}

	normalQuery := map[string]string{
		"per_page": "999999",
	}

	firewallGroupQuery := map[string]string{
		"per_page": "500",
	}

	firewallRulesQuery := map[string]string{
		"per_page": "500",
	}

	_, firewallGroupRes, firewallGroupErrResp, firewallGroupErr := firewall.FirewallGroupList(firewallGroupQuery, config.ApiKey)

	if firewallGroupErr == nil && firewallGroupErrResp == nil {
		for _, v := range firewallGroupRes.Firewalls {
			MappingFirewallGroups(db, firewallGroupRes.Firewalls)

			_, firewallRuleRes, firewallRuleErrResp, firewallRuleErr := firewall.FirewallRulesList(v.ID, firewallRulesQuery, config.ApiKey)

			if firewallRuleErr == nil && firewallRuleErrResp == nil {
				MappingFirewallRules(db, firewallRuleRes.FirewallRules)
			}
		}
	}

	_, instanceRes, instanceErrResp, instanceErr := instance.InstanceList(normalQuery, config.ApiKey)

	if instanceErrResp == nil && instanceErr == nil {
		MappingInstances(db, instanceRes.Instances)
	}

	OsQuery := map[string]string{
		"per_page": "500",
	}

	_, osRes, osErrResp, osErr := os.OsList(OsQuery, config.ApiKey)

	if osErrResp == nil && osErr == nil {
		MappingOs(db, osRes.OsList)
	}

	_, planRes, planErrResp, planErr := plan.PlanList(nil, config.ApiKey)

	if planErrResp == nil && planErr == nil {
		MappingPlans(db, planRes.Plans)
	}

	RegionQuery := map[string]string{
		"per_page": "500",
	}

	_, regionRes, regionErrResp, regionErr := region.RegionList(RegionQuery, config.ApiKey)

	if regionErrResp == nil && regionErr == nil {
		MappingRegions(db, regionRes.Regions)
	}

	ScriptQuery := map[string]string{
		"per_page": "500",
	}

	_, scriptRes, scriptErrResp, scriptErr := script.ScriptList(ScriptQuery, config.ApiKey)

	if scriptErrResp == nil && scriptErr == nil {
		MappingScripts(db, scriptRes.StartScripts)
	}
}

func InitCron(config utils.Config) *cron.Cron {
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		log.Println("Time Zone Load Error:", err)
		return nil
	}

	c := cron.New(cron.WithSeconds(), cron.WithLocation(loc))

	_, err = c.AddJob("0 */1 * * * *", LoadVultr{
		Config: config,
	})

	if err != nil {
		log.Println("LoadVultr Job Add Error:", err)
		return nil
	}

	c.Start()

	return c
}
