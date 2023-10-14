# Socket Programming in Go
Some tinkering with socket programming in Golang

## Components
    1. Client code to create a TCP connection with a server listening on a given port
    2. Server that can accept connection on a new thread (faciliated by Goroutines)
    3. Helper package with shared helper code

## Functionality:
After the client establishes a TCP connection to the server, it will accept input from `stdin`.

The server will accept the message and return a message to the client - the original message in upper-case format.

The server accepts connections in a new Goroutine, which allows it to be multi-threaded with the help of Go's native concurrency model.

## How to run:
1. Start the server. Note the default connection address is `localhost:8080` but we can adjust this using optional flags as demonstrated below:
```
$ go run ./server --host localhost --port 8000
```

2. From a new terminal session - create a new client. Also accepts optional flags for server IP address
```
$ go run ./client --host localhost --port 8000
```

3. You can now enter messages in the client terminal.