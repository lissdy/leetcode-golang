package main

import (
	"net/http"
	"net/http/httptest"
	"bytes"
	"io/ioutil"
	"fmt"
)

type HandlerFunc func(w http.ResponseWriter, response *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request){
	f(w,r)
}

func main(){
	hf := HandlerFunc(HelloHandler)
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET","/",bytes.NewBuffer([]byte("test")))
	hf.ServeHTTP(resp,req)
	bts,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bts))
}

func HelloHandler(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("Hello World"))
}