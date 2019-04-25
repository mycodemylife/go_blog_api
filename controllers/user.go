package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
	"weblog/models"
)

/*
功能：用户注册
路由：/api.user.register
代码：获取邮箱，查询是否已经注册过，加密密码存入数据库
*/
type NewUser struct {
	beego.Controller
}

func (c *NewUser) Post() {
	username := c.GetString("username")
	email := c.GetString("email")
	word := c.GetString("password")
	AvatarUrl := c.GetString("avatar")
	password, err := models.Encrypt(word)
	if err != nil {
		logs.Info(err)
	}
	o := orm.NewOrm()
	us := models.User{Email: email}
	logs.Info(o.Read(&us,"Email"))
	if us.Id==0 {
		user := new(models.User)
		user.Uid = models.Uid()
		user.Username = username
		user.Password = password
		user.Email = email
		user.AvatarUrl = AvatarUrl
		user.Time = time.Now()
		logs.Info(o.Insert(user))
		c.Data["json"] = map[string]interface{}{"result": "ok!","data":"注册成功!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "该邮箱已注册!"}
		c.ServeJSON()
	}
}


/*
功能：用户登录
路由：/api.user.login
代码：获取邮箱，查询用户的密码，与获取的密码加密后对比
*/
type UserLogin struct {
	beego.Controller
}


func (c *UserLogin) Post() {
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
		if user.Password == password {
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

/*
功能：验证用户，退出登陆
路由：/api.user.logout
代码：获取邮箱，查询用户的密码，与获取的密码加密后对比
*/
type UserLogout struct {
	beego.Controller
}

func (c *UserLogout) Post() {
	uid := c.GetString("uid")
	result := models.DeToken(uid)
	if result == "ok!" {
		c.Data["json"] = map[string]interface{}{"result": "ok!"}
		c.ServeJSON()
	}else {
		c.Data["json"] = map[string]interface{}{"result": "err!","data":"用户未登录!"}
		c.ServeJSON()
	}
}
/*
功能：获取全部用户信息
路由：/api.user.list
代码：验证管理员，返回用户数据
*/
type UserList struct {
	beego.Controller
}

func (c *UserList) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true{
		o:=orm.NewOrm()
		user := models.User{Uid:uid}
		logs.Info(o.Read(&user,"Uid"))
		if user.IsAdmin==1{
			var list []*models.User
			num, err := o.QueryTable("user").All(&list)
			if err != nil {
				logs.Info(num)
			}
			c.Data["json"] = map[string]interface{}{"result":"ok!","data":list}
			c.ServeJSON()
		}else {
			c.Data["json"] = map[string]interface{}{"result":"err!","data":"对不起，你无权访问!"}
			c.ServeJSON()
		}
	}else {
		c.Data["json"] = map[string]interface{}{"result":"err!","data":"用户未登录"}
		c.ServeJSON()
	}

}

/*
功能：获取单个用户所有信息
路由：/api.user.data
代码：验证管理员身份，是则返回数据，否则返回异常。
*/
type UserData struct {
	beego.Controller
}

func (c *UserData) Post()  {
	uid:= c.GetString("uid")
	token:=c.GetString("token")
	id:= c.GetString("id")
	result:=models.GetToken(uid,token)
	if result == true{
		o:=orm.NewOrm()
		user := models.User{Uid:uid}
		logs.Info(o.Read(&user,"Uid"))
		if user.IsAdmin==1{
			us:=models.User{Uid:id}
			logs.Info(o.Read(&us,"Uid"))
			if us.Id!=0{
				c.Data["json"] = map[string]interface{}{"result":"ok!","data":us}
				c.ServeJSON()
			}else {
				c.Data["json"] = map[string]interface{}{"result":"err!","data":"没有查询到该用户!"}
				c.ServeJSON()
			}
		}else {
			c.Data["json"] = map[string]interface{}{"result":"err!","data":"对不起，你无权访问!"}
			c.ServeJSON()
		}
	}else {
		c.Data["json"] = map[string]interface{}{"result":"err!","data":"用户未登录"}
		c.ServeJSON()
	}
}
/*
功能：删除用户
*/
type UserDelete struct {
	beego.Controller
}

func (c *UserDelete) Post()  {
	uid:= c.GetString("uid")
	token:=c.GetString("token")
	id:= c.GetString("id")
	result:=models.GetToken(uid,token)
	if result == true{
		o:=orm.NewOrm()
		user := models.User{Uid:uid}
		logs.Info(o.Read(&user,"Uid"))
		if user.IsAdmin==1{
			us:=models.User{Uid:id}
			logs.Info(o.Delete(&us,"Uid"))
			c.Data["json"] = map[string]interface{}{"result":"ok!","data":"删除用户成功!"}
			c.ServeJSON()
		}else {
			c.Data["json"] = map[string]interface{}{"result":"err!","data":"对不起，你无权访问!"}
			c.ServeJSON()
		}
	}else {
		c.Data["json"] = map[string]interface{}{"result":"err!","data":"用户未登录"}
		c.ServeJSON()
	}
}

type UserHeader struct {
	beego.Controller
}

func (c *UserHeader) Post() {
	uid:=c.GetString("uid")
	o:=orm.NewOrm()
	user := models.User{Uid:uid}
	logs.Info(o.Read(&user,"Uid"))
	c.Data["json"] = map[string]interface{}{"result":"ok!","data":user.AvatarUrl}
	c.ServeJSON()

}