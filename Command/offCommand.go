// 3. 구상 커맨드

package main

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}
