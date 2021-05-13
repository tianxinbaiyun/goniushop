package models

import (
	"github.com/astaxie/beego/orm"
)

// GetAreaName GetAreaName
func GetAreaName(regionID int) string {

	o := orm.NewOrm()
	regionTable := new(SysArea)
	var region SysArea
	o.QueryTable(regionTable).Filter("area_id", regionID).One(&region)

	return region.AreaName

}

// GetProvinceName GetProvinceName
func GetProvinceName(regionID int) string {

	o := orm.NewOrm()
	regionTable := new(SysProvince)
	var region SysProvince
	o.QueryTable(regionTable).Filter("province_id", regionID).One(&region)

	return region.ProvinceName

}

// GetCityName GetCityName
func GetCityName(regionID int) string {

	o := orm.NewOrm()
	regionTable := new(SysCity)
	var region SysCity
	o.QueryTable(regionTable).Filter("city_id", regionID).One(&region)

	return region.CityName

}

// GetDistrictName GetDistrictName
func GetDistrictName(regionID int) string {

	o := orm.NewOrm()
	regionTable := new(SysDistrict)
	var region SysDistrict
	o.QueryTable(regionTable).Filter("district_id", regionID).One(&region)

	return region.DistrictName

}
