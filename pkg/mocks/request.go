package mocks

import (
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

type HTTPRequestMock struct {
	CallFunc func (req *retryablehttp.Request) (*http.Response,error)
}

func (h *HTTPRequestMock) Call (req *retryablehttp.Request) (*http.Response,error){
	return h.CallFunc(req)
}
