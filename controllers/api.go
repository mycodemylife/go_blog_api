package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
	"weblog/models"
)

/*
功能：获取所有标签
路由：/api.tag.list
代码：获取标签
*/
type TagListController struct {
	beego.Controller
}

func (c *TagListController) Get() {
	var tag []*models.Tag
	o := orm.NewOrm()
	number, _ := o.QueryTable("tag").All(&tag)
	logs.Info(number)
	c.Data["json"] = map[string]interface{}{"result": "ok!", "data": tag}
	c.ServeJSON()
}

/*
功能：获取单个标签
路由：/api.tag.id/:id([0-9]+)
代码：根据id获取单个标签
*/
type TagController struct {
	beego.Controller
}

func (c *TagController) Get() {
	str := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(str)
	if err != nil {
		logs.Info(err)
	}
	o := orm.NewOrm()
	tag := models.Tag{Id: id}
	logs.Info(o.Read(&tag, "Id"))
	c.Data["json"] = map[string]interface{}{"result": "ok!", "data": tag}
	c.ServeJSON()
}

/*
功能：文章分页
路由：/api.post.list/page/:id([0-9]+)?size=2
代码：根据相关参数对文章进行分页
*/
type PostListPageController struct {
	beego.Controller
}

func (c *PostListPageController) Get() {
	o := orm.NewOrm()
	var posts []*models.Post
	size, err := c.GetInt("size")
	str := c.Ctx.Input.Param(":id")
	page, err := strconv.Atoi(str)
	offset := (page - 1) * size
	num, err := o.QueryTable("post").OrderBy("-Id").Limit(size, offset).All(&posts)
	if err != nil {
		logs.Info(err)
	}
	if int(num) < size {
		next := "无"
		previous := page - 1
		c.Data["json"] = map[string]interface{}{"result": "ok!", "num": num, "previous": previous, "next": next, "data": posts}
		c.ServeJSON()
	} else if int(num) > size {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "没有啦啦啦！"}
		c.ServeJSON()
	} else {
		next := page + 1
		previous := page - 1
		c.Data["json"] = map[string]interface{}{"result": "ok!", "num": num, "previous": previous, "next": next, "data": posts}
		c.ServeJSON()
	}
}

/*
功能：文章按照类型分类分页
路由：/api.post.category/:id([0-9]+)?page=1&size=5
代码：根据相关参数对文章进行分类分页
*/
type CategoryListController struct {
	beego.Controller
}

func (c *CategoryListController) Get() {
	str := c.Ctx.Input.Param(":id")
	tid, err := strconv.Atoi(str)
	if err != nil {
		logs.Info(err)
	}
	o := orm.NewOrm()
	var posts []*models.Post
	size, err := c.GetInt("size")
	page, err := c.GetInt("page")
	offset := (page - 1) * size
	num, err := o.QueryTable("post").OrderBy("-Id").Filter("tid__icontains", tid).Limit(size, offset).All(&posts)
	if err != nil {
		logs.Info(err)
	}
	if int(num) < size {
		next := "无"
		previous := page - 1
		c.Data["json"] = map[string]interface{}{"result": "ok!", "previous": previous, "next": next, "data": posts}
		c.ServeJSON()
	} else if int(num) > size {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "没有啦啦啦！"}
		c.ServeJSON()
	} else if int(num) == 0 {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "没有相关内容!"}
		c.ServeJSON()
	} else {
		next := page + 1
		previous := page - 1
		c.Data["json"] = map[string]interface{}{"result": "ok!", "previous": previous, "next": next, "data": posts}
		c.ServeJSON()
	}
}

/*
功能：获取单篇文章
路由：/api.post.id/:id([0-9]+)
代码：根据id获取文章
*/
type PostController struct {
	beego.Controller
}

func (c *PostController) Get() {
	str := c.Ctx.Input.Param(":id")
	is := c.GetString("is")
	id, err := strconv.Atoi(str)
	if err != nil {
		logs.Info(err)
	}
	o := orm.NewOrm()
	post := models.Post{Id: id}
	logs.Info(o.Read(&post, "Id"))
	if is == "admin" {
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": post}
		c.ServeJSON()
	} else {
		post.Views++
		logs.Info(o.Update(&post))
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": post}
		c.ServeJSON()
	}
}

/*
功能：获取评论
路由：/api.post.comment/:id([0-9]+)
代码：根据文章id获取评论
*/
type PostCommentController struct {
	beego.Controller
}

type Reply struct {
	Id        int       `json:"Id"`
	Uid       string    `json:"Uid"`
	UserName  string    `json:"UserName"`
	UserImage string    `json:"UserImage"`
	CommentId int       `json:"CommentId"`
	Thumbs    int       `json:"Thumbs"`
	Body      string    `json:"Body"`
	Time      time.Time `json:"Time"`
}

type Comment struct {
	Id        int       `json:"Id"`
	Uid       string    `json:"Uid"`
	UserName  string    `json:"UserName"`
	UserImage string    `json:"UserImage"`
	PostId    int       `json:"PostId"`
	Thumbs    int       `json:"Thumbs"`
	Body      string    `json:"Body"`
	Time      time.Time `json:"Time"`
	Reply     []Reply   `json:"Reply"`
}

