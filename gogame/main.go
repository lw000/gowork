// gogame project main.go
package main

import (
	"fmt"
	// "github.com/golang/protobuf/proto"
	// "github.com/golang/protobuf/protoc-gen-go"
)

func main() {
	fmt.Println("Hello World!")

	p := NewPlatform(1, "levi")
	if p != nil {
		p.CreateRoom()
	}
}
