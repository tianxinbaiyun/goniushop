package models

import (
	"github.com/astaxie/beego/orm"
)

func GetAreaName(regionid int) string {

	o := orm.NewOrm()
	regiontable := new(SysArea)
	var region SysArea
	o.QueryTable(regiontable).Filter("area_id", regionid).One(&region)

	return region.AreaName

}

func GetProvinceName(regionid int) string {

	o := orm.NewOrm()
	regiontable := new(SysProvince)
	var region SysProvince
	o.QueryTable(regiontable).Filter("province_id", regionid).One(&region)

	return region.ProvinceName

}
func GetCityName(regionid int) string {

	o := orm.NewOrm()
	regiontable := new(SysCity)
	var region SysCity
	o.QueryTable(regiontable).Filter("city_id", regionid).One(&region)

	return region.CityName

}
func GetDistrictName(regionid int) string {

	o := orm.NewOrm()
	regiontable := new(SysDistrict)
	var region SysDistrict
	o.QueryTable(regiontable).Filter("district_id", regionid).One(&region)

	return region.DistrictName

}
