package main;

import (
	"os/exec"
	"log"
	"flag"
	"bytes"
)

var java bool
var cpp bool
var csharp bool
var js bool
var objectifc bool
var python bool
var ruby bool
var golang bool

var wireJava bool

func main() {
	initialisation()
	chooseOptions()
}

func initialisation() {
	flag.BoolVar(&java, "java", false, "Java Protobuf Version")
	flag.BoolVar(&cpp, "cpp", false, "C++ Protobuf Version")
	flag.BoolVar(&csharp, "csharp", false, "C# Protobuf Version")
	flag.BoolVar(&js, "js", false, "JavaScript Protobuf Version")
	flag.BoolVar(&objectifc, "objc", false, "Objectif C Protobuf Version")
	flag.BoolVar(&python, "py", false, "Python Protobuf Version")
	flag.BoolVar(&ruby, "ruby", false, "Ruby Protobuf Version")
	flag.BoolVar(&golang, "go", false, "Go Protobuf Version")
	flag.BoolVar(&wireJava, "wire", false, "Java Wire Protobuf Version")
	flag.Parse()
}

func createProtobuf(format string) {
	cmd := exec.Command("./protoc", "--"+format+"_out=proto", "hello.proto")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("%s : ", err.Error())
	    log.Println(stderr.String())
	}
}

func createJar() {
	cmd := exec.Command("javac", "-cp", "protobuf-java-3.0.0-beta-3.jar", "proto/proto/hello/HelloOuterClass.java", "-source", "1.6", "-target", "1.6", "-d", "class/" )
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("%s : ", err.Error())
	    log.Println(stderr.String())
	}
	cmd = exec.Command("jar", "cvf", "javaversion.jar", "-C", "class/", ".")
	var stderr2 bytes.Buffer
	cmd.Stderr = &stderr2
	err = cmd.Run()
	if err != nil {
		log.Printf("%s : ", err.Error())
	    log.Println(stderr2.String())
	}
}

func createWireJar() {
	cmd := exec.Command("javac", "-cp", "okio-1.0.0.jar:wire-runtime-2.2.0.jar", "proto/proto/hello/Hello.java", "-source", "1.6", "-target", "1.6", "-d", "classWire/" )
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("%s : ", err.Error())
	    log.Println(stderr.String())
	}
	cmd = exec.Command("jar", "cvf", "javaversionWire.jar", "-C", "classWire/", ".")
	var stderr2 bytes.Buffer
	cmd.Stderr = &stderr2
	err = cmd.Run()
	if err != nil {
		log.Printf("%s : ", err.Error())
	    log.Println(stderr2.String())
	}
}

func chooseOptions() {
	if java {
		createProtobuf("java")
		createJar()
	}
	if golang {
		createProtobuf("go")
	}
	if cpp {
		createProtobuf("cpp")
	}
	if csharp {
		createProtobuf("csharp")
	}
	if js {
		createProtobuf("js")
	}
	if objectifc {
		createProtobuf("objc")
	}
	if python {
		createProtobuf("python")
	}
	if ruby {
		createProtobuf("ruby")
	}
	if wireJava {
		cmd := exec.Command("java", "-jar", "wire-compiler-2.2.0-jar-with-dependencies.jar", "--java_out=proto", "--proto_path=.", "hello.proto")
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			log.Printf("%s : ", err.Error())
		    log.Println(stderr.String())
		}
		createWireJar()
	}
}