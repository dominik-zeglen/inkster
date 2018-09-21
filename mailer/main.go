package mailer

type Mailer interface {
	Send(string, string, string) error
	SendMany([]string, string, string) error
}
