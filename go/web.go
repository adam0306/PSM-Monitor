package main

import (
   "io/ioutil"
   "log"
   "net/http"
)

func main() {
   resp, err := http.Get("http://ifconfig.so/")
   if err != nil {
      log.Fatalln(err)
   }
//We Read the response body on the line below.
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
body, err := ioutil.ReadAll(resp.Body)
  if body = fail {
     log.Fatalln(err) 
}

func main() {
   http.Handle("/", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprint(w, "Hello, World!")
}