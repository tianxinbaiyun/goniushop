package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

type RegionController struct {
	beego.Controller
}

func (this *RegionController) Region_Info() {

	regionId := this.GetString("regionId")
	intregionid := utils.String2Int(regionId)

	o := orm.NewOrm()
	regiontable := new(models.SysCity)
	var region models.SysCity
	o.QueryTable(regiontable).Filter("city_id", intregionid).One(&region)

	utils.ReturnHTTPSuccess(&this.Controller, region)
	this.ServeJSON()

}

func (this *RegionController) Region_List() {

	parentId := this.GetString("parentId")
	intparentid := utils.String2Int(parentId)

	o := orm.NewOrm()
	regiontable := new(models.SysCity)
	var regions []models.SysCity
	o.QueryTable(regiontable).Filter("province_id", intparentid).All(&regions)

	utils.ReturnHTTPSuccess(&this.Controller, regions)
	this.ServeJSON()
}
