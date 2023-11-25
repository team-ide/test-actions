package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Println("GOOS:", runtime.GOOS)
	fmt.Println("GOARCH:", runtime.GOARCH)
	var err error
	terminal := &Terminal{}
	err = terminal.Start()
	if err != nil {
		panic(err)
	}
	go terminal.StartRead()
	go terminal.StartErrRead()
	err = terminal.Write("cd ../\r\n")
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 5)
}
