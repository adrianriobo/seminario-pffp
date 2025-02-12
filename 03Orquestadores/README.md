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


10. Operadaores

kubectl apply -k github.com/zalando/postgres-operator/manifests

kubectl create -f manifests/minimal-postgres-manifest.yaml