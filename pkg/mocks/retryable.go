package mocks

import (
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

type RetryableClientMock struct{
	DoFunc func (req *retryablehttp.Request) (*http.Response, error)
}

func (r *RetryableClientMock) Do(req *retryablehttp.Request) (*http.Response, error){
	return r.DoFunc(req)
}