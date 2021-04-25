package twilio

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendSMS(t *testing.T) {
	var twilio = New(&Config{
		AccountSID:  os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken:   os.Getenv("TWILIO_AUTH_TOKEN"),
		SenderPhone: os.Getenv("TWILIO_SENDER_PHONE"),
	})

	var err = twilio.SendSMS(os.Getenv("TWILIO_RECEIVER_PHONE"), "Test")
	assert.NoError(t, err)

}
