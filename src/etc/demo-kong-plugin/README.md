# Develop Kong plugin using Go

### How to Develop

1. Install [Docker](https://www.docker.com)

2. Install [CompileDaemon](https://github.com/githubnemo/CompileDaemon)
	```
	go get github.com/githubnemo/CompileDaemon

	```

2. Run development script
	```
	sh scripts/start-development.sh
	```


### How to test
1. Visit `curl --location --request GET '0.0.0.0:8000?key=invalidconsumerkey'`, it will return `401`
2. Visit `curl --location --request GET '0.0.0.0:8000?key=mysecretconsumerkey'`, it will return `200` with valid data