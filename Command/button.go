// 1. 발송자

package main

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}
