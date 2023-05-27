package main

import (
	"Project2-Hacktiv8/routers"
)

var PORT = ":302"

func main() {
	routers.StartServer().Run(PORT)
}
