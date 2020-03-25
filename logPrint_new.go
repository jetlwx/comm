package comm

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

type jlog struct{}

//日志是否显示,默认全部显示

var (
	//LogLevel  日志级别 debug|info|warn|error
	//	LogLevel = "debug"
	// 日志路径,若设置则为日志保存路径，不设置则输出到屏幕
	//DebugLogPath, InfoLogPath, WarnLogPath, ErrorLogPath = "", "", "", ""
	Jlog = new(jlog)
)

func (j *jlog) Debug(args ...interface{}) {
	color := string([]byte{27, 91, 51, 55, 109})
	reset := string([]byte{27, 91, 48, 109})
	if LogLevel == "debug" || LogLevel == "info" || LogLevel == "warn" || LogLevel == "error" {
		j.showlog("D", color, reset, args...)
	}

}

func (j *jlog) Info(args ...interface{}) {
	color := string([]byte{27, 91, 51, 50, 109})
	reset := string([]byte{27, 91, 48, 109})
	if LogLevel == "info" || LogLevel == "warn" || LogLevel == "error" {
		j.showlog("D", color, reset, args...)
	}

}
func (j *jlog) Warn(args ...interface{}) {
	color := string([]byte{27, 91, 51, 51, 109})
	reset := string([]byte{27, 91, 48, 109})
	if LogLevel == "warn" || LogLevel == "error" {
		j.showlog("D", color, reset, args...)
	}

}
func (j *jlog) Error(args ...interface{}) {
	color := string([]byte{27, 91, 51, 49, 109})
	reset := string([]byte{27, 91, 48, 109})
	if LogLevel == "error" {
		j.showlog("E", color, reset, args...)
	}

}

func (j *jlog) showlog(loglevel, color, reset string, args ...interface{}) {
	//pc, file, line, ok := runtime.Caller(1)
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		log.Println(color, "log print faild", reset)
	}

	//	f := runtime.FuncForPC(pc)
	_, filename := filepath.Split(file)
	//fn := strings.Split(f.Name(), "/")
	// var funcName string
	// if len(fn) == 1 {
	// 	funcName = f.Name()
	// } else {
	// 	funcName = fn[len(fn)-1]
	// }

	var msg string
	for _, v := range args {
		msg += " " + fmt.Sprintf("%+v", v)
	}
	log.Println(color, "[", loglevel, "]", filename, ":", line, "|", msg, reset)
	return

}

/*
备用
func writeToLogfile() {
	//要整理
	fi, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("E", "记录日志出错：", err)
	}
	defer fi.Close()

	// 查找文件末尾的偏移量
	n, _ := fi.Seek(0, os.SEEK_END)
	// 从末尾的偏移量开始写入内容

	newMsg := time.Now().Format("2006-01-02 15:04:05") + "[ " + logleve + " ]," + filename + ":" + gconv.String(line) + " " + funcName + "|" + msg + "\r\n"
	fi.WriteAt([]byte(newMsg), n)
	return
}
*/
