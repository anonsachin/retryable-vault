package send

import (
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

type RetryableClient interface{
	Do(req *retryablehttp.Request) (*http.Response, error)
}