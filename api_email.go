package elasticemail

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"log"
	"net/http"
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
func (m *ElasticEmailImpl) GetEmailStatus(params GetEmailStatusParams) (status EmailJobStatus, err error) {

	url := fmt.Sprintf("%s/%s/%s", m.apiBase, emailEndpoint, mGetEmailStatus)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	v, _ := query.Values(params)
	v.Add("apikey", m.apiKey)
	req.URL.RawQuery = v.Encode()
	fmt.Println(req.URL.String())

	resp, _ := m.client.Do(req)

	log.Println(resp.Status)

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// jsonStr := string(f)

	fmt.Println(string(f))

	dd := make(map[string]interface{})
	json.Unmarshal(f, &dd)

	success, ok := dd["success"]
	if !ok {
		fmt.Printf("ok is false")
	}
	if success == true {
		fmt.Printf("!!!!!!!!!!!!!!!")
	}

	data, _ := dd["data"]

	original, ok := data.(EmailJobStatus)
	if ok {
		println(original.Status)
		fmt.Printf("ORIGINAL: \n%+v\n", original)
	} else {
		fmt.Printf("can not convert")
	}

	log.Printf("%v", dd)
	fmt.Printf("\n%+v\n", dd)

	return original, nil
}
