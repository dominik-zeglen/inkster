package mailer

import (
	"crypto/tls"
	"fmt"
	"net/mail"
	"net/smtp"
)

type SmtpMailClient struct {
	address  string
	auth     smtp.Auth
	hostname string
	port     string
}

func NewSmtpMailClient(
	login string,
	address string,
	password string,
	hostname string,
	port string,
) SmtpMailClient {
	auth := smtp.PlainAuth("", login, password, hostname)
	return SmtpMailClient{
		address:  address,
		auth:     auth,
		hostname: hostname,
		port:     port,
	}
}

func (client SmtpMailClient) Send(recipent string, subject string, body string) error {
	dialHost := client.hostname + ":" + client.port
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         client.hostname,
	}
	from := mail.Address{"", client.address}
	to := mail.Address{"", recipent}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	conn, err := tls.Dial("tcp", dialHost, tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, client.hostname)
	if err != nil {
		return err
	}
	defer c.Quit()

	err = c.Auth(client.auth)
	if err != nil {
		return err
	}

	err = c.Mail(from.Address)
	if err != nil {
		return err
	}

	err = c.Rcpt(to.Address)
	if err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}
func (client SmtpMailClient) SendMany(recipents []string, subject string, message string) error {
	return smtp.SendMail(
		client.hostname+":"+client.port,
		client.auth,
		client.address,
		recipents,
		[]byte(message),
	)
}
