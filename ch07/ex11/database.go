// Copyright 2016 budougumi0617 All Rights Reserved.
// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "No such item: %q\n", item)
	}
}

// Exemple http://localhost:8000/create?item=test&price=30.0
func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if p, exist := db[item]; exist {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "Exsited item: %q\n", item)
		fmt.Fprintf(w, "%s\n", p)
	} else {
		d, _ := strconv.ParseFloat(price, 32)
		db[item] = dollars(d)
		fmt.Fprintf(w, "Added item: %q\n", item)
		fmt.Fprintf(w, "%s\n", price)
	}
}

// Exemple http://localhost:8000/read?item=test
func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, exist := db[item]; exist {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	} else {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "Not find item: %q\n", item)
	}
}

// Exemple http://localhost:8000/update?item=socks&price=6
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if old, exist := db[item]; exist {
		d, _ := strconv.ParseFloat(price, 32)
		db[item] = dollars(d)
		fmt.Fprintf(w, "Updated item %s: %f\n", item, db[item])
		fmt.Fprintf(w, "old value %s\n", old)
	} else {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "Not find item: %q\n", item)
	}
}

// Exemple http://localhost:8000/delete?item=socks
func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, exist := db[item]; exist {
		delete(db, item)
		fmt.Fprintf(w, "Deleted item: %q\n", item)
	} else {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "Not found item: %q\n", item)
	}
}
