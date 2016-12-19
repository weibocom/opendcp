
# Octans


### Summary

基于Ansible的配置下发通道

### Language & Framework

* Python 2.7
* Flask 0.11
* Ansible 2.1


### Install

* ```pip install -r requirements.txt```
* ```python run.py```

### Config

* ```config.tpl.yml => config.yml```

### Database

* build tables: ```./build_db.py```


### API

* #### run task:
	* URL: ```/api/run```
	* method: ```POST```
	* encode: ``json``:

	```
	{
		"nodes":["xxx.xxx.xxx.xxx"],
		"tasks":[
			"(task json dumps)",
			"...."
		],
		"name":"...",
		"user":"root",
		"parmas":{
			"key":"value"
		}
		"fork_num":5
	}
	```
	* success:

	```
	{
		"code":0,
		"content":{}
	}
	```

	* error:

	```
	{
		"code":-1,
		"message":"param error"
	}
	```



* #### stop task:
	* URL: ```/api/stop```
	* method: ```POST```
	* encode: ``json``:

	(task_name/task_id二选一)
	```
	{
		"task_id":1,
		"task_name":"..."
	}
	```

	* success:

	```
	{
		"code":0,
		"content":{}
	}
	```

	* failed:

	```
	{
		"code":-1,
		"message":"task not found..."
	}
	```
	
* #### check task:
	* URL: ```/api/check```
	* method: ```POST```
	* encode: ``json``:

	(task_name || task_id)
	```
	{
		"task_id":1,
		"task_name":"..."
	}
	```

	* success:

	```
	{
  		"code": 0,
  		"content": {
    		"create_time": "Wed, 17 Aug 2016 18:38:17 GMT",
    		"err": null,
    		"id": 7,
    		"log": "[{...},{...}]",
    		"nodes": "[\"123.207.136.186\"]",
    		"status": 4,
    		"update_time": "Wed, 17 Aug 2016 18:38:22 GMT"
  		}
	}
	```

	* failed:

	```
	{
		"code":-1,
		"message":"task not found..."
	}
	```


### Building & Deploy

* ```sudo docker run -d --name halo-channel -p 8000:8000 xxx/halo-channel```


### Contributing

...
