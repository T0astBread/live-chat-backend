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


## Decisions

I was asked to document my decisions while writing this (also see the
"Decisions" section in the frontend's README):

1. __Mocks: Automatic mocking server or stub implementations?__

Originally I wanted to use an automatic mocking server where I just
have to supply some dummy data to mock the backend out while
developing the frontend. The most likely candidate for this was
[json-graphql-server](https://github.com/marmelab/json-graphql-server)
which provides a GraphQL API from a simple JavaScript object.

However, I also wanted to use GraphQL subscriptions for incoming
messages and no GraphQL mocking server I found (including
json-graphql-server) supported subscriptions. Therefore I decided to
start by setting up the backend with simple stub implementations
instead and use that to work on the frontend.

2. __GraphQL framework for the backend__

I've evaluated a few Go frameworks for GraphQL but it was pretty
clear that gqlgen would fit my use case best. It supports
subscriptions and the code generation approach avoids a lot of
boilerplate code.

Additionally, the open schema would allow me to also generate
frontend code from the same schema, although I decided not to go into
that because I felt the cost to benefit ratio wasn't good enough.
