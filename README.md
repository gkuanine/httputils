httptuils
=======
A simple `HTTP Request` package for golang. `GET` `POST` `DELETE` `PUT` change from github.com/kirinlabs/HttpRequest



### Installation


go get github.com/gkuanine/httputils

### How do we use httptuils?

#### Create request object use http.DefaultTransport

resp, err := httputils.Get("http://127.0.0.1:8000")
resp, err := httputils.SetTimeout(5).Get("http://127.0.0.1:8000")
resp, err := httputils.Debug(true).SetHeaders(map[string]string{}).Get("http://127.0.0.1:8000")

OR

req := httputils.NewRequest()
req := httputils.NewRequest().Debug(true).SetTimeout(5)
resp, err := req.Get("http://127.0.0.1:8000")
resp, err := req.Get("http://127.0.0.1:8000",nil)
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=httputils")
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=httputils","address=beijing")

```

#### Set headers
```go
req.SetHeaders(map[string]string{
    "Content-Type": "application/x-www-form-urlencoded",
    "Connection": "keep-alive",
})

req.SetHeaders(map[string]string{
    "Source":"api",
})
```

#### Set cookies
```go
req.SetCookies(map[string]string{
    "name":"json",
    "token":"",
})

OR

httputils.SetCookies(map[string]string{
    "age":"19",
}).Post()
```

#### Set basic auth
```go
req.SetBasicAuth("username","password")
```

#### Set timeout
```go
req.SetTimeout(5)  //default 30s
```

#### Transport
If you want to customize the Client object and reuse TCP connections, you need to define a global http.RoundTripper or & http.Transport, because the reuse of the http connection pool is based on Transport.
```go
var transport *http.Transport
func init() {   
    transport = &http.Transport{
        DialContext: (&net.Dialer{
            Timeout:   30 * time.Second,
            KeepAlive: 30 * time.Second,
            DualStack: true,
        }).DialContext,
        MaxIdleConns:          100, 
        IdleConnTimeout:       90 * time.Second,
        TLSHandshakeTimeout:   5 * time.Second,
        ExpectContinueTimeout: 1 * time.Second,
    }
}

func demo(){
    // Use http.DefaultTransport
    res, err := httputils.Get("http://127.0.0.1:8080")
    // Use custom Transport
    res, err := httputils.Transport(transport).Get("http://127.0.0.1:8080")
}
```

#### Keep Alives???Only effective for custom Transport
```go
req.DisableKeepAlives(false)

httputils.Transport(transport).DisableKeepAlives(false).Get("http://127.0.0.1:8080")
```

#### Ignore Https certificate validation???Only effective for custom Transport
```go
req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})

httputils.Transport(transport).SetTLSClient(&tls.Config{InsecureSkipVerify: true}).Get("http://127.0.0.1:8080")
```

#### Object-oriented operation mode
```go
req := httputils.NewRequest().
	Debug(true).
	SetHeaders(map[string]string{
	    "Content-Type": "application/x-www-form-urlencoded",
	}).SetTimeout(5)
resp,err := req.Get("http://127.0.0.1")

resp,err := httputils.NewRequest().Get("http://127.0.0.1")
```

### GET

#### Query parameter
```go
resp, err := req.Get("http://127.0.0.1:8000")
resp, err := req.Get("http://127.0.0.1:8000",nil)
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=httputils")
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=httputils","address=beijing")

OR

resp, err := httputils.Get("http://127.0.0.1:8000")
resp, err := httputils.Debug(true).SetHeaders(map[string]string{}).Get("http://127.0.0.1:8000")
```


#### Multi parameter
```go
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=httputils",map[string]interface{}{
    "name":  "jason",
    "score": 100,
})
defer resp.Close()

body, err := resp.Body()
if err != nil {
    return
}

return string(body)
```


### POST

```go
// Send nil
resp, err := httputils.Post("http://127.0.0.1:8000")

// Send integer
resp, err := httputils.Post("http://127.0.0.1:8000", 100)

// Send []byte
resp, err := httputils.Post("http://127.0.0.1:8000", []byte("bytes data"))

// Send io.Reader
resp, err := httputils.Post("http://127.0.0.1:8000", bytes.NewReader(buf []byte))
resp, err := httputils.Post("http://127.0.0.1:8000", strings.NewReader("string data"))
resp, err := httputils.Post("http://127.0.0.1:8000", bytes.NewBuffer(buf []byte))

// Send string
resp, err := httputils.Post("http://127.0.0.1:8000", "title=github&type=1")

// Send JSON
resp, err := httputils.PostJson("http://127.0.0.1:8000", "{\"id\":10,\"title\":\"httputils\"}")

// Send map[string]interface{}{}
resp, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
    "id":    10,
    "title": "httputils",
})
defer resp.Close()

body, err := resp.Body()
if err != nil {
    return
}
return string(body)

