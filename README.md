# Windows Signal For Go

In the windows platform, your golang program can use **winsignal** to wait another process send signal.


## Example
-----------------
```golang
package main

import (
    "flag"
    "log"
    "github.com/c77cc/winsignal"
)

func ExampleWait() {
    stopptr := flag.Bool("stop", false, "stop program")
    flag.Parse()

    if *stopptr {
        // send signal in another process
        winsignal.Send(winsignal.SIGSTOP)
        return
    }

    // wait signal
    winsignal.Wait(winsignal.SIGSTOP, winsignal.SIGTERM)

    doSomethingWhenStop()
}

func doSomethingWhenStop() {
    log.Println("do something...")
}
```

## Related Projects

* [fsnotify](https://github.com/fsnotify/fsnotify)
