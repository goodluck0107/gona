package httpx

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitee.com/andyxt/gona/utils"

	"gitee.com/andyxt/gona/logger"
)

func ProRespDataWrite(Resp http.ResponseWriter, Req *http.Request, requestFunc string, retData interface{}) {
	defer func() {
		Req.Body.Close()
		if recoverErr := recover(); recoverErr != nil {
			logger.Error("ProRespDataWrite :", requestFunc, ", RespErr:", fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	respData, jsonErr := json.Marshal(retData)
	if jsonErr != nil {
		logger.Error("ProRespDataWrite :", requestFunc, ", 响应数据封装错误1 RespErr:", jsonErr)
		errorRespData := "{\"Msg\":\"服务器错误\"}"
		_, writeErr := Resp.Write([]byte(errorRespData))
		logger.Error("ProRespDataWrite :2", string(errorRespData))
		if writeErr != nil {
			logger.Error("ProRespDataWrite :", requestFunc, ", 响应数据写入错误2: writeData:", string(errorRespData), ", RespErr:", writeErr)
		}
	} else {
		_, writeErr := Resp.Write(respData)
		if writeErr != nil {
			logger.Error("ProRespDataWrite :", requestFunc, ", 响应数据写入错误1: writeData:", string(respData), ", RespErr:", writeErr)
		}
	}
}

func ProRespDataWriteMessage(Resp http.ResponseWriter, Req *http.Request, retData string) {
	defer func() {
		Req.Body.Close()
		if recoverErr := recover(); recoverErr != nil {
			logger.Error("ProRespDataWriteMessage,  RespErr:", fmt.Sprint(recoverErr, string(utils.Stack(3))))
		}
	}()
	Header := Resp.Header()
	Header.Set("Access-Control-Allow-Origin", "*")
	Header.Set("Access-Control-Allow-Headers", "Content-Type,Content-Length, Authorization, Accept,X-Requested-With")
	Header.Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	Header.Set("X-Powered-By", "Jetty")
	Header.Set("Content-Type", "json")

	_, writeErr := Resp.Write([]byte(retData))
	if writeErr != nil {
		logger.Error("ProRespDataWriteMessage, net write data:", retData, ", RespErr:", writeErr)
	}
}
