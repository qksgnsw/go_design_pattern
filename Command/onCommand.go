// 3. 구상 커맨드

package main

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}
