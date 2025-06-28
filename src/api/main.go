package main

import "log"

const version = "0.0.1"

func main() {
	configuration := configuration{
		address: "localhost:8080",
	}

	app := application{
		configuration: configuration,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
