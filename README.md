# Live Chat (Backend)

This is the backend for the Live Chat application, a demo app built
on Go, GraphQL, TypeScript and Angular.

Before you can do anything with this, you need to generate the
auto-generated code:

```
go generate
```

Then you can run it for development:

```
go run server.go
```

You can find a helpful GraphQL explorer at http://localhost:8081.
To build for deployment run:

```bash
CGO_ENABLED=0 go build                       # to build the statically-linked binary
sudo docker build --tag live-chat-backend .  # to build the container image
```

Now you can start a container using:

```bash
sudo docker run --rm -d --name backend -p 8081:8081 live-chat-backend
```
