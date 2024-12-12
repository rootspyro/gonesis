# Gonesis
![gonesis](https://github.com/user-attachments/assets/ba357b51-7c84-4a96-8d4e-3921d8ef1ff6)


Terminal tool to generate a boilerplate for a Golang backend projects with the technology stack:

* **Programming Language:** [Golang](https://go.dev)
* **Web Framework:** [Fiber](https://gofiber.io)
* **Database:** [Postgresql](https://www.postgresql.org)
* **Containerization:** [Docker](https://www.docker.com)
* **Database Abstraction Layer:** [Sqlc](https://sqlc.dev)
* **API Documentation:** [Swagger](https://swagger.io)

## Introduction

Gonesis is a terminal tool based on my backend technology stack with Golang. It provides a set of features to help developers quickly create and manage projects.

## Features

* Project creation with a single command
* Automatic setup of Golang, Fiber, Postgresql, and Docker
* Simple and intuitive command-line interface
* Support for Sqlc database abstraction layer
* Swagger API documentation 

## Installation

You download the latest version of Gonesis from the [releases](https://github.com/rootspyro/gonesis/releases).

To install it in a `linux` system, run the following command:

```bash
$ wget https://github.com/rootspyro/gonesis/releases/download/1.0.2/gonesis_linux_[arch].tar.gz | sudo tar -xzf gonesis_linux_[arch].tar.gz -C /usr/local/bin
```

Or, just install it with `go`:
```bash
$ go install github.com/rootspyro/gonesis@latest
```

Try it out by running `gonesis -h`
```bash
$ gonesis -h

   ____ _ ____   ____   ___   _____ (_)_____
  / __  // __ \ / __ \ / _ \ / ___// // ___/
 / /_/ // /_/ // / / //  __/(__  )/ /(__  ) 
 \__, / \____//_/ /_/ \___//____//_//____/  
/____/                                      


 • Version: 1.0.2
 • Author: Spyridon Mihalopoulos
 • Github: https://github.com/rootspyro/gonesis

Usage of gonesis:
  -create string
        Create a new project. Usage: -create <name>
```

## Usage

Init your backend project running `$ gonesis --create [project_name]`. This command will generate the following project structure:
```bash
.
├── cmd
│   └── [project_name] 
│       └── main.go
├── db
│   ├── conn.go
│   ├── migrations
│   │   ├── 000001_first_table.down.sql
│   │   └── ...
│   └── sqlc
│       ├── [project_name] 
│            ├── query.sql
│            └── schema.sql
│ 
├── Dockerfile
├── Makefile
├── pkg
│   ├── config
│   │   └── config.go
│   └── parser
│       └── parser.go
│
├── README.md
├── src
│   ├── api
│   │   ├── handlers
│   │   │   └── common
│   │   │       ├── handler.go
│   │   │       └── pipes.go
│   │   ├── routes.go
│   │   └── services
│   │       └── common.srv.go
│   └── db
│       └── repositories
│           └── [project_name]_repo
│               └── ...
└── sqlc.yaml
```

Once created, run `$ cd [project_name]` and check the `README.md` file for more information about the project.

## Contributing

Contributions are welcome! If you'd like to contribute to Gonesis, please fork the repository and submit a pull request.

## License

Gonesis is licensed under the MIT License. See [LICENSE](LICENSE) for details.
