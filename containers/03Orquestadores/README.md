1. Create kind

curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.26.0/kind-linux-amd64

kind create cluster -n pffp

curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

kind create cluster -n pffp --config 03Orquestadores/01kind-config.yaml

kubectl get pod -A

podman exec -ti pffp-control-plane -- bash 

kubectl apply -f 02Imanges/03cowpod.yaml

cowsay   0/1     ErrImagePull   0          28s

... podemos intentar subir la imagen en multiarch y ver como hace pull de la arch correcta

oc apply -f 03Orquestadores/02pod.yaml

5. Multi tier

kubectl create secret generic mysql-password --from-literal=MYSQL_ROOT_PASSWORD=mypass

kubectl apply -f 03Orquestadores/05tier.yaml

mysql --port=30000 --host=127.0.0.1 --user=root --password=mypass

create database test;

use test;

create table test_on_test(id varchar(20));

insert into test_on_test values ('one');


10. Operadaores

kubectl apply -k github.com/zalando/postgres-operator/manifests

kubectl create -f manifests/minimal-postgres-manifest.yaml