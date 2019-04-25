package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"weblog/models"
)

type AdminUserLogin struct {
	beego.Controller
}

func (c *AdminUserLogin) Post() {
	email := c.GetString("email")
	word := c.GetString("password")
	if len(email)!=0 || len(word)!=0{
		password, err := models.Encrypt(word)
		if err != nil {
			logs.Info(err)
		}
		o := orm.NewOrm()
		user := models.User{Email: email}
		logs.Info(o.Read(&user, "Email"))
		if user.Password == password && user.IsAdmin==1 {
			token:=models.SetToken(user.Uid)
			c.Data["json"] = map[string]interface{}{"result": "ok!","token":token,"uid":user.Uid,"username":user.Username,"AvatarUrl":user.AvatarUrl}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"result": "err!", "data": "密码错误!"}
			c.ServeJSON()
		}
	}else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "账号及密码不能为空!"}
		c.ServeJSON()
	}
}
