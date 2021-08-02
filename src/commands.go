package main

import(
	"fmt"
)
type commandID int

const(
	_ = iota
	USRNAME_CMD  
	JOIN_CMD
	GROUPS_CMD
	MSG_CMD
	QUIT_CMD
)
type command struct{
	id commandID
	client *client
	args []string
}