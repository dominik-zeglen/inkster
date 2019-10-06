package mail

import (
	"fmt"
	"log"

	"github.com/dominik-zeglen/inkster/config"
	"github.com/dominik-zeglen/inkster/utils"
)

type TerminalMailer struct {
	sender string
}

func NewTerminalMailer(config config.Config) TerminalMailer {
	return TerminalMailer{
		sender: config.Mail.Sender,
	}
}

func (mailer *TerminalMailer) printToTerminal(
	recipient string,
	data interface{},
) error {
	prettyData, err := utils.PrintJSON(data)

	if err != nil {
		return err
	}

	message := fmt.Sprintf(
		"Sending message from %s to %s\n-----\n%s\n",
		mailer.sender,
		recipient,
		prettyData,
	)

	log.Println(message)

	return nil
}

func (mailer TerminalMailer) SendPasswordResetToken(
	recipient string,
	data SendPasswordResetTokenTemplateData,
) error {
	return mailer.printToTerminal(recipient, data)
}

func (mailer TerminalMailer) SendUserInvitation(
	recipient string,
	data SendUserInvitationTemplateData,
) error {
	return mailer.printToTerminal(recipient, data)
}
