package controllers

import (
	"github.com/tianxinbaiyun/goniushop/services"
	"github.com/tianxinbaiyun/goniushop/utils"
)

func getLoginUserId() int {
	intuserId := utils.String2Int(services.LoginUserId)
	return intuserId
}
