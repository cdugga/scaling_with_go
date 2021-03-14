package main

import (
	"context"
	"encoding/json"
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
	StartServer()
}

var ctx = context.Background()

func RedisConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		//Addr:     "redis-master:6379",
		Addr: "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}


func WriteKeyHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Connecting to Redis")
	rdb := RedisConnection()

	keyvalue := r.Context().Value(KeyValue{}).(*KeyValue)

	err := rdb.Set(ctx, keyvalue.Key, keyvalue.Value, 0).Err()
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

	writeKeyRouter := sm.Methods(http.MethodPost).Subrouter()
	writeKeyRouter.HandleFunc("/key", WriteKeyHandler)
	writeKeyRouter.Use(MiddleWareProductValidation)

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

type KeyValue struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

func MiddleWareProductValidation(next http.Handler) http.Handler{

	return http.HandlerFunc(func(rw http.ResponseWriter, r*http.Request) {

		keyVal := &KeyValue{}

		e := json.NewDecoder(r.Body)
		err := e.Decode(keyVal)

		if err != nil {
			log.Println("[Error] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// add key value to the context
		ctx := context.WithValue(r.Context(), KeyValue{}, keyVal)
		req := r.WithContext(ctx)

		// call the next handler, which can be another middleware in the chain, or the final handler
		next.ServeHTTP(rw,req)

	})
}

