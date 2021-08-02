package main

type commandID int

const(
	USRNAME_CMD commandID = iota
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