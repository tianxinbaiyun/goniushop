package utils

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

// HTTPData HTTPData
type HTTPData struct {
	ErrNo  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

// ReturnHTTPSuccess ReturnHTTPSuccess
func ReturnHTTPSuccess(c *beego.Controller, val interface{}) {

	rtndata := HTTPData{
		ErrNo:  0,
		ErrMsg: "",
		Data:   val,
	}

	data, err := json.Marshal(rtndata)
	if err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = json.RawMessage(string(data))
	}
}

// ReturnHTTPError ReturnHTTPError
func ReturnHTTPError(c *beego.Controller, errno int, errmsg string) {

	rtndata := HTTPData{
		ErrNo:  errno,
		ErrMsg: errmsg,
		Data:   nil,
	}

	data, err := json.Marshal(rtndata)
	if err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = json.RawMessage(string(data))
	}
}

// GetHTTPRtnJSONData GetHTTPRtnJSONData
func GetHTTPRtnJSONData(errno int, errmsg string) interface{} {

	rtndata := HTTPData{
		ErrNo:  errno,
		ErrMsg: errmsg,
		Data:   nil,
	}
	data, _ := json.Marshal(rtndata)

	return json.RawMessage(string(data))

}
