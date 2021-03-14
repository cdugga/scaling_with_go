package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)


const (
	PORT=8081
)

func main(){
	RedisTest()
	StartServer()
}

var ctx = context.Background()

func RedisConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis-master:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}


func RedisTest(){

	log.Println("Connecting to Redis")
	rdb := RedisConnection()

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

func AddKeyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello from the other side", r.URL.Path)
}


func FetchKeyHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	key, err := vars["key"]

	if err != true {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	log.Printf("Fetching key %s" , key)

	val, err1 := RedisConnection().Get(ctx, key).Result()
	if err1 != nil {
		panic(err1)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Requested value ", val)
}

func StartServer() {

	log.Printf("Staring server on port %d ", PORT )
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", AddKeyHandler)


	getKeyRouter := sm.Methods(http.MethodGet).Subrouter()
	getKeyRouter.HandleFunc("/value/{key}", FetchKeyHandler)

	//CORS header
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	s := http.Server{
		Addr: fmt.Sprint(":" , PORT),
		Handler: ch(sm),
		IdleTimeout: 120*time.Second,
		ReadTimeout: 5 *time.Second,
		WriteTimeout: 10*time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	log.Println("Recieved terminate ,graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown( tc)

}


