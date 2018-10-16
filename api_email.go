package elasticemail

import (
	"fmt"
)

const (
	emailEndpoint   = "email"
	mGetEmailStatus = "getstatus"
)

type GetEmailStatusParams struct {
	TransactionID    string `json:"transactionID" url:"transactionID"`
	ShowAbuse        bool   `json:"showAbuse" url:"showAbuse"`
	ShowClicked      bool   `json:"showClicked" url:"showClicked"`
	ShowDelivered    bool   `json:"showDelivered" url:"showDelivered"`
	ShowErrors       bool   `json:"showErrors" url:"showErrors"`
	ShowFailed       bool   `json:"showFailed" url:"showFailed"`
	ShowMessageIDs   bool   `json:"showMessageIDs" url:"showMessageIDs"`
	ShowOpened       bool   `json:"showOpened" url:"showOpened"`
	ShowPending      bool   `json:"showPending" url:"showPending"`
	ShowSent         bool   `json:"showSent" url:"showSent"`
	ShowUnsubscribed bool   `json:"showUnsubscribed" url:"showUnsubscribed"`
}

// Get email batch status
// Access Level required: ViewReports
func (m *ElasticEmailImpl) GetEmailStatus(params GetEmailStatusParams) (status *EmailJobStatus, err error) {

	url := fmt.Sprintf("%s/%s/%s", m.apiBase, emailEndpoint, mGetEmailStatus)

	out := EmailJobStatus{}
	_, err = sendGetResp(m, url, params, &out)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Result:\n%+v\n", out)

	return &out, nil
}
