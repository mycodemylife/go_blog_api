package main

/*
	作者:ZLOE
	博客:https//zhang18.top
	邮箱:1144620122@qq.com
	QQ群:929724129
	时间:2019-04-24
*/


import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "weblog/routers"
	_ "weblog/models"
)

func main() {
	/*
	 1.添加解决跨域请求问题
	 2.文件下载文件夹，具体看beego官方文档
	*/
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.SetStaticPath("/api.file/download", "download")
	beego.SetStaticPath("/api.file/images", "images")
	beego.Run()
}

