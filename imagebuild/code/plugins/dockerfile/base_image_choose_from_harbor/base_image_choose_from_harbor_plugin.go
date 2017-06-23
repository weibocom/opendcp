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



package main

import (
	"bytes"
	"encoding/json"
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/dullgiulio/pingo"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"weibo.com/opendcp/imagebuild/code/service"
	"weibo.com/opendcp/imagebuild/code/util"
)

/**
从harbor获取基础镜像
 */
type BaseImageChooseFromHarborPlugin struct{}

func (p *BaseImageChooseFromHarborPlugin) Process(params map[string]interface{}, resp *string) error {
	var currentDockerfile string = params["input"].(string)
	config := make(map[string]interface{}, 0)
	json.Unmarshal([]byte(params["config"].(string)), &config)

	handleInput := config["handleInput"].(string)
	newDockerfile, error := service.GetDockerFileOperatorInstance().From(currentDockerfile, handleInput)
	if error != nil {
		return error
	}
	*resp = newDockerfile
	return nil
}

func (p *BaseImageChooseFromHarborPlugin) BaseImageList(params map[string]interface{}, resp *interface{}) error {
	var cluster string = params["cluster"].(string)
	allImages := make([]string, 0)
	harborAddress := util.DefaultValue(params, "harborAddress").(string)
	if harborAddress == "" {
		harborAddress = params["defaultHarborAddress"].(string)
		if harborAddress == "" {
			*resp = allImages
			return errors.New("harborAddress is empty")
		}
	}

	harborUser := util.DefaultValue(params, "user").(string)
	if harborUser == "" {
		harborUser = params["defaultHarborUser"].(string)
		if harborUser == "" {
			*resp = allImages
			return errors.New("harbor user is empty")
		}
	}

	harborPassword := util.DefaultValue(params, "password").(string)
	if harborPassword == "" {
		harborPassword = params["defaultHarborPassword"].(string)
		if harborPassword == "" {
			*resp = allImages
			return errors.New("harbor password is empty")
		}
	}

	var harborAddressWithoutProtocol string

	if strings.HasPrefix(harborAddress, "http://") {
		harborAddressWithoutProtocol = harborAddress[7:]
	} else if strings.HasPrefix(harborAddress, "https://") {
		harborAddressWithoutProtocol = harborAddress[8:]
	} else {
		harborAddressWithoutProtocol = harborAddress
		harborAddress = "http://" + harborAddress
	}

	sessionId := p.doLogin(harborAddress, harborUser, harborPassword)
	if sessionId == "" {
		*resp = allImages
		return errors.New("login harbor error")
	}

	projects := p.projects(harborAddress, sessionId, cluster)

	// 查询每个project中的image
	util.LogInit("/tmp/imagebuild.log")
	for _, project := range projects {
		log.Info("Form harbor get projectName: "+string(project["name"].(string)))
		if strings.Compare(string(project["name"].(string)), cluster) != 0{
			continue
		}
		log.Info("Form harbor get projectName:"+string(project["name"].(string))+ " == cluster: " + cluster)
		id := strconv.Itoa(int(project["project_id"].(float64)))
		images := p.images(harborAddress, id, sessionId)
		for _, image := range images {
			tags := p.tags(harborAddress, image, sessionId)
			for _, tag := range tags {
				log.Infof("baseImage:%s", harborAddressWithoutProtocol+"/"+image+":"+tag)
				allImages = append(allImages, harborAddressWithoutProtocol+"/"+image+":"+tag)
			}
		}
	}
	*resp = allImages
	return nil
}

func (p *BaseImageChooseFromHarborPlugin) doLogin(harborAddress string, user string, password string) string {
	loginForm := make(url.Values)
	loginForm["principal"] = []string{user}
	loginForm["password"] = []string{password}

	resp, error := http.PostForm(harborAddress+"/login", loginForm)
	if error != nil {
		log.Errorf("login error: %s", error)
		return ""
	}

	cookies := resp.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "beegosessionID" {
			return cookie.Value
		}
	}

	return ""
}

