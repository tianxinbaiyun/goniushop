package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tianxinbaiyun/goniushop/utils"
)

// 变量定义
var (
	key         = []byte("adfadf!@#2")
	expireTime  = 20
	LoginUserID string
)

// CustomClaims CustomClaims
type CustomClaims struct {
	UserID string `json:"userid"`
	jwt.StandardClaims
}

// GetUserID GetUserID
func GetUserID(tokenStr string) string {

	token := Parse(tokenStr)
	if token == nil {
		return ""
	}
	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims.UserID
	}
	return ""
}

// Parse Parse
func Parse(tokenStr string) *jwt.Token {

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil
	}
	if token.Valid {
		return token
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("The token is expired or not valid.")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
	return nil

}

// Create Create
func Create(userid string) string {

	claims := CustomClaims{
		userid, jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(expireTime)).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(key)

	if err == nil {
		return tokenStr
	}
	return ""
}

// Verify Verify
func Verify(tokenStr string) bool {

	token := Parse(tokenStr)
	return token != nil

}

// getControllerAndAction getControllerAndAction
func getControllerAndAction(rawvalue string) (controller, action string) {
	values := strings.Split(rawvalue, "/")
	return values[2], values[2] + "/" + values[3]
}

// FilterFunc FilterFunc
func FilterFunc(ctx *context.Context) {

	controller, action := getControllerAndAction(ctx.Request.RequestURI)
	token := ctx.Input.Header("X-GoNiushop-Token")

	if action == "auth/loginByWeixin" {
		return
	}

	if token == "" {
		data := utils.GetHTTPRtnJSONData(401, "need relogin")
		ctx.Output.JSON(data, true, false)
		ctx.Redirect(200, "/")
		return
	}
	LoginUserID = GetUserID(token)

	publicControllerList := beego.AppConfig.String("controller::publicController")
	publicactionList := beego.AppConfig.String("action::publicAction")

	if !strings.Contains(publicControllerList, controller) && !strings.Contains(publicactionList, action) {
		if LoginUserID == "" {
			data := utils.GetHTTPRtnJSONData(401, "need relogin")
			ctx.Output.JSON(data, true, false)
			ctx.Redirect(200, "/")
			//http.Redirect(ctx.ResponseWriter, ctx.Request, "/", http.StatusMovedPermanently)
		}
	}
}
