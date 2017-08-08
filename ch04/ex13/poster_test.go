// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	var tests = []struct {
		args     []string
		filename string
	}{
		{[]string{"poster", "tron"}, "./TRON.jpg"},
	}
	for _, test := range tests {
		t.Skip()
		os.Args = test.args
		main()
		if _, err := os.Stat(test.filename); err != nil {
			t.Errorf("%v is not exist, %v", test.filename, err)
		}
	}
}

func TestSearchPoster(t *testing.T) {
	var tests = []struct {
		keyword  []string
		expected Poster
	}{
	//		{[]string{"tron"}, Poster{"TRON", "http://ia.media-imdb.com/images/M/MV5BMTY0OTM4ODM2MF5BMl5BanBnXkFtZTgwMTI0NDIxMDE@._V1_SX300.jpg"}},
	}
	for _, test := range tests {
		t.Skip()
		result, err := SearchPoster(test.keyword)
		if err != nil {
			t.Errorf("%v", err)
		} else if result.Poster != test.expected.Poster {
			t.Errorf("Result:%q, Expected:%q", result.Poster, test.expected.Poster)
		} else if result.Title != test.expected.Title {
			t.Errorf("Result:%q, Expected:%q", result.Title, test.expected.Title)
		}
	}
}

func TestSavePoster(t *testing.T) {
	var tests = []struct {
		title string
		url   string
		exist bool
	}{
		{"result", "https://github.com/budougumi0617/zsh-presentation/blob/master/image/display.jpg", true},
		{"result", "N/A", false},
		{"result", "https://github.com/budougumi0617/no_exist", false},
	}
	for _, test := range tests {
		t.Skip()
		p := Poster{test.title, test.url}
		SavePoster(p)
		if _, err := os.Stat("./" + test.title + ".jpg"); test.exist && err != nil {
			t.Errorf("%v is not exist, %v", test.title, err)
		}
	}
}
