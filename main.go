package blueutil

import (
	"strings"

	"github.com/satori/go.uuid"
)

func NewV4() string {
	v4 := uuid.NewV4().String()
	return strings.Replace(v4, "-", "", -1)
}
