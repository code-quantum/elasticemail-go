package elasticemail

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/structs"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"reflect"
)

const (
	contactEndpoint = "contact"
	mAddContact     = "add"
)

type AddContactParams struct {
	Email               string   `json:"email" url:"email"`
	PublicAccountID     string   `json:"publicaccountid" url:"publicaccountid"`
	ActivationReturnUrl string   `json:"activationreturnurl" url:"activationreturnurl"`
	ActivationTemplate  string   `json:"activationtemplate" url:"activationtemplate"`
	AlreadyActiveUrl    string   `json:"alreadyactiveurl" url:"alreadyactiveurl"`
	ConsentDate         string   `json:"consentdate" url:"consentdate"`
	ConsentIP           string   `json:"consentip" url:"consentip"`
	ConsentTracking     int      `json:"consenttracking" url:"consenttracking"` // ConsentTracking Enumeration
	Field               []string `json:"field" url:"field"`
	FirstName           string   `json:"firstname" url:"firstname"`
	LastName            string   `json:"lastname" url:"lastname"`
	ListName            []string `json:"listname" url:"listname"`
	NotifyEmail         string   `json:"notifyemail" url:"notifyemail"`
	PublicListID        []string `json:"publiclistid" url:"publiclistid"`
	ReturnUrl           string   `json:"returnurl" url:"returnurl"`
	SendActivation      bool     `json:"sendactivation" url:"sendactivation"`
	Source              int      `json:"source" url:"source"` // ContactSource Enumeration
	SourceUrl           string   `json:"sourceurl" url:"sourceurl"`
}

// Add a new contact and optionally to one of your lists. Note that your API KEY is not required for this call.
// return nil if success
func (m *ElasticEmailImpl) AddContact(params AddContactParams) (err error) {

	url := fmt.Sprintf("%s/%s/%s", m.apiBase, contactEndpoint, mAddContact)

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	contentType := bodyWriter.FormDataContentType()
	paramsMap := structs.Map(params)
	for k, v := range paramsMap {
		switch x := v.(type) {
		case []string:
			for _, mv := range x {
				err = bodyWriter.WriteField(k, mv)
				if err != nil {
					return err
				}
			}
		case bool:
			if x {
				err = bodyWriter.WriteField(k, "true")
				if err != nil {
					return err
				}
			} else {
				err = bodyWriter.WriteField(k, "false")
				if err != nil {
					return err
				}
			}
		case int, int8:
			if str, ok := x.(string); ok {
				err = bodyWriter.WriteField(k, str)
				if err != nil {
					return err
				}
			}

		case string:
			err = bodyWriter.WriteField(k, x)
			if err != nil {
				return err
			}

		default:
			value, ok := v.(string)
			if !ok {
				return errors.New(fmt.Sprintf("not provided param type: %s : %s  (%s)", k, v, reflect.TypeOf(v)))
			}
			err = bodyWriter.WriteField(k, value)
			if err != nil {
				return err
			}
		}
	}
	bodyWriter.Close()

	req, err := http.NewRequest("POST", url, bodyBuf)
	if err != nil {
		return err
	}
	req.Header.Add("content-type", contentType)

	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		f, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		dd := make(map[string]interface{})
		json.Unmarshal(f, &dd)

		success, ok := dd["success"]
		if !ok {
			return errors.New("json response does not contain success field")
		}

		if success == true {

			return nil
		} else {
			errorMsg, ok := dd["error"]
			if !ok {
				return errors.New("json response does not contain 'error' field")
			}

			err = errors.New(errorMsg.(string))
			return err
		}
	} else {
		err = errors.New(resp.Status)
		return err
	}
}
