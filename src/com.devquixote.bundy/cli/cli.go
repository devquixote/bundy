package cli

import (
    "os"
    "syscall"
)

// Configuration data
type Config struct {
    Address string
    Port int
    Path string
    Pids []int
    Signals []os.Signal
    Delay int
    Suicide bool
}

// Construct a config struct w/default values
func NewConfig() Config {
    result := Config{}
    result.Address = "0.0.0.0"
    result.Port = 6666
    result.Path = "/kill"
    result.Pids = []int{1}
    result.Signals = []os.Signal{syscall.SIGTERM, syscall.SIGKILL}
    result.Delay = 20
    result.Suicide = true
    return result
}
