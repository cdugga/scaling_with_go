package http

import (
	"context"
	"fmt"
	"github.com/cdugga/bookmark/env"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"time"
)

type MuxRouter struct {
	router *mux.Router
}

var (
	Env env.Provider = env.NewEnv()
)

func NewMuxRouter() Router {
	return &MuxRouter{router: mux.NewRouter()}
}

func (mx *MuxRouter) Get(url string, f func(w http.ResponseWriter, r *http.Request)) {
	mx.router.HandleFunc(url, f).Methods("GET")
}

func (mx *MuxRouter) GetWithQueryParams(url string, f func(w http.ResponseWriter, r *http.Request), queryParam string) {
	mx.router.HandleFunc(url, f).Queries(queryParam, "{"+queryParam+"}").Methods("GET")
}

func (mx *MuxRouter) Post(url string, f func(w http.ResponseWriter, r *http.Request)) {
	mx.router.HandleFunc(url, f).Methods("POST")
}

func (mx *MuxRouter) RegisterSubRoute(path string) Router {
	subRouter := mx.router.PathPrefix(path).Subrouter()
	return &MuxRouter{router: subRouter}
}

func (mx *MuxRouter) Serve() {
	s := http.Server{
		Addr: ":8080",
		Handler:  mx.router,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 5 *time.Second,
		WriteTimeout: 10*time.Second,
	}

	go func() {
		fmt.Println("Starting host..", Env.Get("HOST"))
		fmt.Printf("Server starting on %v:%v", Env.Get("HOST"), Env.Get("PORT"))
		fmt.Println()

		err := s.ListenAndServe()
		if err != nil {
			log.Fatal("Error starting server on port 8080")
			os.Exit(1)
		}

	}()

	mx.router.PathPrefix("/debug/").Handler(http.DefaultServeMux)

	// trap sigterm or interupt for gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}