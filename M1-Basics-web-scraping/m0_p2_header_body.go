package main

import (
	"fmt"
	"net/http"
	"time"
)

func DisplayError(err error){
	if err != nil{
		fmt.Println(err)
	}
}

func main(){
	url := "http://example.webscraping.com/view/1"
	client := http.Client{Timeout:time.Second*10}
	res, err := client.Get(url)
	DisplayError(err)
	fmt.Println(res.Body)
}