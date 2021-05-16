package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"github.com/cdugga/scaling_with_go/redisclient/data"
	"github.com/cdugga/scaling_with_go/redisclient/handlers"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/fs"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strconv"
	"time"
)


const (
	PORT=8081
)



//go:embed templates
var embededFiles embed.FS


func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}


var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "Status of HTTP response",
	},
	[]string{"status"},
)

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "Number of requests",
	},
	[]string{"path"},
	)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "redisclient_processed_ops_total",
		Help: "The total number of processed events",
	})
)

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

func main(){
	recordMetrics()
	StartServer()
}

var ctx = context.Background()

type Env struct {
	DataSource data.DataAccess
}

func getFileSystem() http.FileSystem{
	log.Print("using embed mode")
	fsys, err := fs.Sub(embededFiles, "templates")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}


func StartServer() {

	log.Printf("Staring server on port %d ", PORT )

	sm := mux.NewRouter()
	sm.Use(PrometheusMiddleWare)

	sm.PathPrefix("/debug/").Handler(http.DefaultServeMux)

	promHandler := sm.Methods(http.MethodGet).Subrouter()
	promHandler.Handle("/metrics", promhttp.Handler())


	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.Handle("/", http.FileServer(getFileSystem()))

	getKeyRouter := sm.Methods(http.MethodGet).Subrouter()
	getKeyRouter.HandleFunc("/value/{key}", handlers.FetchKeyHandler)

	writeKeyRouter := sm.Methods(http.MethodPost).Subrouter()
	writeKeyRouter.HandleFunc("/keyvalue", handlers.WriteKeyHandler)
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



func PrometheusMiddleWare(next http.Handler) http.Handler{
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		w := NewResponseWriter(rw)
		next.ServeHTTP(rw, r)

		statusCode := w.statusCode
		responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
		totalRequests.WithLabelValues(r.RequestURI).Inc()
		timer.ObserveDuration()
	})
}

func init(){
	prometheus.Register(totalRequests)
	prometheus.Register(responseStatus)
	prometheus.Register(httpDuration)

}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}


func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func MiddleWareProductValidation(next http.Handler) http.Handler{

	return http.HandlerFunc(func(rw http.ResponseWriter, r*http.Request) {

		keyVal := &data.KeyValue{}

		e := json.NewDecoder(r.Body)
		err := e.Decode(keyVal)

		if err != nil {
			log.Println("[Error] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// add key value to the context
		ctx := context.WithValue(r.Context(), data.KeyValue{}, keyVal)
		req := r.WithContext(ctx)



		// cGall the next handler, which can be another middleware in the chain, or the final handler
		next.ServeHTTP(rw,req)

	})
}

