package main

import (
	"flag"
	"html/template"
	"net/http"
	"pb_test/ty"
	"tuyue_common/ws/packet"
	"tuyue_common/ws/srv/hub"

	log "github.com/alecthomas/log4go"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var hub *tyhub.Hub

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }} // use default options

func wsEcho(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("upgrade:", err)
		return
	}

	go func() {
		defer func() {
			log.Error("exit")
			c.Close()
		}()

		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Error(err)
				break
			}

			if mt == websocket.TextMessage {
				c.WriteMessage(mt, message)
			} else {
				hub.DispatchMessage(c, mt, message)
			}
		}
	}()
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/ws")
}

func registerFunc() {
	hub = tyhub.NewWSHub()

	hub.Handle(0x0203, 0x0001, func(c *websocket.Conn, mt int, pk *typacket.Packet) {
		var req ty.ReqEcho
		if err := proto.Unmarshal(pk.Bytes(), &req); err != nil {
			log.Error(err)
			return
		}
		log.Info("[%s] [%d] %+v\n", c.RemoteAddr().String(), pk.EventId(), req)

		ack := typacket.NewPacket(pk.Mid(), pk.Sid(), pk.EventId())
		ack.EncodeProto(&ty.AckEcho{Code: req.Id, Data: req.Data})
		if err := c.WriteMessage(mt, ack.Bytes()); err != nil {
			log.Error(err)
		}
	})

	hub.Handle(0x0203, 0x0002, func(c *websocket.Conn, mt int, pk *typacket.Packet) {
		var req ty.ReqLogin
		if err := proto.Unmarshal(pk.Bytes(), &req); err != nil {
			log.Error(err)
			return
		}
		log.Info("[%s] [%d] %+v\n", c.RemoteAddr().String(), pk.EventId(), req)

		ack := typacket.NewPacket(pk.Mid(), pk.Sid(), pk.EventId())
		ack.EncodeProto(&ty.AckLogin{Code: req.Id, Data: req.Data})
		if err := c.WriteMessage(mt, ack.Bytes()); err != nil {
			log.Error(err)
		}
	})

	hub.Handle(0x0203, 0x0003, func(c *websocket.Conn, mt int, pk *typacket.Packet) {
		var req ty.ReqLogout
		if err := proto.Unmarshal(pk.Bytes(), &req); err != nil {
			log.Error(err)
			return
		}

		log.Info("[%s] [%d] %+v\n", c.RemoteAddr().String(), pk.EventId(), req)

		ack := typacket.NewPacket(pk.Mid(), pk.Sid(), pk.EventId())
		ack.EncodeProto(&ty.AckLogout{Code: req.Id, Data: req.Data})
		if err := c.WriteMessage(mt, ack.Bytes()); err != nil {
			log.Error(err)
		}
	})
}

func main() {
	log.LoadConfiguration("../configs/log4go.xml")

	registerFunc()

	flag.Parse()

	http.HandleFunc("/ws", wsEcho)

	http.HandleFunc("/", home)

	log.Error(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
	<!DOCTYPE html>
	<html>
	<head>
	<meta charset="utf-8" />
	<script>  
		window.addEventListener("load", function(evt) {
	
	    var output = document.getElementById("output");
	    var input = document.getElementById("input");
	    var ws;
	
	    var print = function(message) {
	        var d = document.createElement("div");
	        d.innerHTML = message;
	        output.appendChild(d);
	    };
	
	    document.getElementById("open").onclick = function(evt) {
	        if (ws) {
	            return false;
	        }
	        ws = new WebSocket("{{.}}");
	        ws.onopen = function(evt) {
	            print("OPEN");
	        }
	        ws.onclose = function(evt) {
	            print("CLOSE");
	            ws = null;
	        }
	        ws.onmessage = function(evt) {
	            print("RESPONSE: " + evt.data);
	        }
	        ws.onerror = function(evt) {
	            print("ERROR: " + evt.data);
	        }
	        return false;
	    };
	
	    document.getElementById("send").onclick = function(evt) {
	        if (!ws) {
	            return false;
	        }
	        print("SEND: " + input.value);
	        ws.send(input.value);
	        return false;
	    };
	
	    document.getElementById("close").onclick = function(evt) {
	        if (!ws) {
	            return false;
	        }
	        ws.close();
	        return false;
	    };
	
	});
	</script>
	</head>
	<body>
	<table>
	<tr><td valign="top" width="50%">
	<p>Click "Open" to create a connection to the server, 
	"Send" to send a message to the server and "Close" to close the connection. 
	You can change the message and send multiple times.
	<p>
	<form>
	<button id="open">Open</button>
	<button id="close">Close</button>
	<p><input id="input" type="text" value="Hello world!">
	<button id="send">Send</button>
	</form>
	</td><td valign="top" width="50%">
	<div id="output"></div>
	</td></tr></table>
	</body>
	</html>
`))
