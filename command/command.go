package main

import "fmt"

type Command interface {
	Execute()
}

type TV struct {
	Volume int
	Power  bool
}

func (t *TV) PowerOn() {
	t.Power = true
	fmt.Println("tv power on")
}

func (t *TV) PowerOff() {
	t.Power = false
	fmt.Println("tv power off")
}

func (t *TV) VolumeUp() {
	t.Volume++
	fmt.Println("tv volume up")
}

func (t *TV) VolumeDown() {
	t.Volume--
	fmt.Println("tv volume down")
}

type TvOnCommand struct {
	tv *TV
}

func (c *TvOnCommand) Execute() {
	c.tv.PowerOn()
}

type TvOffCommand struct {
	tv *TV
}

func (c *TvOffCommand) Execute() {
	c.tv.PowerOff()
}

type TvVolumeUpCommand struct {
	tv *TV
}

func (c *TvVolumeUpCommand) Execute() {
	c.tv.VolumeUp()
}

type TvVolumeDownCommand struct {
	tv *TV
}

func (c *TvVolumeDownCommand) Execute() {
	c.tv.VolumeDown()
}

type RemoteControle struct {
	command *Command
}

func (r *RemoteControle) ExecuteCommand(command Command) {
	command.Execute()
}
