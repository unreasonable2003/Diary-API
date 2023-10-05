# Diary-API

### Introduction
The Diary API is a backend implementation for a user-centric application that enables individuals to create, manage, and access their personal diary entries. Users can register for an account, log in securely, and perform various actions related to diary management.
The API is only meant for learning purpose, specifically Golang and gin framework. 

### Prerequisites

- **[Go](https://go.dev/)**: any one of the **three latest major** [releases](https://go.dev/doc/devel/release).
- A Code editor/IDE which can help to compile Golang files (example: VS Code, GoLand, etc)
- Replace the `.env` credentials with your credentials.
- Create a database instance in a cloud based PostgreSQL (example: Neon, ElephantSQL).
- Add the cloud based PostgreSQL API url in `database/database.go` file. The URL should be added here:

```go
func Connect() {
	var err error

	dbURL := "YOUR DATABASE URL"

	// Parse the URL
	connector, err := pq.ParseURL(dbURL)
	if err != nil {
		panic(err)
	}

  ...rest of the code
}
```

### Running the API

To execute the the API code, simply open the the folder in an IDE and run the following GO command in the terminal
```sh
$ go run main.go
```
### Articles about Gin

A curated list of awesome Gin framework.

- [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)
