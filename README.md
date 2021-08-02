# A simple Chat Server in go.

This is a simple TCP chat server I implemented in golang.

It would follow the following commands.


- `/username <name>`: get a name or stay anonymous.
- `/join <group name>`: Join existing room or create a new room if it doesn't exist. For now, a user is allowed to use one room at a time.
- `/groups` : show the list of groups that are currently active in the server.
- `/msg <msg>` : to broadcast a message to everyone in the room.
- `/quit` : to quit out of the chat server.
