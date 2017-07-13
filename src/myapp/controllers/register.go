package controllers

import (
	"myapp/models"
	"regexp"
	"time"

	"github.com/astaxie/beego"
)

// 创建控制器
type RegisterController struct {
	beego.Controller
}

// 注册界面
func (this *RegisterController) Index() {
	// 注册界面
	this.TplName = "register/index.tpl"
	// 渲染页面
	this.Render()
}

// 注册方法
func (this *RegisterController) Insert() {

	// 获取参数
	username := this.GetString("username")
	password := this.GetString("password")

	//检测用户
	r := checkUser(username, password)

	if !r {

		return
	}
	// 通过入库
	// 注意：这里要使用取地址符,从底层改变User,否则models.User{}会被当成一种数据类型,即一个副本
	user := &models.User{
		Username:   username,
		Password:   password,
		CreateTime: time.Now().Unix(),
	}

	_, err := user.Insert(user)
	if err != nil {

		return
	}

	// 成功跳转
	this.Redirect("/", 302)

}

func checkUser(username, password string) bool {

	u, _ := regexp.MatchString("^[1-9a-z][a-z0-9_]*", username)

	if !u {

		return false
	}

	p, _ := regexp.MatchString("^.{6,10}$", password)

	if !p {

		return false
	}

	return true

}
