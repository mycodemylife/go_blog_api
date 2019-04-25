### 基于Go语言的开源博客api项目

> 本项目是以beego学习为主，使用的是前后端分离的开发方式，本次开源的也仅仅是后端api这部分，供大家学习参考，也欢迎给我报漏洞，望不要攻击我的博客；集成了用户注册登录，通过token实现会话处理，集成了文件上传下载功能。

项目信息：
* 展示：https://zhang18.top
* 后端：beego框架
* 前端：Vue

#### 一.使用项目该项目

获取项目
```
go get github.com/msterzhang/go_blog_api
```

1.安装相关依赖

后端框架

```
go get github.com/astaxie/beego
```
Mysql驱动

```
go get github.com/go-sql-driver/mysql
```
Redis缓存驱动，用于用户token回话，也可以用于缓存博客数据
```
go get github.com/go-redis/redis
```
邮件发送，具体按照你的需求，见config文件
```
go get gopkg.in/gomail.v2
```
2.打开app.conf，修改host，这个使用来确定你上传文件的下载路径的


```
appname = weblog
httpport = 8003
runmode = dev

host=http://localhost:8003
```
3.打开数据库config.go文件，设置mysql数据库信息，数据库请先创建好，程序会自动创建相关表

```
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "username:password@/数据库名字?charset=utf8&loc=Asia%2FShanghai")
	logs.Info("\n连接数据库成功!")
}
```

#### 二.api文档

> 这个项目的api文档，我使用的是postman创建。

地址：

```
https://documenter.getpostman.com/view/6651807/S1EJYMDb
```

#### 三.关于项目部署

我采用的是Nginx代理的方式部署网站，其中有几处需要注意：

Nginx相关命令:

安装(以Ubuntu为例)

```
sudo apt-get install nginx
```

停止
```
sudo service nginx stop
```
启动
```
sudo service nginx start
```
1.创建Nginx代理转发文件blogapi.conf，放在conf.d目录下，这是开启https的版本，具体如何获取，请看我的博客，不开启的很简单，改改就可以
```
server {
    charset utf-8;
    listen 443 ssl;
    server_name api.zhang18.top;
    ssl_certificate /etc/letsencrypt/live/zhang18.top/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/zhang18.top/privkey.pem;

    access_log  /root/go/src/weblog/access.log;

    location / {
        try_files /_not_exists_ @backend;
    }

    location @backend {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host            $http_host;

        proxy_pass http://127.0.0.1:8003;
    }
}


server {
    listen 80;
    server_name api.zhang18.top;
    return 301 https://$host$request_uri;
}
```

2.给权限,打开nginx.conf,修改如下
```
#user www-data;
user root;
```
3.上传文件大小控制，Nginx默认是不大于20M，所以你需要更改规则,打开nginx.conf,修改如下

```
client_max_body_size 3072m;
```
4.项目运行

服务器编译项目:

```
go build main.go
```
运行项目
```
./main
```

常驻后台运行
```
nohup ./main &
```
#### 四.如何关闭后台任务?

查看后台任务，找到进程号
```
ps -ef
```

杀死进程
```
kill 8999
```

> QQ群：929724129

#### 请作者喝杯咖啡

![](https://api.zhang18.top/api.file/download/mm_facetoface_collect_qrcode_1556171490471.png)