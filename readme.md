# Go Fiber Todo List

Simple todo list app using [Fiber](https://gofiber.io/).

## Requirements

- [Go v15](https://golang.org/)
- [Soda CLI](https://gobuffalo.io/en/docs/db/toolbox) - used for database migrations
- [Yarn v1](https://classic.yarnpkg.com/lang/en/) - Used for Node dependency management and asset compilation
- [Air](https://github.com/cosmtrek/air) - (optional) used for file-watching and re-compilation

## Installation
1. Clone this repo
1. Copy ".env.example" to ".env" and edit to match your environment
1. Copy "database.yml.example" to "database.yml" and edit to match your environment
1. Run `go get`
1. Run `yarn install && yarn build`
1. Run `soda migrate`

## Development Workflow
1. Run `air` (if you installed Air. Otherwise, just run `go run main.go`)

## Project Overview

### Webserver (ie, the Go bits)

The core webserver is split into 3 ordered layers, each grouped as a package. From highest-to-lowest-level, these layers are:

1. Application (*/application*) - This manages the overall application lifecycle, such as managing the database connection.
1. Handlers (*/handler*) - This manages web request handlers (also known as the "controller")
1. Models (*/model*) - This manages the inner domain data.

Each of these packages contains a "context" (abbreviated as "ctx") that is used to pass functionality from one layer to the next.

### Other Artifacts

1. Database migrations (*/migrations*) - Houses the Soda database migration files.
1. Frontend assets (*/frontend*) - Houses frontend asset sources, such as JS and CSS. These are compiled to the */public* folder when the asset pipeline is run.
1. Views (*/view*) - Houses the Go HTML templates used to render pages.
