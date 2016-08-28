// Copyright 2016 budougumi0617 All Rights Reserved.

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 333.

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// InKey is included Key
type InKey struct {
	arrays [2]string
}

// Key is defined for check.
type Key struct {
	i       int
	structs InKey
}

func main() {
	m := make(map[Key]int)
	m[Key{1, InKey{[2]string{"foo", "bar"}}}] = 10
	m[Key{2, InKey{[2]string{"bar", "foo"}}}] = 20
	Display("TestMap", m)
}

//!+Display

// Display provides a means to display structured data.
func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

//!-Display

// formatAtom formats a value without inspecting its internal structure.
// It is a copy of the the function in gopl.io/ch11/format.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

//!+display
func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			switch key.Kind() {
			case reflect.Struct:
				fmt.Printf("%s[", path)
				for i := 0; i < key.NumField(); i++ {
					fieldPath := fmt.Sprintf("%s.%s", key.Type().String(), key.Type().Field(i).Name)
					display(fieldPath, key.Field(i))
				}
				display("]", v.MapIndex(key))
			case reflect.Array:
			case reflect.Slice:
				fmt.Printf("%s[", path)
				for i := 0; i < key.Len(); i++ {
					display(fmt.Sprintf("%s[%d]", key.Type().String(), i), key.Index(i))
				}
				display(key.Type().String(), key)
				display("]", v.MapIndex(key))
			default:
				display(fmt.Sprintf("%s[%s]", path,
					formatAtom(key)), v.MapIndex(key))
			}

		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

//!-display
