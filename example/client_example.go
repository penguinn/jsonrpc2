package main

import (
	"github.com/penguinn/jsonrpc2"
	"log"
)

func main() {
	client := jsonrpc2.NewClient("http://127.0.0.1:8080/v1/jrpc")
	params := make(map[string]interface{})
	params["name"] = "sjy"
	resp, err := client.Call("Add", params, 1)
	if err != nil{
		log.Fatal(err)
	}else {
		log.Println(resp)
	}
}
