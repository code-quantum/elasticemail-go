package elasticemail

import (
	"fmt"
)

const (
	emailEndpoint = "email"     // Send your emails and see their statuses
	mGetStatus    = "getstatus" // Get email batch status
	mStatus       = "status"    // Detailed status of a unique email sent through your account
	mView         = "view"      // View email
	mSend         = "send"      // Submit emails
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

type StatusParams struct {
	MessageID string `json:"messageID" url:"messageID"` // Unique identifier for this email.
}

// Get email batch status
// Access Level required: ViewReports
func (m *ElasticEmailImpl) GetEmailStatus(params GetEmailStatusParams) (status *EmailJobStatus, err error) {

	url := fmt.Sprintf("%s/%s/%s", m.apiBase, emailEndpoint, mGetStatus)

	out := EmailJobStatus{}
	err = sendGetResp(m, url, params, &out)
	if err != nil {
		return nil, err
	}
	fmt.Printf("GetEmailStatus Result:\n%+v\n", out)

	return &out, nil
}

// Detailed status of a unique email sent through your account.
// Returns a 'Email has expired and the status is unknown.' error, if the email has not been fully processed yet.
// Access Level required: ViewReports
func (m *ElasticEmailImpl) Status(params StatusParams) (status *EmailStatus, err error) {
	url := fmt.Sprintf("%s/%s/%s", m.apiBase, emailEndpoint, mStatus)

	out := EmailStatus{}
	err = sendGetResp(m, url, params, &out)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Status result:\n%+v\n", out)

	return &out, nil
}
