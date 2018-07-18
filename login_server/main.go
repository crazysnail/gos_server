package main

import (

)

func main() {
	done := make(chan bool)

	go HttpService()

	go TcpService()

	<-done
}

