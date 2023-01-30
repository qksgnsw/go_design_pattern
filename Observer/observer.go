// 3. 옵저버

package main

type Observer interface {
	update(string)
	getID() string
}
