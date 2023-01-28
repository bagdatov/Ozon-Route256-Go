package resty

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/config"
	"reflect"
)

// cli is client for chgk website,
// implements usecase.Parser interface
type cli struct {
	resty *resty.Client
}

const (
	root       = "tour/xml"
	tourPrefix = "/tour/"
	postfix    = "/xml"
)

// New is a constructor for parser
func New(c *config.Config) (*cli, error) {
	if isNil(c) {
		return nil, fmt.Errorf("config is nil")
	}
	r := resty.New()
	r.SetBaseURL(c.ChgkBase.Url)
	r.SetTimeout(c.ChgkBase.Timeout)

	return &cli{
		resty: r,
	}, nil
}

// isNil is using reflect
// package to determine rather input is nil
func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
