package winsignal
//
//import (
//    "flag"
//    "log"
//    "360.cn/armory/winsignal"
//)
//
//func ExampleWait() {
//    stopptr := flag.Bool("stop", false, "stop program")
//    flag.Parse()
//
//    if *stopptr {
//        // send signal in another process
//        winsignal.Send(winsignal.SIGSTOP)
//        return
//    }
//
//    // wait signal
//    winsignal.Wait(winsignal.SIGSTOP, winsignal.SIGTERM)
//
//    doSomethingWhenStop()
//}
//
//func doSomethingWhenStop() {
//    log.Println("do something...")
//}
