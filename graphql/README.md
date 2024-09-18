# Graphql API using Go

This api exposes the movies data from https://freetestapi.com/api/v1/movies. While bootstraping the application, the 
server caches the movie data in memory to be used as a database. Check [movies.go](./movies/movies.go) for the same.

Use the following commands to run the server
```shell
# Use this command to install dependencies
$ go install

# Use this command to run server
$ go run main.go
```
