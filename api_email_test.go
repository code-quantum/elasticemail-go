package elasticemail

import (
	"github.com/facebookgo/ensure"
	"testing"
)

func TestIncorrectApiKey(t *testing.T) {
	ee := NewElasticEmail("XXXXXXXXXX")
	_, err := ee.GetEmailStatus(GetEmailStatusParams{TransactionID: "800e365b-dd15-4152-805b-ee6d441e78b4"})
	ensure.NotNil(t, err)
	ensure.True(t, err.Error() == "Incorrect apikey")
}

func TestGetStatus(t *testing.T) {
	ee := NewElasticEmailFromEnv()
	_, err := ee.GetEmailStatus(GetEmailStatusParams{TransactionID: "800e365b-dd15-4152-805b-ee6d441e78b4"})
	ensure.Nil(t, err)
}

func TestStatus(t *testing.T) {
	ee := NewElasticEmailFromEnv()
	_, err := ee.Status(StatusParams{MessageID: "OqJjVnGtN2Nm4vr5zmgQRg2"})
	ensure.Nil(t, err)
}

func TestView(t *testing.T) {
	ee := NewElasticEmailFromEnv()
	_, err := ee.View(ViewParams{MessageID: "OqJjVnGtN2Nm4vr5zmgQRg2"})
	ensure.Nil(t, err)
}

func TestSend(t *testing.T) {
	ee := NewElasticEmailFromEnv()
	to := make([]string, 1)
	to[0] = "abc@gmail.com"
	_, err := ee.Send(SendParams{To: to, BodyHtml: "<b>Hello</b>", Subject: "111111", From: "noreply@eeee.eee"})
	ensure.Nil(t, err)
}
