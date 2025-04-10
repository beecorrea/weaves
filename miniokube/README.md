## Prerequisites
* Minikube
* QEMU 
  * SocketVM
* MinIO Client

## Quickstart
Start a local Minikube instance with QEMU as a driver:
```shell
minikube start --driver qemu --network socket_vmnet
```

Deploy the MinIO Operator from their GitHub Release:
```shell
kubectl apply -k "github.com/minio/operator?ref=v7.0.0"
```

Apply the kustomization:
```shell
kubectl kustomize apply kustomization
```

Set the alias for `mc` commands:
```shell
mc alias set miniokube https://<LOADBALANCER_EXTERNAL_IP> minio minio123 --insecure
```

## Stopping
```shell
kubectl scale minio-operator --replicas 0 -ns minio-operator
```
