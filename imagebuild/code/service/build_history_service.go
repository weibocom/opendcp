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
	"database/sql"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
	"weibo.com/opendcp/imagebuild/code/env"
	"weibo.com/opendcp/imagebuild/code/model"
	"weibo.com/opendcp/imagebuild/code/util"
)

var (
	BUILDING = 0
	SUCCESS  = 1
	FAIL     = 2
)
/**
build历史记录servie
 */
type BuildHistoryService struct {
	db *sql.DB

	insertStmt *sql.Stmt
	updateStmt *sql.Stmt

	queryByProjectStmt *sql.Stmt
	queryByIdStmt      *sql.Stmt
	queryLastBuildStmt *sql.Stmt
}

var buildHistoryServiceInstance *BuildHistoryService

var onceForBuildHistoryService sync.Once

func GetBuildHistoryServiceInstance() *BuildHistoryService {
	if buildHistoryServiceInstance == nil {
		onceForBuildHistoryService.Do(func() {
			// 双重检查
			if buildHistoryServiceInstance == nil {
				buildHistoryServiceInstance = &BuildHistoryService{}
				db, err := sql.Open("mysql", env.MYSQL_USER+":"+env.MYSQL_PASSWORD+"@tcp("+env.MYSQL_HOST+":"+env.MYSQL_PORT+")/image_build")
				if err != nil {
					log.Errorf("history service init failed: %s", err)
					buildHistoryServiceInstance = nil
					return
				}
				buildHistoryServiceInstance.db = db

				// 预编译sql语句
				stmt, err := buildHistoryServiceInstance.db.Prepare("insert into t_build_history (project, operator, time, state) values (?, ?, ?, ?)")
				if err != nil {
					log.Errorf("hitory service init failed: %s", err)
					buildHistoryServiceInstance = nil
					return
				}
				buildHistoryServiceInstance.insertStmt = stmt

				stmt, err = buildHistoryServiceInstance.db.Prepare("update t_build_history set state = ? where id = ?")
				if err != nil {
					log.Errorf("hitory service init failed: %s", err)
					buildHistoryServiceInstance = nil
					return
				}
				buildHistoryServiceInstance.updateStmt = stmt

				stmt, err = buildHistoryServiceInstance.db.Prepare("select project, operator, state, time from t_build_history where project = ? limit ?,?")
				if err != nil {
					log.Errorf("hitory service init failed: %s", err)
					buildHistoryServiceInstance = nil
					return
				}
				buildHistoryServiceInstance.queryByProjectStmt = stmt

				stmt, err = buildHistoryServiceInstance.db.Prepare("select project, operator, state, time from t_build_history where id = ?")
				if err != nil {
					log.Errorf("hitory service init failed: %s", err)
					buildHistoryServiceInstance = nil
					return
				}

				buildHistoryServiceInstance.queryByIdStmt = stmt

				stmt, err = buildHistoryServiceInstance.db.Prepare("select project, operator, state, time  from t_build_history where project = ? order by id desc limit 1")
				if err != nil {
					log.Errorf("hitory service init failed: %s", err)
					buildHistoryServiceInstance = nil
					return
				}

				buildHistoryServiceInstance.queryLastBuildStmt = stmt
			}
		})
	}

	return buildHistoryServiceInstance
}

func (s *BuildHistoryService) InsertRecord(operator string, project string) int64 {
	buildTime := time.Now()
	result, err := s.insertStmt.Exec(project, operator, buildTime, BUILDING)
	if err != nil {
		log.Errorf("insert build record error: %s", err)
		return -1
	}

	affected, error := result.RowsAffected()
	if error != nil {
		log.Errorf("insert build record error: %s", err)
		return -1
	}

	if affected == 0 {
		log.Errorf("insert build record error: %s", err)
		return -1
	}

	lastInsertId, error := result.LastInsertId()
	if error != nil {
		log.Errorf("insert build record error: %s", err)
		return -1
	}

	return lastInsertId
}

func (s *BuildHistoryService) UpdateRecord(id int64, state int) {
	_, error := s.updateStmt.Exec(state, id)
	if error != nil {
		util.PrintErrorStack(error)
	}
}

func (s *BuildHistoryService) QueryRecordList(cursor int, offset int, project string) []*model.BuildHistory {
	records := make([]*model.BuildHistory, 0)
	rows, err := s.queryByProjectStmt.Query(project, cursor, offset)
	if err != nil {
		return records
	}

	for rows.Next() {
		var project string
		var operator string
		var state int
		var time time.Time
		rows.Scan(&project, &operator, &state, &time)

		record := model.GetBuildHistory(project, operator, time, state)
		records = append(records, record)
	}

	return records
}

func (s *BuildHistoryService) QueryLastBuildRecord(project string) *model.BuildHistory {
	rows, err := s.queryLastBuildStmt.Query(project)
	if err != nil {
		log.Errorf("query last build record error: %s", err)
		return nil
	}

	if rows.Next() {
		var project string
		var operator string
		var state int
		var time time.Time

		rows.Scan(&project, &operator, &state, &time)

		record := model.GetBuildHistory(project, operator, time, state)
		return record
	}

	return nil
}

func (s *BuildHistoryService) QueryRecord(id int) *model.BuildHistory {
	rows, err := s.queryByIdStmt.Query(id)
	if err != nil {
		log.Errorf("query build record error: %s", err)
		return nil
	}

	if rows.Next() {
		var project string
		var operator string
		var state int
		var time time.Time
		rows.Scan(&project, &operator, &state, &time)

		record := model.GetBuildHistory(project, operator, time, state)
		return record
	}

	return nil
}
