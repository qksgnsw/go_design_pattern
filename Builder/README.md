# 빌더 ( Builder )

## 개념
- 복잡한 객체들을
- 단계별로 생성하는
- **생성 디자인 패턴**
- 공통 인터페이스 필요 x

## 디렉터
- 빌더단계들의 대한 일련의 호출을 디렉터 클래스로 별도로 추출 가능
- **빌더 단계의 순서 등을 정의**
- 빌더클래스는 구현함
- 필수는 아니지만 사용하기 편리하게 만들어 줌
```go
package main

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()

	return d.builder.getHouse()
}
```

## 적용
- 매우 **많은 (선택적)매개변수를 이용한 생성자** 사용시 
- 생성 과정 중 세부 사항만 다른 유사한 단계를 포함하여 사용시
- 단계별 생성, 순서가 필요한 생성, 불완전하지 않은 생성이 필요할 때

## 구현 방법
1. 공통 생성 단계들을 명확하게 파악하기
2. 빌더 인터페이스에 위에서 파악한 단계를 선언하기
```go
type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}
```
3. 각 구상 빌더 클래스 만들기
```go
package main

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}
```
4. 해당 생성 단계 만들기
```go
func (b *NormalBuilder) setWindowType() {
	b.windowType = "Wooden window"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *NormalBuilder) getHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}
```
5. (선택) 디렉터 클래스 만들기

## 장점
- 단계별 생성 가능
- 생성 단계를 연기 및 재귀적 단계 가능
- 다양한 표현 생성시 코드 재사용 가능
- 단일 책임 원칙

## 단점
- 코드가 복잡해짐