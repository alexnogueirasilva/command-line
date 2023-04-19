package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Point start")
	application := app.Generate()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
