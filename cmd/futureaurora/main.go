// cmd/futureaurora/main.go
package main

import (
	"flag"
	"log"
	"os"

	"futureaurora/internal/futureaurora"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := futureaurora.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
