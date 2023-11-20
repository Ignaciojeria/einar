package resty

import (
	"github.com/go-resty/resty/v2"
)

var Client *resty.Client

func init() {
	Client = resty.New()
}
