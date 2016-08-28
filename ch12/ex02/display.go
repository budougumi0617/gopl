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

// Maxdepth is limit of cyclic display.
var Maxdepth = 5

// InKey is included Key
type InKey struct {
	b bool
}

// Key is defined for check.
type Key struct {
	i       int
	structs InKey
	arrays  [2]string
}

// Cycle that points to itself.
type Cycle struct {
	Value int
	Tail  *Cycle
}

func main() {
	var c Cycle
	c = Cycle{42, &c}
	Display("TestCycle", c)
}

func premain() {
	m := make(map[Key]int)
	m[Key{1, InKey{true}, [2]string{"foo", "bar"}}] = 10
	Display("TestMap", m)
	ma := make(map[[2]string]int)
	ma[[2]string{"foo", "bar"}] = 10
	Display("TestMapArrays", ma)
}

//!+Display

// Display provides a means to display structured data.
func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), 0)
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
func display(path string, v reflect.Value, depth int) {
	if Maxdepth < depth {
		fmt.Printf("cyclic limit %d\n", Maxdepth)
		return
	}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), depth)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), depth)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			switch key.Kind() {
			case reflect.Struct:
				fmt.Printf("%s[", path)
				for i := 0; i < key.NumField(); i++ {
					fieldPath := fmt.Sprintf("%s.%s", key.Type().String(), key.Type().Field(i).Name)
					display(fieldPath, key.Field(i), depth)
				}
				display("]", v.MapIndex(key), depth)
			case reflect.Array, reflect.Slice:
				fmt.Printf("%s[", path)
				display(key.Type().String(), key, depth)
				display("]", v.MapIndex(key), depth)
			default:
				display(fmt.Sprintf("%s[%s]", path,
					formatAtom(key)), v.MapIndex(key), depth)
			}

		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			// poor logic...
			d := depth
			// compare address.
			if fmt.Sprintf("%v", v) == fmt.Sprintf("%v", v.Elem().Addr()) {
				d++
			}
			display(fmt.Sprintf("(*%s)", path), v.Elem(), d)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), depth)
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

//!-display
