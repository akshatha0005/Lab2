package main
import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
)
type Req struct{
    Name string `json:"name"`
}
type Response struct{
    Greeting string `json:"greeting"`
}
func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}
func hello1(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
var v Req
defer req.Body.Close()
decoder := json.NewDecoder(req.Body)
err := decoder.Decode(&v)
if err != nil {
panic(err)
}
str := fmt.Sprintf("Hello, %v",v.Name)
b:= &Response{Greeting: str}
rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
rw.WriteHeader(http.StatusOK)
if err := json.NewEncoder(rw).Encode(b); err != nil {
panic(err)
}
}
func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", hello)
    mux.POST("/hello1/",hello1)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}