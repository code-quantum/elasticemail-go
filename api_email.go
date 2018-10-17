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

type ViewParams struct {
	MessageID string `json:"messageID" url:"messageID"` // Unique identifier for this email.
}

type SendParams struct {
	Attachments         []string `json:"attachments"`
	BodyHtml            string   `json:"bodyhtml"`
	BodyText            string   `json:"bodytext"`
	To                  []string `json:"to"`
	Channel             string   `json:"channel"`
	Charset             string   `json:"charset"`
	CharsetBodyHtml     string   `json:"charsetbodyhtml"`
	CharsetBodyText     string   `json:"charsetbodytext"`
	DataSource          string   `json:"datasource"`
	EncodingType        int8     `json:"encodingtype"` // EncodingType Enumeration
	From                string   `json:"from"`
	FromName            string   `json:"fromname"`
	Headers             []string `json:"headers"` // example: headers_xmailer = xmailer: header-value1  Whitespace required!
	IsTransactional     bool     `json:"istransactional"`
	Lists               []string `json:"lists"`
	Merge               string   `json:"merge"` // example: merge_firstname=John, ....
	MergeSourceFilename string   `json:"mergesourcefilename"`
	MsgBcc              []string `json:"msgbcc"`
	MsgCC               []string `json:"msgcc"`
	MsgFrom             string   `json:"msgfrom"`
	MsgFromName         string   `json:"msgfromname"`
	MsgTo               []string `json:"msgto"`
	PoolName            string   `json:"poolname"`
	PostBack            string   `json:"postback"`
	ReplyTo             string   `json:"replyto"`
	ReplyToName         string   `json:"replytoname"`
	Segments            []string `json:"sefments"`
	Sender              string   `json:"sender"`
	SenderName          string   `json:"sendername"`
	Subject             string   `json:"subject"`
	Template            string   `json:"template"`
	TimeOffSetMinutes   string   `json:"timeoffsetminutes"`
	TrackClicks         bool     `json:"trackclicks"`
	TrackOpens          bool     `json:"trackopens"`
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

	return &out, nil
}

// View email
// Access Level required: ViewReports
func (m *ElasticEmailImpl) View(params ViewParams) (status *EmailView, err error) {
	url := fmt.Sprintf("%s/%s/%s", m.apiBase, emailEndpoint, mView)

	out := EmailView{}
	err = sendGetResp(m, url, params, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

// Submit emails.
// The HTTP POST request is suggested.
// The default, maximum (accepted by us) size of an email is 10 MB in total, with or without attachments included.
// For suggested implementations please refer to https://elasticemail.com/support/http-api/
// Access Level required: SendHttp
func (m *ElasticEmailImpl) Send(params SendParams) (status *EmailSend, err error) {
	url := fmt.Sprintf("%s/%s/%s", m.apiBase, emailEndpoint, mSend)

	out := EmailSend{}
	err = sendPostResp(m, url, params, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
