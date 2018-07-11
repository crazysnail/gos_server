package main

import (
)

func main() {
	done := make(chan bool)
	go TcpService()
	<-done
}
