package jlog

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
)

const (
	INFO    = "info"
	ERROR   = "error"
	WARNING = "warning"
	DEBUG   = "debug"
	TRACE   = "trace"
	INIT    = "Init"
	FATAL   = "fatal"
)

type Fields map[string]interface{}

type logData struct {
	Src     string                 `json:"src"`
	FName   string                 `json:"fname"`
	MsgType string                 `json:"type"`
	Msg     string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func Log(msgType string, msg string, data map[string]interface{}) {

	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	var lData logData

	_, lData.Src, _, _ = runtime.Caller(1)
	lData.FName = frame.Function
	lData.MsgType = msgType
	lData.Msg = msg
	lData.Data = data

	data_json, err := json.Marshal(lData)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(string(data_json))
	}

	if msgType == FATAL {
		os.Exit(1)
	}

}

func Info(msg string, data map[string]interface{}) {
	Log(INFO, msg, data)
}

func Error(msg string, data map[string]interface{}) {
	Log(ERROR, msg, data)
}

func Warning(msg string, data map[string]interface{}) {
	Log(WARNING, msg, data)
}

func Debug(msg string, data map[string]interface{}) {
	Log(DEBUG, msg, data)
}

func Trace(msg string, data map[string]interface{}) {
	Log(TRACE, msg, data)
}

func Fatal(msg string, data map[string]interface{}) {
	Log(FATAL, msg, data)
}
