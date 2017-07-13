package main

import (
	// 下划线表示只引用方法
	_ "myapp/initial"
	_ "myapp/routers"

	"github.com/astaxie/beego"
)

func main() {

	// 设置静态文件路径
	beego.SetStaticPath("/static", "static")

	beego.Run()
}
