package main

import (
	"net/http"
	"github.com/golang/protobuf/proto"
	"log"
	"ProtoRepository/proto"
	"strings"
)

func main() {
	http.HandleFunc("/", sendProto)
	http.ListenAndServe(":8002", nil)
}

func sendProto(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	if name == "" {
		name = "Marie-Louise"
	}
	hello := new(hello.Hello)
	hello.Name = strings.ToUpper(name)
	hello.Number = 32
	w.Header().Set("Content-Type", "application/x-protobuf")
	helloCompress, err := proto.Marshal(hello)
	if err != nil {
		log.Fatal("Compress Problem (Marshal error)")
	}
	w.Write(helloCompress)
}