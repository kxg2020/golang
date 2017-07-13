package initial

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm" //引入beego的orm
	_ "github.com/go-sql-driver/mysql"
)

func initSql() {
	dbhost := beego.AppConfig.String("mysqlhost")
	dbport := beego.AppConfig.String("mysqlport")
	dbname := beego.AppConfig.String("mysqldbname")
	dbuser := beego.AppConfig.String("mysqluser")
	dbpass := beego.AppConfig.String("mysqlpass")
	timezone := beego.AppConfig.String("mysqltimezone")

	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpass + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn, 5, 30)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
