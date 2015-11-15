package main

import (
    "fmt"

    "com.devquixote.bundy/api"
)

func main() {
    fmt.Printf("Starting bundy server\n")
    api := api.Api{"127.0.0.1:6666", "/kill"}
    api.Start()
}
