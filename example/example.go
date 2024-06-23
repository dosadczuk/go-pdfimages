package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dosadczuk/go-pdfimages"
)

func main() {
	cmd, err := pdfimages.NewCommand(
		pdfimages.WithSaveRaw(),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Run(context.Background(), "./example.pdf", "./images")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("DONE")
}
