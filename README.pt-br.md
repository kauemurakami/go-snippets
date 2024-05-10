## Go Fast Snippets
[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-snippets/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-snippets/blob/main/README.md)
[![Star on GitHub](https://img.shields.io/github/stars/kauemurakami/go-snippets.svg?style=flat&logo=github&colorB=deeppink&label=stars)](https://github.com/kauemurakami/get_snippets_extension)

![](gofast.png)  

Extensão para vscode com autocompletes para códigos GO.  
*Acelere* o seu processo de desenvolvimento com atalhos que te livram de reescrever códigos rotineiros, como funções, variáveis, structs e etc, veja os snippets disponíveis:<br/><br/>
Aqui estão alguns exemplos de uso de trechos de código. Você pode ver o resto das opções abaixo, neste README.<br/>  

*```gomain```*  
Ao escrever ```gomain``` em um arquivo e seleciona-lo para o autocomplete ele irá gerar o código básico de um arquivo, exemplo, começe escrevendo ```gomain```, ou apenas ```go``` e verá as opções completas dos snippets:  
```go
//gomain gerá está estrutura
package main

func main() {
	
}
```  

*```gofunc```*  
Ao escrever o snippet ```gofunc``` ele irá nos retornar uma função com autocompletes no noma da função, parâmetro e tipo de retorno, basta alternar com a tecla TAB após iniserir um a um, vejamos o que ```gofunc``` nos gera:  
```go
//gofunc
	func name(params type)  returnType {
		return params
	}
```
Lembrando que, inicialmente ```params``` é o primeiro atributo a ser inserido e é espelhado nos dois valors, de um TAB e ira pular para o ```type```, ao preencher de outro TAB e estará em ```returnType```, faça o teste.<br/><br/>

*```gofuncempty```*  
Aqui você tem a opção de criar uma função vazia, sem parâmetros e sem retornos, logo ao digitar ```gofun``` ele aparecerá, mas para o caso de uma função vazia utilizaremos o comando ```gofuncempty``` que nos gerará:  
```go
	func name() {
			
	}
```
Lembrando sempre que ```name``` já vira selecionada para alteração do nome.<br/><br/>

*```gofuncerr```*   
Uma função que pode receber um erro e trata-lo, basta digitar ```gofuncerr``` e teremos o seguinte resultado:  
```go
func name(params) (returnType, error) {
	
}
```
Com o mesmo modus operandi dos outros snippets que possuem mais de um valor para ser alterado, essa função também segue esses padrões, podendo ser alternadas usando TAB ao declarar um valor e ir para o próximo.<br/><br/>

*```govar```*  
Aqui ao digitar ```govar``` e selecionarmos o auto complete, ele nos cria uma variável simples, recebendo o tipo por inferência, recebendo um ```name``` e um ```type```, veja:  
```go
name := val
```
Possuindo o mesmo recurso de estar auto selecionado para alteração de ```name``` e ao dar TAB definir o ```val```.<br/><br/>

*```govartype```*   
Aqui ao digitar ```govar``` já aparecerá a opção ```govartype```, diferente de ```govar```, ```govartype``` cria uma variável com tipo definido, vejamos:  
```go
var name type
```
Possuindo o mesmo comportamento, onde ```name``` está pré selecionado para alteração e ao dar um TAB pode alterar o ```type```.<br/><br/>

*```goif```*   
Ao digitarmos ```goif``` e selecionarmos o auto complete, veremos o seguinte trecho de código simples:  
```go
if var operator val {

}
```
Onde estará préselecionado nossa ```var```, para que altere com o nome de alguma variável ou valor que queira comparar, o ```operator``` que pode ser >, ==, por exemplo, e por fim ```value``` o valor que estamos usando para comparação, tudo podendo ser definido apenas com o teclado usando TAB para alternar de uma para outra.<br/><br/>

*```gofor```*  
Ao digitarmos ```gofor``` ele criará um for simples, com inicialização de variável dentro dele, vejamos:  
```go
for var := initial-value; var operator value ; var factor {
 
}
```
Onde var vem inicialmente pré selecionada para ser definida como quiser, exemplo ```k:= 0```, logo em seguida ao dar um TAB escolherá o ```initial value``` de ```var```, todos os ```var```, serão completados juntamente com a primeira definição de ```var``` ou seja se inicialmente o primeiro ```var``` for ```k```, todos os outros ```var``` também serão, entre isso também, ao dar um TAB, também devemos definir o operador para parar o ```for```, == por exemplo, em seguida definimos o ```factor``` que pode ser um ```++``` ```--``` por exemplo.<br/><br/>

*```goforrange```*   
Aqui unimos nosso ```for``` com ```range``` para recuperar valores, ao digitar ```goforrange``` ele nos gerará este trecho de código:  
```go
for index, obj := range array-or-slice {
				
}
```
Lembrando sempre que esses "nomes" definidos vão poder ser alterado navegando entre eles com TAB.<br/><br/>

*```gostruct```*   
Aqui ao digitar ```gostruct``` você irá criar um modelo de um struct, vejamos:  
```go
type name-struct struct {
	varname type
}
```
Comçando definindo ```name-struct``` com um TAB podemos definir ```varname``` e com outro TAB definimos o ```type```<br/><br/>

*```goinit```*  
Aqui ao digitar ```goinit``` criamos a função ```init()``` que inicia antes de todas do arquivo.go, vejamos:  
```go
func init() {
			
}
```
Aqui não precisamos definir nenhum parâmetro.<br/><br/>

*```goerr```*   
Com ````goerr```` podemos tratar o erro com um ```log.Fatal(err)``` em um simples if, ou como desejar:  
```go
if err != nil {
	log.Fatal(err)
}
```
```log.Fatal()``` foi usado como exemplo, você poderia tratar esse erro de diferentes modos.<br/><br/>

*```goreadinput```*   
Aqui permitimos o usuário inserir dados via terminal com ```goreadinput```, vejamos o trecho de código gerado:  
```go
fmt.Print("Enter input: ")
var input type
fmt.Scanln(&input)
```
Aqui você deve alterar o type apenas, após isso poderá receber um valor que você inserir via terminal.

