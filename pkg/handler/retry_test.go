package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/pkg/mocks"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
)


func TestGet(t *testing.T){
	tt := []struct{
		Name string
		Message string
		StatusCode int
		MockFunc func (req *retryablehttp.Request) (*http.Response,error)
		err bool
	}{
		{
			Name: "Proper Case",
			StatusCode: 200,
			Message: "Test",
			MockFunc:func (req *retryablehttp.Request) (*http.Response, error){
				b := strings.NewReader("Test")

				return &http.Response{
					StatusCode: 200,
					Body: ioutil.NopCloser(b),
				}, nil
			} ,
			err: false,
		},
		{
			Name: "Error Case",
			Message: "Just error",
			StatusCode: http.StatusBadGateway,
			MockFunc:func (req *retryablehttp.Request) (*http.Response, error){
				return nil, fmt.Errorf("Just error")
			} ,
			err: true,
		},
	}

	for _, tc := range tt{
		t.Run(tc.Name, func(t *testing.T){
			mock := &mocks.HTTPRequestMock{}
			mock.CallFunc = tc.MockFunc
			// Setting up request and recorder
			rw := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "http://localhost:9909/volume", nil)

			h := NewRetry(log.New(os.Stdout,"-- unittest --",log.LstdFlags),mock)

			h.Get(rw,req)
			resp := rw.Result()

			if resp.StatusCode != tc.StatusCode {
				t.Fatalf("Expected status code %v but got %v",tc.StatusCode,resp.StatusCode)
			}
			if tc.err {
				str := strings.Split(rw.Body.String(),":")
				test := strings.TrimSpace(str[1])
				if test != tc.Message {
					t.Fatalf("Expected Message %v but got %v",tc.Message,test)
				}
			} else {
				if rw.Body.String() != tc.Message{
					t.Fatalf("Expected Message %v but got %v",tc.Message,rw.Body.String())
				}
			}
			
		})
	}
}


// func TestMakeKV(t *testing.T){
// }
// tt := 