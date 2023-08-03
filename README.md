# A simple Chat Server in go.

This is a simple TCP chat server, that I implemented in golang.

To connect to the server use: 
`telnet localhost 8888` 

It would follow the following commands.


- `/username <name>`: get a name or stay anonymous.
- `/join <group name>`: Join existing room or create a new room if it doesn't exist. For now, a user is allowed to use one room at a time.
- `/groups` : show the list of groups that are currently active in the server.
- `/msg <msg>` : to broadcast a message to everyone in the room.
- `/quit` : to quit out of the chat server.

# Demo

### Server Image
<img width="1246" alt="Server " src="https://github.com/qascade/ChatServer/assets/54154054/f0974a1c-4ccb-4fd7-ad3e-542a021902e6">

### Client 1
<img width="729" alt="Client1" src="https://github.com/qascade/ChatServer/assets/54154054/390783a5-cb15-463e-989d-935f03659487">

### Client 2
<img width="885" alt="Client2" src="https://github.com/qascade/ChatServer/assets/54154054/4a269375-e75d-43a7-b53b-bbb0d0bec8d2">
