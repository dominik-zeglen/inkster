package mail

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

	"github.com/dominik-zeglen/inkster/config"
)

type AwsSesMailer struct {
	awsSes *ses.SES
	sender string
}

func NewAwsSesMailer(config config.Config) AwsSesMailer {
	awsConfig := aws.Config{
		Region: &config.AWS.Region,
	}
	awsConfig.Credentials = credentials.NewStaticCredentials(
		config.AWS.AccessKey,
		config.AWS.SecretAccessKey,
		"",
	)

	awsSession, err := session.NewSession(&awsConfig)
	if err != nil {
		panic(err)
	}

	awsSes := ses.New(awsSession)

	return AwsSesMailer{
		awsSes: awsSes,
		sender: config.Mail.Sender,
	}
}

func (mailer AwsSesMailer) SendPasswordResetToken(
	recipient string,
	data SendPasswordResetTokenTemplateData,
) error {
	templateData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	input := &ses.SendTemplatedEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Template:     aws.String("ResetPassword"),
		TemplateData: aws.String(string(templateData)),
		Source:       aws.String(mailer.sender),
	}

	_, err = mailer.awsSes.SendTemplatedEmail(input)

	return err
}

func (mailer AwsSesMailer) SendUserInvitation(
	recipient string,
	data SendUserInvitationTemplateData,
) error {
	templateData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	input := &ses.SendTemplatedEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Template:     aws.String("UserInvitation"),
		TemplateData: aws.String(string(templateData)),
		Source:       aws.String(mailer.sender),
	}

	_, err = mailer.awsSes.SendTemplatedEmail(input)

	return err
}
