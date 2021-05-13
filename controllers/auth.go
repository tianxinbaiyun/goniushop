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

// AuthController AuthController
type AuthController struct {
	beego.Controller
}

// AuthLoginBody AuthLoginBody
type AuthLoginBody struct {
	Code     string               `json:"code"`
	UserInfo services.ResUserInfo `json:"userInfo"`
}

// AuthLoginByWeixin AuthLoginByWeixin
func (c *AuthController) AuthLoginByWeixin() {

	var alb AuthLoginBody
	body := c.Ctx.Input.RequestBody

	err := json.Unmarshal(body, &alb)
	if err != nil {
		logs.Debug("json Unmarshal,alb:%v, error,err:%v", alb, err)
	}
	logs.Debug("json Unmarshal success,alb:%v", alb)
	clientIP := c.Ctx.Input.IP()
	nowTime := int(utils.GetTimestamp())
	userInfo := services.Login(alb.Code, alb.UserInfo)
	if userInfo == nil {
		logs.Debug("auth ,get user error，userInfo:%v,err:%v,", userInfo, err)
		return
	}
	logs.Debug("Login success,Code:%v userInfo %v", alb.Code, userInfo)
	//开始事务
	rtnInfo := make(map[string]interface{})
	o := orm.NewOrm()
	_ = o.Begin()
	defer func() {
		if err != nil {
			_ = o.Rollback()
			logs.Debug("this has a err ,err:%v", err)
			utils.ReturnHTTPError(&c.Controller, 400, "授权失败")
			c.ServeJSON()
			return
		}
		err = o.Commit()
		if err != nil {
			return
		}
		utils.ReturnHTTPSuccess(&c.Controller, rtnInfo)
		c.ServeJSON()
		return
	}()

	var user models.SysUser
	userTable := new(models.SysUser)
	err = o.QueryTable(userTable).Filter("wx_openid", userInfo.OpenID).One(&user)
	if err == orm.ErrNoRows {
		initPasswd := []byte(userInfo.OpenID)
		passwd := md5.Sum(initPasswd)
		newUser := models.SysUser{
			UserName:         utils.GetUUID(),
			UserPassword:     fmt.Sprintf("%x", passwd),
			WxOpenid:         userInfo.OpenID,
			UserHeadimg:      userInfo.AvatarURL,
			Sex:              int16(userInfo.Gender),
			LastLoginIP:      clientIP,
			LastLoginTime:    nowTime,
			LastLoginType:    1,
			RegTime:          nowTime,
			NickName:         userInfo.NickName,
			CurrentLoginIP:   clientIP,
			CurrentLoginTime: nowTime,
			CurrentLoginType: 1,
			IsMember:         1,
			UserStatus:       1,
			LoginNum:         1,
		}
		_, err = o.Insert(&newUser)
		if err != nil {
			logs.Debug("insert sysUser error,err:%v", err)
			return
		}
		err = o.QueryTable(userTable).Filter("wx_openid", userInfo.OpenID).One(&user)
		if err != nil {
			logs.Debug("insert sysUser success, fetch sysUser error,err:%v", err)
			return
		}
		member := models.NsMember{
			UID:         user.ID,
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
	userinfo["id"] = user.ID
	userinfo["username"] = user.UserName
	userinfo["nickname"] = user.NickName
	userinfo["gender"] = user.Sex
	userinfo["avatar"] = user.UserHeadimg
	userinfo["birthday"] = user.Birthday

	user.LastLoginIP = user.CurrentLoginIP
	user.LastLoginTime = user.CurrentLoginTime
	user.LastLoginType = user.CurrentLoginType
	user.CurrentLoginIP = clientIP
	user.CurrentLoginTime = nowTime
	user.CurrentLoginType = 1
	user.LoginNum = user.LoginNum + 1
	if _, err := o.Update(&user); err == nil {

	}

	sessionKey := services.Create(utils.Int2String(user.ID))
	//fmt.Println("sessionKey==" + sessionKey)

	rtnInfo["token"] = sessionKey
	rtnInfo["userInfo"] = userinfo
	return
}
