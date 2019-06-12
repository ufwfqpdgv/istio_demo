#!/bin/sh

server_name=smh-ticket-server
config_name=${server_name}"-configmap"
istio_config_file=istio-config.yaml
config_path=config/

kubectl delete configmap ${config_name} || \
kubectl create configmap ${config_name} --from-file=${config_path}
kubectl apply -f ${istio_config_file}
