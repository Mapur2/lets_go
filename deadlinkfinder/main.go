package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	html, err := http.Get("http://rupamdev.in/")
	if err != nil {
		log.Fatal("Error in getting the web page ", err)
	}
	fmt.Println("Request is successfull")
	defer html.Body.Close()
	var data []byte
	len,_ := html.Body.Read(data)
	fmt.Println(len, "\n", string(data))
}
