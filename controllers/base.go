package controllers

import (
	"github.com/tianxinbaiyun/goniushop/services"
	"github.com/tianxinbaiyun/goniushop/utils"
)

func getLoginUserID() int {
	intUserID := utils.String2Int(services.LoginUserID)
	return intUserID
}
