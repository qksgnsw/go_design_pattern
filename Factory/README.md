# 팩토리 ( Factory, Factory Method )

## 개념
- 부모 클래스가 객체를 생성 할 수 있는 인터페이스를 제공하지만,
- 자식 클래스들이 생성될 객체들의 유형을 변경 할 수 있도록 하는
- **생성 디자인 패턴**

## 특징
- 객체 직접 생성이 아닌 팩토리 메서드를 활용한 생성
- 생성자 호출을 부모 클래스에 넘기고 상속 받은 자식 클래스가 오버라이딩 함
- 자식 클래스에서 사용 시 공통 클래스 또는 공통 인터페이스가 있는 경우에만 반환 가능
- 공통 인터페이스를 구현한다면 객체들을 손상하지 않고 수행 가능.

## 적용
- 객체들의 유형이나 의존관계를 정확하게 모르는 경우
- 기존의 라이브러리나 컴포넌트를 확장 할 경우
- 기존 객체를 재사용하여 시스템 리소스를 절약해야하는 경우

## 구현 방법
1. 모든 객체가 같은 인터페이스를 따라야하므로 해당 메서드 선언하기
```go
package main

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
```
2. 해당 인터페이스에 맞는 제품들 구현하기
3. 팩토리 메서드 구현하기. if나 switch가 많아 질 수 있음.
```go
package main

import "fmt"

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}

```

## 장점
- 느슨한 결합
- 단일 책임 원칙
- 개방/폐쇄 원칙

## 단점
- 코드가 복잡해짐