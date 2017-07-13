package controllers

// // 使用原生的Go
// import (
// 	"fmt"
// 	"net/http"
// 	"regexp"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// // 创建空结构体

// type Example struct {
// }

// // 创建http服务

// func main() {

// }

// // 结构体的http服务方法
// func (this *Example) ServeHttp(w http.ResponseWriter, r *http.Request) {
// 	// 设置访问的路由
// 	if r.URL.Path == "/" {
// 		fmt.Print("访问的是根目录")
// 	}
// 	if r.URL.Path == "/login" {
// 		fmt.Print("访问的是登录")
// 	}

// 	// 404界面
// 	http.NotFound(w, r)
// 	return
// }

// // 设置方法
// func Hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Print("Hi")
// }

// // 设置方法
// func Login(w http.ResponseWriter, r *http.Request) {
// 	// 请求方式
// 	if r.Method == "GET" {
// 		t, _ := template.Parse("login.tpl")
// 		log.Println(t.Execute(r, nil))
// 	} else {
// 		if len(r.FormValue("username") == 0) {
// 			fmt.Fprint("用户名不能为空")
// 			fmt.Fprint(w, "\npassword", r.FormValue("password"))
// 			return
// 		}
// 	}

// 	// 判断格式
// 	if m, _ := regexp.MatchString("^[0-9]+$", req.Form.Get("age")); !m {
// 		return false
// 	}
// }
