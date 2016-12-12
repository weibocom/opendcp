/*
 *  Copyright 2009-2016 Weibo, Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */
package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/astaxie/beego"
	"fmt"
)

var (
	Http   = &httpUtil{}
	client = &http.Client{
		Timeout: 2 * 60 * time.Second,
	}
)

type httpUtil struct {
}

func (h *httpUtil) Post(url string, data *map[string]interface{}, header *map[string]interface{}) (string, error) {
	return h.Do("POST", url, data, header)
}

func (h *httpUtil) Get(url string, header *map[string]interface{}) (string, error) {
	return h.Do("GET", url, nil, header)

}

func (h *httpUtil) Delete(url string, data *map[string]interface{}, header *map[string]interface{}) (string, error) {
	return h.Do("DELETE", url, data, header)
}

func (h *httpUtil) Do(method string, url string, data *map[string]interface{}, header *map[string]interface{}) (string, error) {
	//log.Debug(method, url, " - ", data)
	jsonBytes, err := json.Marshal(&data)
	if err != nil {
		return "", err
	}
	headerBytes, err := json.Marshal(&header)
	if err != nil {
		return "", err
	}
	log.Debug(method, url, " - ", string(jsonBytes))
	log.Debug("Header:", string(headerBytes))

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
	code := resp.StatusCode
	log.Debug("response Status:", resp.Status)
	log.Debug("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Debug("response Body:", string(body))

	if code != 200 {
		return string(body), errors.New(string(code))
	}

	return string(body), nil
}

/*
func (h *httpUtil) Get(url string) (string, error) {
	log.Debug("GET", url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	return handleResp(resp)
}
*/
