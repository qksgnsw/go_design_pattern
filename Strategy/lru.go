// 구상 전략
package main

import "fmt"

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}
