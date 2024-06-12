package main

import (
	"fmt"
	"log"

	"github.com/dosadczuk/pdfimages"
)

func main() {
	cmd := pdfimages.NewCommand(
		pdfimages.WithSaveRaw(),
	)

	err := cmd.Run("./example.pdf", "./images")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("DONE")
}
