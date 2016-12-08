package models

type RetryOption struct {
	RetryTimes  int  `json:"retry_times"`
	IgnoreError bool `json:"ignore_error"`
}

type ParamValues map[string]interface{}

type StepOption struct {
	Name   string       `json:"name"`
	Values ParamValues  `json:"param_values"`
	Retry  *RetryOption `json:"retry"`
}

//任务流定义
type FlowImpl struct {
	Id    int    `json:"id" orm:"pk;auto"`
	Name  string `json:"name" orm:"size(50);unique"`
	Desc  string `json:"name" orm:"size(255)"`
	Steps string `json:"steps" orm:"type(text)"` //action_name list
}

//任务流对应单步定义
type ActionImpl struct {
	Id     int                    `json:"id" orm:"pk;auto"`
	Name   string                 `json:"name"`
	Desc   string                 `json:"desc"`
	Type   string                 `json:"type"`
	Params map[string]interface{} `json:"params"`
}
