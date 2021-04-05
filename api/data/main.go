package main

import (
	"context"
	"log"

	"github.com/objforce/objforce/api/data/bootstrap"
)

func main() {
	app := bootstrap.App()

	/*
		|--------------------------------------------------------------------------
		| Run The Application
		|--------------------------------------------------------------------------
		|
		| Once we have our application, we can listen for incoming request and send
		| the associated response.
		|
	*/

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
