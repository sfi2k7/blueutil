package blueutil

import (
	"bytes"
	"strings"

	"compress/gzip"

	"encoding/base64"

	"io/ioutil"

	"github.com/satori/go.uuid"
)

//NewV4 - get clean UUID without dashes
//
func NewV4() string {
	v4, _ := uuid.NewV4()
	return strings.Replace(v4.String(), "-", "", -1)
}

func CompressString(s string) (string, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, err := w.Write([]byte(s))
	if err != nil {
		return "", err
	}

	err = w.Flush()
	if err != nil {
		return "", err
	}

	err = w.Close()
	if err != nil {
		return "", err
	}

	str := base64.StdEncoding.EncodeToString(buf.Bytes())
	return str, nil
}

func UncompressString(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	rdata := bytes.NewReader(data)
	r, err := gzip.NewReader(rdata)
	if err != nil {
		return "", err
	}

	out, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
