package utils

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func NewDB(config Config) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     config.DatabaseAddr,
		User:     config.DatabaseUser,
		Password: config.DatabasePassword,
		Database: config.DatabaseSchema,
		PoolSize: 20,
	})

	return db
}

func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*Instance)(nil),
		(*Application)(nil),
		(*FirewallGroup)(nil),
		(*FirewallRule)(nil),
		(*OS)(nil),
		(*Plan)(nil),
		(*Region)(nil),
		(*Script)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:          false,
			IfNotExists:   true,
			FKConstraints: true, // 외래 키 제약 조건 활성화
		})
		if err != nil {
			return fmt.Errorf("failed to create table for model %T: %w", model, err)
		}
	}
	return nil
}

type Instance struct {
	tableName        struct{} `pg:"instance"` // 테이블 이름 지정
	Id               string   `pg:"id,pk,type:varchar(36),notnull"`
	AllowedBandwidth int      `pg:"allowed_bandwidth"`
	AppID            int      `pg:"app_id"`
	DateCreated      string   `pg:"date_created"`
	Disk             int      `pg:"disk"`
	Features         []string `pg:"features"`
	FirewallGroupID  string   `pg:"firewall_group_id"`
	GatewayV4        string   `pg:"gateway_v4"`
	Hostname         string   `pg:"hostname"`
	ImageID          string   `pg:"image_id"`
	InternalIP       string   `pg:"internal_ip"`
	KVM              string   `pg:"kvm"`
	Label            string   `pg:"label"`
	MainIP           string   `pg:"main_ip"`
	NetmaskV4        string   `pg:"netmask_v4"`
	OS               string   `pg:"os"`
	OsID             int      `pg:"os_id"`
	PendingCharges   float64  `pg:"pending_charges"`
	Plan             string   `pg:"plan"`
	PowerStatus      string   `pg:"power_status"`
	RAM              int      `pg:"ram"`
	Region           string   `pg:"region"`
	ServerStatus     string   `pg:"server_status"`
	Status           string   `pg:"status"`
	Tag              string   `pg:"tag"`
	Tags             []string `pg:"tags"`
	UserScheme       string   `pg:"user_scheme"`
	V6MainIP         string   `pg:"v6_main_ip"`
	V6Network        string   `pg:"v6_network"`
	V6NetworkSize    int      `pg:"v6_network_size"`
	VcpuCount        int      `pg:"vcpu_count"`
	Vpcs             []string `pg:"vpcs"`
}

type Application struct {
	tableName  struct{} `pg:"application"`
	Id         int      `pg:"id,pk"`
	Name       string   `pg:"name"`
	ShortName  string   `pg:"short_name"`
	DeployName string   `pg:"deploy_name"`
	Type       string   `pg:"type"`
	Vendor     string   `pg:"vendor"`
	ImageId    string   `pg:"image_id"`
}

type FirewallGroup struct {
	tableName     struct{} `pg:"firewall_group"`
	ID            string   `pg:"id,pk"`
	Description   string   `pg:"description"`
	DateCreated   string   `pg:"date_created"`
	DateModified  string   `pg:"date_modified"`
	InstanceCount int      `pg:"instance_count"`
	RuleCount     int      `pg:"rule_count"`
	MaxRuleCount  int      `pg:"max_rule_count"`
}

type FirewallRule struct {
	tableName  struct{} `pg:"firewall_rule"`
	Id         int64    `pg:"id,pk"`
	Type       string   `pg:"type"`
	IpType     string   `pg:"ip_type"`
	Action     string   `pg:"action"`
	Protocol   string   `pg:"protocol"`
	Port       string   `pg:"port"`
	Subnet     string   `pg:"subnet"`
	SubnetSize int64    `pg:"subnet_size"`
	Source     string   `pg:"source"`
	Notes      string   `pg:"notes"`
}

type OS struct {
	tableName struct{} `pg:"os"`
	Id        int      `pg:"id,pk"`
	Name      string   `pg:"name"`
	Arch      string   `pg:"arch"`
	Family    string   `pg:"family"`
}

type Plan struct {
	tableName   struct{} `pg:"plan"`
	ID          string   `pg:"id,pk"`
	VcpuCount   int      `pg:"vcpu_count"`
	RAM         int      `pg:"ram"`
	Disk        int      `pg:"disk"`
	DiskType    string   `pg:"disk_type"`
	DiskCount   int      `pg:"disk_count"`
	Bandwidth   int      `pg:"bandwidth"`
	MonthlyCost float64  `pg:"monthly_cost"`
	HourlyCost  float64  `pg:"hourly_cost"`
	InvoiceType string   `pg:"invoice_type"`
	Type        string   `pg:"type"`
	Locations   []string `pg:"locations"`
	CpuVendor   string   `pg:"cpu_vendor"`
}

type Region struct {
	tableName struct{} `pg:"region"`
	ID        string   `pg:"id,pk"`
	City      string   `pg:"city"`
	Country   string   `pg:"country"`
	Continent string   `pg:"continent"`
	Options   []string `pg:"options"`
}

type Script struct {
	tableName    struct{} `pg:"script"`
	DateCreated  string   `pg:"date_created"`
	DateModified string   `pg:"date_modified"`
	ID           string   `pg:"id,pk"`
	Name         string   `pg:"name"`
	Type         string   `pg:"type"`
}
