/**
 *    Copyright (C) 2016 Weibo Inc.
 *
 *    This file is part of Opendcp.
 *
 *    Opendcp is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation; version 2 of the License.
 *
 *    Opendcp is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with Opendcp; if not, write to the Free Software
 *    Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA
 */


package service

import (
	"github.com/astaxie/beego/orm"
)

type InitService struct {
	BaseService
}


func (f *InitService) InsertBysql (sql string) (lastId int64,err error) {
	o := orm.NewOrm()
	result,err := o.Raw(sql).Exec()

	if err != nil {
		return -1,err
	}

	lastId,err = result.LastInsertId()

	return lastId,nil
}


/*func (f *InitService) InitAll(biz_id int)  {
	utils.InitConf()
	initConfig := utils.InitConfig

	for i := 0 ; i<len(*initConfig);i++{

		//fmt.Println((*InitConfig)[i].Records)

		tbName := (*initConfig)[i].Table


		records := (*initConfig)[i].Records



		fmt.Println("====================================")

		err := f.CreateTable(biz_id, tbName, records)

		fmt.Println(err)
	}

	fmt.Println("exec here")

}

func (f *InitService) CreateTable(biz_id int, tbName string, records []utils.Record) ( err error ) {
	switch tbName {
	case "remote_action":
		fmt.Println("remote_action")
		f.CreateRemoteAction(biz_id,records)
	case "remote_action_impl":
		fmt.Println("remote_action_impl")
	case "remote_action":
		fmt.Println("remote_action")
	case "remote_step":
		fmt.Println("remote_step")
	case "flow_impl":
		fmt.Println("flow_impl")
	case "service":
		fmt.Println("service")
	case "pool":
		fmt.Println("pool")

	}

	return

}

func (f *InitService) CreateRemoteAction(biz_id int, records []utils.Record)  {
	o := orm.NewOrm()

	result,err := o.Raw("").Exec()

	id,err := result.LastInsertId()

	fmt.Println(id);




	action := make([]models.RemoteAction,0)

	for _,record := range records{
		fmt.Println(record)
		for k,v := range record{
			fmt.Println(k,v)
		}
	}


	action := &models.RemoteAction{
		BizId:  biz_id,
		Name:   "start_docker",
		Desc:   string ,
		Params: string ,

	}

	data := models.RemoteAction{
		Name:   "start_docker",
		BizId:  biz_id,
		Desc:   req.Desc,
		Params: string(paramsStr),
	}
	err = service.Cluster.InsertBase(&data)

	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	count,err := o.Insert(action)

	o.InsertMulti()


}

func (f *InitService) CreateStep(biz_id int)  {
	o := orm.NewOrm()

	action := &models.ActionImpl{}
	err := o.QueryTable(action).Filter("name", name).One(action)
	if err != nil {
		return nil, err
	}

	return action, nil
}

func (f *InitService) CreateFlowImp(biz_id int)  {
	o := orm.NewOrm()

	node := &models.Node{}
	err := o.QueryTable(node).Filter("ip", ip).One(node)
	if err != nil {
		return nil, err
	}

	return node, nil
}
func (f *InitService) CreateService(biz_id int)  {
	o := orm.NewOrm()

	node := &models.Node{}
	err := o.QueryTable(node).Filter("ip", ip).One(node)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (f *InitService) CreatePool(biz_id int)  {
	o := orm.NewOrm()

	node := &models.Node{}
	err := o.QueryTable(node).Filter("ip", ip).One(node)
	if err != nil {
		return nil, err
	}

	return node, nil
}*/



