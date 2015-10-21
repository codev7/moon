package main

import "../../../server"

func main() {
	srvr := server.Server{}
	srvr.ParseConfig()
	srvr.Bootstrap()
}
