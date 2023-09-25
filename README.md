## Golang HTTP Server with Basic Authentication
This is a simple HTTP server that is written in Golang that serves static or dynamic files. Also checks for private resources. 

## Getting Started
1. **prerequisites** : install Go
2. **clone the repo** :
3. **run the server locally**:
```bash
    go run main.go
```
the server will start on port '8080'

## Theory Knowledge

### Http server 
It is a software application responsible for receiving and responding to requests from web clients(eg. web browsers) by delivering web content such as web pages, images, videos etc.

we differentiate file extension between static and dynamic by file extensions. 
static : .html, .css, .js, .jpg, .png
dynamic : .php, .asp. , .cgi

commercial examples of http server: apache, ngnix

more info on golangs http package are given [here] (https://pkg.go.dev/net/http)

### Working of an http server
> listening for requests
> handle requests
> resource retrieval
> response generation
> sending the response
> client rendering


### More Info on how server in golang works
> Request reaches to the server
> Request then goes through a multiplexer
> The multiplexer macthes the route to specific handlers
> handler is a function that runs on every request made