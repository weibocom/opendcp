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
