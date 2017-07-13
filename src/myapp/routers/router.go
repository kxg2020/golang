package routers

import (
	"myapp/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/update/:id", &controllers.MainController{}, "get:Update")
	beego.Router("/login", &controllers.LoginController{}, "get:Index")
	beego.Router("/check", &controllers.LoginController{}, "post:Check")
	beego.Router("/register", &controllers.RegisterController{}, "*:Index")
	beego.Router("/register/insert", &controllers.RegisterController{}, "*:Insert")
}
