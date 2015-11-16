package main

import (
    "os"

    "github.com/jessevdk/go-flags"

    "com.devquixote.bundy/opts"
    "com.devquixote.bundy/api"
)

func main() {
    var options opts.Options
    if _, err := flags.ParseArgs(&options, os.Args); err != nil {
        os.Exit(1)
    }
    // api := api.Api{"127.0.0.1:6666", "/kill"}
    api := api.Api{options}
    api.Start()
}
