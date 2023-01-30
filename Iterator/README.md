# 반복자 ( Iterator )

## 개념
- 컬렉션 요소들의 기본 표현(리스트, 스택, 트리 등)을 노출하지 않고,
- 하나씩 순회할 수 있도록 하는,
- **행동 디자인 패턴**
  
## 특징

## 적용
- 복잡한 데이터 구조를 보안이나 편의상의 이유로 캡슐화하여 숨기고 싶을 때.
- 순회 코드의 중복을 줄이고 싶을 때.
- 데티어 구조를 순회하고 싶은데 유형을 미리 알 수 없을 때.
  
## 구현 방법
1. 반복자 인터페이스 선언. 최소한 다음 요소를 가져오는 메서드 포함.
```go
type Iterator interface {
  hasNext() bool
  getNext() *User
}
```
2. 컬렉션 인터페이스 선언. 반복자 인터페이스를 가져오는 메서드 선언.
```go
type Collection interface {
	createIterator() Iterator
}
```
3. 반복자들이 순회 할 수 있도록 하고 싶은 컬렉션들에 대한 구상 반복자 클래스 구현.
```go
type UserIterator struct {
	index int
	users []*User
}
func (u *UserIterator) hasNext() bool {
	return u.index < len(u.users)

}
func (u *UserIterator) getNext() *User {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
```
4. 컬렉션 인터페이스 구현. (클라이언트에 특정 컬렉션 클래스에 맞게 조정된 반복자들을 생성하기 위한).
```go
type UserCollection struct {
	users []*User
}
func (u *UserCollection) createIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}
```
5. 클라인트에서 실행.
```go
func main() {

	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
```

## 장점
- 단일 책임 원칙.
- 개방/폐쇄 원칙.
- 같은 컬렉션을 병렬로 순회 가능.
- 순회 지연 및 일시 정지 가능.

## 단점
- 단순 컬렉션들만 있을 경우 오버헤드.
- 직접 탐색보다 덜 효율적.