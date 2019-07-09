package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/services"
	"github.com/tianxinbaiyun/goniushop/utils"
)

type AuthController struct {
	beego.Controller
}

type AuthLoginBody struct {
	Code     string               `json:"code"`
	UserInfo services.ResUserInfo `json:"userInfo"`
}

func (this *AuthController) Auth_LoginByWeixin() {

	var alb AuthLoginBody
	body := this.Ctx.Input.RequestBody

	err := json.Unmarshal(body, &alb)
	if err != nil {
		logs.Debug("json Unmarshal,alb:%v, error,err:%v", alb, err)
	}
	logs.Debug("json Unmarshal success,alb:%v", alb)
	clientIP := this.Ctx.Input.IP()
	nowTime := int(utils.GetTimestamp())
	userInfo := services.Login(alb.Code, alb.UserInfo)
	if userInfo == nil {
		logs.Debug("auth ,get user error，userInfo:%v,err:%v,", userInfo, err)
	}
	logs.Debug("Login success,Code:%v userInfo %v", alb.Code, userInfo)
	//开始事务
	rtnInfo := make(map[string]interface{})
	o := orm.NewOrm()
	o.Begin()
	defer func() {
		if err != nil {
			o.Rollback()
			logs.Debug("this has a err ,err:%v", err)
			utils.ReturnHTTPError(&this.Controller, 400, "授权失败")
			this.ServeJSON()
			return
		} else {
			err = o.Commit()
			utils.ReturnHTTPSuccess(&this.Controller, rtnInfo)
			this.ServeJSON()
		}
	}()

	var user models.SysUser
	usertable := new(models.SysUser)
	err = o.QueryTable(usertable).Filter("wx_openid", userInfo.OpenID).One(&user)
	if err == orm.ErrNoRows {
		initPasswd := []byte(userInfo.OpenID)
		passwd := md5.Sum(initPasswd)
		newuser := models.SysUser{
			UserName:         utils.GetUUID(),
			UserPassword:     fmt.Sprintf("%x", passwd),
			WxOpenid:         userInfo.OpenID,
			UserHeadimg:      userInfo.AvatarUrl,
			Sex:              int16(userInfo.Gender),
			LastLoginIp:      clientIP,
			LastLoginTime:    nowTime,
			LastLoginType:    1,
			RegTime:          nowTime,
			NickName:         userInfo.NickName,
			CurrentLoginIp:   clientIP,
			CurrentLoginTime: nowTime,
			CurrentLoginType: 1,
			IsMember:         1,
			UserStatus:       1,
			LoginNum:         1,
		}
		_, err = o.Insert(&newuser)
		if err != nil {
			logs.Debug("insert sysuser error,err:%v", err)
			return
		}
		err = o.QueryTable(usertable).Filter("wx_openid", userInfo.OpenID).One(&user)
		if err != nil {
			logs.Debug("insert sysuser success, fetch sysuser error,err:%v", err)
			return
		}
		member := models.NsMember{
			Uid:         user.Id,
			MemberName:  user.UserName,
			MemberLevel: 47,
			RegTime:     nowTime,
		}
		_, err = o.Insert(&member)
		if err != nil {
			logs.Debug("insert member error,err:%v", err)
			return
		}
	}

	userinfo := make(map[string]interface{})
	userinfo["id"] = user.Id
	userinfo["username"] = user.UserName
	userinfo["nickname"] = user.NickName
	userinfo["gender"] = user.Sex
	userinfo["avatar"] = user.UserHeadimg
	userinfo["birthday"] = user.Birthday

	user.LastLoginIp = user.CurrentLoginIp
	user.LastLoginTime = user.CurrentLoginTime
	user.LastLoginType = user.CurrentLoginType
	user.CurrentLoginIp = clientIP
	user.CurrentLoginTime = nowTime
	user.CurrentLoginType = 1
	user.LoginNum = user.LoginNum + 1
	if _, err := o.Update(&user); err == nil {

	}

	sessionKey := services.Create(utils.Int2String(user.Id))
	//fmt.Println("sessionkey==" + sessionKey)

	rtnInfo["token"] = sessionKey
	rtnInfo["userInfo"] = userinfo
	return
}
