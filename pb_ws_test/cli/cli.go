package main

import (
	"flag"
	"os"
	"os/signal"
	"pb_test/ty"
	"sync/atomic"
	"time"
	"tuyue_common/ws/cli"

	log "github.com/alecthomas/log4go"
	"github.com/golang/protobuf/proto"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

// var addr = flag.String("addr", "192.168.1.168:8830", "http service address")

func main() {
	flag.Parse()

	log.LoadConfiguration("../configs/log4go.xml")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ws := tyws.DefaultClient(30, 15, false, 0)
	err := ws.Open(*addr, "/ws")
	if err != nil {
		return
	}

	ws.Run()

	var reqCount uint32 = 0
	for i := 0; i < 1; i++ {
		go func() {
			// for {
			eventId := atomic.AddUint32(&reqCount, 1)
			req := ty.ReqEcho{Id: 1, Tm: time.Now().Unix(), Data: time.Now().Format("2006-01-02 15:04:05")}
			err := ws.AsyncSendMessage(0x0203, 0x0001, eventId, &req, func(eventId uint32, buf []byte) {
				var ack ty.AckEcho
				if err := proto.Unmarshal(buf, &ack); err == nil {
					log.Info("%+v", ack)
				} else {
					log.Error(err)
					return
				}
			})
			if err != nil {
				log.Error(err)
				return
			}
			time.Sleep(time.Millisecond * time.Duration(1))
			// }
		}()

		go func() {
			// for {
			eventId := atomic.AddUint32(&reqCount, 1)
			req := ty.ReqLogin{Id: 2, Tm: time.Now().Unix(), Data: "login"}
			err := ws.AsyncSendMessage(0x0203, 0x0002, eventId, &req, func(eventId uint32, buf []byte) {
				var ack ty.AckLogin
				if err := proto.Unmarshal(buf, &ack); err == nil {
					log.Info("%+v", ack)
				} else {
					log.Error(err)
					return
				}
			})
			if err != nil {
				log.Error(err)
				return
			}
			time.Sleep(time.Millisecond * time.Duration(1))
			// }
		}()

		go func() {
			// for {
			eventId := atomic.AddUint32(&reqCount, 1)
			req := ty.ReqLogout{Id: 3, Tm: time.Now().Unix(), Data: "logout"}
			err := ws.AsyncSendMessage(0x0203, 0x0003, eventId, &req, func(eventId uint32, buf []byte) {
				var ack ty.AckLogout
				if err := proto.Unmarshal(buf, &ack); err == nil {
					log.Info("%+v", ack)
				} else {
					log.Error(err)
					return
				}
			})
			if err != nil {
				log.Error(err)
				return
			}
			time.Sleep(time.Millisecond * time.Duration(1))
			// }
		}()
	}

	for {
		select {
		case <-interrupt:
			log.Info("interrupt")
			ws.Stop()
			select {
			case <-time.After(time.Second):
			}
			return
		}
	}
}
