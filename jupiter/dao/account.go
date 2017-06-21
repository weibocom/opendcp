package dao

import (
	"weibo.com/opendcp/jupiter/models"
	"errors"
)


func GetAccount (biz_id int, provider string) (*models.Account, error) {
	o := GetOrmer()
	var account models.Account
	err := o.QueryTable(ACCOUNT_TABLE).Filter("biz_id",biz_id).Filter("provider",provider).One(&account)
	if err != nil {
		return nil, err
	}
	return &account, nil

}

func GetAllInAccount (biz_id int) ([]models.Account,error) {
	o := GetOrmer()
	var accounts []models.Account
	_,err := o.QueryTable(ACCOUNT_TABLE).Filter("biz_id",biz_id).All(&accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func GetAccountById(biz_id int, id int64) (*models.Account, error) {
	o := GetOrmer()
	var account models.Account
	err := o.QueryTable(ACCOUNT_TABLE).Filter("biz_id",biz_id).Filter("id",id).One(&account)
	if err != nil {
		return nil, err
	}
	return &account, nil

}

func GetAccountByProvider(biz_id int, provider string) (*models.Account, error) {
	o := GetOrmer()
	var account models.Account
	err := o.QueryTable(ACCOUNT_TABLE).Filter("biz_id",biz_id).Filter("provider",provider).One(&account)
	if err != nil {
		return nil, err
	}
	return &account, nil

}

func UpdateAccount(biz_id int, provider string, spent int64) error {
	o := GetOrmer()
	account,err := GetAccountByProvider(biz_id,provider)
	if err != nil {
		return err
	}

	account.Spent = spent
	_, err = o.Update(account)
	if err != nil {
		return err
	}
	return nil
}

func InsertAccount(account *models.Account) (int64, error) {
	o := GetOrmer()
	id, err := o.Insert(account)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func DeleteAccountByProvider(bizId int, provider string) (bool, error) {
	o := GetOrmer()
	num, err := o.QueryTable(ACCOUNT_TABLE).Filter("biz_id", bizId).Filter("provider", provider).Delete()
	if err != nil {
		return false, err
	}

	if num == 0 {
		return false, errors.New("Account doesn't exists!")
	}
	return true, nil
}
