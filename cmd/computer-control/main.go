package main

import (
	server "computer-control/internal/server"
)

func main() {
	port := ":8888"
	server.StartServerSsdp(port)
	server.StartServerRest(port)
}
