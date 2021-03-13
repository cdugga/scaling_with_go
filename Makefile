.ONESHELL:

build_server_image:
	cd services/server/src; \
	go mod init github.com/cdugga/scaling_with_go/server ;\
	cd ..; \
	docker build -t cdugga/scaling-with-go-server:1.0.0 .;
start_server_image:
	docker run -p 8081:8080 -it cdugga/scaling-with-go-server:1.0.0 ;
push_server_image:
	docker push cdugga/scaling-with-go-server ;
deploy_server_image_k8s:
	cd services/server/; \
	kubectl create -f deployment.yaml