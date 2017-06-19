package account

import (
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/dao"
	"reflect"
	"unsafe"
	"encoding/base64"
	"time"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

const BASE64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"


func GetAccount(bizId int, provider string)  (*models.Account, error){
	account, err := dao.GetAccountByProvider(bizId, provider)
	account.KeySecret, err = Decode(account.KeySecret)
	if err != nil {
		return nil, errors.New("Decrypt keysecret err!")
	}
	if err != nil {
		return nil, err
	}
	return account, nil
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
	account.CreateTime = time.Now()
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
