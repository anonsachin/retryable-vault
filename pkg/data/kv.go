package data

import (
	"io"

	"encoding/json"
)

type KV struct{
	Type string `json:"type"`
	Option Options `json:"options"`
}

type Options struct {
	Version int `json:"version"`
}

func (k *KV) ToJSON(w io.Writer) error{
	e := json.NewEncoder(w)
	return e.Encode(k)
}

func (k *KV) ToBytes() ([]byte,error){
	return json.Marshal(k)
}

func (k *KV) FromJSON(r io.Reader) error{
	d := json.NewDecoder(r)
	return d.Decode(k)
}