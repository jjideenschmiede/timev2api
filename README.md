# timev2 API
The easiest way to load customers or projects from your timev2.

## Install
```console
go get github.com/jjideenschmiede/timev2api
```

## How to use?
You only need your API token and the domain key. With this you get the data back in a struct.

### Clients
```go
clients, err := timev2api.Clients("domainkey", "apitoken")
if err != nil {
    fmt.Println(err)
}
```

### Projects
```go
projects, err := timev2api.Projects("domainkey", "apitoken")
if err != nil {
    fmt.Println(err)
}
```

### Users
```go
users, err := timev2api.Users("domainkey", "apitoken")
if err != nil {
    fmt.Println(err)
}
```