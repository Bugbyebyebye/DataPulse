package util

import (
	"commons/config"
	"math/rand"
	"net/smtp"
	"strconv"
)

func FormEmail(email string) (int, error) {
	var host = config.Conf.SMTP.Host
	var username = config.Conf.SMTP.Username
	var fromemail = config.Conf.SMTP.Password
	var password = config.Conf.SMTP.Password
	validate := rand.Intn(10000) + 10000
	msg := []byte("Subject: 邮箱验证码\n\n您的邮箱验证码为：" + strconv.Itoa(validate))

	auth := smtp.PlainAuth("", username, password, host)
	err := smtp.SendMail(host, auth, fromemail, []string{email}, msg)
	if err != nil {
		return 0, err
	}
	return validate, nil

}
