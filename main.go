package main

import (
	"log"

	"bitbucket.org/isbtotogroup/japan-pools/routers"
)

func main() {
	app := routers.Init()
	log.Fatal(app.Listen(":7072"))
}
