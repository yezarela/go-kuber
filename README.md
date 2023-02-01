# kuba

## Prerequisites

1. [Docker](https://docs.docker.com/desktop/install/mac-install/)
2. [Air](https://github.com/cosmtrek/air#installation) for live reload during development
3. [Migrate CLI](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md#migrate-cli) to manage migrations
4. [Mockgen](https://github.com/golang/mock#go-116) to generate mocks

## Installation
```sh
# install dependencies
make dep

# docker compose up
make setup-local

# run db migrations
make migrate-up
```

## Running locally
Run this command and every changes you make will be automatically rebuild, thanks to Air.
```sh
make run
```

## Building image
```sh
make build-docker
```
Or if you prefer to run the image locally
```sh
docker build --tag simple-app .
docker run --name simple simple-app

# automatically delete existing
docker run --name simple --rm simple-app

# bind to port $docker_host_port:$container_port
docker run -p 0.0.0.0:4001:4001 --name simple --rm simple-app
```

## Using kubernetes locally
```sh
minikube start

kubectl cluster-info

# to point your shell to minikube's docker-daemon, run:
eval $(minikube -p minikube docker-env)

kubectl create -f deployment.yaml
kubectl create -f service.yaml

kubectl port-forward svc/simple-service 4001:4001
curl localhost:4001

# alternative to get the local endpoint url
minikube service simple-app --url
```