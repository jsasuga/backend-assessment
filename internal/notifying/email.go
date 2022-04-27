package notifying

type EmailSender struct {
	apiKey string
	apiUrl string
	sender string
}

type EmailNotifier interface {
	SendEmail(subject string, emailTo []string, templateId string, inputData interface{}) error
}

func CreateEmailSender(apiKey string, apiUrl string, sender string) *EmailSender {
	return &EmailSender{
		apiKey: apiKey,
		apiUrl: apiUrl,
		sender: sender,
	}
}

func (e *EmailSender) SendEmail(subject string, emailTo []string, templateId string, inputData interface{}) error {
	return nil
}
