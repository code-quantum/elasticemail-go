package elasticemail

import (
	"testing"
)

func TestEmailValidation(t *testing.T) {

	ee := NewElasticEmail("XXXXXXXXXX")
	ee.GetEmailStatus(GetEmailStatusParams{TransactionID: "800e365b-dd15-4152-805b-ee6d441e78b4"})
	//ensure.Nil(t, err)
	//
	//ev, err := validator.ValidateEmail("foo@mailgun.com", false)
	//ensure.Nil(t, err)
	//
	//ensure.True(t, ev.IsValid)
	//ensure.DeepEqual(t, ev.MailboxVerification, "")
	//ensure.False(t, ev.IsDisposableAddress)
	//ensure.False(t, ev.IsRoleAddress)
	//ensure.True(t, ev.Parts.DisplayName == "")
	//ensure.DeepEqual(t, ev.Parts.LocalPart, "foo")
	//ensure.DeepEqual(t, ev.Parts.Domain, "mailgun.com")
	//ensure.True(t, ev.Reason == "")
}
