package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// 练习 7.12： 修改/list的handler让它把输出打印成一个HTML的表格而不是文本。html/template包（§4.6）可能会对你有帮助。


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

var itemsTemplate = template.Must(template.New("itemslist").Parse(`
<h1>{{len .}} items</h1>
<table>
<tr style='text-align: left'>
  <th>item</th>
  <th>price</th>
</tr>
{{range $key, $value := .}}
<tr>
  <td>{{$key}}</td>
  <td>{{$value}}</td>
</tr>
{{end}}
</table>
`))

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
	if err := itemsTemplate.Execute(w, db); err != nil {
		log.Fatal(err)
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
