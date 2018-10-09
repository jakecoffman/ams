package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`[{"text":"LEAVE NOW!"},{"text":"false alarm"}]`))
			return
		}

		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// TODO: send an alert
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		fmt.Println(string(data))
		w.WriteHeader(201)
		w.Header().Set("Conent-Type", "application/json")
		w.Write([]byte(`{"ok":"world"}`))
	})
	http.ListenAndServe("0.0.0.0:8081", nil)
}
