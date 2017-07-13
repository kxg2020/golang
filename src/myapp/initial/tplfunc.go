package initial

import (
	"myapp/utils"

	"github.com/astaxie/beego"
)

func initTplFunc() {
	beego.AddFuncMap("substring", utils.SubString)
	beego.AddFuncMap("date_mh", utils.GetDateMH)
	beego.AddFuncMap("avatar", utils.GetGravatar)
}
