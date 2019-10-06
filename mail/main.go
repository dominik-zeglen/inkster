package mail

import "github.com/dominik-zeglen/inkster/core"

type SendPasswordResetTokenTemplateData struct {
	User    core.User    `json:"user"`
	Website core.Website `json:"website"`
	Token   string       `json:"token"`
}

type Mailer interface {
	SendPasswordResetToken(
		string,
		SendPasswordResetTokenTemplateData,
	) error
}
