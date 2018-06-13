package main

import (
   //"fmt"
   "net/http"
   //"strings"
   "log"
)

func homePage(w http.ResponseWriter, r *http.Request) {
   http.ServeFile(w, r, "./static/index.html")
}
func BootstrapPage(w http.ResponseWriter, r *http.Request) {
   http.ServeFile(w, r, "./static/bootstrap.html")
}

func main() {
   http.HandleFunc("/", homePage) // set router
   http.HandleFunc("/bootstrap", BootstrapPage) // set router
   err := http.ListenAndServe(":9090", nil) // set listen port
   if err != nil {
      log.Fatal("ListenAndServe: ", err)
   }
}
