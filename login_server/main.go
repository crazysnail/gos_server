package main

import (

	//"fmt"
	//"github.com/davyxu/cellnet"
	//"github.com/davyxu/cellnet/codec"
	//_ "github.com/davyxu/cellnet/codec/httpform"
	//_ "github.com/davyxu/cellnet/codec/httpjson"
	//"github.com/davyxu/cellnet/peer"
	//httppeer "github.com/davyxu/cellnet/peer/http"
	//"github.com/davyxu/cellnet/proc"
	//_ "github.com/davyxu/cellnet/proc/http"
	//"net/http"
	//"reflect"
	//"testing"
	//"github.com/davyxu/golog"
)

func main() {
	done := make(chan bool)

	go HttpService()
	go TcpService()

	<-done
}

