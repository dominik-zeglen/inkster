package mailer

import "log"

type MockMailClient struct{}

func (_ MockMailClient) Send(recipent string, subject string, message string) error {
	log.Printf("Sending mail to %s:\n", recipent)
	log.Printf("Subject: %s", subject)
	log.Println(message)
	return nil
}
func (_ MockMailClient) SendMany(recipents []string, subject string, message string) error {
	recipentList := recipents[0]
	for _, recipent := range recipents[1:] {
		recipentList += ", " + recipent
	}
	log.Printf("Sending mail to %s:\n", recipentList)
	log.Printf("Subject: %s", subject)
	log.Println(message)
	return nil
}
