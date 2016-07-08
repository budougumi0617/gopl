// Copyright 2016 budougumi0617 All Rights Reserved.
package main

// URL http://xkcd.com
const URL = "http://xkcd.com/"

// JSONFILE name of JSON in http://xkcd.com
const JSONFILE = "/info.0.json"

// ComicSearchResult result
type ComicSearchResult struct {
	TotalCount int
	Items      []*Comic
}

// Comic JSON structure for http://xkcd.com
type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	Transcript string
	SafeTitle  string `json:"safe_title"`
	Alt        string
	Image      string
	Day        string
}
