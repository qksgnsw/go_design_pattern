# 어댑터 ( Wrapper, Adapter )

## 개념
- **구조 디자인 패턴**
- 호환되지 않는 객체들이 협업할 수 있도록 함

## 특징
- 한 객체의 인터페이스를 다른 객체가 이해할 수 있도록 변환하는 객체
- 양방향 어댑터도 가능
- 공통 인터페이스 메서드 래핑기법

## 적용
- 기존 클래스를 사용하고 싶지만 나머지 코드와 호환되지 않을 때
- *TODO 추가 설명 필요*

## 구현 방법
1. 호환되지 않는 인터페이스가 있는 클래스 파악하기
2. 클라이언트 인터페이스를 선언 및 구현
3. 어댑터 클래스 생성 후 클라이언트 인터페이스를 따르게 구현(래핑)
4. 클라이언트는 클라이언트 인터페이스를 사용하여 어댑터를 사용해야함

## 장점
- 단일 책임 원칙
- 개방/폐쇄 원칙

## 단점
- 복잡성 증가