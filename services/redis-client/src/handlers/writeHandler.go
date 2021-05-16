package handlers

import (
	"fmt"
	"github.com/cdugga/scaling_with_go/redisclient/data"
	"log"
	"net/http"
)

type Env struct {
	db data.DataAccess
}


func WriteKeyHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Connecting to Redis")

	keyvalue := r.Context().Value(data.KeyValue{}).(*data.KeyValue)

	data := data.Datasource

	err := data.Set(keyvalue.Key, keyvalue.Value, 0)
	if err != nil {
		panic(err)
	}

	val, err := data.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}