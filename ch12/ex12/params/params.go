// Copyright 2017 budougumi0617 All Rights Reserved.

package params

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Pack build struct by URL parameters.
func Pack(ptr interface{}) (url.URL, error) {
	v := reflect.ValueOf(ptr).Elem()
	if v.Type().Kind() != reflect.Struct {
		return url.URL{}, fmt.Errorf("pack: %v is not a struct", ptr)
	}
	vals := &url.Values{}
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		vals.Add(name, fmt.Sprintf("%v", v.Field(i)))
	}
	return url.URL{RawQuery: vals.Encode()}, nil
}

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	rules := make(map[string]string)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")        // Get value by "http" key
		if name == "" {                // Has not "http" key
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
		if rule := tag.Get("rule"); rule != "" {
			log.Printf("got %v\n", rule)
			rules[name] = rule
		}
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if rules[name] != "" {
				if !validate(rules[name], value) {
					return fmt.Errorf("%s invalid for %s", value, name)
				}
			}
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)
	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

var emailPattern = regexp.MustCompile(`^[a-zA-Z0-9\-.]+@[a-zA-Z0-9\-.]+$`)

func validate(rule, value string) bool {
	switch rule {
	case "email":
		return emailPattern.MatchString(value)
	default:
		return false
	}
}
