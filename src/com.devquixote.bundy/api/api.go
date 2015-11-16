package api

import (
    "fmt"
    "io"
    "net/http"
    "syscall"

    "com.devquixote.bundy/opts"
)

type Api struct {
    Options opts.Options
}

func (api *Api) HandleRequest(writer http.ResponseWriter, response *http.Request) {
    for _, pid := range api.Options.Pids {
        for _, signal := range api.Options.ConvertedSignals() {
            if err := syscall.Kill(pid, signal); err == nil {
                fmt.Printf("Process %v %v\n", pid, signal)
                break
            }
        }

    }
    io.WriteString(writer, "Bundy has killed again!\n")
    if api.Options.Suicide {
        go syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
    }
}

func (api *Api) Start() {
    fmt.Printf("Starting bundy server (%v)\n", api.Options.AddressAndPort())
    http.HandleFunc("/kill", api.HandleRequest)
    http.ListenAndServe(api.Options.AddressAndPort(), nil)
}
