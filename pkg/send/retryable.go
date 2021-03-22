package send

import (
	"log"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

type HTTPRequest interface{
	Call (req *retryablehttp.Request) (*http.Response,error)
}

func NewRetryableRequest(client *retryablehttp.Client, log *log.Logger) *RetryableRequest{
	return &RetryableRequest{client: client,log: log}
}

type RetryableRequest struct{
	client *retryablehttp.Client
	log *log.Logger
}

func  (r *RetryableRequest) Call( req *retryablehttp.Request) (*http.Response,error){
	r.client.Logger = r.log
	response, err := r.client.Do(req)

	if err != nil {
		return nil, err
	}

	return response, nil
}