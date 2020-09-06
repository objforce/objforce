package main

import(
	"context"
	"github.com/objforce/objforce/bot/bootstrap"
	"log"
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