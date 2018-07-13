package config

import (
	"github.com/davyxu/golog"
)

var LogCenterServer = golog.New("CenterServer")
var LogChatServer = golog.New("ChatServer")
var LogGameServer = golog.New("GameServer")
var LogLoginServer = golog.New("LoginServer")

var MongoDBInner = ""
var MongoDBName = ""
var DomainName=""

var CenterServerInner = ""

var LoginServerOut = ""
var ChatServerOut = ""
var GameServerOut = ""

var LogConfig = golog.New("LogConfig")

func init() {
	configs := Get()

	MongoDBInner = configs["MongoDBInner"]
	MongoDBName = configs["MongoDBName"]
	DomainName = configs["DomainName"]

	CenterServerInner = configs["CenterServerInner"]

	LoginServerOut = configs["LoginServerOut"]
	ChatServerOut = configs["ChatServerOut"]
	GameServerOut = configs["GameServerOut"]
	
}