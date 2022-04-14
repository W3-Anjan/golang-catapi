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


#### parse json response 
https://dev.to/billylkc/parse-json-api-response-in-go-10ng
https://mholt.github.io/json-to-go/

