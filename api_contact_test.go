package elasticemail

import (
	"github.com/facebookgo/ensure"
	"testing"
)

func TestAddContact(t *testing.T) {
	ee := NewElasticEmailFromEnv()
	err := ee.AddContact(AddContactParams{
		Email:           "let+1@gmail.com",
		ListName:        []string{"listName1"},
		PublicAccountID: "xxxxxxxx-0000-0a0a-1111-0a0000a00000",
	})
	ensure.Nil(t, err)
}
