package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fanux/lhttp"
)

type ChatProcessor struct {
	*lhttp.BaseProcessor
}

func (this *ChatProcessor) OnMessage(ws *lhttp.WsHandler) {
	log.Printf("on OnMessage: ", ws.GetBody())
	ws.AddHeader("content-type", "image/png")
	ws.SetCommand("auth")
	ws.Send(ws.GetBody())
}

type SubPubProcessor struct {
	*lhttp.BaseProcessor
}

type UpstreamProcessor struct {
	*lhttp.BaseProcessor
}

type UploadProcessor struct {
	*lhttp.BaseProcessor
}

func (this *UploadProcessor) OnMessage(ws *lhttp.WsHandler) {
	for m := ws.GetMultipart(); m != nil; m = m.GetNext() {
		log.Print("multibody:", m.GetBody(), " headers:", m.GetHeaders())
	}
}

func main() {
	lhttp.Regist("chat", &ChatProcessor{&lhttp.BaseProcessor{}})
	lhttp.Regist("subpub", &SubPubProcessor{&lhttp.BaseProcessor{}})
	lhttp.Regist("upstream", &UpstreamProcessor{&lhttp.BaseProcessor{}})
	lhttp.Regist("upload", &UploadProcessor{&lhttp.BaseProcessor{}})

	http.Handle("/echo", lhttp.Handler(lhttp.StartServer))
	http.Handle("/", lhttp.Handler(lhttp.StartServer))
	http.HandleFunc("/https", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", "world")
	})

	http.ListenAndServe(":9096", nil)
}
