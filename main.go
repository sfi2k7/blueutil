package blueutil

import (
	"bytes"
	"strings"

	"gopkg.in/mgo.v2"

	"compress/gzip"

	"encoding/base64"

	"io/ioutil"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

//NewV4 - get clean UUID without dashes
//
func NewV4() string {
	v4, _ := uuid.NewV4()
	return strings.Replace(v4.String(), "-", "", -1)
}

func ID4() string {
	v4 := NewV4()
	return v4[0:4]
}

func ID6() string {
	v4 := NewV4()
	return v4[0:6]
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

func MoveCollection(source, target *mgo.Database, cn string) error {
	count, err := source.C(cn).Count()
	if err != nil {
		return err
	}
	limit := 1000
	skip := 0

	for {
		var all []interface{}
		err = source.C(cn).Find(bson.M{}).Skip(skip).Limit(limit).All(&all)
		if err != nil {
			return err
		}
		if len(all) == 0 {
			return nil
			// return errors.New("Done Copying - 0 in Source")
		}
		skip += len(all)
		err = target.C(cn).Insert(all...)
		if err != nil {
			return err
		}
		if skip == count {
			return nil
		}
	}
}

func MoveCollection2(source, target *mgo.Database, cn string, tcn string) error {
	count, err := source.C(cn).Count()
	if err != nil {
		return err
	}
	limit := 1000
	skip := 0

	for {
		var all []interface{}
		err = source.C(cn).Find(bson.M{}).Skip(skip).Limit(limit).All(&all)
		if err != nil {
			return err
		}
		if len(all) == 0 {
			return nil
			// return errors.New("Done Copying - 0 in Source")
		}
		skip += len(all)
		err = target.C(tcn).Insert(all...)
		if err != nil {
			return err
		}
		if skip == count {
			return nil
		}
	}
}
