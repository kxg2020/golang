package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"     //引入beego的orm
	_ "github.com/go-sql-driver/mysql" //引入beego的mysql驱动
)

//表go_archives的结构
type Article struct {
	Id         int
	Title      string
	Content    string
	UpdateTime int64
	CreateTime int64
	Author     string
}

// 初始化数据库链接
func init() {
	dbprefix := beego.AppConfig.String("mysqlprefix")
	orm.RegisterModelWithPrefix(dbprefix, new(Article))
}

func (this *Article) Index() ([]Article, error) {
	o := orm.NewOrm()
	var data []Article
	_, err := o.QueryTable("xm_article").All(&data)
	return data, err

}

//表go_archives的增加
//func (this *Article) Insert(title, content string, updatetime int, creatime int) (int64, error) {
//	o := orm.NewOrm()
//	arc := Archives{Title: title, Body: body, Typeid: typeid}
//	id, err := o.Insert(&arc)
//	return id, err
//}

func (this *Article) Select(id int) ([]Article, error) {
	o := orm.NewOrm()
	var data []Article
	_, err := o.QueryTable("go_archives").Filter("id", id).All(&data)
	return data, err

}

//func (this *Article) Save(Title string, Body: body, Typeid: typeid, Id: id) error {
//	o := orm.NewOrm()
//	arc := Article{Title: title, Body: body, Typeid: typeid, Id: id}
//	_, err := o.Update(&arc)
//	return err
//}

func (this *Article) Update(id int, updatetime int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("xm_article").Filter("id", id).Update(orm.Params{
		"update_time": updatetime,
	})
	return err
}

//func (this *Article) Delete(id int) error {
//	o := orm.NewOrm()
//	arc := Archives{Id: id}
//	_, err := o.Delete(&arc)
//	return err
//}
