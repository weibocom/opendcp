package sched

import (
	"os"
	"testing"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"weibo.com/opendcp/orion/models"
)

var (
	pools = []*models.Pool{
		{
			Name:    "test_pool_1",
			Desc:    "test_pool_1",
			VmType:  1,
			SdId:    1,
			Tasks:   `{"deploy":3,"expand":1,"shrink":2}`,
			Service: &models.Service{Id: 1},
		},
		{
			Name:    "test_pool_2",
			Desc:    "test_pool_2",
			VmType:  2,
			SdId:    2,
			Tasks:   `{"deploy":6,"expand":4,"shrink":5}`,
			Service: &models.Service{Id: 1},
		},
	}
)

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	db := os.Getenv("TEST_DATABASE")
	if db == "" {
		db = "root:@tcp(localhost:3306)/orion?charset=utf8"
	}
	orm.RegisterDataBase("default", "mysql", db, 10)
	//register model
	orm.RegisterModel(
		&models.Cluster{},
		&models.Service{},
		&models.Pool{},
		&models.Node{},
		&models.CronItem{},
		&models.DependItem{},
		&models.ExecTask{},
	)

	o := orm.NewOrm()
	if _, err := o.QueryTable(&models.Pool{}).Filter("id__gt", 0).Delete(); err != nil {
		panic(err)
	}

	for _, p := range pools {
		if _, err := o.Insert(p); err != nil {
			panic(err)
		}
	}
}

func TestScheduler(t *testing.T) {
	var (
		assert = assert.New(t)
		o      = orm.NewOrm()
		t0     = []*models.ExecTask{
			{
				Pool: pools[0],
				CronItems: []*models.CronItem{
					{
						InstanceNum: 10,
						Time:        "12:00",
					},
					{
						InstanceNum: 20,
						Time:        "14:00",
					},
					{
						InstanceNum: 0,
						Time:        "16:00",
					},
				},
				Type:     "dep",
				ExecType: models.ExecTypeMock,
			},
			{
				Pool: pools[1],
				CronItems: []*models.CronItem{
					{
						InstanceNum: 15,
						Time:        "12:00",
					},
					{
						InstanceNum: 25,
						Time:        "14:00",
					},
					{
						InstanceNum: 5,
						Time:        "16:00",
					},
				},
				Type:     "dep",
				ExecType: models.ExecTypeMock,
			},
		}
	)

	sched, err := NewScheduler()
	if !assert.NoError(err) {
		return
	}

	for i, t := range t0 {
		assert.NoError(sched.Create(t), "%d", i)
	}

	for i, t := range t0 {
		cfg := models.ExecTask{Id: t.Id}
		err := o.QueryTable(&cfg).RelatedSel().Filter("id", cfg.Id).One(&cfg)
		assert.NoError(err, "%d", i)
		_, err = o.LoadRelated(&cfg, "CronItems")
		assert.NoError(err, "%d", i)
		_, err = o.LoadRelated(&cfg, "DependItems")
		assert.NoError(err, "%d", i)
		assert.Equal(len(t.CronItems), len(cfg.CronItems), "%d", i)
	}
}
