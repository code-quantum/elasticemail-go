package elasticemail

import (
	"fmt"
)

const (
	listEndpoint = "list" // API methods for managing your Lists
	mList        = "list" // Shows all your existing lists
)

type ListParams struct {
	From string `json:"from" url:"from"` // Starting date for search in YYYY-MM-DDThh:mm:ss format.
	To   string `json:"to" url:"to"`     // Ending date for search in YYYY-MM-DDThh:mm:ss format.
}

// Shows all your existing lists
// Access Level required: ViewContacts
func (m *ElasticEmailImpl) List(params ListParams) (lists *[]List, err error) {

	url := fmt.Sprintf("%s/%s/%s", m.apiBase, listEndpoint, mList)

	out := make([]List, 0)
	err = sendGetResp(m, url, params, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
