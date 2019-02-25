package comm

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

//日志是否显示,默认全部显示

var LogLevel string = "debug"

//JetLog logleve D--> DEBUG,I-->INFO,W-->WARN,E-->ERROR
//日志默认全部显示，使用时若不想显示某级别的日志，则可将对应全局变量LogLevel=debug|info|warn|eror,默认debug级别
/*
	var (
		//	greenBg      = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
		//whiteBg      = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
		//yellowBg     = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
		redBg = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
		//	blueBg       = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
		//	magentaBg    = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
		//	cyanBg       = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
		// green  = string([]byte{27, 91, 51, 50, 109})
		// white  = string([]byte{27, 91, 51, 55, 109})
		// yellow = string([]byte{27, 91, 51, 51, 109})
		// red    = string([]byte{27, 91, 51, 49, 109})
		//blue         = string([]byte{27, 91, 51, 52, 109})
		//	magenta      = string([]byte{27, 91, 51, 53, 109})
		//cyan         = string([]byte{27, 91, 51, 54, 109})
		reset = string([]byte{27, 91, 48, 109})
		//disableColor = false
	)


*/
func JetLog(logleve string, args ...interface{}) {
	var color []byte

	switch logleve {
	case "D":
		// white  = string([]byte{27, 91, 51, 55, 109})
		color = []byte{27, 91, 51, 55, 109}
	case "I":
		// green  = string([]byte{27, 91, 51, 50, 109})
		color = []byte{27, 91, 51, 50, 109}
	case "W":
		// yellow = string([]byte{27, 91, 51, 51, 109})
		color = []byte{27, 91, 51, 51, 109}
	case "E":
		// red    = string([]byte{27, 91, 51, 49, 109})
		color = []byte{27, 91, 51, 49, 109}
	default:
		//blue         = string([]byte{27, 91, 51, 52, 109})
		color = []byte{27, 91, 51, 52, 109}
	}
	switch LogLevel {
	case "debug":
		if logleve == "D" || logleve == "I" || logleve == "W" || logleve == "E" {
			goto SHOWLOG
		}
		return
	case "info":
		if logleve == "I" || logleve == "W" || logleve == "E" {
			goto SHOWLOG
		}
		return
	case "warn":
		if logleve == "W" || logleve == "E" {
			goto SHOWLOG

		}
		return
	case "error":
		if logleve == "E" {
			goto SHOWLOG

		}
		return
	default:
		return
	}
SHOWLOG:
	reset := string([]byte{27, 91, 48, 109})
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		log.Println(color, "log print faild", reset)
	}

	f := runtime.FuncForPC(pc)
	_, filename := filepath.Split(file)
	fn := strings.Split(f.Name(), "/")
	var funcName string
	if len(fn) == 1 {
		funcName = f.Name()
	} else {
		funcName = fn[len(fn)-1]
	}
	msg := fmt.Sprintf("%+v", args)
	log.Println(string(color), "[", logleve, "]", filename, ":", funcName, line, "|", msg, reset)

}
