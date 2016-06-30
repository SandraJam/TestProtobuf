package main

import (
	"fmt"
	"net/http"
	"flag"
	"log"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"ProtoRepository/proto"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "", "Name for the server")
	flag.Parse()

	resp, err := http.Get("http://localhost:8002/"+name)
	if err != nil {
		log.Fatal("Get request Problem: "+err.Error())
	}

	defer resp.Body.Close()

	hello := new(hello.Hello)
	byteProto, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Read Problem")
	}
	err = proto.Unmarshal(byteProto, hello)
	if err != nil {
		log.Fatal("Unmarshal Problem "+err.Error())
	}
	fmt.Println(hello.Name)
}
