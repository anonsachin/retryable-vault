// +build integration

package handler

import (
	"log"
	"main/pkg/handler"
	"main/pkg/send"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
)

func TestGetInt(t *testing.T) {
	tt := []struct {
		Name       string
		StatusCode int
		err        bool
	}{
		{
			Name:       "Proper Case",
			StatusCode: 200,
			err:        false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			// Setting up request and recorder
			rw := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "http://localhost:9909/volume", nil)

			client := retryablehttp.NewClient()
			l := log.New(os.Stdout, "-- unittest --", log.LstdFlags)

			c := send.NewRetryableRequest(client, l)

			h := handler.NewRetry(l, c)

			h.Get(rw, req)
			resp := rw.Result()

			if resp.StatusCode != tc.StatusCode {
				t.Fatalf("Expected status code %v but got %v", tc.StatusCode, resp.StatusCode)
			}

		})
	}
}
