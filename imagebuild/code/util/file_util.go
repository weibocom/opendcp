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
