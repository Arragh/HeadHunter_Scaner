package httphandler

import "hhscaner/configuration"

//go:generate mockgen -source=httpclient.go -destination=../../test/mock/httpclient_mock.go -package=mock

type HttpClient interface {
	Get(url string, params *[]configuration.UrlParameter) ([]byte, error)
}
