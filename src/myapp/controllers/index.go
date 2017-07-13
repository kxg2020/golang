package controllers

import (
	"myapp/models"
	"myapp/utils"
	"time"

	"github.com/astaxie/beego"
)

var arc models.Article

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Name"] = "XX"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["TrueCond"] = false
	this.TplName = "index.tpl"
	//	type User struct {
	//		Name string
	//		Sex  string
	//		Age  int
	//	}
	//	user := &User{Name: "XxX", Sex: "男", Age: 20}
	//	this.Data["user"] = user

	//	nums := [...]int{11, 22, 3, 4, 5, 6, 7}
	//	this.Data["nums"] = nums

	//	// 定义一个集合
	//	mape := make(map[string]string)
	//	mape["name"] = "小明"
	//	this.Data["student"] = mape

	//	this.Data["html"] = "<div><h1>这是一段html代码</h1></div>"
	//	this.Data["pipe"] = "<div><h1>这还是一段html代码</h1></div>"
}

func (this *MainController) Index() {

	//>>  查询数据库
	data, _ := arc.Index()

	this.Data["article"] = data

	//>> 截取字符串
	for k, _ := range data {

		data[k].Content = utils.SubString(data[k].Content, 0, 80) + "..."
	}

	this.TplName = "index/index.tpl"
	this.Render()

}

//func (this *MainController) Select() {
//	data, _ := arc.Select()
//	beego.Debug(data[0].Id)
//}

//func (this *MainController) Add() {
//	title := "测试标"
//	body := "测试内"
//	typeid := 1
//	id, err := arc.Add(title, body, typeid)
//	fmt.Println(id, err)
//}

//func (this *MainController) Save() {
//	title := "测试标题"
//	body := "测试内容"
//	typeid := 1
//	id := 1
//	err := arc.Save(title, body, typeid, id)
//	fmt.Println(err)
///}

func (this *MainController) Update() {

	updatetime := time.Now().Unix()
	// 获取参数,类型为string
	id, _ := this.GetInt(":id")

	err := arc.Update(id, updatetime)

	if err != nil {

		beego.Debug(err)
		return
	}
	this.TplName = "index/index.tpl"
}

//func (this *MainController) Delete() {

//	id, _ := this.GetInt(":id")
//	err := arc.Delete(id)

//	beego.Debug(err)

//}
