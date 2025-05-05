## Installation
Before installing, download and install go version 1.18 and above from **[golang's official website](https://go.dev/doc/install)**
Ensure you have docker if you are running it on docker. 
You can also set up postgres DB locally

## Install make 
for MacOS run

```brew install make```

### Make available the correct set of dependencies 
```go mod tidy``` 

### starting server
- create a `.env` file in the base directory with sample fields from `.env.sample`
- run `make postgres` to provision a postgres DB on docker
- run `make run` to run the service/app with the provisioned DB locally.

You can also start up the app with the docker build using `docker compose up` or `make docker-up`. 
This will build the app in the dockerfile present in the base directory.
- NOTE: You will need to change the DB_HOST in your env to `postgres` in order to run the service using `docker`
