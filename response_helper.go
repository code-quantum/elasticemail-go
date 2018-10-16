package elasticemail

import (
	"encoding/json"
	"errors"
	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
	"net/http"
)

func sendGetResp(m *ElasticEmailImpl, url string, params interface{}, out interface{}) (result *interface{}, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	v, _ := query.Values(params)
	v.Add("apikey", m.apiKey)
	req.URL.RawQuery = v.Encode()

	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {

		defer resp.Body.Close()

		f, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		dd := make(map[string]interface{})
		json.Unmarshal(f, &dd)

		log.Printf("DD:\n%+v\n", dd)

		success, ok := dd["success"]
		if !ok {
			return nil, errors.New("json response does not contain success field")
		}

		if success == true {
			data, ok := dd["data"]
			if !ok {
				return nil, errors.New("json response contain success true, but data field does not exists")
			}

			err := mapstructure.Decode(data, out)
			if err != nil {
				return nil, err
			}

			log.Printf("OUT:\n%+v\n", out)

			return &out, nil
		} else {
			err = errors.New(dd["error"].(string))
			return nil, err
		}
	}

	err = errors.New(resp.Status)
	return nil, err
}