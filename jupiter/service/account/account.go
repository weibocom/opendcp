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
	"fmt"
	"strconv"
	"weibo.com/opendcp/jupiter/conf"
)

const BASE64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"


func GetAccount(bizId int, provider string)  (*models.Account, error){
	theAccount, err := dao.GetAccountByProvider(bizId, provider)
	if err != nil {
		return nil, err
	}
	//theAccount.KeySecret, err = Decode(theAccount.KeySecret)
	if err != nil {
		return nil, errors.New("Decrypt keysecret err!")
	}
	return theAccount, nil
}


func ListAccounts(bizId int) ([]models.Account, error) {
	accounts, err := dao.GetAllInAccount(bizId)
	/*for i, _ := range accounts {
		accounts[i].KeySecret, err = Decode(accounts[i].KeySecret)
		if err != nil {
			return nil,errors.New("Decrypt keysecret err!")
		}
	}*/
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

func UpdateAccountInfo(obj interface{}, column []string) (err error) {
	err = dao.UpdateAccountInfo(obj ,column)
	return  err
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

func ComputeCostNew(begin time.Time, end time.Time, instance models.Instance) (float64) {
	beego.Info(fmt.Sprintf("$$$$instance=%v ; begin=%v ,end=%v",instance.Id,begin,end))
	var rt1 time.Time = begin;
	var rt2 time.Time = end;
	var totalHour int = 0;
	if begin.Minute() !=0 || begin.Second() !=0  {
		totalHour ++
		m1 := begin.Minute()
		s1 := begin.Second()
		cm1 := 59-m1
		cs1 := 60-s1
		var totalS int = cm1*60 + cs1
		str := strconv.Itoa(totalS)+"s"
		d,_ := time.ParseDuration(str)
		rt1 =begin.Add(d)

	}
	if end.Minute() !=0 || end.Second() !=0  {
		totalHour++
		m2 := end.Minute()
		s2 := end.Second()
		var totalS int = m2*60 + s2
		str := "-"+strconv.Itoa(totalS)+"s"
		d,_ := time.ParseDuration(str)
		rt2_tmp :=end.Add(d)

		strT2 := rt2_tmp.Format("2006-01-02 15:04:05")

		//rt,_ := time.Parse("2006-01-02 15:04:05",strT2)
		rt,_ := time.ParseInLocation("2006-01-02 15:04:05",strT2, time.Local)

		rt2 = rt

	}
	beego.Info(fmt.Sprintf("$$$$ rt1=%v ,rt2=%v",rt1,rt2))
	durationH := rt2.Sub(rt1).Hours()
	totalHour = totalHour+int(durationH)
	beego.Info(fmt.Sprintf("$$$$ durationH=%v ,totalHour=%v",durationH,totalHour))
	cpuNum := instance.Cpu
	ramNum := instance.Ram
	if cpuNum == 0 {
		cpuNum = 1
	}

	if ramNum == 0 {
		ramNum  = 1
	}
	return float64(totalHour*(cpuNum+ramNum)/2)
}

/**
生成额度信息
 */
func GenerateMultiCost() error{
	beego.Info(fmt.Sprintf("###########begin compute cost for all biz"))
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
	beego.Info(fmt.Sprintf("###########finish compute cost for all biz"))
	return nil

}

func GenerateOneCost(biz_id int) error {
	beego.Info(fmt.Sprintf("#######begin compute cost for biz %d",biz_id))
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
	beego.Info(fmt.Sprintf("$$$$exist account for biz %d,content is %v",biz_id,existAccount))
	//2、获取此业务方的所有实例
	instances,err := dao.GetAllInstance(biz_id)
	if err != nil {
		beego.Error(err)
		return err
	}

	//3、计算额度并且更新库表
	now := time.Now()
	//var duration time.Duration

	spendMap := make(map[string]float64)

	for _,instance := range instances {
		//3.1去除存在云厂商账户的
		provider := instance.Provider
		if _, ok := existAccount[provider]; ok {
			continue
		}

		ctime := instance.CreateTime
		var rtime time.Time
		/*if instance.Status == models.Deleted {
			rtime := instance.ReturnTime
			duration = rtime.Sub(ctime)
		}else{
			duration = now.Sub(ctime)

		}
		spendTime := duration.Minutes()
		cost := ComputeCost(spendTime,instance)*/
		if instance.Status == models.Deleted {
			rtime = instance.ReturnTime
		}else {
			rtime = now
		}
		cost := ComputeCostNew(ctime,rtime,instance)

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


	}
	beego.Info(fmt.Sprintf("$$$$compute result for biz %d,content is %v",biz_id,spendMap))
	beego.Info(fmt.Sprintf("#######finish compute cost for biz %d",biz_id))
	return nil

}

func GetLatestCost(bizId int, provider string) (map[string]float64, error) {
	err := GenerateOneCost(bizId)
	if err != nil{
		return nil, err
	}
	return instance.GetCost(bizId, provider)
}

func GetTotalCosts() (map[string]float64, error) {
	instances, err := dao.GetAllBidAndProviderInInstance()
	if err != nil {
		return nil, err
	}

	totalCosts := make(map[string]float64)
	for _, ins := range instances {
		costs, err := instance.GetCost(ins.BizId, ins.Provider)
		if err != nil {
			return nil, err
		}
		for k, v := range costs{
			totalCosts[k] += v
		}
	}
	return totalCosts, nil
}

func CheckCredit() error {
	err := GenerateMultiCost()
	if err != nil {
		return err
	}

	instances, err := dao.GetAllBidAndProviderInInstance()
	if err != nil {
		return err
	}

	for _, ins := range instances {
		//查账户的额度，额度不足时删除机器
		costs, err := instance.GetCost(ins.BizId, ins.Provider)
		if err != nil {
			return  err
		}
		if instance.GreaterOrEqual(costs["spent"], costs["credit"] + 1) {
			beego.Info(fmt.Sprintf("$$$$delete instance for biz %d from %s",ins.BizId, ins.Provider))
			instances, err := dao.GetTestingInstances(ins.BizId, ins.Provider)
			if err != nil {
				return  err
			}
			go instance.DeleteTestingInstances(instances, ins.BizId)
		}
	}
	return nil
}

func SendEmail(data models.EmailData) error {
	emailName := conf.Config.EmailName
	emailPassword:= conf.Config.EmailPassword
	emailServer := conf.Config.EmailServer
	sender := emailService.NewSender(emailName, emailPassword, emailServer)
	subject := "用户注册提醒"
	body :=  `
	<html>
		<body>
			<h3>尊敬的用户` + data.UserName + `:</h3>
			<p>` + data.Content + `</p>
		</body>
	</html>
	`
	receiver := data.Receiver
	beego.Warn("Email data",data)
	email := emailService.NewEmailDate(sender.EmailName,"", receiver, subject, body, "html")
	return emailService.SendMail(sender, email)
}


