package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

type AddressController struct {
	beego.Controller
}

type AddressListRtnJson struct {
	models.NsMemberExpressAddress
	ProviceName  string `json:"provice_name"`
	CityName     string `json:"city_name"`
	DistrictName string `json:"district_name"`
	FullRegion   string `json:"full_region"`
}

func (this *AddressController) Address_List() {
	var err error
	defer func() {
		if err != nil {
			logs.Debug("this has a err ,err:%v", err)
		}
		this.ServeJSON()
		return
	}()
	o := orm.NewOrm()
	addresstable := new(models.NsMemberExpressAddress)
	var addresses []models.NsMemberExpressAddress

	_, err = o.QueryTable(addresstable).Filter("uid", getLoginUserId()).All(&addresses)
	if err != nil {
		logs.Debug("get address error,err:%v", err)
		utils.ReturnHTTPError(&this.Controller, 400, "地址不存在")
		return
	}
	rtnaddress := make([]AddressListRtnJson, 0)

	for _, val := range addresses {

		provicename := models.GetProvinceName(val.ProvinceId)
		cityname := models.GetCityName(val.CityId)
		distinctname := models.GetDistrictName(val.DistrictId)
		rtnaddress = append(rtnaddress, AddressListRtnJson{
			NsMemberExpressAddress: val,
			ProviceName:            provicename,
			CityName:               cityname,
			DistrictName:           distinctname,
			FullRegion:             provicename + cityname + distinctname,
		})

	}

	utils.ReturnHTTPSuccess(&this.Controller, rtnaddress)
	return
}
func (this *AddressController) Address_Detail() {
	//请求参数
	id := this.GetString("id")
	intid := utils.String2Int(id)
	//错误返回
	var err error
	defer func() {
		if err != nil {
			logs.Debug("this address id[%v] has a error ,err:%v", id, err)
		}
		this.ServeJSON()
		return
	}()
	//查询地址
	o := orm.NewOrm()
	addresstable := new(models.NsMemberExpressAddress)
	var address models.NsMemberExpressAddress

	err = o.QueryTable(addresstable).Filter("id", intid).Filter("uid", getLoginUserId()).One(&address)

	var val AddressListRtnJson

	if err != orm.ErrNoRows {
		provicename := models.GetProvinceName(address.ProvinceId)
		cityname := models.GetCityName(address.CityId)
		distinctname := models.GetDistrictName(address.DistrictId)
		val = AddressListRtnJson{
			NsMemberExpressAddress: address,
			ProviceName:            provicename,
			CityName:               cityname,
			DistrictName:           distinctname,
			FullRegion:             provicename + cityname + distinctname,
		}
	}
	utils.ReturnHTTPSuccess(&this.Controller, val)
	// this.ServeJSON()
}

type AddressSaveBody struct {
	Address    string `json:"address"`
	CityId     int    `json:"city_id"`
	DistrictId int    `json:"district_id"`
	IsDefault  bool   `json:"is_default"`
	Mobile     string `json:"mobile"`
	Name       string `json:"name"`
	ProvinceId int    `json:"province_id"`
	AddressId  int    `json:"address_id"`
}

func (this *AddressController) Address_Save() {

	var asb AddressSaveBody
	body := this.Ctx.Input.RequestBody
	json.Unmarshal(body, &asb)

	address := asb.Address
	name := asb.Name
	mobile := asb.Mobile
	provinceid := asb.ProvinceId
	cityid := asb.CityId
	distinctid := asb.DistrictId
	isdefault := asb.IsDefault
	addressid := asb.AddressId
	userid := getLoginUserId()
	var intisdefault int
	if isdefault {
		intisdefault = 1
	} else {
		intisdefault = 0
	}

	intcityid := cityid
	intprovinceid := provinceid
	intdistinctid := distinctid

	addressdata := models.NsMemberExpressAddress{
		Address:    address,
		CityId:     intcityid,
		DistrictId: intdistinctid,
		ProvinceId: intprovinceid,
		Name:       name,
		Mobile:     mobile,
		UserId:     userid,
		IsDefault:  intisdefault,
	}
	o := orm.NewOrm()
	addresstable := new(models.NsMemberExpressAddress)

	var intid int64
	if addressid == 0 {
		id, err := o.Insert(&addressdata)
		if err == nil {
			intid = id
		}
	} else {
		o.QueryTable(addresstable).Filter("id", intid).Filter("uid", userid).Update(orm.Params{
			"is_default": 0,
		})
	}

	if isdefault {
		_, err := o.Raw("UPDATE ns_member_express_address SET is_default = 0 where id <> ? and uid = ?", intid, userid).Exec()
		if err == nil {
			//res.RowsAffected()
			//fmt.Println("mysql row affected nums: ", num)
		}
	}
	var addressinfo models.NsMemberExpressAddress
	o.QueryTable(addresstable).Filter("id", intid).One(&addressinfo)

	utils.ReturnHTTPSuccess(&this.Controller, addressinfo)
	this.ServeJSON()

}

func (this *AddressController) Address_Delete() {

	addressid := this.GetString("id")
	intaddressid := utils.String2Int(addressid)

	o := orm.NewOrm()
	addresstable := new(models.NsMemberExpressAddress)
	o.QueryTable(addresstable).Filter("id", intaddressid).Filter("uid", getLoginUserId()).Delete()

	return

}
