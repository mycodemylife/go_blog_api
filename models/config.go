package models

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofrs/uuid"
	"gopkg.in/gomail.v2"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "username:password@/db_name?charset=utf8&loc=Asia%2FShanghai")
	logs.Info("\n连接数据库成功!")
}

//加密密码
func Encrypt(text string) (string, error) {
	key := []byte{0xBA, 0x47, 0x2F, 0x02, 0xC8, 0x92, 0x1F, 0x7D,
		0x2A, 0x3D, 0x8F, 0x06, 0x41, 0x9B, 0x6F, 0x2D,
		0xBA, 0x36, 0x6F, 0x07, 0xC7, 0x52, 0x1F, 0x7D,
		0x4A, 0x5D, 0x4F, 0x06, 0x45, 0x8B, 0x3F, 0x4D,
	}
	var iv = key[:aes.BlockSize]
	encrypted := make([]byte, len(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(encrypted, []byte(text))
	return hex.EncodeToString(encrypted), nil
}

//解密密码
func Decrypt(encrypted string) (string, error) {
	key := []byte{0xBA, 0x47, 0x2F, 0x02, 0xC8, 0x92, 0x1F, 0x7D,
		0x2A, 0x3D, 0x8F, 0x06, 0x41, 0x9B, 0x6F, 0x2D,
		0xBA, 0x36, 0x6F, 0x07, 0xC7, 0x52, 0x1F, 0x7D,
		0x4A, 0x5D, 0x4F, 0x06, 0x45, 0x8B, 0x3F, 0x4D,
	}
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	src, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	var iv = key[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var block cipher.Block
	block, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}

/*
获取UUID
*/
func Uid() string {
	uid, err := uuid.NewV4()
	if err != nil {
		logs.Info(err)
	}
	return uid.String()
}

/*
 beego.AppConfig.String 获取app.conf相关参数，具体见beego官方文档
*/
func SendEmail(user string, msg string) string {
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", user)
	m.SetHeader("Subject", "通知邮件！如有打扰，请多包涵！")
	m.SetBody("text/html", msg)
	d := gomail.NewDialer("smtp.qq.com", 587, username, password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
		return "err!"
	}
	return "ok"
}

/*
 计算文件哈希值
*/
func GetSha1(path string) string {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return ""
	}
	h := sha1.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return ""
	}
	result := fmt.Sprintf("%x", h.Sum(nil))
	return result
}

/*
 添加token进入Redis缓存，用于用户认证
*/
func SetToken(uid string) string {
	red := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	token := Uid()
	err := red.Set(uid, token, 90*24*60*time.Minute).Err()
	if err != nil {
		panic(err)
	}
	return token
}

/*
 根据token判断是否是已登录的用户
*/
func GetToken(uid string, token string) bool {
	red := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	to, err := red.Get(uid).Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		panic(err)
		return false
	} else {
		if to == token {
			return true
		} else {
			return false
		}

	}
}

/*
 删除token，用户登出
*/
func DeToken(uid string) string {
	red := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := red.Del(uid).Result()
	if err == redis.Nil {
		return "err!"
	} else if err != nil {
		panic(err)
		return "err!"
	} else {
		return "ok!"
	}

}

/*
 功能：产生随机图片，供给网站头像使用
*/
func RandomHead() string {
	b, err := ioutil.ReadFile("Head.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	Img := strings.Split(str, "\n")
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(Img)
	IMG := Img[n]
	file := beego.AppConfig.String("host") + "/api.file" + strings.Replace(IMG, "\r\n", "", 1)
	return file
}

/*
 生成随机背景图
*/
func RandomBack() string {
	b, err := ioutil.ReadFile("IMG.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	Img := strings.Split(str, "\n")
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(Img)
	IMG := Img[n]
	return IMG
}
