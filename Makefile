.PHONY:hint test test2 build run clean all push run2 deploy cancel_deploy

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

repo=smh
image_name=smh-ticket-server
tag=v0.1
image_whole_name=${repo}/${image_name}:${tag}
port=11114
server_name=${image_name}
server_path=server/smh_ticket_server.go
online_docker_addr=registry.cn-shenzhen.aliyuncs.com
namespace=smh-demo
online_docker_image_addr=${online_docker_addr}/${namespace}/${image_name}:${tag}
istio_config_file=istio-config.yaml
config_path=recommend/config
config_name=${server_name}"-configmap"

hint:
	@echo "test:docker test,test2:docker-compose test"

test:
	make clean || make build && make run

test2:
	make clean || make build && make push && make run2

build:
	go build -o ${server_name} ${server_path}
	docker build --no-cache -t ${image_whole_name} .

run:
	docker run --name ${image_name} -p ${port}:${port} ${image_whole_name} #需要后台运行就加上-d

clean:
	docker rm -f ${image_name}

push:
	docker tag ${image_whole_name} ${online_docker_image_addr}
	docker push ${online_docker_image_addr}

run2:
	docker-compose down
	docker-compose up

deploy:
	kubectl delete configmap ${config_name} || \
	kubectl create configmap ${config_name} --from-file=${config_path}
	kubectl apply -f ${istio_config_file}

cancel_deploy:
	kubectl delete -f ${istio_config_file}
