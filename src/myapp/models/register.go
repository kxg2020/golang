package models

import (
	"github.com/astaxie/beego/orm" //引入beego的orm
	_ "github.com/go-sql-driver/mysql"
)

// 添加用户
func (this *User) Insert(user *User) (int64, error) {

	// orm对象
	o := orm.NewOrm()
	// 插入数据
	id, err := o.Insert(user)
	return id, err
}
