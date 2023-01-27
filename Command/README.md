# 커맨드 ( Action, Transaction, Command )

## 개념
- **행동 디자인 패턴**
- 요청을 독립 실행형 객체로 변환
- 요청에 대한 모든 정보가 포함됨
- 지연, 대기열등록, 실행 취소 등의 기능 구현 가능
  
## 특징
- **관심사 분리의 원칙** 기반
- 요청시 비지니스 논리 객체에 직접 접근하지 말아야 할 경우,
- 해당 요청을 작동 시키는 단일 메서드를 가진 별도의 커맨드클래스가 있는 것이 필요
- 커맨드 클래스는 결국 요청자와 비즈니스 논리 객체 사이의 링크 역할을 함
- 요청시 인수는 호출자가 전달하는 것이 아니고,
- 커맨드 메서드에 미리 설정, 또는 자체적으로 가져오게 해야 함

## 적용
- 객체를 매개변수화 하려는 경우
- 실행 예약, 대기열 등록, 원격 실행등에 사용
- 되돌릴 수 있는 작업 구현 시
  
## 구현 방법
1. 단일 실행 메서드로 커맨드 인터페이스 선언
2. 위의 인터페이스를 구현하는 구상 커맨드 클래스 구현
3. 발송자 역할을 할 클래스를 선별,
```go
type Button struct {
	command Command
}
func (b *Button) press() {
	b.command.execute()
}
```
4. **커맨드 인터페이스로만 통신해야함**
```go
type Command interface {
	execute()
}
type OnCommand struct {
	device Device
}
func (c *OnCommand) execute() {
	c.device.on()
}
type OffCommand struct {
	device Device
}
func (c *OffCommand) execute() {
	c.device.off()
}
```
5. **수신자에게 직접 요청하는 것이 아니다!**
6. 다음과 같은 순서로 객체들을 초기화해야함
   1. 수신자를 만든다.
   2. 커맨드를 만들고 필요한 경우 수신자들과 연관시킨다.
   3. 발송자들을 만들고 특정 커맨드과 연관 시킨다.
```go
package main

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
```

## 장점
- 단일 책임 원칙
- 개방/폐쇄 원칙
- 실행 취소/다시 실행 구현 가능
- 지연 실행 구현 가능
- 간단한 커맨드 조합 가능

## 단점
- 발송자와 수신자 사이에 새로운 레이어가 도입 되는 것이므로,
- 코드의 복잡성 증가