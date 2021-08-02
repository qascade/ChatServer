package main

import(
	"net"
	"bufio"
	"strings"
	"fmt"
)

type client struct{
	conn net.Conn
	username string
	currGroup *group
	commands chan<- command
}

func (c* client) readInput(){
	for{
		msg,err := bufio.NewReader(c.conn).ReadString('\n')
		if err!=nil {
			return 
		}
		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg," ")
		cmd := strings.TrimSpace(args[0])//Removes all the whitespaces left and right to the string.

		switch cmd{
			case "/username":
				c.commands <- command{
					id: USRNAME_CMD,
					client: c,
					args: args,
				}
			case "/join":
				c.commands <- command{
					id: JOIN_CMD,
					client: c,
					args: args,
				}
			case "/groups":
				c.commands <- command{
					id: GROUPS_CMD,
					client: c,
					args: args,
				}
			case "/msg":
				c.commands <- command{
					id: MSG_CMD,
					client: c,
					args: args,
				}
			case "/quit":
				c.commands <- command{
					id: QUIT_CMD,
					client: c,
					args: args,
				}
			default:
				c.err(fmt.Errorf("%s Command not found.",cmd))
		}
	}
}

func (c* client) err(err error){
	c.conn.Write([]byte("Err: "+ err.Error() +"\n"))
}

func (c* client) msg(msg string){
	c.conn.Write([]byte("> "+ msg +"\n"))
}