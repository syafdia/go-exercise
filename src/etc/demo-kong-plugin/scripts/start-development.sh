#!/bin/bash

# Run docker
docker-compose up -d

# Hot Reload
CompileDaemon \
-build='docker-compose exec -T kong-plugin-builder make all' \
-command='docker-compose exec -T kong kong reload'