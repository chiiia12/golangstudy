package main

import (
	"testing"
	"net/http"
	"net/url"
)

func TestUnpack1(t *testing.T) {
	var data struct {
		Labels     []string `http:"label"`
		MaxResults int      `http:"maxvalue"`
		Exact      bool     `http:"exact"`
		Email      string   `http:"email"`
	}
	var req http.Request
	url1, _ := url.Parse("http://localhost:12345/search?label=hoge&maxvalue=4&exact=true&email=hoge")
	req.URL = url1
	if err := Unpack(&req, &data); err == nil {
		t.Errorf("expect error occured.but it doesn't occur.")
	}
}
func TestUnpack2(t *testing.T) {
	var data struct {
		Labels     []string `http:"label"`
		MaxResults int      `http:"maxvalue"`
		Exact      bool     `http:"exact"`
		Email      string   `http:"email"`
	}
	var req2 http.Request
	url2, _ := url.Parse("http://localhost:12345/search?label=hoge&maxvalue=4&exact=true&email=fuga@fuga.com")
	req2.URL = url2
	if err := Unpack(&req2, &data); err != nil {
		t.Errorf("error occured.%v", err)
	}

}
