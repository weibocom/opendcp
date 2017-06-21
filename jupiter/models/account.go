package models

type Account struct {
	Id   		int64 `orm:"pk;auto"`
	BizId 		int
	Provider        string
	KeyId     	string
	KeySecret       string
	Spent		int64
	Credit		int64
}

