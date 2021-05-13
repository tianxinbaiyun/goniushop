package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

// RegionController RegionController
type RegionController struct {
	beego.Controller
}

// RegionInfo RegionInfo
func (c *RegionController) RegionInfo() {

	regionID := c.GetString("regionId")
	intRegionID := utils.String2Int(regionID)

	o := orm.NewOrm()
	regionTable := new(models.SysCity)
	var region models.SysCity
	o.QueryTable(regionTable).Filter("city_id", intRegionID).One(&region)

	utils.ReturnHTTPSuccess(&c.Controller, region)
	c.ServeJSON()

}

// RegionList RegionList
func (c *RegionController) RegionList() {

	parentID := c.GetString("parentId")
	intParentID := utils.String2Int(parentID)

	o := orm.NewOrm()
	regionTable := new(models.SysCity)
	var regions []models.SysCity
	o.QueryTable(regionTable).Filter("province_id", intParentID).All(&regions)

	utils.ReturnHTTPSuccess(&c.Controller, regions)
	c.ServeJSON()
}
