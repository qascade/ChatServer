package main

import(
	"net"
)

type group struct{
	groupName string
	members map[net.Addr]*client
}

func (g *group) broadcast(sender *client, msg string){
	for addr, m:= range g.members{
		if addr !=sender.conn.RemoteAddr(){
			m.msg(msg)
		}
	}
}		