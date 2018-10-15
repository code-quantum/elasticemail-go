package elasticemail

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
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
func (m *ElasticEmailImpl) GetEmailStatus(params GetEmailStatusParams) {

	req, err := m.client.NewRequest("GET", m.apiBase, nil)
	if err != nil {
		log.Fatalln(err)
	}
	v, _ := query.Values(params)
	v.Add("apikey", m.apiKey)
	req.URL.RawQuery = v.Encode()
	fmt.Println(req.URL.String())

	resp, _ := m.client.Do(r)

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(f))
}
