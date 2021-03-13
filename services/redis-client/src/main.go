package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
)

func main(){
	StartServer()
	RedisConnection()

}

var ctx = context.Background()

func RedisConnection(){
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-master:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}

func StartServer() {

	fmt.Println("Starting redis client")

	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello from the other side", r.URL.Path)
}
