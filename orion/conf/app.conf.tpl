appname = orion
httpport = 8080
runmode = dev

database_url = root:@tcp(localhost:3306)/orion?charset=utf8

vm_mgr_addr = xx.xx.xx.xx:8888
vm_create_url = /v1/cluster/%d/expand/%d
vm_return_url = /v1/instance/%s
vm_check_url =  /v1/instance/status/%s

sd_mgr_addr = 
sd_register_url = /v1/adaptor/auto_alteration/add
sd_unregister_url = /v1/adaptor/auto_alteration/remove
sd_check_url = /v1/adaptor/auto_alteration/check_state
sd_add_nginx_node_url = /v1/adaptor/auto_alteration/addNode_post/
sd_delete_nginx_node_url = /v1/adaptor/auto_alteration/deleteNode_delete/
sd_appkey = 6741bc42-9e21-4763-977c-ac3a1fc0bdd8

remote_mgr_addr = xx.xx.xx.xx:8000
remote_command_url = /api/run
remote_check_url =  /api/check
