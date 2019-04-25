package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/gofrs/uuid"
	"time"
	"weblog/models"
)

/*
功能：新建标签
路由：/api.tag.new
代码：验证token后新建标签
*/
type NewTag struct {
	beego.Controller
}

func (c *NewTag) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true {
		name := c.GetString("name")
		o := orm.NewOrm()
		ta := models.Tag{Name: name}
		logs.Info(o.Read(&ta, "Name"))
		if ta.Id == 0 {
			tag := new(models.Tag)
			tag.Name = name
			tag.Time = time.Now()
			logs.Info(o.Insert(tag))
			c.Data["json"] = map[string]interface{}{"result": "ok!", "data": "新建标签成功!"}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"result": "err!", "data": "标签已经存在，新建失败!"}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "你未登录，新建失败!"}
		c.ServeJSON()
	}
}

/*
功能：更新标签
路由：/api.tag.update
代码：根据标签名查询更改标签
*/
type UpdateTag struct {
	beego.Controller
}

func (c *UpdateTag) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true {
		name := c.GetString("name")
		id, err := c.GetInt("id")
		if err != nil {
			logs.Info(err)
		}
		o := orm.NewOrm()
		tag := models.Tag{Id: id}
		logs.Info(o.Read(&tag, "Id"))
		tag.Name = name
		logs.Info(o.Update(&tag, "Name"))
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": "更新成功!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "你未登录，更新失败!"}
		c.ServeJSON()
	}
}

/*
功能：删除标签
路由：/api.tag.delete
代码：根据标签名字删除标签
*/
type DeleteTag struct {
	beego.Controller
}

func (c *DeleteTag) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true {
		name := c.GetString("name")
		o := orm.NewOrm()
		tag := models.Tag{Name: name}
		logs.Info(o.Delete(&tag, "Name"))
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": "删除成功!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "你未登录，删除失败!"}
		c.ServeJSON()
	}
}

/*
功能：新建文章
路由：/api.post.new
代码：验证登录，提交保存文章
*/
type NewPost struct {
	beego.Controller
}

func (c *NewPost) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true {
		o := orm.NewOrm()
		title := c.GetString("title")
		Tag := c.GetString("tag")
		Tid, err := c.GetInt("tid")
		if err != nil {
			logs.Info(err)
		}
		image := c.GetString("image")
		content := c.GetString("content")
		mark := c.GetString("markdown")
		po := models.Post{Title: title}
		logs.Info(o.Read(&po, "Title"))
		if po.Id == 0 {
			post := new(models.Post)
			uid, err := uuid.NewV4()
			if err != nil {
				logs.Info(err)
			}
			post.Uid = uid.String()
			post.Title = title
			post.Tid = Tid
			post.Tag = Tag
			post.Views = 0
			post.Thumbs = 0
			post.Image = image
			post.Markdown=mark
			post.Content = content
			post.Time = time.Now()
			logs.Info(o.Insert(post))
			c.Data["json"] = map[string]interface{}{"result": "ok!", "data": "新建文章成功!"}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"result": "err!", "data": "文章已经存在，新建失败!"}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "你未登录，新建失败!"}
		c.ServeJSON()
	}
}

/*
功能：更新文章
路由：/api.post.update
代码：根据文章的uuid查询并且更新文章
*/
type UpdatePost struct {
	beego.Controller
}

func (c *UpdatePost) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true {
		id := c.GetString("id")
		o := orm.NewOrm()
		post := models.Post{Uid: id}
		logs.Info(o.Read(&post, "Uid"))
		title := c.GetString("title")
		views,err:=c.GetInt("views")
		Tag := c.GetString("tag")
		Tid,err := c.GetInt("tid")
		if err!=nil{
			logs.Info(err)
		}
		image := c.GetString("image")
		content := c.GetString("content")
		mark:=c.GetString("markdown")
		post.Title = title
		post.Tag = Tag
		post.Tid = Tid
		post.Views=views
		post.Image = image
		post.Markdown=mark
		post.Content = content
		logs.Info(o.Update(&post))
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": "更新成功!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "你未登录，更新失败!"}
		c.ServeJSON()
	}

}

/*
功能：删除文章
路由：/api.post.delete
代码：根据标签名字删除文章
*/
type DeletePost struct {
	beego.Controller
}

func (c *DeletePost) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true {
		id := c.GetString("id")
		o := orm.NewOrm()
		post := models.Post{Uid: id}
		logs.Info(o.Read(&post, "Uid"))
		var comments []*models.Comment
		number, err := o.QueryTable("comment").OrderBy("-Id").Filter("post_id", post.Id).All(&comments)
		if err != nil {
			logs.Info(number)
		}
		for _,comment := range comments {
			numbers, err := o.QueryTable("comment_reply").Filter("comment_id", comment.Id).Delete()
			if err != nil {
				logs.Info(numbers)
			}
		}
		num, err := o.QueryTable("comment").OrderBy("-Id").Filter("post_id", post.Id).Delete()
		if err != nil {
			logs.Info(num)
		}
		logs.Info(o.Delete(&post, "Uid"))
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": "删除成功!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "你未登录，删除失败!"}
		c.ServeJSON()
	}
}

