# Golang GRPC Server

## Description

This is a simple GRPC server written in Golang which performs aimed to explore the GRPC framework and its features.

## Requirements

Create API where you can submit a purchase for a ticket. Details included in the receipt are:
From, To, User , price paid.
User should include first and last name, email address
The user is allocated a seat in the train. Assume the train has only 2 sections, section A and section B.
An API that shows the details of the receipt for the user
An API that lets you view the users and seat they are allocated by the requested section
An API to remove a user from the train
An API to modify a user's seat

## Tech Stack

-   Golang
-   GRPC
-   Protobuf
-   Docker (Future Scope)
-   Kubernetes (Future Scope)
-   MongoDB (Future Scope)
-   Redis (Future Scope)

## How to run

-   Clone the repo
-   Run `go mod tidy` to install the dependencies
-   Run `go run server.go` to start the server
-   Run `cd client && go run client.go` to start the client

## Future Scope

-   Add Dockerfile
-   Add MongoDB

Screenshot of the client

![Screenshot](./assets/grpcSetup.png)
