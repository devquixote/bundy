package cli

import (
    "testing"
    "reflect"
    "os"
    "syscall"
)

var c Config = NewConfig()
var errorMsg = "Expected %v for %v but got %v"

func TestNewConfig(t *testing.T) {
    if c.Address != "0.0.0.0" {
      t.Errorf(errorMsg, "0.0.0.0", "Address", c.Address)
    }
    if c.Port != 6666 {
        t.Errorf(errorMsg, 6666, "Port", c.Port)
    }
    if c.Path != "/kill" {
        t.Errorf(errorMsg, "/kill", "Path", c.Path)
    }
    if !reflect.DeepEqual(c.Pids, []int{1}) {
        t.Errorf(errorMsg, []int{1}, "Pids", c.Pids)
    }
    var expectedSignals []os.Signal = []os.Signal{syscall.SIGTERM, syscall.SIGKILL}
    if !reflect.DeepEqual(c.Signals, expectedSignals) {
        t.Errorf(errorMsg, expectedSignals, "Signals", c.Signals)
    }
    if c.Delay != 20 {
        t.Errorf(errorMsg, 20, "Delay", c.Delay)
    }
    if !c.Suicide {
        t.Errorf(errorMsg, true, "Suicide", c.Suicide)
    }
}
