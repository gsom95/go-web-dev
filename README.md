# What it is
Repo for Jon Calhoun [Web Development with Go](https://www.usegolang.com/) course.

## Useful commands

### Start docker compose 
`docker compose up`

---
### Start docker compose in detached mode (terminal not stuck with container output)
`docker compose up -d`

---
### Stop services
`docker compose stop`

---
### For executing psql inside db container
`docker compose exec -it db psql -U baloo -d lenslocked`

`docker compose exec` — This lets us execute a binary inside of a container that is being run by docker compose.

`-it` — flag, combination of two flags `-i` and `-t` for interacting with terminal after **running docker compose exec**.

---
### For listing active **compose** containers
`docker compose ls`

### List containers
`docker ps`

### Execute a command on running container (with docker)
`docker exec -it <CONTAINER_NAME> psql -U baloo -d lenslocked`

---
### Check who is using port
`lsof -i tcp:3000`
