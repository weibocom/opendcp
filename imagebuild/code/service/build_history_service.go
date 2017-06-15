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
				stmt, err := buildHistoryServiceInstance.db.Prepare("insert into t_build_history (project, operator, time, state, logs) values (?, ?, ?, ?, ?)")
				if err != nil {
					log.Errorf("hitory service init failed: %s", err)
					buildHistoryServiceInstance = nil
					return
				}
				buildHistoryServiceInstance.insertStmt = stmt

				stmt, err = buildHistoryServiceInstance.db.Prepare("update t_build_history set state = ?, logs = ? where id = ?")
				if err != nil {
					log.Errorf("hitory service init failed: %s", err)
					buildHistoryServiceInstance = nil
					return
				}
				buildHistoryServiceInstance.updateStmt = stmt

				stmt, err = buildHistoryServiceInstance.db.Prepare("select project, operator, state, time, logs from t_build_history where project = ? limit ?,?")
				if err != nil {
					log.Errorf("hitory service init failed: %s", err)
					buildHistoryServiceInstance = nil
					return
				}
				buildHistoryServiceInstance.queryByProjectStmt = stmt

				stmt, err = buildHistoryServiceInstance.db.Prepare("select project, operator, state, time, logs from t_build_history where id = ?")
				if err != nil {
					log.Errorf("hitory service init failed: %s", err)
					buildHistoryServiceInstance = nil
					return
				}

				buildHistoryServiceInstance.queryByIdStmt = stmt

				stmt, err = buildHistoryServiceInstance.db.Prepare("select project, operator, state, time, logs  from t_build_history where project = ? order by id desc limit 1")
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

	result, err := s.insertStmt.Exec(project, operator, buildTime, BUILDING, "")
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

func (s *BuildHistoryService) UpdateRecord(id int64, logs string, state int) {
	_, error := s.updateStmt.Exec(state, logs, id)
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
		var time_bytes []byte
		var logs string

		rows.Scan(&project, &operator, &state, &time_bytes, &logs)
		build_time,_:= time.Parse("2006-01-02 15:04:05", string(time_bytes))

		record := model.GetBuildHistory(project, operator, build_time, state, logs)
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
		var time_bytes []byte
		var logs string

		rows.Scan(&project, &operator, &state, &time_bytes, &logs)

		build_time,_:= time.Parse("2006-01-02 15:04:05", string(time_bytes))

		record := model.GetBuildHistory(project, operator, build_time, state, logs)

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
		var logs string

		rows.Scan(&project, &operator, &state, &time, &logs)

		record := model.GetBuildHistory(project, operator, time, state, logs)
		return record
	}

	return nil
}
