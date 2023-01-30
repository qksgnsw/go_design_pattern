# 옵저버 ( Observer )

## 개념
- **행동 디자인 패턴**
- 내가 관찰 중인 객체에 발생하는 이벤트를,
- 구독하는 매커니즘을 정의 할 수 있도록 하는 패턴
  
## 특징
- 구독자 객체들에 대한 참조의 리스트 배열 필드 필요
- 구독자를 추가 및 제거하는 메서드들 필요
  
## 적용
- 한 객체의 상태가 변경되어 다른 객체들을 변경해야 할 필요성이 있을 때
- 실제 객체 집합들을 미리 알 수 없거나, 이런 집합들이 동적으로 변할 때
- 일부 객체들이 제한 시간, 또는 특정 경우에만 다른 객체들을 관찰해야 할 때

## 구현 방법
1. 비즈니스 로직을 파악하고, 핵심기능은 출판사, 나머지는 구독자 클래스 집합으로 변경
2. 구독자 인터페이스 선언. 최소한 하나의 update 메서드 선언
```go
type Observer interface {
	update(string)
	getID() string
}
```
3. 출판사 인터페이스 선언. 구독자 객체를 구독자 리스트에 추가 및 제거하는 메서드 구현
```go
type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}
```
4. 구독자 메서드 구현하기.
```go
type Customer struct {
	id string
}
func (c *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}
func (c *Customer) getID() string {
	return c.id
}
```
5. 출판사 메서드 구현하기.
6. 구독자 생성 및 등록하기.
```go
func main() {
	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
```

## 장점
- 개방/폐쇄 원칙
- 런타임에 객체 간의 관계를 형성 할 수 있음.

## 단점
- 구독자는 무작위로 알림을 받을 수 있음