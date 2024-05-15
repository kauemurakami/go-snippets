// package gof

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"path/filepath"

// 	"github.com/urfave/cli"
// )

// // return the application command line ready to use
// func Gerar() *cli.App {
// 	app := cli.NewApp()
// 	app.Name = "Go fast Snippets CLI (gof)"                     //nome
// 	app.Usage = "Create new package with files and basic codes" //utilização
// 	flags := []cli.Flag{
// 		cli.StringFlag{
// 			Name:  "pkgname",
// 			Usage: "Create a package folder and internal files .go, and _test.go, with basic code inside files",
// 		},
// 	}
// 	// Add commands use this property, it is slice of the commands
// 	app.Commands = []cli.Command{
// 		{
// 			Name:   "create",
// 			Usage:  "Create a new directory with a .go file inside",
// 			Flags:  flags,
// 			Action: createDirectory,
// 		},
// 	}
// 	return app
// }

// // bonus function
// func createDirectory(c *cli.Context) {
// 	dirName := c.String("pkgname") // Obtém o nome do diretório da flag --dirname

// 	if dirName == "" {
// 		fmt.Println("Please provide a directory name using --pkgname flag.")
// 		return
// 	}

// 	err := os.Mkdir(dirName, 0755) // Cria o diretório
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fileName1 := filepath.Join(dirName, dirName+".go") // Cria o nome do arquivo com extensão .go
// 	file1, err := os.Create(fileName1)                 // Cria o arquivo
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fileName2 := filepath.Join(dirName, dirName+"_test.go") // Cria o nome do arquivo com extensão .go
// 	file2, err := os.Create(fileName2)                      // Cria o arquivo
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Conteúdo do arquivo .go
// 	content1 := fmt.Sprintf(`package %s

// import "fmt"

// func Myfunction() {
//     fmt.Println("Hello, world!")
// }`, dirName)

// 	// Conteúdo do arquivo .go
// 	content2 := fmt.Sprintf(`package %s

// import (
// 	"fmt"
// 	"testing"
// )

// func TestYourFunc(t *testing.T) {

// }`, dirName)
// 	// Escreve o conteúdo no arquivo
// 	_, err = file1.WriteString(content1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_, err = file2.WriteString(content2)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file1.Close()
// 	defer file2.Close()

//		fmt.Printf("Directory '%s', file '%s' and '%s' created successfully.\n", dirName, fileName1, fileName2)
//	}
package gof

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Go fast Snippets CLI (gof)"
	app.Usage = "Create new package with files and basic codes"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Usage: "Create a package folder and internal files .go, and _test.go, with basic code inside files",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "create:p",
			Usage:  "Create a new pacakge directory with a .go and _test.go files inside",
			Flags:  flags,
			Action: createDirectory,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	return app
}

func createDirectory(c *cli.Context) {
	dirName := c.String("name")

	if dirName == "" {
		fmt.Println("Please provide a directory name using --name flag.")
		return
	}

	// Verifique se o diretório já existe
	if _, err := os.Stat(dirName); !os.IsNotExist(err) {
		// fmt.Printf("Directory '%s' already exists.\n", dirName)
		return
	}

	err := os.Mkdir(dirName, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Verifique se os arquivos já existem antes de criar cada um deles
	fileName1 := filepath.Join(dirName, dirName+".go")
	if _, err := os.Stat(fileName1); !os.IsNotExist(err) {
		fmt.Printf("File '%s' already exists.\n", fileName1)
		return
	}

	file1, err := os.Create(fileName1)
	if err != nil {
		log.Fatal(err)
	}

	fileName2 := filepath.Join(dirName, dirName+"_test.go")
	if _, err := os.Stat(fileName2); !os.IsNotExist(err) {
		// fmt.Printf("File '%s' already exists.\n", fileName2)
		return
	}

	file2, err := os.Create(fileName2)
	if err != nil {
		log.Fatal(err)
	}

	// Escreva o conteúdo nos arquivos e feche-os
	content1 := fmt.Sprintf(`package %s

import "fmt"

func MyFunction() {
    fmt.Println("Hello, world!")
}`, dirName)

	content2 := fmt.Sprintf(`package %s

import (
	"testing"
)

func TestMyFunction(t *testing.T) {
    
}`, dirName)

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

	fmt.Printf("Package '%s', file '%s' and '%s' createds successfully.\n", dirName, fileName1, fileName2)
}
