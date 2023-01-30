// 전략 인터페이스
package main

type EvictionAlgo interface {
	evict(c *Cache)
}
