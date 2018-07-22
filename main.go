package main

import(
	"github.com/amazingfly/rtmpServer/functions"
	)

func main(){
	addr := "8888"
	sKey := "123"
	go rtmpServer.ServeRTMP(&addr, &sKey)
	addr2 := "8089"
	go rtmpServer.ServeRTMP(&addr2, &sKey)
	var forever = make(chan bool)
	<-forever
}
