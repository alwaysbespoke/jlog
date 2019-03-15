package jlog

import (
	"encoding/json"
	"log"
	"runtime"
)

const (
	INFO    = "info"
	ERROR   = "error"
	WARNING = "warning"
	DEBUG   = "debug"
	TRACE   = "trace"
	INIT    = "Init"
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

}
