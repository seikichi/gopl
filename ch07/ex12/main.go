package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"text/template"
)

var mu sync.RWMutex

func main() {
	mu = sync.RWMutex{}
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)

	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/destroy", db.destroy)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var listTable = template.Must(template.New("listtable").Parse(`
<h1>Database</h1>
<table>
<tr style='text-align: left'>
  <th>Item</th>
  <th>Price</th>
</tr>
{{range $key, $value := . }}
<tr>
  <td>{{ $key }}</td>
  <td>{{ $value }}</td>
</tr>
{{end}}
</table>
`))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	listTable.Execute(w, db)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	item := req.URL.Query().Get("item")

	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "already exists: %q\n", item)
		return
	}

	price := req.URL.Query().Get("price")
	floatPrice, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}

	db[item] = dollars(floatPrice)
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	price := req.URL.Query().Get("price")
	floatPrice, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}

	db[item] = dollars(floatPrice)
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) destroy(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	item := req.URL.Query().Get("item")
	delete(db, item)
}