/*
功能：新建文章评论
路由：/api.comment.new
代码：验证登录添加评论
*/
type NewComment struct {
	beego.Controller
}

func (c *NewComment) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true {
		username := c.GetString("username")
		userimage := c.GetString("userimage")
		body := c.GetString("body")
		post, err := c.GetInt("post")
		if err != nil {
			logs.Info(err)
		}
		id, err := uuid.NewV4()
		if err != nil {
			logs.Info(err)
		}
		o := orm.NewOrm()
		comment := new(models.Comment)
		comment.Uid = id.String()
		comment.UserName = username
		comment.UserImage = userimage
		comment.PostId = post
		comment.Body = body
		comment.Time = time.Now()
		logs.Info(o.Insert(comment))
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": "评论成功!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "你未登录，评论失败!"}
		c.ServeJSON()
	}
}

/*
功能：删除评论
路由：/api.comment.delete
代码：根据评论的UUID，删除评论
*/
type DeleteComment struct {
	beego.Controller
}

func (c *DeleteComment) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true {
		id := c.GetString("id")
		comment := models.Comment{Uid: id}
		o := orm.NewOrm()
		logs.Info(o.Delete(&comment, "Uid"))
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": "删除评论成功!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "你未登录，删除评论失败!"}
		c.ServeJSON()
	}

}

/*
功能：新建评论回复
路由：/api.replay.new
代码：验证用户，新建评论的回复
*/
type NewCommentReply struct {
	beego.Controller
}

func (c *NewCommentReply) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true {
		username := c.GetString("username")
		userimage := c.GetString("userimage")
		body := c.GetString("body")
		comment, err := c.GetInt("comment")
		if err != nil {
			logs.Info(err)
		}
		id, err := uuid.NewV4()
		if err != nil {
			logs.Info(err)
		}
		o := orm.NewOrm()
		reply := new(models.CommentReply)
		reply.Uid = id.String()
		reply.UserName = username
		reply.UserImage = userimage
		reply.CommentId = comment
		reply.Body = body
		reply.Time = time.Now()
		logs.Info(o.Insert(reply))
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": "回复成功!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "你未登录，回复失败!"}
		c.ServeJSON()
	}
}

/*
功能：删除评论的回复
路由：/api.replay.delete
代码：删除某条评论的回复
*/
type DeleteCommentReply struct {
	beego.Controller
}

func (c *DeleteCommentReply) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true {
		id := c.GetString("id")
		reply := models.CommentReply{Uid: id}
		o := orm.NewOrm()
		logs.Info(o.Delete(&reply, "Uid"))
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": "删除回复成功!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "你未登录，删除回复失败!"}
		c.ServeJSON()
	}
}

/*
功能：新建友情链接
路由：/api.link.new
代码：新建友情链接
*/
type NewLink struct {
	beego.Controller
}

func (c *NewLink) Post() {
	o := orm.NewOrm()
	email := c.GetString("email")
	li := models.Links{Email: email}
	logs.Info(o.Read(&li, "Email"))
	if li.Id == 0 {
		name := c.GetString("hostname")
		image := c.GetString("hostimage")
		url := c.GetString("hosturl")
		body := c.GetString("body")
		link := new(models.Links)
		uid, err := uuid.NewV4()
		if err != nil {
			logs.Info(err)
		}
		link.Uid = uid.String()
		link.Is = 0
		link.Email = email
		link.HostName = name
		link.HostImage = image
		link.HostUrl = url
		link.Body = body
		link.Time = time.Now()
		logs.Info(o.Insert(link))
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": "新建友链成功！"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "您已经创建过了，请等待审核！"}
		c.ServeJSON()
	}
}

/*
功能：更新友情链接
路由：/api.link.update
代码：更新友情链接，根据UUID。
*/
type UpdateLink struct {
	beego.Controller
}

func (c *UpdateLink) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)

	if result == true {
		id := c.GetString("id")
		name := c.GetString("hostname")
		image := c.GetString("hostimage")
		url := c.GetString("hosturl")
		body := c.GetString("body")
		is, err := c.GetInt("is")
		if err != nil {
			logs.Info(err)
		}
		o := orm.NewOrm()
		link := models.Links{Uid: id}
		logs.Info(o.Read(&link, "Uid"))
		link.HostName = name
		link.Is = is
		link.HostImage = image
		link.HostUrl = url
		link.Body = body
		logs.Info(o.Update(&link))
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": "更新友链成功!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "你未登录，更新友链失败!"}
		c.ServeJSON()
	}

}

/*
功能：删除友情链接
路由：/api.link.delete
代码：根据UUID删除友情链接
*/
type DeleteLink struct {
	beego.Controller
}

func (c *DeleteLink) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true {
		id := c.GetString("id")
		o := orm.NewOrm()
		link := models.Links{Uid: id}
		logs.Info(o.Delete(&link, "Uid"))
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": "删除友链成功!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "你未登录，删除友链失败!"}
		c.ServeJSON()
	}
}
