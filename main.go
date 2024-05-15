package main

import (
	gof "app/cli"
	"log"
	"os"
)

func main() {
	application := gof.Gerar()
	application.Run(os.Args)
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
