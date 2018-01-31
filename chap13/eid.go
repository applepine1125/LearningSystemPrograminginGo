package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("uid: %d\n", os.Getuid())
	fmt.Printf("gid: %d\n", os.Getegid())
	fmt.Printf("euid: %d\n", os.Geteuid())
	fmt.Printf("egid: %d\n", os.Getegid())
}
