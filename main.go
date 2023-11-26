package main

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
	"os"
	"runtime"
	"time"
)

func main() {

	var err error
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println("hostname:", hostname)
	fmt.Println("GOOS:", runtime.GOOS)
	fmt.Println("GOARCH:", runtime.GOARCH)

	info, err := host.Info()
	if err != nil {
		panic(err)
	}
	bs, _ := json.Marshal(info)
	fmt.Println("info:", string(bs))
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
