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

deploy_redis_master:
	cd services/redis/; \
   	kubectl create -f redis-master.yaml

build_redisclient_image:
	cd services/redis-client/src; \
	go mod init github.com/cdugga/scaling_with_go/redisclient ;\
	cd ..; \
	docker build -t cdugga/scaling-with-go-redisclient:1.0.3 .;
start_redisclient_image:
	docker run -p 8081:8080 -it cdugga/scaling-with-go-redisclient:1.0.3 ;
push_redisclient_image:
	docker push cdugga/scaling-with-go-redisclient:1.0.3 ;
deploy_redisclient_image_k8s:
	cd services/redis-client/; \
	kubectl create -f deployment.yaml

redisclient_all: build_redisclient_image push_redisclient_image deploy_redisclient_image_k8s