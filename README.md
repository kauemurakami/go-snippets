## Go Fast Snippets
[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-snippets/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-snippets/blob/main/README.md)
[![Star on GitHub](https://img.shields.io/github/stars/kauemurakami/go-snippets.svg?style=flat&logo=github&colorB=deeppink&label=stars)](https://github.com/kauemurakami/get_snippets_extension)

<marquee direction="right" scrollamount=40>![](gofast.png)</marquee> 

#### Extension for vscode with autocomplete snippets for GO codes.
*Speed ​​up* your development process with shortcuts that free you from rewriting routine code, such as functions, variables, structs, etc., see the available snippets:<br/>
Here are some examples of using code snippets. You can see the rest of the options below, in this README.<br/> 

### gomain
![](assets/gomain.gif)  
When writing ```gomain``` in a file and selecting it for autocomplete it will generate the basic code of a file, for example, start by writing ```gomain```, or just ```go``` and you will see the full snippet options:   
```go
package main

func main() {
	
}
```

### goinit
![](assets/goinit.gif)  
Here, by typing ```goinit``` we create the ```init()``` function that starts before all the functions in the .go file, let's see:   
```go
func init() {
			
}
```
Aqui não precisamos definir nenhum parâmetro.<br/>

### gofile
![](assets/gofile.gif)   
When typing ```gofile``` in an empty file it generates the basic code of a ```.go``` file, with ```name``` pre-selected for change:  
```go
package name

func init() {
	
}

func main() {
	
}
```  
```gofile``` generates a code that is a little more complete than ```gomain```, also providing us with ```goinit``` <br/>

### gotestfile  
When writing ```gotestfile``` it will return us a code for a file that contains the import of ```testing``` and creation of the package with the initial test function, keeping the word ```Test``` for go to recognize it as a test function, the first and only attribute to be defined is the remaining name of the function after ```Test```, remembering that the second word after ```Test``` also starts with capital letter example:   
```go
package name

import "testing" 

func TestNameFunction(t *testing.T) {
	
}
```  

### gofunc
![](assets/gofunc.gif)  
When writing ```gofunc``` it will return us a function with autocompletes in the function name, parameter and return type, just switch with the TAB key after entering one by one, let's see what ```gofunc``` generates us:  
```go
//gofunc
func name(params type)  returnType {
  return params
}
```
Remembering that, initially ```name``` is the first attribute to be inserted, from a TAB and will jump to the ```params```, being able to insert one or more parameters separated by a comma, from another TAB and it will be in ```type``` which is the type of ```params``` in the last TAB it will define the return type ```returnType```, do the test.<br/>

### gofunctest
It generates a test function with the default parameters, see:  
```go
func TestNameFunction(t *testing.T) {
	
}
```   
### gofuncempty  
![](assets/gofuncempty.gif)  
Here you have the option of creating an empty function, without parameters and without returns, so when typing ```gofun``` it will appear, but in the case of an empty function we will use the command ```gofuncempty``` which gives us will generate:   
```go
func name() {
		
}
```
Always remember that ```name``` has already been selected for name change.<br/>

### gofuncerr
![](assets/gofuncerr.gif)  
A function that can receive an error and handle it, just start typing ```gofuncerr``` and we will have the following result:  
```go
func name(params) (returnType, error) {
	
}
```
With the same modus operandi as other snippets that have more than one value to be changed, this function also follows these patterns and can be switched using TAB when declaring a value and going to the next.<br/>

### gofunchttp
Typing ```gofunchttp``` will generate a basic http function like:   
```go
func FunctionName(w http.ResponseWriter, r *http.Request) {
	
}
```
Crtl+S auto-imports the packages that will be used but have not yet been imported.  

### gostruct
![](assets/gostruct.gif)  
Here, when typing ```gostruct``` you will create a model of a struct, let's see:  
```go
type name struct {
  varname type
}
```
Starting by defining ```name``` with a TAB we can define ```varname``` and with another TAB we define the ```type```<br/>

### gointerface  
![](assets/gointerface.gif)  

Here, when we type ```gointerface``` it will give us a snippet of an interface: 
```go
type name interface {
  
}
```
Leaving the name pre-selected for change.<br/>

### gointerfacegeneric
![](assets/gointerfacegeneric.gif)  

Here, when we type ```gointer..``` it will show us the snippet we are looking for ```gointerfacegeneric```, when selecting it, it will generate the following code:   
```go
func name(nameinterface interface{}) {
  
}
```
As the first pre-selected attribute for renaming, ```name``` and with a TAB we also define ```nameinterface``` <br/>

### goroutine
Returns a function with the clause ```go``` to become a ```goroutine```:   
```go
go func(params){
	
}
```
### goroutineanon
Creates an anonymous ```goroutine```:    
```go
go func(){
	
}()
```  

### goselect
A structure very similar to a switch in decision making, but focused on use with competition, just use our ```goselect``` snippet to create a basic structure:   
```go
select {
 case obj := <-channel1:
		fmt.Println("a")
	case obj := <-channel2:
		fmt.Println("b")
}
```  

### gomap
Creates a declaration of a map variable with key type and value type, in this case it is preselected for changing the ```varname```, ```keytype``` and ```valuetype```:  
```go
var varname map[keytype]valuetype
```  

### gomapvalues
With ```gomapvalues``` you will create a map by inference, already being able to assign values, let's see the code generated when using the snippet:  
```go
mapname := map[keytype]valuetype{
  "key1" : "Value1", 
  "key2" : "Value1", 
}
```
Following the same pattern in order, ```mapname``` is the first pre-selected, when navigating with TAB you go to ```keytpe``` -> ```valuetype``` -> ```key1 ``` -> ```value1``` and so on and you can change them effortlessly.<br/>


### govar  
Here, when we type ```govar``` and select auto complete, it creates a simple variable, receiving the type by inference, receiving a ```name``` and a ```type```, see:  
```go
name := val
```
Having the same feature of being self-selected to change the ```name``` and when TAB defining the ```val```.<br/>

### govartype
Here, when typing ```govar``` the option ```govartype``` will appear, unlike ```govar```, ```govartype``` creates a variable with a defined type, let's see:  
```go
var name type
```
Having the same behavior, where ```name``` is pre-selected for change and by giving a TAB you can change the ```type```.<br/>

### gowriteheader
When you start typing ```gowriteheader```, the snippet will appear, select it and it will return the following code:  
```go
w.WriteHeader(http.StatusBadRequest)
```
### gowriter
```gowriter``` brings us a complete solution to return a json response to the client, using ```json.Marshal(obj)``` to be converted into json for the variable ```varJSON`` `, we also check if anything goes wrong during this, and return an error message, if everything went well, we can return our json at the end of the function.    
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
Remembering that the variables pre-selected for change are:  
 1. varJson (TAB) 
 2. objectToConvert (TAB)
 3. ErrorMessage   

### gomarshal
When you start typing ```gomarshal``` and select the snippet, it will generate code to convert a struct or other type into a ```json``` with the ```json.Marshal``` function, see the example:   
```go
var varname *varType
varJSON, err := json.Marshal(*varname)
if err != nil {
  w.WriteHeader(http.StatusBadRequest)
  w.Write([]byte("ErrorMessage"))
}
```  
Remembering that the variables pre-selected for change are:  
 1. varname (TAB) 
 2. varType (TAB)
 3. ErrorMessage  

### gounmarshal
When you start typing ```gounmarshal``` we generate code that uses ```json.Unmarshal``` to "translate" our ```json``` object as a ```body``` of a request into a struct, or expected variable, example:   
```go
	var varname *varType
	if err = json.Unmarshal(body, &varname); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(http_error.ResponseError("ErrorMessage"))
		return
	}
```  
Remembering that the variables pre-selected for change are:  
 1. varname (TAB) 
 2. varType (TAB)
 3. ErrorMessage  

### goroutespackage
As I use a modular architecture that I have been developing from [another project of mine](https://github.com/kauemurakami/getx_pattern), my packages are my modules, each package contained in ```/internal```, must contain the same files and definitions, I will speak briefly here to exemplify why I treat each route per module and centralize its initialization in the next snippets.  
Imagine that we have our ```/internal``` directory that is responsible for all our packages, imagine a basic ```/auth``` package, we would have these files:  
 + auth.go -> package auth
 + auth_test.go -> package auth
 + auth_routes.go -> package auth
 All using the same package name, where inside ```auth_routes.go``` we will start by typing ```goroutespackage```, before even typing it completely, the suggestion to use it will appear.  
 When selected, it returns this code:   
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
 Later I will show where the functions come from...  
 Remembering that the variables pre-selected for change are:  
 1. packageName (TAB)
 2. FuncName is between Setup{FuncName}Routes
 3. Note that the same package name you choose will be used in the initial route, in this case as it is auth, the first route is /auth.  
 This should be pretty clear from the following snippets.  

### goroute
Create a simple route by typing ```goroute``` it generates the following code:    
```go
router.HandleFunc("/routeName", Function).Methods(http.MethodPost)
```  
This allows you to quickly add routes to your previously created route file.  
Remembering that the variables pre-selected for change are:  
 1. routeName (TAB)
 2. Function (TAB)
 3. MethodHTTP
  
### gosetuproutes
When you start typing ```gosetuproutes``` it will show us a snippet to generate code in an empty file that will bring together all the things we added in the packages for a single initialization, then following the example of the ```/ routes auth``` let's see what our ```routes.go``` file looks like,
remembering that for this I created a directory in the root of the project called ```/routes``` which contains our file ```routes.go```, I did this because this file will serve all our packages, so it should be separated in a certain way, to serve the application.  
When creating the repository and the.go file, we will see the result of our snippet:    
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
After that, just use this ```SetupAppRoutes``` function in main like this:   
```go
func main(){
  ...
  router := routes.SetupAppRoutes()
  ...
}
```
And all your routes will be available  

### goif  
When we type ```goif``` and select auto complete, we will see the following simple code snippet:   
```go
if var operator val {

}
```
Where our ```var``` will be preselected, so that you can change it with the name of any variable or value you want to compare, the ```operator``` which can be >, ==, for example, and finally ` ``value``` the value we are using for comparison, all of which can be defined just with the keyboard using TAB to switch from one to another.<br/>

### gofor  
![](assets/gofor.gif)  

When we type ```gofor``` it will create a simple for, with variable initialization inside it, let's see:  
```go
for var := initial-value; var operator value ; var factor {
 
}
```
Where var is initially pre-selected to be defined as you wish, for example ```k:= 0```, then when giving a TAB it will choose the ```initial value``` of ```var```, all ```var```, will be completed together with the first definition of ```var```, that is, if initially the first ```var``` is ```k```, all others ```var``` will also be, among this too, when giving a TAB, we must also define the operator to stop ```for```, == for example, then we define the ```factor``` which can be ```++``` ```--``` for example.<br/>

### goforrange   
![](assets/goforrange.gif)  

Here we join our ```for``` with ```range``` to retrieve values, when typing ```goforrange``` it will generate this code snippet:  
```go
for index, obj := range array-or-slice {
				
}
```
Always remember that these defined "names" can be changed by navigating between them with TAB.<br/>


### goerr
With ```goerr``` we can treat the error with ```log.Fatal(err)``` in a simple if, or as you wish:  
```go
if err != nil {
  log.Fatal(err)
}
```
```log.Fatal()``` was used as an example, you could handle this error in different ways.<br/>

### goreadinput
Here we allow the user to enter data via terminal with ```goreadinput```, let's see the generated code snippet:   
```go
fmt.Print("Enter input: ")
var input type
fmt.Scanln(&input)
```
Here you must only change the type, after which you can receive a value that you enter via the terminal.
