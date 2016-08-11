// Copyright 2016 budougumi0617 All Rights Reserved.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

// Package defines result of "go list -json".
type Package struct {
	Dir        string // directory containing package sources
	ImportPath string // import path of package in dir
	Name       string // package name
	Root       string // Go root or Go path dir containing this package

	// Dependency information
	Imports []string // import paths used by this package
	Deps    []string // all (recursively) imported dependencies
}

func getpackages(targetPaths []string) []string {
	var result []string
	for _, path := range targetPaths {
		args := []string{"list", "-json"}
		args = append(args, path)
		out, err := exec.Command("go", args...).Output()
		if err != nil {
			fmt.Printf("Error %v\n", err)
			return result
		}
		var pkg Package
		if err = json.NewDecoder(bytes.NewReader(out)).Decode(&pkg); err != nil {
			fmt.Printf("Error %v\n", err)
			return result
		}
		result = append(result, pkg.Deps...)
	}
	return result
}

func main() {

	pkgs := getpackages(os.Args[1:])
	args := []string{"list", `-f={{join .Deps " "}}`}
	args = append(args, pkgs...)
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	gots := strings.Fields(string(out))
	pkgs = append(pkgs, gots...)
	targets := make(map[string]bool)
	for _, pkg := range pkgs {
		targets[pkg] = true
	}
	var results []string
	for p := range targets {
		results = append(results, p)
	}
	sort.Strings(results)
	fmt.Println(results)
}
