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

package service

import (
	"fmt"
	stackError "github.com/go-errors/errors"
	"io/ioutil"
	"net/http"
	"sync"
	"weibo.com/opendcp/imagebuild/code/util"
)

/**
http 服务
 */
type HttpService struct {
}

var httpServiceInstance *HttpService
var httpServiceOnce sync.Once

func GetHttpServiceInstance() *HttpService {
	httpServiceOnce.Do(func() {
		httpServiceInstance = &HttpService{}
	})
	return httpServiceInstance
}

func (service *HttpService) Get(url string) ([]byte, *stackError.Error) {
	resp, error := http.Get(url)
	if error != nil {
		util.PrintErrorStack(error)
		return nil, util.ErrorWrapper(error)
	}

	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		util.PrintErrorStack(error)
		return nil, util.ErrorWrapper(error)
	}

	fmt.Println(string(body))
	return body, nil
}
