package elasticemail

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/structs"
	"github.com/google/go-querystring/query"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"reflect"
)

func sendGetResp(m *ElasticEmailImpl, url string, params interface{}, out interface{}) (err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	v, _ := query.Values(params)
	v.Add("apikey", m.apiKey)
	req.URL.RawQuery = v.Encode()

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
			data, ok := dd["data"]
			if !ok {
				return errors.New("json response contain success true, but data field does not exists")
			}

			err := mapstructure.Decode(data, out)
			if err != nil {
				return err
			}

			return nil
		} else {
			err = errors.New(dd["error"].(string))
			return err
		}
	}

	err = errors.New(resp.Status)
	return err
}

func sendPostResp(m *ElasticEmailImpl, url string, params interface{}, out interface{}) (err error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	contentType := bodyWriter.FormDataContentType()
	err = bodyWriter.WriteField("apikey", m.apiKey)
	if err != nil {
		return err
	}

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
			}

		case string:
			err = bodyWriter.WriteField(k, x)

		default:
			value, ok := v.(string)
			if !ok {
				return errors.New(fmt.Sprintf("not provided param type: %s : %s  (%s)", k, v, reflect.TypeOf(v)))
			}
			err = bodyWriter.WriteField(k, value)

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

	if resp.StatusCode == http.StatusOK {

		defer resp.Body.Close()

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
			data, ok := dd["data"]
			if !ok {
				return errors.New("json response contain success true, but data field does not exists")
			}

			err := mapstructure.Decode(data, out)
			if err != nil {
				return err
			}

			return nil
		} else {
			err = errors.New(dd["error"].(string))
			return err
		}
	}

	err = errors.New(resp.Status)
	return err
}
