package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("uid: %d\n", os.Getuid())
	fmt.Printf("gid:%d\n", os.Getegid())
	groups, _ := os.Getgroups()
	fmt.Printf("subgid: %v\n", groups)
}
