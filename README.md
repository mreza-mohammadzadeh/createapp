# CreateApp CLI

## Overview
CreateApp is a command-line tool for generating a structured Golang project boilerplate. It helps developers quickly scaffold a project with predefined directories and files.

## Features
- Generates a modular Golang project structure.
- Includes predefined handlers, services, repositories, and migrations.
- Supports HTTP and gRPC delivery layers.
- Provides an alias setup for easy command execution.

## Installation
Clone the repository and build the binary:

```sh
git clone https://github.com/mreza-mohammadzadeh/createapp.git
cd createapp

# Build the binary
go build -o bin/createapp cmd/cli/createapp/main.go
```

## Setup
To make `createapp` globally accessible, add an alias in your shell configuration file:

```sh
echo 'alias createapp="/path/to/createapp/bin/createapp"' >> ~/.zshrc
source ~/.zshrc
```
_Replace `/path/to/createapp` with the actual path to your `createapp` directory._

## Usage
Run the following command to generate a new project structure:

```sh
createapp create app <AppName> [PathDirectory]
```
- `<AppName>`: The name of the new application.
- `[PathDirectory]` (optional): The directory where the project should be created. If omitted, the project is created in the current directory.

### Example
```sh
createapp create app myProject ~/go/src
```
This will generate the following structure inside `~/go/src/myProject`:

```
myProject/
├── cmd/
│   ├── build.go
├── delivery/
│   ├── httpserver/
│   │   ├── login.go
│   │   ├── route.go
│   │   ├── handler.go
│   ├── grpcserver/
├── entity/
│   ├── user.go
├── migrator/
│   ├── migrations/
│   │   ├── 1702146900_add_users_table.sql
│   ├── provider.go
├── param/
│   ├── get_by_id.go
├── repository/
│   ├── db.go
│   ├── get_by_id.go
├── service/
│   ├── service.go
│   ├── get_by_id.go
```

## Contributing
Feel free to submit issues and pull requests to improve `createapp`.

## License
This project is licensed under the MIT License.

## Author
[Your Name](https://github.com/mreza-mohammadzadeh)
