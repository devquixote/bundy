package cli

import (
    "testing"
    "reflect"
    "github.com/jessevdk/go-flags"
)

var errorMsg = "Expected %v for %v but got %v"

func TestDefaultOptions(t *testing.T) {
    args := []string{}
    opts := Options{}
    args, err := flags.ParseArgs(&opts, args)

    if err != nil {
        t.Errorf("unexpectedf error %v", err)
    }
    if opts.Address != "0.0.0.0" {
        t.Errorf(errorMsg, "0.0.0.0", "Address", opts.Address)
    }
    if opts.Port != 6666 {
        t.Errorf(errorMsg, 6666, "Port", opts.Port)
    }
    if !reflect.DeepEqual(opts.Pids, []int{1}) {
        t.Errorf(errorMsg, []int{1}, "Pids", opts.Pids)
    }
    if !reflect.DeepEqual(opts.Signals, []string{"TERM","KILL"}) {
        t.Errorf(errorMsg, []string{"TERM","KILL"}, "Signals", opts.Signals)
    }
    if opts.Delay != 20 {
        t.Errorf(errorMsg, 20, "Delay", opts.Delay)
    }
    if !opts.Suicide {
        t.Errorf(errorMsg, true, "Suicide", opts.Suicide)
    }
}
