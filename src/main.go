package main

import(
	"log"
	"net"
)

func main(){
	s := newServer() // Initialising a new server 
	listener,err := net.Listen("tcp",":8888")
	if err!=nil {
		log.Fatalf("unable to start server: %s",err.Error())
		//An error will be thrown if not able to connect to the server.
	}

	defer listener.Close()
	log.Printf("Started Server on port 8888.")
	
	for{
		conn,err := listener.Accept()
		if err:=nil {
			log.Printf("Unable to accept a connection %s", err.Error())
			continue
		}
		go s.newClient(conn)
	}
}