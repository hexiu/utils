package log

import (
	"log"
	"os"
	"path/filepath"
)

const logPath = "logs"
const logFile = "error.log"
const Prefix = "[StandAlone]"

const LogHigh = 1
const LogMid = 2
const LogSma = 3

// 支持i，2，3，4，5等标识，主要表现是时间记录不一致，可以看源代码的注释详细了解。
const logflag = 5

var logger *log.Logger

func init() {
	initlog()
}

func JudPathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func initlog() {
	PrintAbsPath()
	pathexist := JudPathExist(logPath)
	if pathexist != true {
		os.Mkdir(logPath, os.ModePerm)
	}
	filename := filepath.Join(logPath, logFile)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Panic("create log file error! info:", err)
	}
	logger = log.New(file, "[StandAlone]", logflag)
}

func PrintAbsPath() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir)
}

func CheckErr(err error, logtype int) {
	if err != nil {
		logger.Println(err)
		switch logtype {
		case LogHigh:
			panic(err)
		default:
			break
		}
	}
	return
}
