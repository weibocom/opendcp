package models

type Cluster struct {
	Id   int    `json:"id" orm:"pk;auto"`
	Name string `json:"name" orm:"size(50)"`
	Desc string `json:"desc" orm:"size(255);null"`
	Biz  string `json:"biz"` //产品线
}

type Service struct {
	Id          int      `json:"id" orm:"pk;auto"`
	Name        string   `json:"name" orm:"size(50)"`
	Desc        string   `json:"desc" orm:"size(255);null"`
	ServiceType string   `json:"service_type"` //服务类型
	DockerImage string   `json:"docker_image"` //Docker镜像
	Cluster     *Cluster `json:"-" orm:"rel(fk);on_delete(cascade)"`
}

type Pool struct {
	Id      int      `json:"id" orm:"pk;auto"`
	Name    string   `json:"name" orm:"size(50)"`
	Desc    string   `json:"desc" orm:"size(255);null"`
	VmType  int      `json:"vm_type"` //VM类型
	SdId    int      `json:"sd_id"`   //服务发现ID
	Tasks   string   `json:"tasks"`   //对应任务(task_name arr)
	Service *Service `json:"-" orm:"rel(fk);on_delete(cascade)"`
}

type Node struct {
	Id     int    `json:"id" orm:"pk;auto"`
	Ip     string `json:"ip" orm:"null"`
	VmId   string `json:"vm_id" orm:"null"`
	Status int    `json:"status"`
	Pool   *Pool  `json:"-" orm:"rel(fk);on_delete(cascade)"`
	//Cluster*Cluster`json:"-" orm:"rel(fk);on_delete(cascade)"`
}
