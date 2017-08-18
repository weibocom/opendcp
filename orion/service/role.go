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
	"bufio"
	"fmt"
	"github.com/mholt/archiver"
	"os"
	"strconv"
	"strings"
	"weibo.com/opendcp/orion/models"
)

const (
	ROLES_URL           = "./roles/"
	ROLES_REPO          = "./tmp/"
	TASKS               = "tasks/"
	TEMPLATES           = "templates/"
	VARS                = "vars/"
	MAIN 				= "main"
	ANSIBLE_INCLUDE     = "- include: "
	DEFAULT_FILE_PERM   = 0666
	DEFAULT_FOLDER_PERM = 0755
	YAML_SUFFIX         = ".yml"
)

type RoleService struct {
	BaseService
}

var (
	roleService = &RoleService{}
)

func (f *RoleService) CheckRoleParams(params []string) (string, error) {
	for _, param := range params {
		if len(param) == 0 {
			continue
		}

		ids := strings.Split(param, ",")
		for _, id := range ids {
			if _, err := strconv.Atoi(id); err != nil {
				return param, err
			}
		}
	}
	return "", nil
}

func (f *RoleService) PackRoles(stepName string, rolesName []string) error {
	for _, roleName := range rolesName {
		roleName = ROLES_URL + roleName
	}
	packageName := ROLES_REPO + stepName

	err := archiver.TarGz.Make(packageName, rolesName)

	return err
}

func (f *RoleService) BuildRoleFile(role *models.Role) error {
	resoures := []models.RoleResource{}

	idsString := []string{}
	if len(role.Templates) > 0 {
		idsString = append(idsString, role.Templates)
	}
	if len(role.Tasks) > 0 {
		idsString = append(idsString, role.Tasks)
	}
	if len(role.Vars) > 0 {
		idsString = append(idsString, role.Vars)
	}

	ids := strings.Split(strings.Join(idsString, ","), ",")
	if len(ids) < 1 {
		return fmt.Errorf("The role must contain some resource.")
	}


	for _, id := range ids {
		idInt, _ := strconv.Atoi(id)
		resource := &models.RoleResource{Id: idInt}
		err := f.GetBase(resource)
		if err != nil {
			return err
		}

		resoures = append(resoures, *resource)
	}

	if err := f.WriteRoleFile(role.Name, resoures); err != nil {
		return err
	}

	return nil
}

func (f *RoleService) WriteRoleFile(roleName string, resources []models.RoleResource) error {

	if len(roleName) < 1 {
		return fmt.Errorf("the roleName cant be empty.")
	}

	if err := os.MkdirAll(ROLES_URL+roleName+"/"+TASKS, DEFAULT_FOLDER_PERM); err != nil {
		return err
	}

	if err := os.MkdirAll(ROLES_URL+roleName+"/"+TEMPLATES, DEFAULT_FOLDER_PERM); err != nil {
		return err
	}

	if err := os.MkdirAll(ROLES_URL+roleName+"/"+VARS, DEFAULT_FOLDER_PERM); err != nil {
		return err
	}

	// write resource
	taskNames := []string {}
	varNames := []string {}
	if len(resources) > 0 {
		for _, resource := range resources {
			suffix := ""
			switch resource.ResourceType {
			case "template":
				suffix = ""
			case "var":
				suffix = YAML_SUFFIX
				varNames = append(varNames, resource.Name)
			case "task":
				suffix = YAML_SUFFIX
				taskNames = append(taskNames, resource.Name)
			default:
				suffix = YAML_SUFFIX
			}
			outputFile, err := os.OpenFile(ROLES_URL+roleName+"/"+resource.ResourceType+"s/"+resource.Name+suffix, os.O_WRONLY|os.O_CREATE, DEFAULT_FILE_PERM)
			if err != nil {
				return err
			}
			defer outputFile.Close()
			outputWriter := bufio.NewWriter(outputFile)
			outputWriter.WriteString(resource.ResourceContent)
			outputWriter.Flush()
		}

		// generate the main.yml of resource
		if len(taskNames) > 0 {
			taskOutputFile, err := os.OpenFile(ROLES_URL+roleName+"/"+TASKS+MAIN+YAML_SUFFIX, os.O_WRONLY|os.O_CREATE, DEFAULT_FILE_PERM)
			defer taskOutputFile.Close()

			taskOutputWrite := bufio.NewWriter(taskOutputFile)
			taskOutputWrite.WriteString("---\n")
			for _, name := range taskNames {
				if err != nil {
					return err
				}
				taskOutputWrite.WriteString(ANSIBLE_INCLUDE+name+YAML_SUFFIX+"\n")
			}
			taskOutputWrite.Flush()
		}

		if len(varNames) > 0 {
			varOutputFile, err := os.OpenFile(ROLES_URL+roleName+"/"+VARS+MAIN+YAML_SUFFIX, os.O_WRONLY|os.O_CREATE, DEFAULT_FILE_PERM)
			defer varOutputFile.Close()

			varOutputWrite := bufio.NewWriter(varOutputFile)
			varOutputWrite.WriteString("---\n")
			for _, name := range varNames {
				if err != nil {
					return err
				}
				varOutputWrite.WriteString(ANSIBLE_INCLUDE+name+YAML_SUFFIX+"\n")
			}
			varOutputWrite.Flush()
		}

	} else {
		return fmt.Errorf("A role must contain some resources.")
	}

	return nil
}

func (f *RoleService) RemoveRoleFile(role *models.Role) error {
	roleName := role.Name

	_, err := os.Stat(ROLES_URL + roleName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	err = os.RemoveAll(ROLES_URL + roleName)
	if err != nil {
		return err
	}

	return nil
}

func (f *RoleService) UpdateRoleFile(role *models.Role) error {
	if err := f.RemoveRoleFile(role); err != nil {
		return err
	}

	if err := f.BuildRoleFile(role); err != nil {
		_ = f.RemoveRoleFile(role)
		return err
	}

	return nil
}
