.ONESHELL:

build_postgres_image:
	cd services/postgres; \
	docker build -t cdugga/scaling-with-go-postgres:1.0.0 .
start_postgres_image:
	docker run -p 5432:5432 -d -e POSTGRES_USER=docker -e POSTGRES_PASSWORD=docker -e POSTGRES_DB=docker -it cdugga/scaling-with-go-postgres:1.0.0 ;
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
prometheus_build:
	cd services/prometheus; \
	docker build -t cdugga/scaling-with-go-prometheus:1.0.0 .;
prometheus_deploy: prometheus_build
	cd services/prometheus; \
	docker push cdugga/scaling-with-go-prometheus:1.0.0;
deploy_prometheus_k8s:
	cd services/prometheus; \
	kubectl create -f deployment.yaml

build_nginx:
	cd services/nginx; \
	docker build -t cdugga/scaling-with-go-nginx:1.0.5 .;
	docker push cdugga/scaling-with-go-nginx:1.0.5 ;
deploy_nginx_k8s: build_nginx
	cd services/nginx; \
	kubectl create -f deployment.yaml

build_grafana-sv:
	cd services/grafana/shared-volume; \
	docker build -t cdugga/scaling-with-go-grafana-svolume:1.0.0 .; \
	docker push cdugga/scaling-with-go-grafana-svolume:1.0.0
deploy_grafana_k8s: build_grafana-sv
	cd services/grafana; \
	kubectl create -f deployment.yml

redisclient_all: build_redisclient_image push_redisclient_image deploy_redisclient_image_k8s
prometheus_all: prometheus_deploy deploy_prometheus_k8s

delete_k8s_deployments:
	kubectl delete deploy prometheus-deployment; \
	kubectl delete deploy redisclient-deployment; \
	kubectl delete deploy redis-master; \
	kubectl delete deploy grafana; \
	kubectl delete deploy nginx-deployment;
delete_k8s_services:
	kubectl delete svc redis-client-svc; \
	kubectl delete svc prometheus-main; \
	kubectl delete svc redis-master;

spin-up: deploy_redis_master redisclient_all prometheus_all deploy_grafana_k8s deploy_nginx_k8s
teardown: delete_k8s_services delete_k8s_deployments