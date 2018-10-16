package elasticemail

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"
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
func (m *ElasticEmailImpl) GetEmailStatus(params GetEmailStatusParams) (status *EmailJobStatus, err error) {

	url := fmt.Sprintf("%s/%s/%s", m.apiBase, emailEndpoint, mGetEmailStatus)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	v, _ := query.Values(params)
	v.Add("apikey", m.apiKey)
	req.URL.RawQuery = v.Encode()
	fmt.Println(req.URL.String())

	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	log.Println(resp.Status)

	if resp.StatusCode == http.StatusOK {

		defer resp.Body.Close()

		f, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		dd := make(map[string]interface{})
		json.Unmarshal(f, &dd)

		success, ok := dd["success"]
		if !ok {
			return nil, errors.New("json response does not contain success field")
		}

		if success == true {
			data, ok := dd["data"]
			if !ok {
				return nil, errors.New("json response contain success true, but data field does not exists")
			}

			st := EmailJobStatus{}
			err := mapstructure.Decode(data, &st)
			if err != nil {
				return nil, err
			}
			fmt.Printf("\n%+v\n", st)
			return &st, nil
		} else {
			err = errors.New(dd["error"].(string))
			return nil, err
		}
	}

	err = errors.New(resp.Status)
	return nil, err
}