func (p *BaseImageChooseFromHarborPlugin) projects(harborAddress string, sessionId string, cluster string) []map[string]interface{} {

	//　查询harbor中的project
	projects := make([]map[string]interface{}, 0)
	urlToGetProjectsInHarbor := harborAddress + "/api/projects?is_public=0&project_name=" + url.QueryEscape(cluster)

	client := http.Client{}
	req, error := http.NewRequest("GET", urlToGetProjectsInHarbor, nil)
	if error != nil {
		log.Errorf("get projects in harbor error: %s", error)
		return projects
	}
	req.AddCookie(&http.Cookie{Name: "beegosessionID", Value: sessionId})

	resp, error := client.Do(req)

	if error != nil {
		log.Errorf("get projects in harbor error: %s", error)
		return projects
	}

	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		log.Errorf("get projects in harbor error: %s", error)
		return projects
	}

	error = json.Unmarshal(body, &projects)
	if error != nil {
		log.Errorf("%s", error)
		return projects
	}

	return projects
}

func (p *BaseImageChooseFromHarborPlugin) images(harborAddress string, id string, sessionId string) []string {
	images := make([]string, 0)
	urlToGetImages := harborAddress + "/api/repositories?project_id=" + id + "&q="
	client := http.Client{}
	req, error := http.NewRequest("GET", urlToGetImages, nil)
	req.AddCookie(&http.Cookie{Name: "beegosessionID", Value: sessionId})

	resp, error := client.Do(req)
	if error != nil {
		log.Errorf("get images in harbor error: %s", error)
		return images
	}

	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		log.Errorf("get images in harbor error: %s", error)
		return images
	}

	if string(body) == "null" {
		return images
	}

	error = json.Unmarshal(body, &images)
	if error != nil {
		log.Errorf("get images in harbor error: %s", error)
		return images
	}

	return images
}

func (p *BaseImageChooseFromHarborPlugin) tags(harborAddress string, image string, sessionId string) []string {
	tags := make([]string, 0)
	urlToGetTags := harborAddress + "/api/repositories/tags?repo_name=" + url.QueryEscape(image)

	client := http.Client{}
	req, error := http.NewRequest("GET", urlToGetTags, nil)
	req.AddCookie(&http.Cookie{Name: "beegosessionID", Value: sessionId})

	resp, error := client.Do(req)
	if error != nil {
		log.Errorf("get tags in harbor error: %s", error)
		return tags
	}

	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		log.Errorf("get tags in harbor error: %s", error)
		return tags
	}

	if string(body) == "null" {
		return tags
	}

	error = json.Unmarshal(body, &tags)
	if error != nil {
		log.Errorf("get tags in harbor error: %s", error)
		return tags
	}

	return tags
}

func (p *BaseImageChooseFromHarborPlugin) View(configs map[string]interface{}, resp *interface{}) error {
	// 查询所有镜像
	// 生成view
	configsNew := make(map[string]interface{}, 0)
	for key, value := range configs {
		configsNew[key] = value
	}

	// index.html
	configPage := "./index.html"
	if !util.IsFileExists(configPage) {
		*resp = ""
		return errors.New("")
	}

	configPageContent, error := ioutil.ReadFile("./index.html")
	if error != nil {
		*resp = ""
		return errors.New("")
	}

	t := template.New("")

	t.Funcs(template.FuncMap{"defaultV": util.DefaultValue})
	t.Funcs(template.FuncMap{"defaultA": util.DefaultEmptyArray})
	t.Funcs(template.FuncMap{"startwith": util.StartWith})
	t.Funcs(template.FuncMap{"isArray": util.IsArray})

	t, _ = t.Parse(string(configPageContent))
	var htmlContent bytes.Buffer

	t.Execute(&htmlContent, configsNew)
	*resp = htmlContent.String()

	return nil
}

func main() {
	plugin := &BaseImageChooseFromHarborPlugin{}
	pingo.Register(plugin)
	pingo.Run()
}
