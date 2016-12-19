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
