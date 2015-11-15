package api

import (
    "io"
    "net/http"
)

type Api struct {
    Address string
    Path string
}

func (api *Api) HandleRequest(writer http.ResponseWriter, response *http.Request) {
    io.WriteString(writer, "Hello, World!")
}

func (api *Api) Start() {
    http.HandleFunc(api.Path, api.HandleRequest)
    http.ListenAndServe(api.Address, nil)
}
