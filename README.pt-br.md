## Go Fast Snippets
[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-snippets/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-snippets/blob/main/README.md)
[![Star on GitHub](https://img.shields.io/github/stars/kauemurakami/go-snippets.svg?style=flat&logo=github&colorB=deeppink&label=stars)](https://github.com/kauemurakami/get_snippets_extension)

<marquee direction="right" scrollamount=40>![](gofast.png)</marquee>

#### Extensão para vscode com snippets de autocompletes para códigos em GO.
*Acelere* o seu processo de desenvolvimento com atalhos que te livram de reescrever códigos rotineiros, como funções, variáveis, structs e etc, veja os snippets disponíveis:<br/><br/>
Aqui estão alguns exemplos de uso de trechos de código. Você pode ver o resto das opções abaixo, neste README.<br/>  

### gomain
![](assets/gomain.gif)  
Ao escrever ```gomain``` em um arquivo e seleciona-lo para o autocomplete ele irá gerar o código básico de um arquivo, exemplo, começe escrevendo ```gomain```, ou apenas ```go``` e verá as opções completas dos snippets:  
```go
package main

func main() {
	
}
```
Simples assim <br/>

### goinit
![](assets/goinit.gif)  
Aqui ao digitar ```goinit``` criamos a função ```init()``` que inicia antes de todas do arquivo.go, vejamos:  
```go
func init() {
			
}
```
Aqui não precisamos definir nenhum parâmetro.<br/>

### gofile
![](assets/gofile.gif)   
Ao digitar ```gofile``` em um arquivo vazio ele gera o código básico de um arquivo ```.go```, com ```name``` pré selecionado para alteração:  
```go
package name

func init() {
	
}

func main() {
	
}
```  
```gofile``` nos gera um código um pouco mais completo que ```gomain```, nos fornecendo também ```goinit``` <br/>

### gotestfile  
Ao escrever ```gotestfile``` ele irá nos retornar um código para um arquivo que contém a importação de ```testing``` e criação do package com a função test inicial, mantendo a palavra ```Test``` para o go reconhecer como uma função de test, o primeiro atributo e único a ser defitino é o nome restante da função aopós ```Test```, lembrando que a segunda palavra aops ```Test``` Também de iniciar com letra maiúscula exemplo:  
```go
package name

import "testing" 

func TestNameFunction(t *testing.T) {
	
}
```  
### gofunc
![](assets/gofunc.gif)  
Ao escrever ```gofunc``` ele irá nos retornar uma função com autocompletes no nome da função, parâmetro e tipo de retorno, basta alternar com a tecla TAB após iniserir um a um, vejamos o que ```gofunc``` nos gera:  
```go
//gofunc
func name(params type)  returnType {
  return params
}
```
Lembrando que, inicialmente ```name``` é o primeiro atributo a ser inserido, de um TAB e ira pular para o/os ```params```, podendo inserir um ou mais parâmetros separados por vírgula, de outro TAB e estará em ```type``` que é o tipo do ```params``` no ultimo TAB irá definir o tipo do retorno ```returnType```, faça o teste.<br/>

### gofunctest
Nos gera uma função test com os parâmetros padrão veja:  
```go
func TestNameFunction(t *testing.T) {
	
}
```  

### gofuncempty  
![](assets/gofuncempty.gif)  
Aqui você tem a opção de criar uma função vazia, sem parâmetros e sem retornos, logo ao digitar ```gofun``` ele aparecerá, mas para o caso de uma função vazia utilizaremos o comando ```gofuncempty``` que nos gerará:  
```go
func name() {
		
}
```
Lembrando sempre que ```name``` já vira selecionada para alteração do nome.<br/>

### gofuncerr
![](assets/gofuncerr.gif)  
Uma função que pode receber um erro e trata-lo, basta começar a digitar ```gofuncerr``` e teremos o seguinte resultado:  
```go
func name(params) (returnType, error) {
	
}
```
Com o mesmo modus operandi dos outros snippets que possuem mais de um valor para ser alterado, essa função também segue esses padrões, podendo ser alternadas usando TAB ao declarar um valor e ir para o próximo.<br/>

### gofunchttp
Ao digitar ```gofunchttp``` ira nos gerar uma função básica de http como:  
```go
func FunctionName(w http.ResponseWriter, r *http.Request) {
	
}
```
Crtl+S faz um auto importe dos packages que serão usados mas ainda não foram importados.  

### gostruct
![](assets/gostruct.gif)  
Aqui ao digitar ```gostruct``` você irá criar um modelo de um struct, vejamos:  
```go
type name struct {
  varname type
}
```
Comçando definindo ```name``` com um TAB podemos definir ```varname``` e com outro TAB definimos o ```type```<br/>

### gointerface  
![](assets/gointerface.gif)  

Aqui ao digitarmos ```gointerface``` ele nos dara um snippet de uma interface:  
```go
type name interface {
  
}
```
Deixando o nome pré-selecionado para sua alteração.<br/>

### gointerfacegeneric
![](assets/gointerfacegeneric.gif)  

Aqui ao digitarmos  ```gointer..``` já nos mostrará o snippet que buscamos ```gointerfacegeneric```, ao seleciona-lo, nos gerará o seguinte código:  
```go
func name(nameinterface interface{}) {
  
}
```
Como primeiro primeiro atributo pré-selecionado para renomeação o ```name``` e com um TAB definimos também ```nameinterface``` <br/>

### goroutine
Nos retorna uma função com a clásula ```go``` para se tornar uma ```goroutine```:  
```go
go func(params){
	
}
```
### goroutineanon
Cria uma ```goroutine``` anônima:  
```go
go func(){
	
}()
```  

### goselect
Uma estrutura muito parecida com switch na tomada de decisão, mas com foco em uso com concorrência, basta usar nosso snippet ```goselect``` para criar uma estrutura básica:  
```go
select {
 case obj := <-channel:
		fmt.Println("a")
	case obj := <-channel:
		fmt.Println("a")
}
```  

### gomap
Cria uma declaração de uma variável map com tipo da chave e tipo do valor, neste caso segue preselecionado para alteração da ```varname```, ```keytype``` e ```valuetype```:  
```go
var varname map[keytype]valuetype
```  

### gomapvalues
Com ```gomapvalues``` você irá criar um map por inferência, ja podendo atribuir valores, vejamos o código gerado ao usar o snippet:  
```go
mapname := map[keytype]valuetype{
  "key1" : "Value1", 
  "key2" : "Value1", 
}
```
Seguindo o mesmo padrão em ordem, ```mapname``` é o primeiro pré-selecionado, ao navegar com TAB você segue para ```keytpe``` -> ```valuetype``` -> ```key1``` -> ```value1``` e assim por diante podendo alterá-los sem esforço.<br/>


### govar  
Aqui ao digitar ```govar``` e selecionarmos o auto complete, ele nos cria uma variável simples, recebendo o tipo por inferência, recebendo um ```name``` e um ```type```, veja:  
```go
name := val
```
Possuindo o mesmo recurso de estar auto selecionado para alteração de ```name``` e ao dar TAB definir o ```val```.<br/>

### govartype
Aqui ao digitar ```govar``` já aparecerá a opção ```govartype```, diferente de ```govar```, ```govartype``` cria uma variável com tipo definido, vejamos:  
```go
var name type
```
Possuindo o mesmo comportamento, onde ```name``` está pré selecionado para alteração e ao dar um TAB pode alterar o ```type```.<br/>

### gowriteheader
Ao começar a digitar ```gowriteheader```, logo surgira o snippet, selecione ele e ele retornará o seguinte código:  
```go
w.WriteHeader(http.StatusBadRequest)
```
### gowriter
```gowriter``` nos traz uma solução completa para retornar uma resposta em json para o cliente, usando ```json.Marshal(obj)``` para ser convertido em um json para a variavel ```varJSON```, também verificamos se algo sair errado durante isso, e retornamos uma mensagem de erro, caso tudo tenha ocorrido bem, podemos retornar nosso json ao fim da função.  
```go
varJson, err := json.Marshal(objectToConvert)
if err != nil {
 w.WriteHeader(http.StatusBadRequest)
 w.Write([]byte("ErrorMessage"))
 return
}

w.WriteHeader(http.StatusOK)
w.Write(varJson)
```  
Lembrando que as variáveis pré selecionadas para alteração são:  
 1. varJson (TAB) 
 2. objectToConvert (TAB)
 3. ErrorMessage   

### gomarshal
Ao começar a digitar ```gomarshal``` e selecionar o snippet, ele nos gerará um código para converter um struct ou outro tipo em um ```json``` com a função ```json.Marshal```, veja no exemplo:  
```go
var varname *varType
varJSON, err := json.Marshal(*varname)
if err != nil {
  w.WriteHeader(http.StatusBadRequest)
  w.Write([]byte("ErrorMessage"))
}
```  
Lembrando que as variáveis pré selecionadas para alteração são:  
 1. varname (TAB) 
 2. varType (TAB)
 3. ErrorMessage  

### gounmarshal
Ao começara digitar ```gounmarshal```geramos um código que usa ```json.Unmarshal``` para "traduzir" nosso objeto ```json``` como um ```body``` de uma requisição em um struct, ou variável esperados, exemplo:  
```go
	var varname *varType
	if err = json.Unmarshal(body, &varname); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(http_error.ResponseError("ErrorMessage"))
		return
	}
```  
Lembrando que as variáveis pré selecionadas para alteração são:  
 1. varname (TAB) 
 2. varType (TAB)
 3. ErrorMessage  

### goroutespackage
Como uso uma arquitetura modular que venho desenvolvendo a partir de [outro projeto meu](https://github.com/kauemurakami/getx_pattern), meus packages são meus modulos, cada package contido em ```/internal```, deve conter os mesmo arquivos e definições, falarei por alto aqui para exemplificar o porquê de tratar cas rotas por módulo e centralizar sua inicialização nos próximos snippets.  
Imagine que temos nosso diretório ```/internal``` que é responsável por todos os nossos packages, imagine um package ```/auth``` básico, teríamos esses arquivos:  
 + auth.go -> package auth
 + auth_test.go -> package auth
 + auth_routes.go -> package auth  
 Todos usando o mesmo package name, onde dentro de ```auth_routes.go``` vamos começar digitando ```goroutespackage```, antes mesmo de digitar por completo, já irá aparecer a sugestão de usá-lo.  
 Ao selecionar ele nos retorna esse código:  
 ```go
 package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupAuthRoutes(router *mux.Router) {
	router.HandleFunc("/auth/signin", Signin).Methods(http.MethodPost)
	router.HandleFunc("/auth/signup", Signup).Methods(http.MethodPost)
}
 ```  
 Mais adiante mostrarei de onde vem as funções...  
 Lembrando que as variáveis pré selecionadas para alteração são:  
 1. packageName (TAB) 
 2. FuncName no caso fica entre Setup{FuncName}Routes  
 3. Repare que o mesmo package name que escolher será usado na rota inicial, neste caso como é auth, a primeira rota fica /auth.  
 Isso deve ficar bem claro para os snippets a seguir

### goroute
Cria uma rota simples ao digitar ```goroute``` ele nos gera o seguinte código:  
```go
router.HandleFunc("/routeName", Function).Methods(http.MethodPost)
```  
Isso serve pra você adicionar rapidamente rotas ao seu arquivo de rotas criado anteriormente.
Lembrando que as variáveis pré selecionadas para alteração são:  
 1. routeName (TAB) 
 2. Function (TAB) 
 3. MethodHTTP  
  
### gosetuproutes
Ao começar a digitar ```gosetuproutes``` ele nos mostrará um snippet para gerar uma código em um arquivo vazio que irá reunir todas as todas que adicionamos nos packages para uma inicialização única, então seguindo o exemplo das rotas de ```/auth``` vamo ver como fica nosso arquivo ```routes.go```,
lembrando que pra isso eu criei um diretório na raiz do projeto chamado ```/routes``` que contém o nosso arquivo ```routes.go```, fiz isso pois esse arquivo irá servir todos os nossos packages,portanto deve estar separado de certa forma, para servir a aplicação.  
Ao criar o repositório e o arquivo.go vamos ver o resultado do nosso snippet:  
```go
// /routes/routes.go
package routes
// imports automaticos com crtl + s
func SetupAppRoutes() *mux.Router {
	// users.SetupUserRoutes(router)
	router := mux.NewRouter()
	auth.SetupAuthRoutes(router)
	return router
}
```
Após isso, basta usar essa função ```SetupAppRoutes``` na main dessa forma:  
```go
func main(){
  ...
  router := routes.SetupAppRoutes()
  ...
}
```
E todas as suas rotas estarão disponíveis

### goif  
Ao digitarmos ```goif``` e selecionarmos o auto complete, veremos o seguinte trecho de código simples:  
```go
if var operator val {

}
```
Onde estará préselecionado nossa ```var```, para que altere com o nome de alguma variável ou valor que queira comparar, o ```operator``` que pode ser >, ==, por exemplo, e por fim ```value``` o valor que estamos usando para comparação, tudo podendo ser definido apenas com o teclado usando TAB para alternar de uma para outra.<br/>

### gofor  
![](assets/gofor.gif)  

Ao digitarmos ```gofor``` ele criará um for simples, com inicialização de variável dentro dele, vejamos:  
```go
for var := initial-value; var operator value ; var factor {
 
}
```
Onde var vem inicialmente pré selecionada para ser definida como quiser, exemplo ```k:= 0```, logo em seguida ao dar um TAB escolherá o ```initial value``` de ```var```, todos os ```var```, serão completados juntamente com a primeira definição de ```var``` ou seja se inicialmente o primeiro ```var``` for ```k```, todos os outros ```var``` também serão, entre isso também, ao dar um TAB, também devemos definir o operador para parar o ```for```, == por exemplo, em seguida definimos o ```factor``` que pode ser um ```++``` ```--``` por exemplo.<br/>

### goforrange   
![](assets/goforrange.gif)  

Aqui unimos nosso ```for``` com ```range``` para recuperar valores, ao digitar ```goforrange``` ele nos gerará este trecho de código:  
```go
for index, obj := range array-or-slice {
				
}
```
Lembrando sempre que esses "nomes" definidos vão poder ser alterado navegando entre eles com TAB.<br/>


### goerr
Com ```goerr``` podemos tratar o erro com um ```log.Fatal(err)``` em um simples if, ou como desejar:  
```go
if err != nil {
  log.Fatal(err)
}
```
```log.Fatal()``` foi usado como exemplo, você poderia tratar esse erro de diferentes modos.<br/>

### goreadinput
Aqui permitimos o usuário inserir dados via terminal com ```goreadinput```, vejamos o trecho de código gerado:  
```go
fmt.Print("Enter input: ")
var input type
fmt.Scanln(&input)
```
Aqui você deve alterar o type apenas, após isso poderá receber um valor que você inserir via terminal.

