package main

import (
	"github.com/op/go-logging"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type task struct {
	ImageName string
	Tag       string
	TaskName  string
	Priority  string
	Ports     string
	IPs       string
}

type database struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	DatabaseName string `json:"database_name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

// 日志
var log = logging.MustGetLogger("example")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05} %{shortfile} %{shortfunc} ▶ %{level:.4s} %{color:reset}  %{message}`,
)

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index+1]
}
func init() {
	// 日志初始化配置
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")
	logging.SetBackend(backend1Leveled, backend2Formatter)
}

func main() {
	b, err := ioutil.ReadFile(GetAppPath() + "scanservice.txt")
	if err != nil {
		log.Error(err.Error())
		return
	}
	id, err := strconv.Atoi(string(b[:len(b)-1]))
	if err != nil {
		log.Error(err.Error())
		return
	}
	idStr := strconv.Itoa(LoadIP(id))
	err = ioutil.WriteFile(GetAppPath()+"scanservice.txt", []byte(idStr), 0644)
	if err != nil {
		log.Error(err.Error())
	}
}
