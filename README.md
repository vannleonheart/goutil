### Utilities for Go
#### Installation
```go
go get -u github.com/vannleonheart/goutil
```

<br />

#### Random String
Generate Using Default Character Set
```go
randomString := goutil.NewRandomString("")

// random string with exactly 100 characters in length
length := 100
exactRandomString := randomString.Generate(length)

// random string with random length between 10 and 25 characters
minLength := 10
maxLength := 25
rangedRandomString := randomString.GenerateRange(minLength, maxLength)

fmt.Println(exactRandomString)
fmt.Println(rangedRandomString)
```
Set Custom Character Set
```go
newCharset := "abc123@#$"

randomString.SetCharset(newCharset)

// random string with exactly 10 characters in length
// using only character set abc123@#$
fmt.Println(randomString.Generate(10))
```
Use From Predefined Character Set
```go
// random string with exactly 15 characters in length
// using hexadecimal character set
fmt.Println(randomString.WithCharset(goutil.HexadecimalCharset).Generate(15))
```
Available Predefined Character Set

| Const              | Character Set                                          |
|--------------------|--------------------------------------------------------|
| AlphaCharset       | abcdefghijklmnopqrstuvwxyz                             |
| AlphaUCharset      | ABCDEFGHIJKLMNOPQRSTUVWXYZ                             |
| NumCharset         | 0123456789                                             |
| AlphaNumCharset    | AlphaCharset + NumCharset                              |
| AlphaUNumCharset   | AlphaUCharset + NumCharset                             |
| AlphaAllNumCharset | AlphaCharset + AlphaUNumCharset                        |
| HexadecimalCharset | NumCharset + "abcdef"                                  |
| SymbolCharset      | \~\!\@\#\$\%\^\&\*\(\)\_\-\+\=\[\{\}\]\|\;\:\,\<\.\>\? |

<br />

#### Config
Load From JSON File
```go
jsonFilePath := "{your_json_file_path}"

var output map[string]interface{}

ptrByteFileContent, err := goutil.LoadJsonFile(jsonFilePath, &output)

if err != nil {
    // handle error
}

// convert pointer of []byte to string
// will return json file content in string
stringFileContent := string(*ptrByteFileContent)

fmt.Println(stringFileContent)


// will unmarshall json string to desired data type
// in this case to map[string]interface{}
fmt.Println(output)
```

<br />

#### HTTP Request
Generate Query String
```go
toQueryString := map[string]interface{}{
	"page": 1,
	"search": "keyword",
}

queryString, err := goutil.GenerateQueryString(toQueryString)

if err != nil {
    // handle error
}

// will print page=1&search=keyword
fmt.Println(*queryString)
```
Sending Http Request
```go
var result map[string]interface{}

targetUrl := "{your_target_url}"

requestData := map[string]interface{}{
	"page": 1,
}

requestHeaders := map[string]string{
	"Content-Type": "application/json",
}

ptrByteResponseBody, err := goutil.SendHttpRequest(http.MethodGet, targetUrl, &requestData, &requestHeaders, &result)

if err != nil {
    // handle error
}

// convert pointer of []byte to string
// will return response body in string
stringResponseBody := string(*ptrByteResponseBody)

fmt.Println(stringResponseBody)

// will unmarshall json string to desired data type
// in this case to map[string]interface{}
fmt.Println(result)
```
Sending HTTP Get Request
```go
ptrByteResponseBody, err := goutil.SendHttpGet(targetUrl, &requestData, &requestHeaders, &result)
```
Sending HTTP Post Request
```go
ptrByteResponseBody, err := goutil.SendHttpPost(targetUrl, &requestData, &requestHeaders, &result)
```
Sending HTTP Put Request
```go
ptrByteResponseBody, err := goutil.SendHttpPut(targetUrl, &requestData, &requestHeaders, &result)
```
Sending HTTP Patch Request
```go
ptrByteResponseBody, err := goutil.SendHttpPatch(targetUrl, &requestData, &requestHeaders, &result)
```
Sending HTTP Delete Request
```go
ptrByteResponseBody, err := goutil.SendHttpDelete(targetUrl, &requestHeaders, &result)
```