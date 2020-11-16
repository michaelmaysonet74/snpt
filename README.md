# snpt API

snpt API is a RESTful API built with the golang std library, meant to be used as a POC/case study.

## Models
Snippet - a small and encapsulated string of code, data or script.

User - entity that can see, create, update and/or delete a Snippet or a list of Snippets.

## Version and Dependencies

The project is built in golang@1.15.3.

We use Docker for packaging `snpt API` container.

MongoDB is the database used by the `snptAPI`. 

## Getting Started

### Run App Locally
```
$ make dev
```
To run the app locally, you will need to have intalled the proper golang version.

### Build and Run Docker Container
```
$ make docker-build
...
...
$ make docker-run
```
