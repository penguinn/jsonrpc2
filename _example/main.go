package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/penguinn/jsonrpc2"
)
type Echo int

type (
	EchoParams struct {
		Name string `json:"name"`
	}
	EchoResult struct {
		Message string `json:"message"`
	}
)

func(p *Echo)Add(c context.Context, params *json.RawMessage) (interface{}, *jsonrpc2.Error) {
	log.Println(1)
	var param EchoParams
	if err := jsonrpc2.Unmarshal(params, &param); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return EchoResult{
		Message: "Hello, " + param.Name,
	}, nil
}



func main() {
	p := new(Echo)
	jsonrpc2.RegisterMethod("Add", p.Add, EchoParams{}, EchoResult{})
	http.HandleFunc("/v1/jrpc", jsonrpc2.Handler)
	http.HandleFunc("/v1/jrpc/debug", jsonrpc2.DebugHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
