package main

import (
	"github.com/guigoebel/client-server-api/client"
	"github.com/guigoebel/client-server-api/server"
)

func main() {
	server.Start()

	client.GetQuotation()
}
