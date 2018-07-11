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
	"github.com/davyxu/golog"
	
	"gos_server/config"
	"gos_server/proto"
)

var httpLog = golog.New("LoginSeverHttpService")
var httpAddr = config.LoginServerIP+":"+config.LoginServerPort

func HttpService() {
	queue := cellnet.NewEventQueue()
	p := peer.NewGenericPeer("http.Acceptor", "LoginSeverHttpService", httpAddr, nil)

	proc.BindProcessorHandler(p, "http", func(raw cellnet.Event) {
	
		switch msg := raw.Message().(type) {
		case *proto.HttpEchoREQ:
			httpLog.Debugln(reflect.TypeOf(msg),msg)
			raw.Session().Send(&httppeer.MessageRespond{
				StatusCode: http.StatusOK,
				Msg: &proto.HttpEchoACK{
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

