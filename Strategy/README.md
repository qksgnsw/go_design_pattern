# 전략 ( Strategy )

## 개념
- 알고리즘들의 패밀리를 정의하고,
- 각 패밀리를 별도의 클래스에 넣은 후
- 그 객체들을 상호교환 할 수 있도록 하는
- **행동 디자인 패턴**
  
## 특징
- 특정 작업을 다양한 방식으로 수행하는 클래스를 선택 후,
- 모든 알고리즘을 별도의 클래스로 추출.
- 콘텍스트는 전략 중 하나에 대한 참조를 저장하기 위한 필드가 존재해야 함.
```go
package main

type EvictionAlgo interface {
	evict(c *Cache)
}
type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo // 전락 참조 저장 필드
	capacity     int
	maxCapacity  int
}
```
- 콘텍스트는 전략 선택에 대한 책임이 없음.
- 대신, 클라이언트가 원하는 전략을 전달.

## 적용
- 걱체 내에서 한 알고리즘의 다양한 변형들을 사용하고 싶을 때.
- 런타임 중 알고리즘을 전환하고 싶을 때.
- 일부 행동을 실행하는 방식에서만 차이가 있는 유사한 클래스가 많은 경우.

## 구현 방법
1. 콘텍스트 클래스에서 자주 변경되는 알고리즘 파악. 조건문인 경우가 많음.
2. 알고리즘에서 공통 메서드를 전략 인터페이스로 선언.
```go
type EvictionAlgo interface {
	evict(c *Cache)
}
```
3. 알고리즘을 자체 클래스로 추출. 2의 인터페이스 구현해야 함.
```go
type Fifo struct {}
func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}

type Lru struct {}
func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}

type Lfu struct {}
func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strtegy")
}
```
4. 콘텍스트에서 전략개체 참조에 필요한 필드 추가 후,
5. 해당 필드 값을 수정하기 위한 메서드 구현
```go
type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo // 4
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

// 5
func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}
```
6. 콘텍스트의 클라이언트들은 적절한 전략과 연관 시켜야함.

## 장점
- 런타임에 객체 내부에서 사용되는 알고리즘 변경 가능
- 알고리즘 세부 구현 고립
- 상속을 합성으로 대체
- 개방/폐쇄 원칙

## 단점
- 알고리즘의 수가 적고, 거의 변하지 않는다면 해당 패턴은 오버헤드
- 전략간의 차이를 확실히 알아야 함.
