package elasticemail

import (
	"fmt"
	"github.com/facebookgo/ensure"
	"testing"
)

func TestList(t *testing.T) {
	ee := NewElasticEmailFromEnv()
	res, err := ee.List(ListParams{})
	ensure.Nil(t, err)
	fmt.Printf("\n\nLists:\n %+v \n", res)
}
