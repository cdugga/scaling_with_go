package main

import "testing"

func TestLoadHtml(t *testing.T){

	r := LoadPage()

	if r.Body == nil {
		t.Error("Expected content, return nothing")
	}

	t.Log("Response", r.StatusCode)

}

func TestScrapeHTML(t *testing.T) {

	ScrapeHTML()

}