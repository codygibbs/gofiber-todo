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
