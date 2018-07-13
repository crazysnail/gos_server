package main

import (

	_ "fmt"
	"github.com/davyxu/cellnet"
	_ "github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/httpform"
	_ "github.com/davyxu/cellnet/codec/httpjson"
	"github.com/davyxu/cellnet/peer"
	httppeer "github.com/davyxu/cellnet/peer/http"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/http"
	"net/http"
	"reflect"

	"gos_server/config"
	"gos_server/proto/c2s_proto"
)


func HttpService() {
	queue := cellnet.NewEventQueue()
	p := peer.NewGenericPeer("http.Acceptor", "GameSeverHttpService", config.GameServerOut, nil)

	proc.BindProcessorHandler(p, "http", func(raw cellnet.Event) {
	
		switch msg := raw.Message().(type) {
		case *c2s_proto.HttpEchoREQ:
			config.LogGameServer.Debugln(reflect.TypeOf(msg),msg)
			raw.Session().Send(&httppeer.MessageRespond{
				StatusCode: http.StatusOK,
				Msg: &c2s_proto.HttpEchoACK{
					Status: 0,
					Token:  msg.UserName,
				},
				CodecName: "httpjson",
			})

		}

	})

	p.Start()
	queue.StartLoop()
	queue.Wait()
}

