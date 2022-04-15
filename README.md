#### Enable Golang in VSCode 
https://docs.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code

#### create new beego project inside project folders.
Make sure ``bin folder`` is added in the path. It is the golang path. Then run:
```bee new golang-catapi```
This will create a new golang project with beego
Run ``go mod tidy``
Run ``go get``
Now enter inside project folder and ``bee run`` it will open http://localhost:8080/

#### GOLang basics
1. Variadic functions can be called with any number of trailing arguments. For example, ``fmt.Println`` is a common variadic function.
2. In Go the concept of char is called ``Rune``
3. Go supports methods defined on struct types.

#### parse json response 
https://dev.to/billylkc/parse-json-api-response-in-go-10ng
https://mholt.github.io/json-to-go/

#### Project Structure
1. This project contains the golang features 
2. GET and POST request from backend go project to frontend and parse the response

#### Project output 
<img align="left" src="master/static/img/1.png" width=100 >![Search Cat and Make Favorite] (https://github.com/W3-Anjan/golang-catapi/blob/master/static/img/1.PNG?raw=true)
![Get all Favorite Cats] (https://github.com/W3-Anjan/golang-catapi/blob/master/static/img/2.PNG?raw=true)

