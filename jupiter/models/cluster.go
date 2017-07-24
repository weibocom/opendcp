package models

import (
	"time"
)

type Cluster struct {
	Id   int64 `orm:"pk;auto"`
	Name string
	//OrganizationId     int64
	Provider           string
	LastestPartNum     int
	//PartOfInstances    []*InstancesResp `orm:"-"`
	Desc               string
	CreateTime         time.Time
	DeleteTime         time.Time `orm:"null"`
	Cpu                int
	Ram                int
	InstanceType       string
	ImageId            string
	PostScript         string
	KeyName            string
	Network            *Network      `orm:"rel(fk)"`
	Zone               *Zone         `orm:"rel(fk)"`
	Replication        []Replication `orm:"-"`
	SystemDiskCategory string
	DataDiskSize       int
	DataDiskNum        int
	DataDiskCategory   string
}

type Replication struct {
	Id      int `orm:"pk;auto"`
	PartNum int
	Cluster *Cluster `orm:"rel(fk)"`
	// The ID of the instance.
	InstanceId string `locationName:"instanceId" type:"string"`
	CreateTime time.Time
}

type Network struct {
	Id            int64 `orm:"pk;auto"`
	VpcId         string
	SubnetId      string
	SecurityGroup string
	InternetChargeType string
	InternetMaxBandwidthOut int
}

// Describes an Availability Zone.
type Zone struct {
	Id int64 `orm:"pk;auto"`
	// The name of the region.
	RegionName string
	// The name of the Availability Zone.
	ZoneName string
}

type Volume struct {
	DeviceName string
	Size       string
	Type       string
}

// 多字段唯一键
func (u *Zone) TableUnique() [][]string {
	return [][]string{
		[]string{"region_name", "zone_name"},
	}
}


type Detail struct {
	Id             int64      	`orm:"pk;auto"`
	InstanceNumber string 		`orm:"type(text);null"`
	RunningTime    time.Time 	`orm:"auto_now_add;type(datetime)"`
}

type InstanceDetail struct {
	InstanceNumber   map[string]int	`json:"number"`
	RunningTime		 string			`json:"time"`
}
