package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"     //引入beego的orm
	_ "github.com/go-sql-driver/mysql" //引入beego的mysql驱动
)

// user表的结构

type User struct {
	Id         int
	Username   string
	Password   string
	CreateTime int64
}

// 初始化数据库

func init() {
	dbprefix := beego.AppConfig.String("mysqlprefix")
	// 注册数据库模型
	orm.RegisterModelWithPrefix(dbprefix, new(User))
}

// 数据库模型结构的方法
func (this *User) Check(username, password string) ([]User, error) {

	//根据用户名和密码查询
	o := orm.NewOrm()
	var data []User
	err := o.QueryTable("xm_user").Filter("username", username).Filter("password", password).One(&data)
	return data, err
}
