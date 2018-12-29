package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path"
	"path/filepath"
	"runtime/debug"
	"sync"
	"time"
)

const (
	// VERSION 版本
	VERSION = "1.0.0"
	// DATAFORMAT 日期格式（用于文件命名）
	DATAFORMAT = "2006-01-02"
	// TIMEFORMAT 时间格式（日志时间格式）
	TIMEFORMAT = "2016/01/02 15:04:05"
	// SPACE 参数分割
	SPACE = " "
	// TABLE 日志文件行分割符
	TABLE = "\t"
	// JOIN 参数连接符
	JOIN = "&"
	// FILEOPERATMODE 文件操作权限
	FILEOPERATMODE = 0644
	// FILECREATEMODE 文件创建权限
	FILECREATEMODE = 0666
	// LABEL 标签
	LABEL = "[_logger_]"
)

// 日志文件存储模式
const (
	// 普通模式，不分割
	LOGFILESAVEUSUAL = 1
	// 文件大小分割模式
	LOGFILESAVESIZE = 2
	// 日期分割模式
	LOGFILESAVEDATE = 3
)

// 文件大小单位
const (
	_ = iota
	// KB kb
	KB int64 = 1 << (iota * 10)
	// MB mb
	MB
	// GB gb
	GB
	// TB tb
	TB
)

const (
	// EXTENNAME 日志文件扩展名
	EXTENNAME = ".log"
	// CHECKTIME 定时检测是否分割周期
	CHECKTIME = 900 * time.Millisecond
	// WRITETIME 定时写入文件周期
	WRITETIME = 1300 * time.Millisecond
)

var (
	// ISDEBUG 调试模式
	ISDEBUG = false
	// TIMEERWRITE 定时写入文件
	TIMEERWRITE = false
)

// 接口信息

// Logger 日志结构信息
type Logger struct {
	// 日志类型
	logType uint
	// 日志文件存储路径
	path string
	// 目录
	dir string
	// 文件名
	filename string
	// 文件大小
	maxFileSize int64
	// 文件个数
	maxFileCount int64
	// 日分割
	dailyRolling bool
	// 大小分割
	sizeRolling bool
	// 普通模式 不分割
	normalRolling bool
	// 大小分割文件的当前序号
	suffix int
	// 文件时间
	date *time.Time
	// 缓冲锁
	muBuf *sync.Mutex
	// 文件锁
	muFile *sync.Mutex
	// 文件句柄
	logfile *os.File
	// 坚实定时器
	timer *time.Timer
	// 批量写入定时器
	writeTimer *time.Timer
	// 缓冲区
	buf *bytes.Buffer
}

// New 获取日志对象
func New() *Logger {
	this := &Logger{}
	this.buf = &bytes.Buffer{}
	this.muBuf = new(sync.Mutex)
	this.muFile = new(sync.Mutex)
	return this
}

// Printf 格式化输出
func (l *Logger) Printf(format string, a ...interface{}) {
	defer func() {
		if !TIMEERWRITE {
			go l.bufWrite()
		}
	}()
	tp := fmt.Sprintf(format, a...)
	l.muBuf.Lock()
	defer l.muBuf.Unlock()
	l.buf.WriteString(
		fmt.Sprintf(
			"%s\t%d\t%s\n",
			time.Now().Format(TIMEFORMAT),
			l.logType,
			tp,
		),
	)
	return
}

// Println 按行输出
func (l *Logger) Println(a ...interface{}) {
	defer func() {
		if !TIMEERWRITE {
			go l.bufWrite()
		}
	}()
	tp := fmt.Sprintln(a...)
	l.muBuf.Lock()
	defer l.muBuf.Unlock()
	l.buf.WriteString(
		fmt.Sprintf(
			"%s\t%d\t%s",
			time.Now().Format(TIMEFORMAT),
			l.logType,
			tp,
		),
	)
}

// SetDeBug 测试模式
func (l *Logger) SetDeBug(isDebug bool) {
	ISDEBUG = isDebug
}

// SetTimeWrite 定时写入
func (l *Logger) SetTimeWrite(timeWrite bool) {
	TIMEERWRITE = timeWrite
}

// SetType 日志类型
func (l *Logger) SetType(tp uint) {
	l.logType = tp
}

