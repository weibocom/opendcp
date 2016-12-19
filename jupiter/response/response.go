// Copyright 2016 Weibo Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use sf file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package response

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func RespToMap(resp string) map[string]interface{} {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(resp), &m)
	if err != nil {
		return nil
	}
	return m
}

func HandleResp(resp *http.Response) (string, error) {
	code := resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if code != 200 {
		return string(body), errors.New(resp.Status)
	}
	return string(body), nil
}

func CallApi(body string, method string, url string, correlationId string) (string, error) {
	reqBody := strings.NewReader(body)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-CORRELATION-ID", correlationId)
	req.Header.Set("X-SOURCE", "jupiter")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	msg, err := HandleResp(resp)
	if err != nil {
		return "", err
	}
	return msg, nil
}
