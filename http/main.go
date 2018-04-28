package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
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

func Post(url string, data []byte) ([]byte, error) {
	c := &http.Client{}
	rdr := bytes.NewReader(data)
	req, err := http.NewRequest("POST", url, rdr)
	if err != nil {
		return nil, err
	}
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

func Put(url string, data []byte) ([]byte, error) {
	c := &http.Client{}
	rdr := bytes.NewReader(data)
	req, err := http.NewRequest("PUT", url, rdr)
	if err != nil {
		return nil, err
	}
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
