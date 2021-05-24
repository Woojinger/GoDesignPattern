# GoDesignPattern

https://www.youtube.com/watch?v=UEjsbd3IZvA&list=PLsoscMhnRc7pPsRHmgN4M8tqUdWZzkpxY
자바를 베이스로한 위 디자인 패턴 유튜브 강의를 Go로 바꿔 구현

## 1. Strategy Pattern
여러 알고리즘을 하나의 추상적인 접근점을 만들어 접근점에서 서로 교환 가능하도록 하는 패턴
Strategy Pattern은 runtime에 behavior가 결정되게 해준다.

Ex)
요구 사항
- 게임에서 캐릭터와 무기 구현
- 캐릭터는 무기를 쉽게 바꿀 수 있어야 한다

## 2. Template Method Pattern
어떤 작업을 처리하는 일부분을 서브 클래스로 캡슐화해 전체 일을 수행하는 구조는 바꾸지 않으면서 특정 단계에서 수행하는 내역을 바꾸는 패턴
- 구현하려는 알고리즘이 일정한 프로세스가 있다
- 구현하려는 알고리즘이 변경 가능성이 있다

구현 step
1. 알고리즘을 여러 단계로 나눈다
2. 나눠진 알고리즘의 단계를 메소드로 선언한다
3. 알고리즘을 수행할 템플릿 메소드를 만든다
4. 하위 클래스에서 나눠진 메소드들을 구현한다

Ex)
요구 사항
1) 신작 게임의 접속을 구현
requestConnection(str string)
2) 유저가 게임 접속시 다음을 고려해야 한다
- 보안 과정 : 보안 관련 부분을 처리. doSecurity(string string) string
- 인증 과정 : user name과 password가 일치하는지 확인. authentication(id string, password string) bool
- 권한 과정 : 접속자가 유료 회원인지 무료 회원인지 게임 마스터인지 확인. authorization(userName string) int
- 접속 과정 : 접속자에게 커넥션을 정보로 넘겨줌. connection(info string) string

접속 알고리즘을 단계로 나눔(보안, 인증, 권한, 접속). 템플릿 메소드 : requestConnection

## 3. Factory Method Pattern
Factory Method Pattern에서 Template Method Pattern이 사용이 된다.

Ex) 요구사항
게임 아이템과 아이템 생성 구현
- 아이템을 생성하기 전에 데이터 베이스에서 아이템 정보를 요청
- 아이템을 생성 후 아이템 복제 등의 불법을 방지하기 위해 데이터 베이스에 아이템 생성 정보를 남김
- 아이템을 생성하는 주체를 ItemCreator로 이름 짓는다
- 아이템은 item이라는 interface로 다룬다. item은 use 함수를 기본 함수로 갖고 잇다
- 아이템의 종류는 체력 회복 물약, 마력 회복 물약이 있다

## 4. Singleton Pattern
싱글톤 패턴을 이용해 하나의 인스턴스만 생성하도록 구현할 수 있다

Ex) 요구사항
개발 주에 시스템에서 스피커에 접근할 수 있는 클래스를 만들어 달라

