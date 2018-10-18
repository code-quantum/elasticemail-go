package elasticemail

import (
	"fmt"
	"github.com/facebookgo/ensure"
	"testing"
)

func TestGetTemplateList(t *testing.T) {
	ee := NewElasticEmailFromEnv()
	res, err := ee.GetList(TemplateGetListParams{Limit: 100, Offset: 0})
	ensure.Nil(t, err)
	fmt.Printf("\n\nTemplateList:\n %+v \n", res)
}
