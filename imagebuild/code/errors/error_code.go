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

package errors

var (
	OK = 10000

	CREATE_PROJECT_ALREADY_EXIST int = 10001 // already exist

	CLONE_SRC_PROJECT_NOT_EXIST int = 10002 // src project not exist

	DELETE_PROJECT_NOT_EXIST int = 10003 // project to delete not exist

	EXTENSION_INTERFACE_PARAM_ERROR int = 10004 //

	EXTENSION_INTERFACE_CALL_ERROR int = 10005 //

	PARAMETER_INVALID = 10006

	BUILD_PROJECT_NOT_EXIST = 10007

	PROJECT_NOT_EXIST = 10008

	PLUGIN_NOT_EXIST = 10009

	INTERNAL_ERROR = 10100

	errorToMessage = map[int]string{
		OK: "ok",
		CREATE_PROJECT_ALREADY_EXIST:    "project already exist",
		CLONE_SRC_PROJECT_NOT_EXIST:     "src project not exist",
		DELETE_PROJECT_NOT_EXIST:        "project to delete not exist",
		EXTENSION_INTERFACE_PARAM_ERROR: "param error",
		EXTENSION_INTERFACE_CALL_ERROR:  "call error",
		PARAMETER_INVALID:               "parameter invalid",
		BUILD_PROJECT_NOT_EXIST:         "project to build not exist",
		PROJECT_NOT_EXIST:               "project not exist",
		PLUGIN_NOT_EXIST:                "plugin not exist",
		INTERNAL_ERROR:                  "server internal error"}
)

func ErrorCodeToMessage(errorCode int) string {
	return errorToMessage[errorCode]
}
