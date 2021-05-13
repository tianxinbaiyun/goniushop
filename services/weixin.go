package services

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/objcoding/wxpay"
	"github.com/tianxinbaiyun/goniushop/utils"
)

// WXLoginResponse 返回结构体
type WXLoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

//https://developers.weixin.qq.com/miniprogram/dev/api/wx.getUserInfo.html

// Watermark Watermark
type Watermark struct {
	AppID     string `json:"appid"`
	TimeStamp int64  `json:"timestamp"`
}

// WXUserInfo 微信用户
type WXUserInfo struct {
	OpenID    string    `json:"openId,omitempty"`
	NickName  string    `json:"nickName"`
	AvatarURL string    `json:"avatarUrl"`
	Gender    int       `json:"gender"`
	Country   string    `json:"country"`
	Province  string    `json:"province"`
	City      string    `json:"city"`
	UnionID   string    `json:"unionId,omitempty"`
	Language  string    `json:"language"`
	Watermark Watermark `json:"watermark,omitempty"`
}

// ResUserInfo 用户信息
type ResUserInfo struct {
	UserInfo      WXUserInfo `json:"userInfo"`
	RawData       string     `json:"rawData"`
	Signature     string     `json:"signature"`
	EncryptedData string     `json:"encryptedData"`
	IV            string     `json:"iv"`
}

// Login 登陆
func Login(code string, fullUserInfo ResUserInfo) *WXUserInfo {

	secret := beego.AppConfig.String("weixin::secret")
	appid := beego.AppConfig.String("weixin::appid")

	//https://developers.weixin.qq.com/miniprogram/dev/api-backend/auth.code2Session.html
	req := httplib.Get("https://api.weixin.qq.com/sns/jscode2session")
	req.Param("grant_type", "authorization_code")
	req.Param("js_code", code)
	req.Param("secret", secret)
	req.Param("appid", appid)

	var res WXLoginResponse
	err := req.ToJSON(&res)
	if err != nil {
		logs.Debug("WXLogin error,err:%v", err)
	}
	logs.Debug("WXLogin success,req:%v, res:%v", req, res)
	s := sha1.New()
	s.Write([]byte(fullUserInfo.RawData + res.SessionKey))
	sha1 := s.Sum(nil)
	sha1hash := hex.EncodeToString(sha1)

	logs.Debug("fullUserInfo:%v,\n,SessionKey:%v", fullUserInfo.RawData, res.SessionKey)
	logs.Debug("fullUserInfo.Signature:%v", fullUserInfo.Signature)
	logs.Debug("sha1hash:%v", sha1hash)

	if fullUserInfo.Signature != sha1hash {
		// return nil
	}
	userinfo := DecryptUserInfoData(res.SessionKey, fullUserInfo.EncryptedData, fullUserInfo.IV)
	logs.Debug("get weixin userIngo:%v", userinfo)
	return userinfo

}

// DecryptUserInfoData 数据解密
func DecryptUserInfoData(sessionKey string, encryptedData string, iv string) *WXUserInfo {

	sk, _ := base64.StdEncoding.DecodeString(sessionKey)
	ed, _ := base64.StdEncoding.DecodeString(encryptedData)
	i, _ := base64.StdEncoding.DecodeString(iv)

	decryptedData, err := utils.AesCBCDecrypt(ed, sk, i)

	if err != nil {
		return nil
	}

	var wxuserinfo WXUserInfo
	//fmt.Println(string(decryptedData))
	err = json.Unmarshal(decryptedData, &wxuserinfo)
	if err != nil {

	}
	return &wxuserinfo
}

// PayInfo 支付信息
type PayInfo struct {
	OpenID         string
	Body           string
	OutTradeNo     string
	TotalFee       int64
	SpbillCreateIP string
}

// CreateUnifiedOrder 创建订单
func CreateUnifiedOrder(payInfo PayInfo) (wxpay.Params, error) {

	appid := beego.AppConfig.String("weixin::appid")
	mchid := beego.AppConfig.String("weixin::mch_id")
	apikey := beego.AppConfig.String("weixin::apikey")
	notifyurl := beego.AppConfig.String("weixin::notify_url")
	account := wxpay.NewAccount(appid, mchid, apikey, false)
	client := wxpay.NewClient(account)
	params := make(wxpay.Params)
	params.SetString("body", payInfo.Body).
		SetString("out_trade_no", payInfo.OutTradeNo).
		SetInt64("total_fee", payInfo.TotalFee).
		SetString("spbill_create_ip", payInfo.SpbillCreateIP).
		SetString("notify_url", notifyurl).
		SetString("trade_type", "APP")
	return client.UnifiedOrder(params)

}
