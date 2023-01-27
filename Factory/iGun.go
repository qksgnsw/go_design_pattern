// 1. 제품 인터페이스

package main

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
