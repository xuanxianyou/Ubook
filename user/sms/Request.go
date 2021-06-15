package sms

import "encoding/json"

type Request struct {
	method     string
	bizContent []byte
}

func NewRequest() *Request {
	return &Request{}
}

func (r *Request) SetMethod(method string) *Request {
	r.method = method
	return r
}

func (r *Request) GetMethod() string {
	return r.method
}

func (r *Request) SetBizContent(bizContent interface{}) *Request {
	if bizContent != nil {
		r.bizContent, _ = json.Marshal(bizContent)
	}
	return r
}

func (r *Request) GetBizContent() string {
	return string(r.bizContent)
}
