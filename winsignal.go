package winsignal

import (
    "os"
	"log"
    "strconv"
    "io/ioutil"
	"github.com/fsnotify/fsnotify"
)

type WinSignal int

const (
    SIGABRT   = WinSignal(1)
    SIGCHLD   = WinSignal(3)
    SIGCLD    = WinSignal(4)
    SIGINT    = WinSignal(5)
    SIGKILL   = WinSignal(6)
    SIGQUIT   = WinSignal(7)
    SIGSTOP   = WinSignal(8)
    SIGTERM   = WinSignal(9)
    SIGWINCH  = WinSignal(10)
)

var shmfile = ".shm"

func Wait(sigs ...WinSignal) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
        panic(err)
	}
	defer watcher.Close()

    if f, err := os.Create(shmfile); err != nil {
        panic(err)
    } else {
        f.Close()
    }

	if err = watcher.Add(shmfile); err != nil {
        panic(err)
    }

    skip := false
    for {
        if skip {
            break
        }

        select {
        case event := <-watcher.Events:
            if event.Op&fsnotify.Write == fsnotify.Write {
                if sig := readSignalFrom(shmfile); inSignals(sig, sigs) {
                    skip = true
                }
            } else if event.Op&fsnotify.Remove == fsnotify.Remove ||
                event.Op&fsnotify.Rename == fsnotify.Rename {
                skip = true
            }
        case err := <-watcher.Errors:
            skip = true
            log.Println("winsignal error:", err)
        }
    }

    watcher.Remove(shmfile)
    os.Remove(shmfile)
}

func Send(s WinSignal) {
    i := strconv.FormatInt(int64(s), 10)
    ioutil.WriteFile(shmfile, []byte(i), 0666)
}

func readSignalFrom(f string) (s WinSignal) {
    buf, err := ioutil.ReadFile(f)
    if err != nil {
        return
    }
    i, _ := strconv.ParseInt(string(buf), 10, 32)
    return WinSignal(int(i))
}

func inSignals(s WinSignal, sigs []WinSignal) bool {
    for i, _ := range sigs {
        if sigs[i] == s {
            return true
        }
    }
    return false
}
