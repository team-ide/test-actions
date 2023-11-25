package main

import (
	"fmt"
	"io"
)

type Terminal struct {
	isStopped bool
	reader    io.Reader
	errReader io.Reader
	writer    io.Writer
}

func (this_ *Terminal) StartRead() {
	var bs = make([]byte, 1024)
	for {
		n, err := this_.reader.Read(bs)
		if n > 0 {
			fmt.Println("on read:", string(bs[0:n]))
		}
		if err != nil {
			fmt.Println("on read error:", err.Error())
			this_.isStopped = true
			return
		}
	}
}

func (this_ *Terminal) StartErrRead() {
	if this_.errReader == nil {
		return
	}
	var bs = make([]byte, 1024)
	for {
		n, err := this_.errReader.Read(bs)
		if n > 0 {
			fmt.Println("on err read:", string(bs[0:n]))
		}
		if err != nil {
			fmt.Println("on err read error:", err.Error())
			this_.isStopped = true
			return
		}
	}
}

func (this_ *Terminal) Write(cmd string) (err error) {
	_, err = this_.writer.Write([]byte(cmd))
	if err != nil {
		return
	}
	return
}
