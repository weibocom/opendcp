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



package util

import (
	"io/ioutil"
	"os"
)

/**
file工具
 */
func NewFile(baseDir string, name string, isFolder bool) bool {
	if !IsDirExists(baseDir) {
		return false
	}

	newFilePath := baseDir + "/" + name
	if isFolder {
		error := os.Mkdir(newFilePath, 0700) // only the owner has permission to access the file
		if error != nil {
			PrintErrorStack(error)
			return false
		}

		return true
	}

	file, error := os.Create(newFilePath)
	if error != nil {
		PrintErrorStack(error)
		return false
	}

	defer file.Close()
	return true

}

func DeleteFile(path string) bool {
	// recursive delete file
	error := os.RemoveAll(path)
	if error != nil {
		PrintErrorStack(error)
		return false
	}

	return true
}

func ClearFolder(path string) {
	fileInfos, error := ioutil.ReadDir(path)

	if error != nil {
		PrintErrorStack(error)
		return
	}

	for _, fileInfo := range fileInfos {
		DeleteFile(path + "/" + fileInfo.Name())
	}
}

func CopyFile(sourcePath string, newPath string) bool {
	return true
}

func IsDirExists(dir string) bool {
	fileInfo, error := os.Stat(dir)
	if error != nil {
		return os.IsExist(error)
	} else {
		return fileInfo.IsDir()
	}
}

func IsFileExists(file string) bool {
	fileInfo, error := os.Stat(file)
	if error != nil {
		return os.IsExist(error)
	} else {
		return !fileInfo.IsDir()
	}
}