// SetRollingFile 大小分割
func (l *Logger) SetRollingFile(dir, file string, maxn int64, maxs int64, u int64) {
	// 合法性判断
	if l.sizeRolling ||
		l.dailyRolling ||
		l.normalRolling {
		log.Println(LABEL, "mode can't be changed! ")
		return
	}

	// 设置各个标识符
	l.sizeRolling = true
	l.dailyRolling = false
	l.normalRolling = false

	// 设置日志器各个参数
	l.maxFileCount = maxn
	l.maxFileSize = maxs * u
	l.dir = dir
	l.filename = file
	for i := 1; i <= int(maxn); i++ {
		sizeFile := fmt.Sprint(
			filepath.Join(l.dir, l.filename),
			EXTENNAME,
			".",
			fmt.Sprintf("%5d", i),
		)
		if IsExist(sizeFile) {
			l.suffix = i
		} else {
			break
		}
	}
	// 文件实时写入
	l.path = fmt.Sprint(
		filepath.Join(l.dir, l.filename),
		EXTENNAME,
	)
	l.startLogger(l.path)
}

// SetRollingDaily 日期分割
func (l *Logger) SetRollingDaily(dir, file string) {
	// 输入合法性
	if l.sizeRolling ||
		l.dailyRolling ||
		l.normalRolling {
		log.Println(LABEL, "mode can't be changed! ")
		return
	}

	// 设置模式标识符
	l.sizeRolling = false
	l.dailyRolling = true
	l.normalRolling = false

	// 设置日志器的各个参数
	l.dir = dir
	l.filename = file
	l.date = getNowFormatDate(DATAFORMAT)
	l.startLogger(
		fmt.Sprint(
			filepath.Join(l.dir, l.filename),
			EXTENNAME,
			".",
			l.date.Format(DATAFORMAT),
		),
	)
}

// SetRollingNormal 普通模式
func (l *Logger) SetRollingNormal(dir, file string) {
	// 输入合法性
	if l.sizeRolling ||
		l.dailyRolling ||
		l.normalRolling {
		log.Println(LABEL, "mode can't be changed! ")
		return
	}

	// 设置模式标识符
	l.sizeRolling = false
	l.dailyRolling = false
	l.normalRolling = true

	// 设置日志器的各个参数
	l.dir = dir
	l.filename = file
	l.startLogger(
		fmt.Sprint(
			filepath.Join(l.dir, l.filename),
			EXTENNAME,
		),
	)
}

// Close 关闭日志器
func (l *Logger) Close() {
	l.muFile.Lock()
	defer l.muFile.Unlock()
	l.muBuf.Lock()
	defer l.muBuf.Unlock()

	if l.timer != nil {
		l.timer.Stop()
	}
	if l.writeTimer != nil {
		l.writeTimer.Stop()
	}

	if l.logfile != nil {
		err := l.logfile.Close()
		if err != nil {
			log.Println(LABEL, "file close error !", err)
		}
	} else {
		log.Println(LABEL, "file has been close! ")
	}

	// 清理
	l.sizeRolling = false
	l.dailyRolling = false
	l.normalRolling = false
}

// 初始化日志器
func (l *Logger) startLogger(fp string) {
	defer func() {
		if e, ok := recover().(error); ok {
			log.Println(LABEL, "WARN: panic - %v", e)
			log.Println(LABEL, string(debug.Stack()))
		}
	}()

	// 初始化空间
	var err error
	l.buf = &bytes.Buffer{}
	l.muBuf = new(sync.Mutex)
	l.muFile = new(sync.Mutex)
	l.path = fp
	checkFileDir(fp)
	l.logfile, err = os.OpenFile(
		fp,
		os.O_RDWR|os.O_APPEND|os.O_CREATE,
		FILECREATEMODE,
	)
	if err != nil {
		log.Println(LABEL, "\n\n\n\nn\n\n\n\n\nOpenFile error! ", fp, err)
	}

	// 初始化监视线程
	go func() {
		l.timer = time.NewTimer(CHECKTIME)
		l.writeTimer = time.NewTimer(WRITETIME)
		if !TIMEERWRITE {
			l.writeTimer.Stop()
		}
		for {
			select {
			case <-l.timer.C:
				l.fileCheck()
				if ISDEBUG && false {
					log.Printf("*")
				}
				break
			case <-l.writeTimer.C:
				l.bufWrite()
				if ISDEBUG && false {
					log.Printf(".")
				}
				break
			}
		}
	}()

	if ISDEBUG {
		jstr, err := json.Marshal(l)
		if err == nil {
			log.Println(LABEL, VERSION, string(jstr))
		}
	}
}

