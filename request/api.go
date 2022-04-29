package request

import (
	"errors"
	"strconv"
)

const (
	//Metaphorpsum API:  http://metaphorpsum.com/sentences/3
	Metaphorpsum   = "http://metaphorpsum.com/sentences/3"
	Itsthisforthat = "https://itsthisforthat.com/api.php"

//Itsthisforthat API https://itsthisforthat.com/api.php?text

)

var (
	MetaphorpsumMethod   *Method = NewMethod(Metaphorpsum)
	ItsthisforthatMethod *Method = func() *Method {

		m := NewMethod(Itsthisforthat)
		m.AddQueryParam("text", "")
		return m
	}()
)

type (
	Daily struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}
)

func NewDaily() *Daily {
	return &Daily{}
}

//GetSentence
func (d *Daily) GetSentence(m *Method) error {

	if m.URL == "" {
		return errors.New("url is nul")
	}

	response, statusCode, err := m.Get()
	if err != nil {
		return err
	}

	if statusCode != 200 {
		d.Status = statusCode
		return errors.New("http status code: " + strconv.Itoa(statusCode))
	}

	d.Message = string(response)
	d.Status = 200
	return nil
}
