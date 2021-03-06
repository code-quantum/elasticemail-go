package elasticemail

type successResponse struct {
	success string
	data    interface{}
}

type errorResponse struct {
	success string
	error   string
}

type EmailJobFailedStatus struct {
	Address   string `json:"address"`
	Category  string `json:"dategory"`
	Error     string `json:"error"`
	ErrorCode int    `json:"errorcode"`
}

type EmailJobStatus struct {
	AbuseReports      *[]string               `json:"abusereports"`
	AbuseReportsCount int                     `json:"abusereportscount"`
	Clicked           *[]string               `json:"clicked"`
	ClickedCount      int                     `json:"clickedcount"`
	Delivered         *[]string               `json:"delivered"`
	DeliveredCount    int                     `json:"deliveredcount"`
	Failed            *[]EmailJobFailedStatus `json:"failed"`
	FailedCount       int                     `json:"failedcount"`
	ID                string                  `json:"id"`
	MessageIDs        *[]string               `json:"messageids"`
	Opened            *[]string               `json:"opened"`
	OpenedCount       int                     `json:"openedcount"`
	Pending           *[]string               `json:"pending"`
	PendingCount      int                     `json:"pendingcount"`
	RecipientsCount   int                     `json:"recipientscount"`
	Sent              *[]string               `json:"sent"`
	SentCount         int                     `json:"sentcount"`
	Status            string                  `json:"status"`
	Unsubscribed      *[]string               `json:"unsubscribed"`
	UnsubscribedCount int                     `json:"unsubscribedcount"`
}

type EmailStatus struct {
	Date             string  `json:"date"`
	DateClicked      *string `json:"dateclicked"`
	DateOpened       *string `json:"dateopened"`
	DateSent         *string `json:"datesent"`
	ErrorMessage     string  `json:"errormessage"`
	From             string  `json:"from"`
	Status           uint8   `json:"status"` // value from LogJobStatus Enumeration
	StatusChangeDate string  `json:"statuschangedate"`
	StatusName       string  `json:"statusname"`
	To               string  `json:"to"`
	TransactionID    string  `json:"transactionid"`
}

type EmailView struct {
	Body    string `json:"body"`
	From    string `json:"from"`
	Subject string `json:"subject"`
}

type EmailSend struct {
	MessageID     string `json:"messageid"`
	TransactionID string `json:"transactionid"`
}

type List struct {
	AllowUnsubscribe bool   `json:"allowunsubscribe"`
	Count            int    `json:"count"`
	DateAdded        string `json:"dateadded"`
	ListID           int    `json:"listid"`
	ListName         string `json:"listname"`
	PublicListID     string `json:"publiclistid"`
	Rule             string `json:"rule"`
}

type Template struct {
	BodyHtm            string `json:"bodyhtml"`
	BodyText           string `json:"bodytext"`
	Css                string `json:"css"`
	DateAdded          string `json:"dateadded"`
	FromEmail          string `json:"fromemail"`
	FromName           string `json:"fromname"`
	Name               string `json:"name"`
	OriginalTemplateID int    `json:"originaltemplateid"`
	Subject            string `json:"subject"`
	TemplateID         int    `json:"templateid"`
	TemplateScope      int    `json:"templatescope"`
	TemplateType       int    `json:"templatetype"`
}

type TemplateList struct {
	DraftTemplate []Template `json:"drafttemplate"`
	Templates     []Template `json:"templates"`
}
