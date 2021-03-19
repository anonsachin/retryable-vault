package send

import (
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

func Call(client *retryablehttp.Client, req *retryablehttp.Request) (*http.Response,error){
	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return response, nil
}