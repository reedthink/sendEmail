//main
package main

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/go-gomail/gomail"
)

//SendRegisterEmail 发送注册邮件。address:目标邮箱。body: html文件位置
func SendRegisterEmail(address string, body string) error {
	//配置发送端
	mailConn := map[string]string{

		"user": "reedthink@qq.com",

		"pass": "***",
		//这里的密码不是邮箱的密码，是邮箱给的一个授权码

		"host": "smtp.qq.com",

		"port": "587",
	}
	port, _ := strconv.Atoi(mailConn["port"])

	m := gomail.NewMessage()
	m.SetAddressHeader("From", mailConn["user"], "admin")
	m.SetAddressHeader("To", address, "You")
	m.SetHeader("Subject", "欢迎加入MeetDown") //设置邮件主题

	buf, err := ioutil.ReadFile(body) //此函数返回[]byte,而不是*FILE
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}

	m.SetBody("text/html", string(buf))

	d := gomail.NewPlainDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err = d.DialAndSend(m)
	return err
}
func GetTime() string {
	now := time.Now()      //获取当前时间
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	return fmt.Sprintf("%d年%02d月%02d日%02d点%02d分%02d秒", year, month, day, hour, minute, second)
}
func GetDate() string {
	now := time.Now()    //获取当前时间
	year := now.Year()   //年
	month := now.Month() //月
	day := now.Day()     //日
	return fmt.Sprintf("%d年%02d月%02d日", year, month, day)
}

//UserInfo 是用户信息。首字母记得大写，害
type UserInfo struct {
	Name    string
	UID     string
	Mail    string
	RegTime string //注册时间
	Date    string
}

func GenerateHTML(name string, uid string, mail string) {
	U := UserInfo{name, uid, mail, GetTime(), GetDate()}
	tmpl, _ := template.ParseFiles("registerTmpl.html") //读取模板
	F, _ := os.Create("register.html")
	defer F.Close()

	tmpl.Execute(F, U) //替换
}
func main() {

	address := "xxxxxxxxxxxx@163.com"
	GenerateHTML("律政先锋张益达", "00000666", address) //替换模板

	err := SendRegisterEmail(address, "register.html")

	if err != nil {
		fmt.Printf("发送失败，你翻车了。%v\n", err)
		return
	}
	fmt.Println("发送成功！")

}
