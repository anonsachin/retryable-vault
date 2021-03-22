package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/pkg/data"
	"main/pkg/send"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)
const vault = "http://127.0.0.1:8200/v1/sys/audit"
const token = "myroot"
const kv = "http://127.0.0.1:8200/v1/sys/mounts/store"

type Retry struct{
	log *log.Logger
	client *retryablehttp.Client
}

func NewRetry(log *log.Logger, client *retryablehttp.Client) *Retry{
	return &Retry{log: log,client: client}
}

func (ret *Retry) Get(w http.ResponseWriter, r *http.Request){
	ret.client.Logger = ret.log
	req, err := retryablehttp.NewRequest(http.MethodGet,vault,nil)

	if err !=nil{
		message := fmt.Sprintf("[ERROR] Unable to create request : %s",err.Error())
		ret.log.Println(message)
		http.Error(w,message,http.StatusInternalServerError)
		return
	}

	req.WithContext(r.Context())
	req.Header.Add("X-Vault-Token",token)
	resp, err := send.Call(ret.client,req)

	if err != nil{
		message := fmt.Sprintf("[ERROR] Unable to make request : %s",err.Error())
		ret.log.Println(message)
		http.Error(w,message,http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(w,resp.Body)

	if err != nil{
		message := fmt.Sprintf("[ERROR] Unable to write request : %s",err.Error())
		ret.log.Println(message)
		http.Error(w,message,http.StatusBadGateway)
		return
	}
}

func (ret *Retry) MakeKV(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	input := &data.KV{}

	err := input.FromJSON(r.Body)

	if err != nil{
		msg := fmt.Sprintf("Unable to get request data: %v",err)
		ret.log.Println(msg)
		http.Error(w,msg,http.StatusBadRequest)
		return
	}
	b, err := json.Marshal(input)

	if err != nil{
		msg := fmt.Sprintf("Unable to get marshal vault data: %v",err)
		ret.log.Println(msg)
		http.Error(w,msg,http.StatusBadRequest)
		return
	}

	req, err := retryablehttp.NewRequest(http.MethodPost,kv,bytes.NewBuffer(b))

	if err != nil{
		msg := fmt.Sprintf("Unable to create request: %v",err)
		ret.log.Println(msg)
		http.Error(w,msg,http.StatusBadRequest)
		return
	}
	req.Header.Set("X-Vault-Token",token)
	req.Header.Set("Content-Type", "application/json")
	req.WithContext(r.Context())

	resp, err := send.Call(ret.client,req)

	if err != nil{
		message := fmt.Sprintf("[ERROR] Unable to make request : %s",err.Error())
		ret.log.Println(message)
		http.Error(w,message,http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(w,resp.Body)

	if err != nil{
		message := fmt.Sprintf("[ERROR] Unable to write request : %s",err.Error())
		ret.log.Println(message)
		http.Error(w,message,http.StatusBadGateway)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}