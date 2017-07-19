package controllers

import (
	"github.com/astaxie/beego"
)

// 控制器
type PhotoController struct {
	beego.Controller
}

//控制器方法
func (this *PhotoController)Index()  {

	this.TplName="photo/index.html"
	this.Render()
}