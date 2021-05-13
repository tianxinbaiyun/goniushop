package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

// AddressController 地址控制器
type AddressController struct {
	beego.Controller
}

// AddressListRtnJSON 地址列表返回结构体
type AddressListRtnJSON struct {
	models.NsMemberExpressAddress
	ProvinceName string `json:"provice_name"`
	CityName     string `json:"city_name"`
	DistrictName string `json:"district_name"`
	FullRegion   string `json:"full_region"`
}

// AddressList 地址列表
func (c *AddressController) AddressList() {
	var err error
	defer func() {
		if err != nil {
			logs.Debug("c has a err ,err:%v", err)
		}
		c.ServeJSON()
		return
	}()
	o := orm.NewOrm()
	addressTable := new(models.NsMemberExpressAddress)
	var addresses []models.NsMemberExpressAddress

	_, err = o.QueryTable(addressTable).Filter("uid", getLoginUserID()).All(&addresses)
	if err != nil {
		logs.Debug("get address error,err:%v", err)
		utils.ReturnHTTPError(&c.Controller, 400, "地址不存在")
		return
	}
	rtnAddress := make([]AddressListRtnJSON, 0)

	for _, val := range addresses {

		provinceName := models.GetProvinceName(val.ProvinceID)
		cityName := models.GetCityName(val.CityID)
		distinctName := models.GetDistrictName(val.DistrictID)
		rtnAddress = append(rtnAddress, AddressListRtnJSON{
			NsMemberExpressAddress: val,
			ProvinceName:           provinceName,
			CityName:               cityName,
			DistrictName:           distinctName,
			FullRegion:             provinceName + cityName + distinctName,
		})

	}

	utils.ReturnHTTPSuccess(&c.Controller, rtnAddress)
	return
}

// AddressDetail AddressDetail
func (c *AddressController) AddressDetail() {
	//请求参数
	id := c.GetString("id")
	intID := utils.String2Int(id)
	//错误返回
	var err error
	defer func() {
		if err != nil {
			logs.Debug("c address id[%v] has a error ,err:%v", id, err)
		}
		c.ServeJSON()
		return
	}()
	//查询地址
	o := orm.NewOrm()
	addressTable := new(models.NsMemberExpressAddress)
	var address models.NsMemberExpressAddress

	err = o.QueryTable(addressTable).Filter("id", intID).Filter("uid", getLoginUserID()).One(&address)

	var val AddressListRtnJSON

	if err != orm.ErrNoRows {
		provinceName := models.GetProvinceName(address.ProvinceID)
		cityName := models.GetCityName(address.CityID)
		distinctName := models.GetDistrictName(address.DistrictID)
		val = AddressListRtnJSON{
			NsMemberExpressAddress: address,
			ProvinceName:           provinceName,
			CityName:               cityName,
			DistrictName:           distinctName,
			FullRegion:             provinceName + cityName + distinctName,
		}
	}
	utils.ReturnHTTPSuccess(&c.Controller, val)
	// c.ServeJSON()
}

// AddressSaveBody 地址保存结构体
type AddressSaveBody struct {
	Address    string `json:"address"`
	CityID     int    `json:"city_id"`
	DistrictID int    `json:"district_id"`
	IsDefault  bool   `json:"is_default"`
	Mobile     string `json:"mobile"`
	Name       string `json:"name"`
	ProvinceID int    `json:"province_id"`
	AddressID  int    `json:"address_id"`
}

// AddressSave AddressSave
func (c *AddressController) AddressSave() {

	var asb AddressSaveBody
	body := c.Ctx.Input.RequestBody
	_ = json.Unmarshal(body, &asb)

	address := asb.Address
	name := asb.Name
	mobile := asb.Mobile
	ProvinceID := asb.ProvinceID
	CityID := asb.CityID
	distinctID := asb.DistrictID
	isDefault := asb.IsDefault
	addressID := asb.AddressID
	userID := getLoginUserID()
	var intIsDefault int
	if isDefault {
		intIsDefault = 1
	} else {
		intIsDefault = 0
	}

	intCityID := CityID
	intProvinceID := ProvinceID
	intDistinctID := distinctID

	addressData := models.NsMemberExpressAddress{
		Address:    address,
		CityID:     intCityID,
		DistrictID: intDistinctID,
		ProvinceID: intProvinceID,
		Name:       name,
		Mobile:     mobile,
		UserID:     userID,
		IsDefault:  intIsDefault,
	}
	o := orm.NewOrm()
	addressTable := new(models.NsMemberExpressAddress)

	var intID int64
	if addressID == 0 {
		id, err := o.Insert(&addressData)
		if err == nil {
			intID = id
		}
	} else {
		_, _ = o.QueryTable(addressTable).Filter("id", intID).Filter("uid", userID).Update(orm.Params{
			"is_default": 0,
		})
	}

	if isDefault {
		_, err := o.Raw("UPDATE ns_member_express_address SET is_default = 0 where id <> ? and uid = ?", intID, userID).Exec()
		if err == nil {
			//res.RowsAffected()
			//fmt.Println("mysql row affected nums: ", num)
		}
	}
	var addressInfo models.NsMemberExpressAddress
	_ = o.QueryTable(addressTable).Filter("id", intID).One(&addressInfo)

	utils.ReturnHTTPSuccess(&c.Controller, addressInfo)
	c.ServeJSON()

}

// AddressDelete AddressDelete
func (c *AddressController) AddressDelete() {

	addressID := c.GetString("id")
	intAddressID := utils.String2Int(addressID)

	o := orm.NewOrm()
	addressTable := new(models.NsMemberExpressAddress)
	_, _ = o.QueryTable(addressTable).Filter("id", intAddressID).Filter("uid", getLoginUserID()).Delete()

	return

}