func (c *PostCommentController) Get() {
	str := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(str)
	if err != nil {
		logs.Info(err)
	}
	o := orm.NewOrm()
	var jsoncomment []Comment
	var comments []*models.Comment
	number, err := o.QueryTable("comment").OrderBy("-Id").Filter("post_id", id).All(&comments)
	if err != nil {
		logs.Info(number)
	}
	for _, comment := range comments {
		var jsonCommentReply []Reply
		var CommentReplys []*models.CommentReply
		numbers, err := o.QueryTable("comment_reply").Filter("comment_id", comment.Id).All(&CommentReplys)
		if err != nil {
			logs.Info(numbers)
		}
		for _, CommentReply := range CommentReplys {

			jsonCommentReply = append(jsonCommentReply, Reply{Id: CommentReply.Id, Uid: CommentReply.Uid, UserName: CommentReply.UserName, UserImage: CommentReply.UserImage, CommentId: CommentReply.CommentId, Thumbs: CommentReply.Thumbs, Body: CommentReply.Body, Time: CommentReply.Time})
		}
		jsoncomment = append(jsoncomment, Comment{Id: comment.Id, Uid: comment.Uid, UserName: comment.UserName, UserImage: comment.UserImage, PostId: comment.Id, Thumbs: comment.Thumbs, Body: comment.Body, Time: comment.Time, Reply: jsonCommentReply})
	}
	logs.Info(id)
	c.Data["json"] = map[string]interface{}{"result": "ok!", "data": jsoncomment}
	c.ServeJSON()
}

/*
功能：获取评论
路由：/api.post.comment.list
*/
type CommentListController struct {
	beego.Controller
}

func (c *CommentListController) Get() {
	o := orm.NewOrm()
	var comment []*models.Comment
	num, err := o.QueryTable("comment").OrderBy("-Id").All(&comment)
	if err != nil {
		logs.Info(num)
	}
	c.Data["json"] = map[string]interface{}{"result": "ok!", "data": comment}
	c.ServeJSON()
}

/*
功能：获取回复
路由：/api.post.reply.list
*/
type ReplyListController struct {
	beego.Controller
}

func (c *ReplyListController) Get() {
	o := orm.NewOrm()
	var CommentReply []*models.CommentReply
	numbers, err := o.QueryTable("comment_reply").All(&CommentReply)
	if err != nil {
		logs.Info(numbers)
	}
	c.Data["json"] = map[string]interface{}{"result": "ok!", "data": CommentReply}
	c.ServeJSON()
}

/*
功能：获取友链
路由：/api.post.link.list
代码：获取全部友链
*/
type LinkListController struct {
	beego.Controller
}

func (c *LinkListController) Get() {
	var links []*models.Links
	o := orm.NewOrm()
	is := c.GetString("is")
	if is == "admin" {
		number, err := o.QueryTable("links").OrderBy("-Id").All(&links)
		if err != nil {
			logs.Info(number)
		}
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": links}
		c.ServeJSON()
	} else {
		number, err := o.QueryTable("links").OrderBy("-Id").Filter("is", 1).All(&links)
		if err != nil {
			logs.Info(number)
		}
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": links}
		c.ServeJSON()
	}
}

/*
功能：获取单条友链
路由：/api.post.link.id/:id([0-9]+)
代码：根据id获取单独友链
*/
type LinkController struct {
	beego.Controller
}

func (c *LinkController) Get() {
	str := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(str)
	if err != nil {
		logs.Info(err)
	}
	o := orm.NewOrm()
	link := models.Links{Id: id}
	logs.Info(o.Read(&link, "Id"))
	c.Data["json"] = map[string]interface{}{"result": "ok!", "data": link}
	c.ServeJSON()
}

/*
功能：搜索文章
路由：/api.search
代码：根据关键词搜索文章，标题
*/
type SearchController struct {
	beego.Controller
}

func (c *SearchController) Get() {
	q := c.GetString("q")
	var Post []*models.Post
	o := orm.NewOrm()
	num, err := o.QueryTable("post").OrderBy("-Id").Filter("title__icontains", q).All(&Post)
	if err != nil {
		logs.Info(num)
	}
	if int(num) > 0 {
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": Post}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "没有搜索到相关内容!"}
		c.ServeJSON()
	}
}

/*
 功能：随机头像图片
 路由：/api.random.image
 功能：用户注册的时候，给用户提供随机图片选择
*/
type RandomHeadController struct {
	beego.Controller
}

func (c *RandomHeadController) Post() {
	file := models.RandomHead()
	c.Data["json"] = map[string]interface{}{"result": "ok!", "data": file}
	c.ServeJSON()
}

/*
 功能：随机背景图片
 路由：/api.random.back
 功能：获取随机背景图片，用于博客背景图
*/
type RandomBackController struct {
	beego.Controller
}

func (c *RandomBackController) Post() {
	file := models.RandomBack()
	c.Data["json"] = map[string]interface{}{"result": "ok!", "data": file}
	c.ServeJSON()
}

/*
 功能：获取上传的所有图片
 路由：/api.image.list
 功能：验证用户，获取上传的所有文件
*/
type UploadImageListController struct {
	beego.Controller
}

func (c *UploadImageListController) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true {
		o := orm.NewOrm()
		var images []*models.UploadImage
		num, err := o.QueryTable("upload_image").OrderBy("-Id").All(&images)
		if err != nil {
			logs.Info(num)
		}
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": images}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "你没有访问权限!"}
	}
}

/*
 功能：获取上传的所有文件
 路由：/api.file.list
 功能：验证用户，获取上传的所有文件
*/
type UploadFileListController struct {
	beego.Controller
}

func (c *UploadFileListController) Post() {
	uid := c.GetString("uid")
	token := c.GetString("token")
	result := models.GetToken(uid, token)
	if result == true {
		o := orm.NewOrm()
		var files []*models.UploadFile
		num, err := o.QueryTable("upload_file").OrderBy("-Id").All(&files)
		if err != nil {
			logs.Info(num)
		}
		c.Data["json"] = map[string]interface{}{"result": "ok!", "data": files}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"result": "err!", "data": "你没有访问权限!"}
	}
}
