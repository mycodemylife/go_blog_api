package routers

import (
	"weblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/api.user.register",&controllers.NewUser{})
    beego.Router("/api.user.login",&controllers.UserLogin{})
    beego.Router("/api.user.admin.login",&controllers.AdminUserLogin{})
    beego.Router("/api.user.logout",&controllers.UserLogout{})
	beego.Router("/api.user.list",&controllers.UserList{})
    beego.Router("/api.user.data",&controllers.UserData{})
    beego.Router("/api.user.delete",&controllers.UserDelete{})
    beego.Router("/api.user.header",&controllers.UserHeader{})

    beego.Router("/api.tag.new",&controllers.NewTag{})
    beego.Router("/api.tag.update",&controllers.UpdateTag{})
    beego.Router("/api.tag.delete",&controllers.DeleteTag{})
    beego.Router("/api.post.new",&controllers.NewPost{})
    beego.Router("/api.post.update",&controllers.UpdatePost{})
    beego.Router("/api.post.delete",&controllers.DeletePost{})
    beego.Router("/api.comment.new",&controllers.NewComment{})
    beego.Router("/api.comment.delete",&controllers.DeleteComment{})
    beego.Router("/api.reply.new",&controllers.NewCommentReply{})
    beego.Router("/api.reply.delete",&controllers.DeleteCommentReply{})
    beego.Router("/api.link.new",&controllers.NewLink{})
    beego.Router("/api.link.update",&controllers.UpdateLink{})
    beego.Router("/api.link.delete",&controllers.DeleteLink{})

    beego.Router("/api.tag.list",&controllers.TagListController{})
    beego.Router("/api.tag.id/:id([0-9]+)",&controllers.TagController{})
    beego.Router("/api.post.list",&controllers.PostListPageController{})
    beego.Router("/api.post.list/page/:id([0-9]+)",&controllers.PostListPageController{})
    beego.Router("/api.post.category/:id([0-9]+)",&controllers.CategoryListController{})
    beego.Router("/api.post.id/:id([0-9]+)",&controllers.PostController{})
    beego.Router("/api.post.comment/:id([0-9]+)",&controllers.PostCommentController{})
    beego.Router("/api.link.list",&controllers.LinkListController{})
    beego.Router("/api.link.id/:id([0-9]+)",&controllers.LinkController{})

    beego.Router("/api.comment.list",&controllers.CommentListController{})
    beego.Router("/api.reply.list",&controllers.ReplyListController{})

    beego.Router("/api.search",&controllers.SearchController{})
    beego.Router("/api.upload.image",&controllers.UploadImage{})
    beego.Router("/api.upload.file",&controllers.UploadFile{})
    beego.Router("/api.image.list",&controllers.UploadImageListController{})
    beego.Router("/api.file.list",&controllers.UploadFileListController{})
    beego.Router("/api.random.image",&controllers.RandomHeadController{})
    beego.Router("/api.random.back",&controllers.RandomBackController{})
}
