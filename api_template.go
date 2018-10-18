package elasticemail

import (
	"fmt"
)

const (
	templateEndpoint = "template" // Managing and editing templates of your emails
	mGetList         = "getlist"  // Shows all your existing lists
)

type TemplateGetListParams struct {
	Limit  int `json:"limit" url:"limit"`   // Maximum of loaded items.
	Offset int `json:"offset" url:"offset"` // How many items should be loaded ahead.
}

// Lists your templates
// Access Level required: ViewTemplates
func (m *ElasticEmailImpl) GetList(params TemplateGetListParams) (lists *TemplateList, err error) {

	url := fmt.Sprintf("%s/%s/%s", m.apiBase, templateEndpoint, mGetList)

	out := TemplateList{}
	err = sendGetResp(m, url, params, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
