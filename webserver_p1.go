//go 1.10.4

package main
import {
  "fmt"
  "log"
  "net/http"
}

func helloHandler(w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/hello" {
    http.Error(w, "404 not found", http.StatusNotFound){
      return
    }
  if r.Method != "GET"{
    http.Error(w, "method is not supported", http.StatusNotFound){
      return
    }
  fmt.Fprintf(w, "hello!")
  }

func main() {
  // Using terminal change directory acc to video and then code
  // Code
  // You have to make 3 files, index.html, home.html and main.go
  // Here we will see mainly the go code
  fileServer := http.FileServer(http.Dir("./static"))
  http.Handle("/", fileServer)
  http.HandleFunc("/hello", helloHandler)
  http.HandleFunc("/form", formHandler)
  
  fmt.Printf("Starting the server at port 8080 \n")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}