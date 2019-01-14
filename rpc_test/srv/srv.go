package srv

import (
	"fmt"
	"rpc_test/echo"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type EchoServerImp struct {
}

func (e *EchoServerImp) Echo( /*ctx context.Context, */ req *echo.EchoReq) (*echo.EchoRes, error) {
	fmt.Printf("message from client[%d]: %v\n", req.GetTag(), req.GetMsg())

	res := &echo.EchoRes{
		Msg: req.GetMsg(),
		Tag: req.GetTag(),
	}

	return res, nil
}

func RunSrv() {
	transport, err := thrift.NewTServerSocket(":9898")
	if err != nil {
		panic(err)
	}

	processor := echo.NewEchoProcessor(&EchoServerImp{})
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		thrift.NewTBufferedTransportFactory(8192),
		thrift.NewTCompactProtocolFactory(),
	)

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
