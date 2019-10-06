package config

import "fmt"

type MailBackend string

const MailAwsSes = "ses"
const MailTerm = "term"

func getMailBackend(str MailBackend) (MailBackend, error) {
	switch str {
	case MailTerm:
		return MailTerm, nil
	case MailAwsSes:
		return MailAwsSes, nil
	default:
		return MailBackend(""), fmt.Errorf("Unknown mail backend: %s", str)
	}
}

type mailConfig struct {
	Backend MailBackend `toml:"backend"`
	Sender  string      `toml:"sender"`
}
