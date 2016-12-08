package service

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/go-errors/errors"
	"fmt"
)

type BaseService struct {
}

var (
	Cluster = &ClusterService{}
	Flow    = &FlowService{}
	Remote  = &RemoteStepService{}
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
		return errors.New("fail to delete: " + fmt.Sprintf("%v", obj))
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
		qr = o.QueryTable(obj).OrderBy(sortstr[0])
	case 2:
		qr = o.QueryTable(obj).OrderBy(sortstr[0], sortstr[1])
	case 3:
		qr = o.QueryTable(obj).OrderBy(sortstr[0], sortstr[1], sortstr[2])
	default:
		qr = o.QueryTable(obj)
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