// 文件检查 会锁定文件
func (l *Logger) fileCheck() {
	// 边界检查
	if nil == l.muFile ||
		nil == l.logfile ||
		"" == l.path {
		return
	}
	defer func() {
		if e, ok := recover().(error); ok {
			log.Println(LABEL, "WARN: panic - %v", e)
			log.Println(LABEL, string(debug.Stack()))
		}
	}()

	// 重命名判断
	var RanameFlag = false
	var CheckTime = CHECKTIME
	l.timer.Stop()
	defer l.timer.Reset(CheckTime)
	if l.dailyRolling {
		// 日分割模式
		now := getNowFormatDate(DATAFORMAT)
		if now != nil &&
			l.date != nil &&
			now.After(*l.date) {
			RanameFlag = true
		} else {
			du := l.date.UnixNano() - now.UnixNano()
			abs := math.Abs(float64(du))
			CheckTime = CheckTime * time.Duration(abs/abs)
		}
	} else if l.sizeRolling {
		// 文件大小模式
		if l.path != "" &&
			l.maxFileCount >= 1 &&
			fileSize(l.path) >= l.maxFileSize {
			RanameFlag = true
		}
	} else if l.normalRolling {
		// 普通模式
		RanameFlag = false
	}

	// 重名操作
	if RanameFlag {
		l.muFile.Lock()
		defer l.muFile.Unlock()
		if ISDEBUG {
			log.Println(LABEL, l.path, "is need rename. ")
		}
		l.fileRename()
	}
	return
}

// 重命名文件
func (l *Logger) fileRename() {
	var err error
	var newName string
	var oldName string
	defer func() {
		if ISDEBUG {
			log.Println(
				LABEL,
				oldName,
				"->",
				newName,
				":",
				err,
			)
		}
	}()

	if l.dailyRolling {
		// 日期分割模式(文件不重命名)
		oldName = l.path
		newName = l.path
		l.date = getNowFormatDate(DATAFORMAT)
		l.path = fmt.Sprint(
			filepath.Join(l.dir, l.filename),
			EXTENNAME,
			".",
			l.date.Format(DATAFORMAT),
		)
	} else if l.sizeRolling {
		// 大小分割模式
		suffix := int(l.suffix%int(l.maxFileCount) + 1)
		oldName = l.path
		newName = fmt.Sprint(
			filepath.Join(l.dir, l.filename),
			".",
			fmt.Sprintf("%05d", suffix),
		)
		l.suffix = suffix

	} else if l.normalRolling {
		// 普通模式
	}

	// 处理旧文件
	l.logfile.Close()
	if oldName != "" &&
		newName != "" &&
		oldName != newName {
		if IsExist(newName) {
			err := os.Remove(newName)
			if err != nil {
				log.Println(LABEL, "remove file err ", err.Error())
			}
			l.path = newName
		}
		err := os.Rename(oldName, newName)
		if err != nil {
			log.Println(LABEL, "rename file error ", err.Error())
		}
	}

	// 创建新文件
	l.logfile, err = os.OpenFile(
		l.path,
		os.O_RDWR|os.O_APPEND|os.O_CREATE,
		FILECREATEMODE,
	)
	if err != nil {
		log.Println(LABEL, "create file error! ")
	}
	return
}

// 缓冲区写入文件
func (l *Logger) bufWrite() {
	if nil == l.buf ||
		"" == l.path ||
		nil == l.logfile ||
		nil == l.muBuf ||
		nil == l.muFile {
		return
	}
	var writeTime = WRITETIME
	if nil != l.writeTimer {
		l.writeTimer.Stop()
		defer l.writeTimer.Reset(writeTime)
	}
	l.muBuf.Lock()
	defer l.muBuf.Unlock()
	l.muFile.Lock()
	defer l.muFile.Unlock()
	defer l.buf.Reset()
	n, err := io.WriteString(l.logfile, l.buf.String())
	if err != nil {
		checkFileDir(l.path)
		l.logfile, err = os.OpenFile(
			l.path,
			os.O_RDWR|os.O_APPEND|os.O_CREATE,
			FILEOPERATMODE,
		)
		if err != nil {
			log.Println(LABEL, "log bufWrite() error! ")
		}
	}
	if n == 0 {
		writeTime = WRITETIME
	} else {
		writeTime = WRITETIME * time.Duration(n/n)
	}
}

// 检查文件目录，不存在创建
func checkFileDir(fp string) {
	p, _ := path.Split(fp)
	d, err := os.Stat(p)
	if err != nil || !d.IsDir() {
		if err := os.MkdirAll(p, FILECREATEMODE); err != nil {
			log.Println(LABEL, "CheckFileDir() Create dir failed! ", fp)
		}
	}
}

// IsExist 判断文件否存在
func IsExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

// 获取当前制定格式的日期
func getNowFormatDate(fm string) *time.Time {
	t, err := time.Parse(fm, time.Now().Format(fm))
	if err != nil {
		log.Println(LABEL, "getNowFormatDate() ", err.Error())
		t = time.Time{}
	}
	return &t
}

// 获取文件大小
func fileSize(fp string) int64 {
	f, err := os.Stat(fp)
	if err != nil {
		if ISDEBUG {
			log.Println(LABEL, err.Error())
		}
		return 0
	}
	return f.Size()
}
