package main

import "trandung2k1/server/routes"

func main() {
	server := routes.NewAPIServer(":8080")
	server.Run()
}
