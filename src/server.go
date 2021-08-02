package main

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