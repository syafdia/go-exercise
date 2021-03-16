# Develop Kong plugin using Go

### How to build

1. Build docker image 
   ```
   docker build -t kong-demo
   ```
  
2. Run container

 	```
  	docker run -ti --rm --name kong-go-plugins \
  	-e "KONG_DATABASE=off" \
  	-e "KONG_DECLARATIVE_CONFIG=/tmp/config.yml" \
  	-e "KONG_PLUGINS=bundled,key-checker" \
  	-e "KONG_PLUGINSERVER_NAMES=key-checker" \
  	-e "KONG_PLUGINSERVER_KEY_CHECKER_START_CMD=/usr/local/bin/key-checker" \
  	-e "KONG_PLUGINSERVER_KEY_CHECKER_QUERY_CMD=/usr/local/bin/key-checker -dump" \
  	-e "KONG_PROXY_LISTEN=0.0.0.0:8000" \
  	-e "KONG_LOG_LEVEL=debug" \
  	-p 8000:8000 \
     kong-demo
   ```


### How to test
1. Visit `curl --location --request GET '0.0.0.0:8000?key=invalidconsumerkey'`, it will return `401`
2. Visit `curl --location --request GET '0.0.0.0:8000?key=mysecretconsumerkey', it will return `200` with valid data