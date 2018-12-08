package comm

import (
	"log"
	"path/filepath"
	"runtime"
)

//日志是否显示,默认全部显示

var LogLevel string = "debug"

//logleve D--> DEBUG,I-->INFO,W-->WARN,E-->ERROR
//日志默认全部显示，使用时若不想显示某级别的日志，则可将对应全局变量LogLevel=debug|info|warn|eror,默认debug级别
func JetLog(logleve string, args ...interface{}) {
	// switch logleve {
	// case "D":
	// 	if LogLevel == "debug" || LogLevel == "info" || LogLevel == "error" || LogLevel == "warn" {
	// 		goto SHOWLOG
	// 	}
	// 	return
	// case "I":
	// 	if LogLevel == "info" || LogLevel == "error" || LogLevel == "warn" {
	// 		goto SHOWLOG
	// 	}
	// 	return
	// case "W":
	// 	if LogLevel == "warn" || LogLevel == "error" {
	// 		goto SHOWLOG
	// 	}
	// 	return
	// case "E":
	// 	if LogLevel == "error" {
	// 		goto SHOWLOG
	// 	}
	// 	return

	// default:
	// 	return
	// }
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
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		log.Println("log print faild")
	}

	f := runtime.FuncForPC(pc)
	_, filename := filepath.Split(file)

	log.Println("[", logleve, "]", filename, ":", line, "func:", f.Name(), "|", args)
}
