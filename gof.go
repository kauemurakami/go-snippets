package gof

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

// return the application command line ready to use
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Commanda Line Applications"    //nome
	app.Usage = "Search IP's and servers name" //utilização
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Usage: "Name of the directory to create",
		},
	}
	// Add commands use this property, it is slice of the commands
	app.Commands = []cli.Command{

		{
			Name:   "create",
			Usage:  "Create a new directory with a .go file inside",
			Flags:  flags,
			Action: createDirectory,
		},
	}
	return app
}

// bonus function
func createDirectory(c *cli.Context) {
	dirName := c.String("dirname") // Obtém o nome do diretório da flag --dirname

	if dirName == "" {
		fmt.Println("Please provide a directory name using --dirname flag.")
		return
	}

	err := os.Mkdir(dirName, 0755) // Cria o diretório
	if err != nil {
		log.Fatal(err)
	}

	fileName1 := filepath.Join(dirName, dirName+".go") // Cria o nome do arquivo com extensão .go
	file1, err := os.Create(fileName1)                 // Cria o arquivo
	if err != nil {
		log.Fatal(err)
	}
	fileName2 := filepath.Join(dirName, dirName+"_test.go") // Cria o nome do arquivo com extensão .go
	file2, err := os.Create(fileName2)                      // Cria o arquivo
	if err != nil {
		log.Fatal(err)
	}

	// Conteúdo do arquivo .go
	content1 := fmt.Sprintf(`package %s

import "fmt"

func Myfunction() {
    fmt.Println("Hello, world!")
}`, dirName)

	// Conteúdo do arquivo .go
	content2 := fmt.Sprintf(`package %s

import (
	"fmt"
	"testing"
)

func TestYourFunc(t *testing.T) {
    
}`, dirName)
	// Escreve o conteúdo no arquivo
	_, err = file1.WriteString(content1)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file2.WriteString(content2)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()
	defer file2.Close()

	fmt.Printf("Directory '%s', file '%s' and '%s' created successfully.\n", dirName, fileName1, fileName2)
}
