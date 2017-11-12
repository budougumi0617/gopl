// Copyright 2017 budougumi0617 All Rights Reserved.

package params

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/budougumi0617/gopl/ch12/ex11/params"
)

func TestPack(t *testing.T) {
	s := struct {
		Name string `http:"n"`
		Age  int    `http:"a"`
	}{"Arugula", 35}
	u, err := Pack(&s)
	if err != nil {
		t.Errorf("Pack(%#v): %s", s, err)
	}
	want := "a=35&n=Arugula"
	got := u.RawQuery
	if got != want {
		t.Errorf("Pack(%#v): got %q, want %q", s, got, want)
	}
}

func TestUnpackWithRule(t *testing.T) {
	type Data struct {
		// email needs email address format as a XXXX@YYYY
		Email string `http:"m" rule:"email"`
	}

	for _, test := range []struct {
		url  string
		err  bool
		data Data
	}{
		{`http://localhost:3000/search?mm=`, false, Data{""}},
		{`http://localhost:3000/search?m=success@gmail.com`, false,
			Data{Email: "success@gmail.com"}},
		{`http://localhost:3000/search?m=@`, true, Data{}},
		{`http://localhost:false/search?m=failed@`, true, Data{}},
		{`http://localhost:flase/search?m=@gmail.com`, true, Data{}},
	} {
		var data Data

		req, err := newRequest(test.url)
		if err != nil {
			t.Errorf("Parse: %v\n", err)
			continue
		}

		if err := params.Unpack(req, &data); err != nil {
			if test.err == true {
				continue
			}

			t.Errorf("Unpack: %v\n", err)
			continue
		}

		if !reflect.DeepEqual(data, test.data) {
			t.Errorf("%q => \n%+v but want %+v\n", test.url, data, test.data)
		}
	}
}

func newRequest(rawurl string) (*http.Request, error) {
	var req http.Request
	url, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	req.URL = url
	return &req, nil
}
