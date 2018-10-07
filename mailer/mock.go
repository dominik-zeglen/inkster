package mailer

import "log"

type MockMailClient struct {
	log []string
}

func (client *MockMailClient) Send(recipent string, subject string, message string) error {
	log.Printf("Sending mail to %s:\n", recipent)
	log.Printf("Subject: %s", subject)
	log.Println(message)

	client.log = append(client.log, message)

	return nil
}
func (client *MockMailClient) SendMany(recipents []string, subject string, message string) error {
	recipentList := recipents[0]
	for _, recipent := range recipents[1:] {
		recipentList += ", " + recipent
	}
	log.Printf("Sending mail to %s:\n", recipentList)
	log.Printf("Subject: %s", subject)
	log.Println(message)

	client.log = append(client.log, message)

	return nil
}
func (client *MockMailClient) Reset() {
	client.log = []string{}
}
func (client MockMailClient) Last() string {
	return client.log[len(client.log)-1]
}
