package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dosadczuk/go-pdfimages"
)

func main() {
	cmd := pdfimages.NewCommand(
		pdfimages.WithSaveRaw(),
	)

	err := cmd.Run(context.Background(), "./example.pdf", "./images")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("DONE")
}
