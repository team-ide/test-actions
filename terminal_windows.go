package main

import "os/exec"

func (this_ *Terminal) Start() (err error) {
	cmd := exec.Command("conhost", "--headless")
	this_.reader, err = cmd.StdoutPipe()
	if err != nil {
		return
	}
	this_.errReader, err = cmd.StderrPipe()
	if err != nil {
		return
	}
	this_.writer, err = cmd.StdinPipe()
	if err != nil {
		return
	}
	err = cmd.Start()
	if err != nil {
		return
	}
	return
}