resp, err := httputils.Post("http://127.0.0.1:8000")
resp, err := httputils.JSON().Post("http://127.0.0.1:8000",map[string]interface{}{"title":"github"})
resp, err := httputils.Debug(true).SetHeaders(map[string]string{}).JSON().Post("http://127.0.0.1:8000","{\"title\":\"github\"}")
```

### Jar
```go
j, _ := cookiejar.New(nil)
j.SetCookies(&url.URL{
	Scheme: "http",
	Host:   "127.0.0.1:8000",
}, []*http.Cookie{
	&http.Cookie{Name: "identity-user", Value: "83df5154d0ed31d166f5c54ddc"},
	&http.Cookie{Name: "token_id", Value: "JSb99d0e7d809610186813583b4f802a37b99d"},
})
resp, err := httputils.Jar(j).Get("http://127.0.0.1:8000/city/list")
defer resp.Close()

if err != nil {
	log.Fatalf("Request error???%v", err.Error())
}
```

### Proxy
```go
proxy, err := url.Parse("http://proxyip:proxyport")
if err != nil {
	log.Println(err)
}

resp, err := httputils.Proxy(http.ProxyURL(proxy)).Get("http://127.0.0.1:8000/ip")
defer resp.Close()

if err != nil {
	log.Println("Request error???%v", err.Error())
}

body, err := resp.Body()
if err != nil {
	log.Println("Get body error???%v", err.Error())
}
log.Println(string(body))
```

### Upload
Params: url, filename, fileinput

```go
resp, err := req.Upload("http://127.0.0.1:8000/upload", "/root/demo.txt","uploadFile")
body, err := resp.Body()
defer resp.Close()
if err != nil {
    return
}
return string(body)
```


### Debug
#### Default false

```go
req.Debug(true)
```

#### Print in standard output???
```go
[httputils]
-------------------------------------------------------------------
Request: GET http://127.0.0.1:8000?name=iceview&age=19&score=100
Headers: map[Content-Type:application/x-www-form-urlencoded]
Cookies: map[]
Timeout: 30s
ReqBody: map[age:19 score:100]
-------------------------------------------------------------------
```


## Json
Post JSON request

#### Set header
```go
 req.SetHeaders(map[string]string{"Content-Type": "application/json"})
```
Or
```go
req.JSON().Post("http://127.0.0.1:8000", map[string]interface{}{
    "id":    10,
    "title": "github",
})

req.JSON().Post("http://127.0.0.1:8000", "{\"title\":\"github\",\"id\":10}")
```

#### Post request
```go
resp, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
    "id":    10,
    "title": "httputils",
})
```

#### Print formatted JSON
```go
str, err := resp.Export()
if err != nil {
   return
}
```

#### Unmarshal JSON
```go
var u User
err := resp.Json(&u)
if err != nil {
   return err
}

var m map[string]interface{}
err := resp.Json(&m)
if err != nil {
   return err
}
```

### Response

#### Response() *http.Response
```go
resp, err := req.Post("http://127.0.0.1:8000/") //res is a http.Response object
```

#### StatusCode() int
```go
resp.StatusCode()
```

#### Body() ([]byte, error)
```go
body, err := resp.Body()
log.Println(string(body))
```

#### Close() error
```go
resp.Close()
```

#### Time() string
```go
resp.Time()  //ms
```

#### Print formatted JSON
```go
str, err := resp.Export()
if err != nil {
   return
}
```

#### Unmarshal JSON
```go
var u User
err := resp.Json(&u)
if err != nil {
   return err
}

var m map[string]interface{}
err := resp.Json(&m)
if err != nil {
   return err
}
```

#### Url() string
```go
resp.Url()  //return the requested url
```

#### Headers() http.Header
```go
resp.Headers()  //return the response headers
resp.Headers().Get("Content-Type")
```

#### Cookies() []*http.Cookie
```go
resp.Cookies()  //return the response cookies
```

### Advanced
#### GET
```go
import "github.com/gkuanine/httputils"
   
resp,err := httputils.Get("http://127.0.0.1:8000/")
resp,err := httputils.Get("http://127.0.0.1:8000/","title=github")
resp,err := httputils.Get("http://127.0.0.1:8000/?title=github")
resp,err := httputils.Debug(true).JSON().Get("http://127.0.0.1:8000/")
```

#### POST
```go
import "github.com/gkuanine/httputils"
   
resp,err := httputils.Post("http://127.0.0.1:8000/")
resp,err := httputils.SetHeaders(map[string]string{
	"title":"github",
}).Post("http://127.0.0.1:8000/")
resp,err := httputils.Debug(true).JSON().Post("http://127.0.0.1:8000/")
```


### Example
```go
import "github.com/gkuanine/httputils"
   
resp,err := httputils.Get("http://127.0.0.1:8000/")
resp,err := httputils.Get("http://127.0.0.1:8000/","title=github")
resp,err := httputils.Get("http://127.0.0.1:8000/?title=github")
resp,err := httputils.Get("http://127.0.0.1:8000/",map[string]interface{}{
	"title":"github",
})
resp,err := httputils.Debug(true).JSON().SetHeaders(map[string]string{
	"source":"api",
}).SetCookies(map[string]string{
	"name":"httputils",
}).Post("http://127.0.0.1:8000/")


//Or
req := httputils.NewRequest()
req := req.Debug(true).SetHeaders()
resp,err := req.Debug(true).JSON().SetHeaders(map[string]string{
    "source":"api",
}).SetCookies(map[string]string{
    "name":"httputils",
}).Post("http://127.0.0.1:8000/")
```
# httputils
