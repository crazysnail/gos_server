package main

import (
	_ "github.com/davyxu/golog"

	"gos_server/center_server/db"
)

func main() {
	done := make(chan bool)
	db.DBService()
	go TcpService()
	<-done
}
