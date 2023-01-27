// 1. 프로토타입 인터페이스

package main

type Inode interface {
	print(string)
	clone() Inode
}
