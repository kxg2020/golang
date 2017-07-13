package controllers

import (
	"myapp/models"

	"github.com/astaxie/beego"
)

// 用变量保存模型对象
var aru models.User

// 登录控制器
type LoginController struct {
	
	beego.Controller
}

// 控制器方法


func (this *LoginController) Index() {

	// 展示登录界面
	this.TplName = "login/index.tpl"
	this.Render()

}

func (this *LoginController) Check() {

	// 获取post或put参数的集合
	username := this.GetString("username")
	password := this.GetString("password")

	// 查询数据库
	_, err := aru.Check(username, password)
	if err != nil {
		// 没有查找到用户
		this.Data["json"] = map[string]interface{}{"status": 0, "message": "登录失败"}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"status": 1, "message": "登录成功"}
	beego.Debug(username)
}
