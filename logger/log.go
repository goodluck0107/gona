package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"

	"gitee.com/andyxt/gona/utils"
)

var logLevel int = LogLevelDebug
var printLevel int = LogLevelError

func SetLogLevel(targetLevel int) {
	logLevel = targetLevel
}

func SetPrintLevel(targetLevel int) {
	printLevel = targetLevel
}

type MyLog struct {
	typeName   string
	myLogger   *log.Logger
	m_init     int
	fileDay    int
	fileHour   int
	fileMinute int
	logFile    *os.File
	isMinute   bool
	withTime   bool
}

func NewMyLog(logType string) *MyLog {
	MyLoger := &MyLog{}
	MyLoger.typeName = logType
	MyLoger.m_init = 0
	MyLoger.fileDay = 0
	MyLoger.fileMinute = 0
	MyLoger.isMinute = false
	MyLoger.withTime = true
	return MyLoger
}

func NewMyLogMinute(logType string) *MyLog {
	MyLoger := &MyLog{}
	MyLoger.typeName = logType
	MyLoger.m_init = 0
	MyLoger.fileDay = 0
	MyLoger.fileMinute = 0
	MyLoger.isMinute = true
	MyLoger.withTime = true
	return MyLoger
}

func NewMyLogNoTime(logType string) *MyLog {
	MyLoger := &MyLog{}
	MyLoger.typeName = logType
	MyLoger.m_init = 0
	MyLoger.fileDay = 0
	MyLoger.fileMinute = 0
	MyLoger.isMinute = false
	MyLoger.withTime = false
	return MyLoger
}
func (mylog *MyLog) logInit(fileName string) {
	cdir, err := os.Getwd()
	utils.CheckError(err)
	var path string
	path = cdir

	if runtime.GOOS == "windows" {
		path = path + "\\log\\"
	} else {
		path = path + "/log/"
	}
	err = os.MkdirAll(path, 0777)
	utils.CheckError(err)
	if mylog.m_init == 1 {
		_ = mylog.logFile.Close()
	}
	mylog.logFile, err = os.OpenFile(path+mylog.typeName+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Printf("open file error=%s\r\n", err.Error())
		os.Exit(-1)
	}
	writers := []io.Writer{
		mylog.logFile,
		//os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if mylog.withTime {
		gLogger := log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime)
		mylog.myLogger = gLogger
	} else {
		gLogger := log.New(fileAndStdoutWriter, "", 0)
		mylog.myLogger = gLogger
	}

}
func (mylog *MyLog) SPrintln(str string) {
	now := time.Now()
	d := now.Day()
	h := now.Hour()
	m := now.Minute()
	if mylog.isMinute {
		if mylog.fileDay != d || mylog.fileHour != h || mylog.fileMinute != m {
			s := fmt.Sprintf("%02d%02d%02d%02d.log", now.Month(), d, h, m)
			mylog.logInit(s)
			mylog.m_init = 1
			mylog.fileDay = d
			mylog.fileHour = h
			mylog.fileMinute = m
		} else if mylog.m_init != 1 {
			s := fmt.Sprintf("%02d%02d%02d%02d.log", now.Month(), d, h, m)
			mylog.logInit(s)
			mylog.m_init = 1
			mylog.fileDay = d
			mylog.fileHour = h
			mylog.fileMinute = m
		}
	} else {
		if mylog.fileDay != d || mylog.fileHour != h {
			s := fmt.Sprintf("%02d%02d%02d.log", now.Month(), d, h)
			mylog.logInit(s)
			mylog.m_init = 1
			mylog.fileDay = d
			mylog.fileHour = h
		} else if mylog.m_init != 1 {
			s := fmt.Sprintf("%02d%02d%02d.log", now.Month(), d, h)
			mylog.logInit(s)
			mylog.m_init = 1
		}
	}
	mylog.myLogger.Printf("%v\n", str) // .Println(v...)
}
func (mylog *MyLog) Println(v ...interface{}) {
	now := time.Now()
	d := now.Day()
	h := now.Hour()
	m := now.Minute()
	if mylog.isMinute {
		if mylog.fileDay != d || mylog.fileHour != h || mylog.fileMinute != m {
			s := fmt.Sprintf("%02d%02d%02d%02d.log", now.Month(), d, h, m)
			mylog.logInit(s)
			mylog.m_init = 1
			mylog.fileDay = d
			mylog.fileHour = h
			mylog.fileMinute = m
		} else if mylog.m_init != 1 {
			s := fmt.Sprintf("%02d%02d%02d%02d.log", now.Month(), d, h, m)
			mylog.logInit(s)
			mylog.m_init = 1
			mylog.fileDay = d
			mylog.fileHour = h
			mylog.fileMinute = m
		}
	} else {
		if mylog.fileDay != d || mylog.fileHour != h {
			s := fmt.Sprintf("%02d%02d%02d.log", now.Month(), d, h)
			mylog.logInit(s)
			mylog.m_init = 1
			mylog.fileDay = d
			mylog.fileHour = h
		} else if mylog.m_init != 1 {
			s := fmt.Sprintf("%02d%02d%02d.log", now.Month(), d, h)
			mylog.logInit(s)
			mylog.m_init = 1
		}
	}
	mylog.myLogger.Printf("%v\n", v) // .Println(v...)
}

var startUpLogger *MyLog = NewMyLog("startupLog")
var infoLogger *MyLog = NewMyLog("infoLog")
var errorLogger *MyLog = NewMyLog("errorLog")
var warnLogger *MyLog = NewMyLog("warnLog")
var debugLogger *MyLog = NewMyLog("debugLog")
