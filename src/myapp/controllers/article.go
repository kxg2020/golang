package controllers

import (
	"github.com/astaxie/beego"
	//"myapp/models"
	"myapp/models"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController)Edit()  {
	id := this.GetString(":id")
	iid,_ := strconv.Atoi(id)
	// 修改数据
	var arc models.Article

	err := arc.Edit(iid)

	if err != nil{
		return
	}

}

func (this *ArticleController)Delete()  {

	id := this.GetString(":id")

	iid,_ := strconv.Atoi(id)

	var arc models.Article

	e := arc.Delete(iid)

	if e != nil{

		return
	}
	this.Redirect("/",302)
}

