# Coverage metric
go test -cover

# Coverage
go tool cover -html=cpu.out

# benchmark
go test -bench=. -benchmem -v

# Generate CPU profile  
go test -run=NONE -bench=. -cpuprofile=cpu.log

# View CPU profile report
go tool pprof -text -nodecount=10 ./service.test.exe cpu.log

# Generate MEM profile
go test -run=NONE -bench=. -memprofile=mem.log

# View MEM profile report
go tool pprof -text -nodecount=10 ./service.test.exe mem.log

# PPROF interactive mode
go tool pprof service.test cpu.log
pprof>top10
pprof>list <funcname>

# Flamegraph ( and other graph options)
go tool pprof -http=":8001" ./service.test.exe cpu.log 

# Load test 
go-wrk -d 5 http://localhost:8080/rockwoo

# Trace
1. Import: _ "net/http/pprof"
2. Add path for debug info: mx.router.PathPrefix("/debug/").Handler(http.DefaultServeMux)
3. Download trace info: http://localhost:8080/debug/pprof/trace
4. Use go tool to view trace information , analyse core usage: go tool trace .\trace