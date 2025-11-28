package httphandler

import "hhscaner/configuration"

type HttpClient interface {
	Get(url string, params *[]configuration.UrlParameter) ([]byte, error)
}
