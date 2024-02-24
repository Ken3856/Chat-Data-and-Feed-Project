This is a Go program that implements a Web socket server.

1. The program listens for incoming web socket connections on port 3000,
   When a client establishes a web socket connection, the server then logs the client adderess and adds to connection map.

2. Server then reads the clients messages continuously.

3. When a message is received, it is acknowledged by sending an acknowledgement message back to the client and logged.
   If there are any errors when reading or writing they are logged and the connection is closed.
