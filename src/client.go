package main

import(
	"net"
)
type client struct{
	conn net.Addr
	username string
	currRoom *room
	command chan<- commandID
}