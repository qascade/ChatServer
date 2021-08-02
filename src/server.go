package main

import(
	"log"
	"net"
	"fmt"
	"strings"
)

type server struct{
	groups map[string]*group
	commands chan command
}

func newServer() *server{
	return &server{
		groups: make(map[string]*group),
		commands: make(chan command),
	}
}

func (s* server) run(){
	for cmd:= range s.commands{
		switch cmd.id{
			case USRNAME_CMD:
				s.username(cmd.client,cmd.args)
			case JOIN_CMD:
				s.join(cmd.client,cmd.args)
			case GROUPS_CMD:
				s.groupsList(cmd.client,cmd.args)
			case MSG_CMD:
				s.msg(cmd.client,cmd.args)
			case QUIT_CMD:
				s.quit(cmd.client,cmd.args)	

		}
	}
}
func (s* server) newClient(conn net.Conn) {
	log.Printf("New Client has connected the server, %s", conn.RemoteAddr().String())
	c := &client{
		conn: conn,
		username: "anonymous",
		commands: s.commands,
	}
	c.readInput()
}

func (s*server) username(c *client, args []string){
	c.username = args[1]
	c.msg(fmt.Sprintf("Welcome, Your Username has been set to %s",c.username))  
	//Sprint used for creating string
}

func (s*server) join(c *client, args []string){
	group_name := args[1]
	g,ok := s.groups[group_name]
	if !ok {
		g = &group{
			groupName: group_name,
			members: make(map[net.Addr]*client),
		}
		s.groups[group_name] = g
	}
	g.members[c.conn.RemoteAddr()] = c
	s.quitCurrGroup(c)
	c.currGroup=g

	g.broadcast(c, fmt.Sprintf("%s has joined the group",c.username))
	c.msg(fmt.Sprintf("Welcome to Group %s!",g.groupName))
}
func (s*server) groupsList(c *client, args []string){
	var groups []string
	for groupname := range s.groups{
		groups=append(groups,groupname)
	}
	c.msg(fmt.Sprintf("Available Groups: %s",strings.Join(groups,", ")))
}
func (s*server) msg(c *client, args []string){
	if c.currGroup == nil {
		c.err(fmt.Errorf("You must join a room before sending a message."))
		return 
	}
	c.currGroup.broadcast(c,c.username + ": "+ strings.Join(args[1:len(args)]," "))
}
func (s*server) quit(c *client, args []string){
	log.Printf("Client has disconnected: %s",c.conn.RemoteAddr().String())
	s.quitCurrGroup(c)
	c.msg("Have a good day, Hope to see you again.")
	c.conn.Close()
}
func (s* server) quitCurrGroup(c *client){
	if c.currGroup != nil {
		delete(c.currGroup.members, c.conn.RemoteAddr())
		c.currGroup.broadcast(c, fmt.Sprintf("%s user has left the group",c.username))
	}
}