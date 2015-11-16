package cli

type Options struct {
    Address string `short:"a" long:"address" default:"0.0.0.0" description:"Network address to bind to"`
    Port int `short:"p" long:"port" default:"6666" description:"Port to listen on"`
    Pids []int `short:"P" long:"pids" default:"1" description:"Pids to kill"`
    Signals []string `short:"s" long:"signals" default:"TERM" default:"KILL" description:"Signals to send in order"`
    Delay int `short:"d" long:"delay" default:"20" description:"Seconds to wait between signals"`
    Suicide bool `short:"S" long:"suicide" default:"true" description:"If bundy should kill itself after killing the target process"`
}
