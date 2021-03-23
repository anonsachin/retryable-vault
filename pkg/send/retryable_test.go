// +build unittest

package send

import (
	"fmt"
	"log"
	"main/pkg/mocks"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
)

func TestCall (t *testing.T){
	tt := []struct{
		Name string
		Resp *http.Response
		MockFunc func (req *retryablehttp.Request) (*http.Response, error)
		err bool
	}{
		{
			Name: "Proper Case",
			Resp: &http.Response{
				StatusCode: 200,
			},
			MockFunc:func (req *retryablehttp.Request) (*http.Response, error){
				return &http.Response{
					StatusCode: 200,
				}, nil
			} ,
			err: false,
		},
		{
			Name: "Error Case",
			Resp: &http.Response{
				StatusCode: 200,
			},
			MockFunc:func (req *retryablehttp.Request) (*http.Response, error){
				return nil, fmt.Errorf("Just error")
			} ,
			err: true,
		},
	}

	for _, tc := range tt{
		t.Run(tc.Name,func(t *testing.T){
			mock := &mocks.RetryableClientMock{}
			mock.DoFunc = tc.MockFunc
			r := &RetryableRequest{
				log: log.New(os.Stdout,"-- unittest --",log.LstdFlags),
				client: mock,
			}
			if tc.err{
				_, err := r.Call(&retryablehttp.Request{})
				if err == nil{
					t.Fatalf("Excected err but got nil")
				}
			} else {
				resp, err := r.Call(&retryablehttp.Request{})
				if err != nil {
					t.Fatalf("Excected nil but got err : %v",err)
				}
				if resp.StatusCode != tc.Resp.StatusCode {
					t.Fatalf("Excected %v but got : %v",tc.Resp.StatusCode,resp.StatusCode)
				}
			}
		})
	}
	
}