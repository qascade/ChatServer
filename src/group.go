package main

type group struct{
	roomName string
	members map[net.Addr]*client
}