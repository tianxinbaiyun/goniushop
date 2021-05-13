package services

import (
	"encoding/json"

	"github.com/astaxie/beego/httplib"

	"github.com/astaxie/beego"
	"github.com/tianxinbaiyun/goniushop/utils"
)

// Traces Traces
type Traces struct {
	AcceptTime    string `json:"accept_time"`
	AcceptStation string `json:"accept_station"`
	Remark        string `json:"remark"`
}

// ExpressRtnInfo ExpressRtnInfo
type ExpressRtnInfo struct {
	Success      bool     `json:"success"`
	ShipperCode  string   `json:"shipper_code"`
	ShipperName  string   `json:"shipper_name"`
	LogisticCode string   `json:"logistic_code"`
	IsFinish     int      `json:"is_finish"`
	Traces       []Traces `json:"traces"`
	RequestTime  int64    `json:"request_time"`
}

// ExpressResult ExpressResult
type ExpressResult struct {
	Success bool     `json:"success"`
	State   int      `json:"state"`
	Traces  []Traces `json:"traces"`
}

// QueryExpress QueryExpress
func QueryExpress(shippercode, logisticcode string, ordercode string) ExpressRtnInfo {
	var expressInfo = ExpressRtnInfo{
		Success:      false,
		ShipperCode:  shippercode,
		ShipperName:  "",
		LogisticCode: logisticcode,
		IsFinish:     0,
		Traces:       make([]Traces, 0),
	}
	fromdata := GenerateFromData(shippercode, logisticcode, ordercode)

	posturl := beego.AppConfig.String("express::request_url")

	req := httplib.Post(posturl)
	req.Header("content-type", "application/x-www-form-urlencoded")
	//need fix
	jsondata, _ := json.Marshal(fromdata)
	req.Param("form", string(jsondata))

	var res ExpressResult
	req.ToJSON(&res)
	expressInfo.Success = res.Success
	if res.State == 3 {
		expressInfo.IsFinish = 1
	}
	expressInfo.Traces = append(expressInfo.Traces, res.Traces...)

	return expressInfo

}

// ExpressFromData ExpressFromData
type ExpressFromData struct {
	RequestData string
	EBusinessID string
	RequestType string
	DataSign    string
	DataType    string
}

// GenerateFromData GenerateFromData
func GenerateFromData(shippercode, logisticcode, ordercode string) ExpressFromData {
	requestdata := GenerateRequestData(shippercode, logisticcode, ordercode)
	encoderequestdata, _ := utils.URLEncode(requestdata)
	return ExpressFromData{
		RequestData: encoderequestdata,
		EBusinessID: beego.AppConfig.String("express::appid"),
		RequestType: "1002",
		DataSign:    GenerateDataSign(requestdata),
		DataType:    "2"}

}

// ExpressRequestData ExpressRequestData
type ExpressRequestData struct {
	OrderCode    string
	ShipperCode  string
	LogisticCode string
}

// GenerateRequestData GenerateRequestData
func GenerateRequestData(shippercode, logisticcode, ordercode string) string {

	data, err := json.Marshal(ExpressRequestData{ordercode, shippercode, logisticcode})
	if err != nil {
		return ""
	}
	return string(data)

}

// GenerateDataSign GenerateDataSign
func GenerateDataSign(requestdata string) string {

	md5str := utils.Md5(requestdata)
	appkey := beego.AppConfig.String("express::appkey")
	base64str := utils.Base64Encode(md5str + appkey)
	rv, err := utils.URLEncode(base64str)
	if err == nil {
		return ""
	}
	return rv

}
