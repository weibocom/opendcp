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
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BaseService struct {
}

var (
	Cluster = &ClusterService{}
	Flow    = &FlowService{}
	Remote  = &RemoteStepService{}
	Logs    = &LogsService{}
	Task    = &TaskService{}
)

func (b *BaseService) InsertBase(obj interface{}) error {
	o := orm.NewOrm()

	_, err := o.Insert(obj)

	return err
}

func (b *BaseService) UpdateBase(obj interface{}) error {
	o := orm.NewOrm()

	_, err := o.Update(obj)

	return err
}

func (b *BaseService) DeleteBase(obj interface{}) error {
	o := orm.NewOrm()

	n, err := o.Delete(obj)
	if err != nil {
		beego.Error("Error when deleting", obj, ", err:", err)
	}

	if n == 0 {
		return fmt.Errorf("fail to delete: %v", obj)
	}

	return nil
}

func (b *BaseService) GetBase(obj interface{}) error {
	return b.GetBy(obj, "Id")
}

func (b *BaseService) GetBy(obj interface{}, field string) error {
	o := orm.NewOrm()

	err := o.Read(obj, field)
	if err != nil {
		return err
	}

	return nil
}

func (b *BaseService) GetByMultiFieldValue(obj interface{}, conditions map[string]interface{}) error {
	o := orm.NewOrm()

	qs := o.QueryTable(obj)

	for k, v := range conditions {
		qs = qs.Filter(k, v)
	}

	err := qs.One(obj)
	if err != nil {
		return err
	}

	return nil
}

func (b *BaseService) DeleteByMultiFieldValue(obj interface{}, conditions map[string]interface{}) error {
	o := orm.NewOrm()

	qs := o.QueryTable(obj)

	for k, v := range conditions {
		qs = qs.Filter(k, v)
	}

	_, err := qs.Delete()
	if err != nil {
		return err
	}

	return nil
}

/*
*	load data by value list of specified field
 */
func (b *BaseService) GetByStringValues(obj interface{}, list interface{}, field string, values []string) error {

	o := orm.NewOrm()

	expression := field + "__" + "in"
	_, err := o.QueryTable(obj).Filter(expression, values).All(list)
	if err != nil {
		return err
	}

	return nil
}

/*
 *	Get by objects by ids
 */
func (b *BaseService) GetByIds(obj interface{}, list interface{}, ids []int) error {
	return b.GetByMultiIds(obj, list, "Id", ids)
}

/*
 *	Get by objects by field & ids
 */
func (b *BaseService) GetByMultiIds(obj interface{}, list interface{},
	field string, ids []int) error {

	o := orm.NewOrm()

	expression := field + "__in"
	c, err := o.QueryTable(obj).Filter(expression, ids).All(list)
	if err != nil {
		return err
	}

	beego.Debug("GetByMultiIds", expression, ids, " num =", c)

	return nil
}

func (b *BaseService) ListByPage(page, pageSize int, obj interface{}, list interface{}) (int, error) {
	o := orm.NewOrm()

	qr := o.QueryTable(obj)

	count, err := qr.Count()
	if err != nil {
		return 0, err
	}
	_, err = qr.Limit(pageSize, (page-1)*pageSize).All(list)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

/**
*   query table by page with sort
*	page: 		page index to query
*   pageSize: 	page size per page
*   obj:    	object to query
*   list: 		query result
*   sortstr: 	sort field names, multiple fields sorting is available. default is asc and field name with previous '-' means
 */
func (b *BaseService) ListByPageWithSort(page, pageSize int, obj interface{}, list interface{}, sortstr ...string) (int, error) {
	o := orm.NewOrm()
	var qr orm.QuerySeter
	switch len(sortstr) {
	case 1:
		qr = o.QueryTable(obj).OrderBy(sortstr[0]).RelatedSel()
	case 2:
		qr = o.QueryTable(obj).OrderBy(sortstr[0], sortstr[1]).RelatedSel()
	case 3:
		qr = o.QueryTable(obj).OrderBy(sortstr[0], sortstr[1], sortstr[2]).RelatedSel()
	default:
		qr = o.QueryTable(obj).RelatedSel()
	}
	count, err := qr.Count()
	if err != nil {
		return 0, err
	}
	_, err = qr.Limit(pageSize, (page-1)*pageSize).All(list)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (b *BaseService) ListByPageWithFilter(page int, pageSize int, obj interface{}, list interface{},
	filterkey string, filtervalue interface{}) (int, error) {

	o := orm.NewOrm()

	qr := o.QueryTable(obj).Filter(filterkey, filtervalue)

	count, err := qr.Count()
	if err != nil {
		return 0, err
	}
	_, err = qr.Limit(pageSize, (page-1)*pageSize).All(list)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (b *BaseService) GetCount(obj interface{}, filterkey string, filtervalue interface{}) (int, error) {

	o := orm.NewOrm()

	qr := o.QueryTable(obj).Filter(filterkey, filtervalue)

	count, err := qr.Count()

	return int(count), err
}
