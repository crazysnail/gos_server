package config

import (
	"github.com/davyxu/golog"
)

var LogCenterServer = golog.New("CenterServer")
var LogChatServer = golog.New("ChatServer")
var LogGameServer = golog.New("GameServer")
var LogLoginServer = golog.New("LoginServer")

var DBInner = ""
var DBName = ""
var DomainName=""

var CenterServerInner = ""

var LoginServerOut = ""
var ChatServerOut = ""
var GameServerOut = ""

var LogConfig = golog.New("LogConfig")

func init() {
	configs := Get()

	DBInner = configs["DBInner"]
	DBName = configs["DBName"]
	DomainName = configs["DomainName"]

	CenterServerInner = configs["CenterServerInner"]

	LoginServerOut = configs["LoginServerOut"]
	ChatServerOut = configs["ChatServerOut"]
	GameServerOut = configs["GameServerOut"]
	
}