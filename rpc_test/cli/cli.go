package cli

import (
	"fmt"
	"log"
	"net"
	"os"
	"rpc_test/echo"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func RunCli() {
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9898"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := echo.NewEchoClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:9898", " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	for i := 0; i < 10; i++ {
		req := &echo.EchoReq{Msg: "You are welcome.", Tag: int32(i)}
		res, err := client.Echo(req)
		if err != nil {
			log.Println("Echo failed:", err)
			return
		}

		log.Printf("response[%d]: %v", res.GetTag(), res.Msg)
	}
	fmt.Println("well done")
}
