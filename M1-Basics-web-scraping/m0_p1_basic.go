package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main(){
	url := "http://example.webscraping.com/view/1"
	response, err := http.Get(url)
	if err != nil{
		log.Fatal(err)
	}
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil{
		log.Fatal(err)
	}
	response.Body.Close()
	fmt.Println("bytes written: ", n)
}