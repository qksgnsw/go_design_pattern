// 1. 컬렉션

package main

type Collection interface {
	createIterator() Iterator
}
