package main

import (
	server "volume-control/internal/server"
)

func main() {
	port := ":8888"
	server.StartServerSsdp(port)
	server.StartServerRest(port)
}
