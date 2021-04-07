package main

import (
	"goination/chtest/uniontest/common"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	common.Routes()
}

func TestSendJSON(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/sendjson", nil)
	if err != nil {
		t.Fatal("create Request failure!")
	}

	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)

	log.Println("code:", rw.Code)

	log.Println("body:", rw.Body.String())
}
