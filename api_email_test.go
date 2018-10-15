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
