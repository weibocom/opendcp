package models

type RemoteStep struct {
	Id      int    `json:"id" orm:"pk;auto"`
	Name    string `json:"name" orm:"size(50);unique"`
	Desc    string `json:"desc" orm:"size(255);null"`
	Actions string `json:"actions" orm:"type(text)"`
}

type RemoteAction struct {
	Id     int    `json:"id" orm:"pk;auto"`
	Name   string `json:"name" orm:"size(50);unique"`
	Desc   string `json:"desc" orm:"size(255);null"`
	Params string `json:"params" orm:"type(text)"`
}

type RemoteActionImpl struct {
	Id       int    `json:"id" orm:"pk;auto"`
	Type     string `json:"type" orm:"size(50)"`
	Template string `json:"template" orm:"type(text)"`
	ActionId int    `json:"action_id"`
}
