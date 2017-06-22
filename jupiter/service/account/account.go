package account

import (
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/dao"
	"reflect"
	"unsafe"
	"encoding/base64"
	"time"
	"errors"
	"github.com/astaxie/beego"
	"weibo.com/opendcp/jupiter/service/instance"
)

const BASE64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"


func GetAccount(bizId int, provider string)  (*models.Account, error){
	theAccount, err := dao.GetAccountByProvider(bizId, provider)
	if err != nil {
		return nil, err
	}
	theAccount.KeySecret, err = Decode(theAccount.KeySecret)
	if err != nil {
		return nil, errors.New("Decrypt keysecret err!")
	}
	return theAccount, nil
}


func ListAccounts(bizId int) ([]models.Account, error) {
	accounts, err := dao.GetAllInAccount(bizId)
	for i, _ := range accounts {
		accounts[i].KeySecret, err = Decode(accounts[i].KeySecret)
		if err != nil {
			return nil,errors.New("Decrypt keysecret err!")
		}
	}
	if err != nil {
		return nil, err
	}
	return accounts, err
}

func CreateAccount(account *models.Account) (int64, error) {
	account.KeySecret = Encode(account.KeySecret)
	id, err := dao.InsertAccount(account)
	if err != nil {
		return 0, err
	}
	return id, err
}

func DeleteAccount(bizId int, provider string) (bool, error) {
	result, err := dao.DeleteAccountByProvider(bizId, provider)
	return result, err
}

func Encode(data string) string {
	content := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data))))
	coder := base64.NewEncoding(BASE64Table)
	return coder.EncodeToString(content)
}

func Decode(data string) (string, error)  {
	coder := base64.NewEncoding(BASE64Table)
	result, err := coder.DecodeString(data)
	if err != nil {
		return "", err
	}
	return *(*string)(unsafe.Pointer(&result)), nil
}



/**
计算额度算法
 */
func ComputeCost (time float64, instance models.Instance) ( float64 ) {
	cpu := float64(instance.Cpu)
	mem := float64(instance.Ram)

	cpuWeight:= float64(2.0/3.0)
	memWeight:= float64(1.0/3.0)


	return (cpu*cpuWeight+mem*memWeight)*(time/60.0)

}

/**
生成额度信息
 */
func GenerateMultiCost() error{
	instances,err := dao.GetAllBIdInInstance()
	if err != nil {
		beego.Error(err)
		return err
	}

	bizInInstance := make([]int,len(instances))
	for i,instance := range instances {
		bizInInstance[i]= instance.BizId
	}

	for _,biz_id := range bizInInstance {
		err := GenerateOneCost(biz_id)
		if err != nil {
			beego.Error(err)
			return err
		}
	}
	return nil

}

func GenerateOneCost(biz_id int) error {
	//1、获取此业务方的账户信息
	accounts,err := dao.GetAllInAccount(biz_id)
	if err != nil {
		beego.Error(err)
		return err
	}

	//biz_id provider
	existAccount := make(map[string]interface{})
	for _,account := range accounts {
		if account.KeyId != "" || account.KeySecret != ""{
			existAccount[account.Provider] = account
		}
	}
	//2、获取此业务方的所有实例
	instances,err := dao.GetAllInstance(biz_id)
	if err != nil {
		beego.Error(err)
		return err
	}

	//3、计算额度并且更新库表
	now := time.Now()
	var duration time.Duration

	spendMap := make(map[string]float64)

	for _,instance := range instances {
		//3.1去除存在云厂商账户的
		provider := instance.Provider
		if _, ok := existAccount[provider]; ok {
			continue
		}

		ctime := instance.CreateTime
		if instance.Status == models.Deleted {
			rtime := instance.ReturnTime
			duration = rtime.Sub(ctime)
		}else{
			duration = now.Sub(ctime)

		}
		spendTime := duration.Minutes()
		cost := ComputeCost(spendTime,instance)
		if v, ok := spendMap[provider]; ok {
			spendMap[provider] = v+cost
		}else{
			spendMap[provider] = cost
		}

	}

	//更新account数据库表
	for k,v := range spendMap {
		err := dao.UpdateAccount(biz_id,k,v)
		if err != nil {
			beego.Error(err)
			return err
		}
		//重新检查账户的额度，额度不足时删除机器
		costs, err := instance.GetCost(biz_id, k)
		if err != nil {
			return  err
		}
		if instance.GreaterOrEqual(costs["spent"], costs["credit"]) {
			instances, err := dao.GetTestingInstances(biz_id, k)
			if err != nil {
				return  err
			}
			go instance.DeleteInstances(instances, biz_id)
		}

	}
	beego.Info(spendMap)
	return nil

}

func UpdateAccountInfo(obj interface{}, column []string) (err error) {
	err = dao.UpdateAccountInfo(obj ,column)
	return  err
}


