package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/gofrs/uuid"
	"io"
	"time"
	"weblog/models"
)

type UploadImage struct {
	beego.Controller
}

/*
 功能：获取保存图片
 路由：/api.upload.image
 功能：接收上传的文件，计算哈希值作为文件名，防止产生重复文件；保存图片相关信息进入数据库
*/
func (c *UploadImage) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true {
		f, _, _ := c.GetFile("file")
		hash := sha1.New()
		io.Copy(hash, f)
		mid := fmt.Sprintf("%x", hash.Sum(nil))
		path := "images/" + mid + ".jpg"
		defer f.Close()
		c.SaveToFile("file", path)
		file := beego.AppConfig.String("host") + "/api.file/" + path
		id, err := uuid.NewV4()
		if err != nil {
			logs.Info(err)
		}
		o:=orm.NewOrm()
		fi := new(models.UploadImage)
		fi.UserId = uid
		fi.Uid = id.String()
		fi.Md = mid
		fi.Download = file
		fi.Time = time.Now()
		logs.Info(o.Insert(fi))
		c.Data["json"] = map[string]interface{}{"result": "ok!", "download": file}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "您未登录！上传失败！"}
		c.ServeJSON()
	}
}


/*
 功能：获取保存文件
 路由：/api.upload.file
 功能：接收上传的文件，保存文件，并把文件相关信息保存到数据库
*/
type UploadFile struct {
	beego.Controller
}

func (c *UploadFile) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true {
		f, h, _ := c.GetFile("file")
		path := "download/" + h.Filename
		defer f.Close()
		c.SaveToFile("file", path)
		file := beego.AppConfig.String("host") + "/api.file/" + path
		id, err := uuid.NewV4()
		if err != nil {
			logs.Info(err)
		}
		o:=orm.NewOrm()
		image := new(models.UploadFile)
		image.UserId = uid
		image.Uid = id.String()
		image.Md = h.Filename
		image.Download = file
		image.Time = time.Now()
		logs.Info(o.Insert(image))
		c.Data["json"] = map[string]interface{}{"result": "ok!", "download": file}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "您未登录！上传失败！"}
		c.ServeJSON()
	}
}
