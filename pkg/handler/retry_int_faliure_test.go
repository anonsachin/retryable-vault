// +build !unit_test

package handler

import (
	"log"
	"main/pkg/send"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
)

func TestGetIntFailure(t *testing.T){
	tt := []struct{
		Name string
		StatusCode int
	}{
		{
			Name: "Proper Case",
			StatusCode: 502,
		},
	}

	for _, tc := range tt{
		t.Run(tc.Name, func(t *testing.T){
			// Setting up request and recorder
			rw := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "http://localhost:9909/volume", nil)

			client := retryablehttp.NewClient()
			client.RetryMax = 1
			l := log.New(os.Stdout,"-- unittest --",log.LstdFlags)

			c := send.NewRetryableRequest(client,l)

			h := NewRetry(l,c)

			h.Get(rw,req)
			resp := rw.Result()

			if resp.StatusCode != tc.StatusCode {
				t.Fatalf("Expected status code %v but got %v",tc.StatusCode,resp.StatusCode)
			}
			
		})
	}
}