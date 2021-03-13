.ONESHELL:

build_server_image:
	cd services/server/src; \
	go mod init github.com/cdugga/scaling_with_go/server ;\
	cd ..; \
	docker build -t swg_server .;
start_server_image:
	docker run -p 8081:8080 -it swg_server ;