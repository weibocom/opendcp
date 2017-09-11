/**
 *    Copyright (C) 2016 Weibo Inc.
 *
 *    This file is part of Opendcp.
 *
 *    Opendcp is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation; version 2 of the License.
 *
 *    Opendcp is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with Opendcp; if not, write to the Free Software
 *    Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA
 */

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"time"

	"bytes"
	"errors"
	log "github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"weibo.com/opendcp/jupiter/logstore"
)

var (
	ORION_ADDR = beego.AppConfig.String("orion_mgr_addr")

	ADD_NODE_URL = "http://%s/pool/%d/add_phy_dev"
)

var (
	client = &http.Client{
		Timeout: 5 * 60 * time.Second,
	}
)

type AppendPhyDevContent struct {
	Success     int      `json:"success"`
	Failed      int      `json:"failed"`
	SuccessList []string `json:"successes"`
	ErrorList   []string `json:"errors"`
	ErrorMsg    []string `json:"error_msg"`
}

type sdCmdResp struct {
	Code    int
	Message string              `json:"msg"`
	Content AppendPhyDevContent `json:"data"`
}

func AddPhyDevToPool(ips []string, pool_id int, label string, correlationId string, finalInstanceId []string) (*sdCmdResp, error) {

	return do(ADD_NODE_URL, ips, pool_id, label, correlationId, finalInstanceId)
}

func do(action string, ips []string, pool_id int, label string, correlationId string, finalInstanceId []string) (*sdCmdResp, error) {

	for _, id := range finalInstanceId {
		logstore.Info(correlationId, id, "begin to add to pool")
	}

	data := make(map[string]interface{})
	data["nodes"] = ips
	data["label"] = label
	data["instanceId"] =finalInstanceId

	header := make(map[string]interface{})

	resp := &sdCmdResp{}
	url := fmt.Sprintf(action, ORION_ADDR, pool_id)
	_, err := callAPI("POST", url, &data, &header, resp)
	if err != nil {
		return nil, err
	}
	beego.Info("hhhhhhhhhhhh")
	beego.Info(resp.Content)
	//beego.Info(resp.Content["errors"])
	//beego.Info(resp.Content.(AppendPhyDevContent))
	if resp.Code != 0 {
		return resp, errors.New("add failed")
	} else {
		return resp, nil
	}
}

func callAPI(method string, url string,
data *map[string]interface{}, header *map[string]interface{}, obj interface{}) (string, error) {

	msg, err := Do(method, url, data, header)
	if err != nil {
		beego.Error("Fail to ", method, url, ": ", err)
		return "", err
	}

	err = json.Unmarshal([]byte(msg), obj)
	if err != nil {
		beego.Error("Fail to unmarshal", msg, "err:", err)
		beego.Error("Bad resp:", msg)
		return "", errors.New("Bad resp: " + msg)
	}
	return msg, nil
}

func Do(method string, url string, data *map[string]interface{}, header *map[string]interface{}) (string, error) {

	jsonBytes, err := json.Marshal(&data)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}
	log.Info(method, url, " - ", string(jsonBytes))
	//logstore.Debug(method, url, " - ", string(jsonBytes))
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	if header != nil {
		for k, v := range *header {
			req.Header.Set(k, fmt.Sprintf("%s", v))
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return handleResp(resp)
}

func handleResp(resp *http.Response) (string, error) {
	//code := resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
