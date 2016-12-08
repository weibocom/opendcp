package service

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"weibo.com/opendcp/orion/models"
)

type ClusterService struct {
	BaseService
}

func (c *ClusterService) AppendIpList(ips []string, pool *models.Pool) []int {
	o := orm.NewOrm()

	respDatas := make([]int, 0, len(ips))

	for _, ip := range ips {
		data := models.Node{
			Ip:   ip,
			Pool: pool,
		}
		_, err := o.Insert(&data)
		if err != nil {
			beego.Error("insert node failed, error: ", err)
			return nil
		}
		respDatas = append(respDatas, data.Id)
	}

	return respDatas
}

// SearchPoolByIP returs the pool id of the ip given, if it exists
func (c* ClusterService) SearchPoolByIP(ips []string) map[string]int {

	result := make(map[string]int)
	for _, ip := range ips {
		node := &models.Node{Ip: ip}
		err := c.GetBy(node, "ip")
		id := -1
		if err == nil {
			id = node.Pool.Id
		}
		result[ip] = id
	}

	return result
}
