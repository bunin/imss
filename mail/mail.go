package mail

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

var (
	_host      string
	_port      int
	_fromName  string
	_fromEmail string
	_login     string
	_password  string
)

func Setup(host string, port int, fromName string, fromEmail string, login string, password string) {
	_host = host
	_port = port
	_fromName = fromName
	_fromEmail = fromEmail
	_login = login
	_password = password
}

func Send(to, name, link string) error {
	dialer := gomail.NewDialer(_host, _port, _login, _password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	msg := gomail.NewMessage()
	msg.SetHeaders(map[string][]string{
		"From":    {_fromName + " <" + _fromEmail + ">"},
		"Subject": {name},
		"To":      {to},
	})
	msg.AddAlternative("text/plain", Text(name, link))
	msg.AddAlternative("text/html", Html(name, link))
	return dialer.DialAndSend(msg)
}
