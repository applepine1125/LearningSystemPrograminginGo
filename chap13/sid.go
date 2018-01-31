package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Printf("gid: %d sid: %d\n", syscall.Getpgrp(), sid)
}
