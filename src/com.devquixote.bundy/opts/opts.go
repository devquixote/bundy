package opts

import (
    "strconv"
    "syscall"
)

var signalMap = map[string]syscall.Signal{
    "TERM": syscall.SIGTERM,
    "KILL": syscall.SIGKILL,
    // TODO expand this
}

type Options struct {
    Address string `short:"a" long:"address" default:"0.0.0.0" description:"Network address to bind to"`
    Port int `short:"p" long:"port" default:"6666" description:"Port to listen on"`
    Pids []int `short:"P" long:"pids" default:"1" description:"Pids to kill"`
    Signals []string `short:"s" long:"signals" default:"TERM" default:"KILL" description:"Signals to send in order"`
    Delay int `short:"d" long:"delay" default:"20" description:"Seconds to wait between signals"`
    Suicide bool `short:"S" long:"suicide" default:"true" description:"If bundy should kill itself after killing the target process"`
}

func (opts *Options) AddressAndPort() string {
    return opts.Address + ":" + strconv.Itoa(opts.Port)
}

func (opts *Options) ConvertedSignals() []syscall.Signal {
    var result []syscall.Signal = []syscall.Signal{}
    for _, sigName := range opts.Signals {
        result = append(result, signalMap[sigName])
    }
    return result
}
