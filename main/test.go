package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func main() {
	auth := smtp.PlainAuth("", "1069351083@qq.com", "tkqiwsqjtfndbeja", "smtp.qq.com")
	to := []string{"1807326334@qq.com"}
	nickname := "test"
	user := "1069351083@qq.com"
	subject := "test mail"
	content_type := "Content-Type: text/html; charset=UTF-8"
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="iso-8859-15">
			<title>MMOGA POWER</title>
		</head>
		<body>
			GO 发送邮件，官方连包都帮我们写好了，真是贴心啊！！！
		</body>
		</html>`

	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
}
