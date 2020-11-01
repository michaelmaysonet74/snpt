# snpt API

snpt API is a RESTful API built in golang meant to be used by the snpt UI project.

## Version and Dependencies
---

The project is built in golang@1.15.3. We use Docker for packaging `snpt API` container.

## Getting Started
---

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
