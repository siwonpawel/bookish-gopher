# bookish-gopher
REST service implemented usign Golang

## Purpose of this project

First try on REST service written in Golang. Application provides basic CRUD operations on Book repository.  
Java is fun but why not to try Go?

## What I've learnt

- Golang project structure,
- How to work with packages in Go,
- Interactions with SQL Database,
- How to build types and data structures,
- Dependency management,
- Error management,
- Sending data over HTTP

## Technologies and libraries

- [Golang](https://golang.org/)
  - [Gorilla MUX](https://github.com/gorilla/mux)
  - [Gotoenv](github.com/subosito/gotenv)
  - [PostgreSQL client for Go](github.com/lib/pq)
- [PostgreSQL](https://www.postgresql.org/)

## JSON format for book entity
```json
{
  "id":     "int",  
  "title":  "string",  
  "author": "string",  
  "year":   "string"  
}
```
