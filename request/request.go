package request

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type (
	Method struct {
		URL     string      `json:"url"`
		Header  http.Header `json:"header"`
		Query   url.Values
		Timeout int
	}
)

//NewMethod new Method structure
//@u url string
func NewMethod(u string) *Method {

	return &Method{URL: u, Timeout: 10, Query: url.Values{}}
}

//AddQueryParam
func (m *Method) AddQueryParam(key, value string) {
	m.Query.Add(key, value)
}

//Get
func (m *Method) Get() ([]byte, int, error) {

	client := http.Client{Timeout: time.Duration(m.Timeout) * time.Second}

	req, err := http.NewRequest(http.MethodGet, m.URL, nil)
	if err != nil {

		return nil, 0, err
	}

	req.Header = m.Header
	req.URL.RawQuery = m.Query.Encode()
	response, err := client.Do(req)
	if err != nil {

		return nil, 0, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {

		return nil, 0, err
	}

	return body, response.StatusCode, nil
}
