package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)
// 练习 7.11： 增加额外的handler让客户端可以创建，读取，更新和删除数据库记录。例如，一个形如 /update?item=socks&price=6
// 的请求会更新库存清单里一个货品的价格并且当这个货品不存在或价格无效时返回一个错误值。（注意：这个修改会引入变量同时更新的问题）

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) create(w http.ResponseWriter, req *http.Request) {
	quryMap := req.URL.Query()
	item := quryMap.Get("item")
	price := quryMap.Get("price")
	if item == "" || price =="" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Price: %q or Item: %q is invalid.\n", price, item)
	}

	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusConflict) // 409
		fmt.Fprintf(w, "The item: %q already exists.\n", item)
	} else {
		price64, err := strconv.ParseFloat(price, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Price: %q is invalid.", price)
		} else {
			db[item] = dollars(price64)
			fmt.Fprintf(w, "%+v", db)
		}
	}
}

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
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "%+v", db)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	newPrice := req.URL.Query().Get("price")
	if price, ok := db[item]; ok {
		newPrice64 , err := strconv.ParseFloat(newPrice, 32)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Price: %q is invalid.\n", price)
		} else {
			db[item] = dollars(newPrice64)
			fmt.Fprintf(w, "%+v", db)
		}
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

