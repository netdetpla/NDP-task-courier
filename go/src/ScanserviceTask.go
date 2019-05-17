package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func PostTask(taskInfo task) {
	//log.Info("Post: " + string(taskInfo))
	res, err := http.PostForm(
		"http://10.0.21.229:8080/task/",
		url.Values{
			"image-name": {taskInfo.ImageName},
			"tag": {taskInfo.Tag},
			"task-name": {taskInfo.TaskName},
			"priority": {taskInfo.Priority},
			"params[]": {taskInfo.IPs, taskInfo.Ports},
		},
	)
	if err != nil {
		log.Error(err.Error())
		return
	}
	if res == nil {
		log.Warning("res is empty")
		return
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Warning(err.Error())
		}
	}()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Info(string(content))
}
func LoadIP(id int) (ipID int) {
	ipID = id
	databaseInfo := new(database)
	b, err := ioutil.ReadFile(GetAppPath() + "config.json")
	if err != nil {
		log.Error(err.Error())
		return
	}
	err = json.Unmarshal(b, databaseInfo)
	if err != nil {
		log.Error(err.Error())
		return
	}
	// 连接建立
	databaseURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?timeout=20s",
		databaseInfo.Username,
		databaseInfo.Password,
		databaseInfo.Host,
		databaseInfo.Port,
		databaseInfo.DatabaseName)
	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		log.Warning(err.Error())
		return
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Warning(err.Error())
		}
	}()
	// 测试数据库连接
	if err = db.Ping(); err != nil {
		log.Warning(err.Error())
		return
	}
	selectIPSQL := "select distinct id, ip from ip where id >= ?"
	rows, err := db.Query(selectIPSQL, id)
	if err != nil {
		log.Warning(err.Error())
		return
	}
	count := 0
	nameCount := 0
	i := new(task)
	i.ImageName = "scanservice"
	i.Tag = "1.0.3"
	i.Priority = "5"
	i.Ports = "21,22,23,25,53,80,110,161,443,8080,3306,1433"
	var ip string
	var ipList string
	for rows.Next() {
		err = rows.Scan(&ipID, &ip)
		if err != nil {
			return
		}
		ipList = ipList + ip + ","
		count++
		if count >= 100 {
			i.TaskName = "task-courier-" + strconv.Itoa(nameCount)
			i.IPs = ipList
			PostTask(*i)
			i.IPs = ""
			count = 0
			nameCount++
		}
	}
	err = rows.Close()
	if err != nil {
		log.Warning(err.Error())
	}
	return
}
