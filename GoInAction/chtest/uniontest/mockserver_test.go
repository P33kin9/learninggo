package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockServer() *httptest.Server {
	sendJson := func(rw http.ResponseWriter, r *http.Request) {
		u := struct {
			Name string
		}{
			Name: "张三",
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(u)
	}
	return httptest.NewServer(http.HandlerFunc(sendJson))
}

func TestSendJSON(t *testing.T) {

	server := mockServer()
	defer server.Close()

	resq, err := http.Get(server.URL)
	if err != nil {
		t.Fatal("创建Get失败")
	}
	defer resq.Body.Close()

	log.Println("code:", resq.StatusCode)
	json, err := ioutil.ReadAll(resq.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("body:%s\n", json)
}
