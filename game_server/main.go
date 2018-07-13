package main

import (
	"gos_server/game_server/db"
)

func main() {
	done := make(chan bool)

	db.DBService()
	go HttpService()
	go TcpService()

	<-done
}
