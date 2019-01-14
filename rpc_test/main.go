// rpc_test project main.go
package main

import (
	"flag"
	"fmt"
	"rpc_test/cli"
	"rpc_test/srv"
)

var (
	t int
)

func main() {
	flag.IntVar(&t, "t", 1, "please input -t=1 or -t=0")
	flag.Parse()
	// if flag.NFlag() == 0 {
	// 	flag.PrintDefaults()
	// 	return
	// }
	fmt.Printf("args: %d\n", t)
	if t == 1 {
		fmt.Println("server")
		srv.RunSrv()
	} else if t == 0 {
		fmt.Println("client")
		cli.RunCli()
	} else {

	}
}
