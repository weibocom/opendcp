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

package util

import (
	log "github.com/Sirupsen/logrus"
	stackError "github.com/go-errors/errors"
)
/**
打印日志
 */
func PrintErrorStack(error interface{}) {
	log.Error(stackError.New(error).ErrorStack())
}

func ErrorWrapper(error interface{}) *stackError.Error {
	return stackError.New(error)
}

func StackString(error interface{}) string {
	return stackError.New(error).ErrorStack()
}
