package models

import (

	"github.com/astaxie/beego/orm"     //引入beego的orm
	_ "github.com/go-sql-driver/mysql" //引入beego的mysql驱动
	"time"
)

func (this *Article)Edit(id int) error {

	o:= orm.NewOrm()
	_,err := o.QueryTable("xm_article").Filter("id",id).Update(orm.Params{
		"update_time":time.Now().Unix(),
	})

	return err

}

func (this *Article)Delete(id int) (error) {

	o := orm.NewOrm()

	_,err := o.QueryTable("xm_article").Filter("id",id).Delete()

	return  err
}