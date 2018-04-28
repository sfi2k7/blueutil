package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

const (
	contentType = "application/json"
)

type Client struct {
	Headers     map[string]string
	ContentType string
	r           *http.Request
}

func (cl *Client) addHeaders(r *http.Request) {
	if cl.Headers != nil {
		for k, v := range cl.Headers {
			r.Header.Add(k, v)
		}
	}
}

func (cl *Client) Get(url string) ([]byte, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	cl.addHeaders(req)
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (cl *Client) Post(url string, data []byte) ([]byte, error) {
	c := &http.Client{}
	rdr := bytes.NewReader(data)
	req, err := http.NewRequest("POST", url, rdr)
	if err != nil {
		return nil, err
	}
	cl.addHeaders(req)
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (cl *Client) Put(url string, data []byte) ([]byte, error) {
	c := &http.Client{}
	rdr := bytes.NewReader(data)
	req, err := http.NewRequest("PUT", url, rdr)
	if err != nil {
		return nil, err
	}
	cl.addHeaders(req)
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
