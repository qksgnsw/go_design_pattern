// 3. 반복자

package main

type Iterator interface {
	hasNext() bool
	getNext() *User
}
