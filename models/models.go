package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

/*
用户:用户ID 用户UID 是否是管理员(0 否 1 是) 用户名 密码 头像 邮箱(用于回复提醒) 创建时间
*/
type User struct {
	Id        int
	Uid       string
	IsAdmin   int
	Username  string
	Password  string
	AvatarUrl string
	Email     string
	Time      time.Time
}

/*
博客文章:文章ID 文章UID 标题 分类 图片 浏览量 点赞量 正文 创建时间
*/
type Post struct {
	Id       int
	Uid      string
	Title    string
	Tid      int
	Tag      string
	Image    string
	Views    int
	Thumbs   int
	Content  string `orm:"type(text)"`
	Markdown string `orm:"type(text)"`
	Time     time.Time
}

/*
博客文章标签:标签ID 标签名字 创建时间
*/
type Tag struct {
	Id   int
	Name string
	Time time.Time
}

/*
博客文章评论:评论ID 评论UID 评论人的用户名 评论人的头像 评论文章的ID 点赞数 正文 创建时间
*/
type Comment struct {
	Id        int
	Uid       string
	UserName  string
	UserImage string
	PostId    int
	Thumbs    int
	Body      string
	Time      time.Time
}

/*
博客文章回复:回复的ID 回复的UID  回复人的用户名 回复人的头像 回复评论的ID 点赞数 正文 创建时间
*/
type CommentReply struct {
	Id        int
	Uid       string
	UserName  string
	UserImage string
	CommentId int
	Thumbs    int
	Body      string
	Time      time.Time
}

/*
友情链接：网站ID Uid 站长邮箱 网站名字 网站的图片 网站链接 网站介绍 时间
*/
type Links struct {
	Id        int
	Uid       string
	Is        int
	Email     string
	HostName  string
	HostImage string
	HostUrl   string
	Body      string `orm:"type(text)"`
	Time      time.Time
}

/*
上传的图片
*/
type UploadImage struct {
	Id       int
	UserId   string
	Uid      string
	Md       string
	Download string
	Time     time.Time
}

/*
上传的文件
*/
type UploadFile struct {
	Id       int
	UserId   string
	Uid      string
	Md       string
	Download string
	Time     time.Time
}

func init() {
	orm.RegisterModel(new(User), new(Tag), new(Post), new(Comment), new(CommentReply), new(Links),new(UploadImage),new(UploadFile))
	re := orm.RunSyncdb("default", false, true)
	logs.Info(re)
	logs.Info("创建数据表成功！")
}
