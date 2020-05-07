# golangTraining

golang 복습(https://www.youtube.com/channel/UCZp_ftx6UB_32VfVmlS3o_A)

----------------------------------------------

<hr>

1. 트랜지스터
   1. 컴퓨터의 기본 요소
   2. CPU는 트랜지스터의 덩어리
   3. 기능
      1. 스위치(전기적 신호)
      2. 증폭
   4. 1bit = 트랜지스터 1
2. 논리소자
   1. 다음 연산을 트랜지스터를 만들 수 있다 = ``논리소자(가산기)`` = ``계산기``를 만들 수 있다는 의미
      1. AND
      2. OR
      3. XOR
         1. 0 1 ``1``
         2. 1 0 ``1``
         3. 0 0 ``0``
         4. 1 1 ``0``
      4. NOT
         1. 0 ``1``
         2. 1 ``0``
3. 컴퓨터
   1. 모래 -> 규소 -> 실리콘 -> 컴퓨터
   2. 튜링 테스트(튜링기계)
   3. 폰 노이만(우주 최강의 두뇌ㅋㅋ) 기계
      1. 메모리를 따로 쓰자(현대 컴퓨터)
4. 컴퓨터의 원리
   1. 컴퓨터는 명령을 순서대로 수행
   2. ``1``과 ``0``으로 이미지, 문자 등 모두 표현할 수 있다.
   3. ASCII CODE
      1. 255개 문자를 표현(256 = 0 ~ 255)
      2. 한 문자당 1byte
   4. UNICODE
      2. UTF-8(영문자/숫자: 1byte, 이외의 문자: 3byte)
      3. UTF-16(한 문자당 2byte)
   5. HDD -> CACHE -> CPU -> REGISTER(연산)
5. 프로그래밍 언어
   1. 프로그램
      1. 명령과 순서를 쓴 문서(글)
      2. 사람의 언어 -(코딩)-> 프로그래밍 언어 -(컴파일)-> 기계어
   2. 고급언어(디컴파일이 어렵다)
      1. 컴파일 언어
         1. 코딩 -> 빌드 -> 기계어
         2. C, C++
      2. 동적 언어(인터프리트 언어)
         1. 코딩 -> 빌드 -> 중간상태 -(프로그램이 동작할 때)-> 기계어
         2. JAVA, C#, PYTHON, JAVASCRIPT
   3. 어셈블리어(디컴파일이 쉽다)
   4. **프로그래머는 한가지 언어에 집착하면 안된다.**
6. 컴파일언어와 동적언어의 차이
   1. 컴파일언어
      1. 속도 빠름
      2. platform 별로 변환해야 함
   2. 동적언어
      1. 속도 느림
      2. platform 별로 변환필요 없음
7. Golang
   1. 2009년 google이 만듬
   2. system programing을 하기 위해 만듬
   3. 아버지
      1. 켄 톰슨(70세 현역)
         1. B언어의 아버지
      2. 롭 파이크
         1. UTF-8의 아부지
8. Hello World
   1. C#과 가깝다
   2. **공부를 하는 것은 페인트 칠하는 것과 비슷하다**
   3. package main
      1. 묶음
         1. 라이브러리
            1. 기능을 묶어놓은 것
         2. 모듈
            1. 기능을 묶어놓은 것(좀 더 포커싱)
         3. 패키지
            1. 기능을 묶어놓은 것
         4. 프레임 웍
            1. 기능을 묶어놓은 것(절차까지 포함)
         5. 엔진
            1. 기능과 다른 프로그램 툴까지 모두 묶어놓은 것
      2. package로 선언
      3. main
         1. 프로그램의 시작점
   4. func
      1. 함수
9. Go의 변수
   1.  속성
       1.  이름
       2.  값
       3.  TYPE
           1.  int
               1.  int32(4byte): -21억 ~ 21억
                   1. uint32: 42억
               2.  int64(8byte): 크다~~
                   1.  uint64: 크다~~
               3.  int8(1byte): -128 ~ 127
                   1.  uint8: 0 ~ 255
               4.  int16(2byte): 2에 16제곱 - 1(65535)
                   1.  uint16: 0 ~ 65535
           2.  float
               1.  숫자부와 지수부
               2.  float32
                   1.  숫자부는 7개까지 표현
               3.  float64
                   1.  15개까지
           3.  string
               1.  문자의 길이에 따라 다름
               2.  한 글자: rune: 1 ~ 3byte
           4.  bool
       4.  메모리 주소
           1.  어디서?
       5.  사이즈
           1.  어디까지 읽어올 것인가?
   2.  선언
10. Go의 변수2
11. Go의 연산자
    1.  연산자의 종류
        1.  이항 연산자
        2.  단항 연산자
    2.  Go의 연산자
        1.  산술 연산자
        2.  비트 연산자
            1.  4 & 2 = 0
            2.  3 & 2 = 2
            3.  4 | 2 = 6
        3.  논리 연산자
        4.  그 외
            1.  shift 연산자(``>>``, ``<<``)
            2.  clear 연산자
12. 조건연산자와 조건문
13. Switch-case와 반복문
    1.  tour.golang.org/welcome/1
14. For 문 예제
    1.  ch14 코드 참고
15. 함수
    1.  기능의 모듈화
    2.  참조호출, 복사호출
    3.  함수를 호출할 때 값이 복사된다
        1. **go에서 모든 함수의 콜은 복사호출이다**
16. 함수의 호출과정과 재귀호출
    1.  함수 -> call -> jump -> 대입, 연산 -> return
    2.  재귀호출
        1.  **항상 끝나는 조건을 만들어라**
        2.  모든 재귀호출은 반복문으로 바꿀 수 있다
        3.  수학정의는 재귀호출이 쉬운 경우가 많다
    3.  **응집성은 높이고 종속성은 낮추는 것이 좋은 코딩**
17. 배열과 문자열
    1.  문자는 배열이고, 배열은 메모리다
        1.  go는 utf-8을 쓰므로 영어는 1byte, 한글은 3byte
    2.  문자열은 byte배열과 rune배열(utf-8)로 나타낼 수 있다
        1.  rune배열로 나타내면 한글도 한 글자씩 찍어낼 수 있다.
18. 배열, Radix Sort
    1.  배열의 복사
    2.  Radix sort
        1.  모든 경우에 사용할 수 없다
            1.  특정 조건
                1.  **원소 값의 범위가 한정적이고 작을 때 사용(N개)**
                    1.  이름순으로 정렬 
        2.  숫자를 작은 순부터 차례대로 카운트해서 뽑아낸다
19. 구조체
    1.  type 이름 struct{}
    2.  **객체**(개념을 한곳에 모아놓은 것)
        1.  속성
        2.  기능
    3.  goLang은 객체 밖에서 함수가 정의된다.
20. **포인터**
    1. 메모리의 ``주소``를 가르키는 것
    2. 변수가 가지는 모든 ``type``는 ``포인터``가 될 수 있다
    3. ``&``: 메모리 주소 값, ``*``: 메모리 주소가 가르키는 값
    4. 포인터가 쓰이는 이유
    5. 객체를 만들 때마다 함수의 메모리가 생성되는데 이것은 메모리의 낭비다
    6. 값을 함수인자로 받으면 함수를 복사되고, 메모리 주소를 복사하면 메모리의 주소만 복사하기 때문에 메모리가 더 적게 든다(javascript의 prototype를 이용한 객체 함수와 비슷한 개념)
21. 숫자야구 프로그래밍(code 참고)
    1. 컴퓨터가 숫자 3개를 뽑고 사용자가 숫자 3개를 입력하여서 몇 번만에 사용자가 맞추는지 계산하는 프로그램
    2. 순서도(어떤 순서대로 프로그램이 진행될지 나타내는 그림)
        1. 컴퓨터가 무작위 숫자 3개 뽑는다
        2. 사용자가 입력
        3. 비교
            1. 게임 끝
            2. 돌아가서 리게임
    3. **프로그램 만들 때 빌드가 되는 상황까지 만드는 것이 중요**
    4. 랜덤 숫자를 뽑을 때
       1. 컴퓨터는 정해진 코드대로 명령을 수행하기 때문에 랜덤 값을 뽑을 수 없다
          1. Seed 값을 설정해 줘야 한다.
             1. Time(항상 변하는 Seed 값)
22. 숫자야구2
    1.  123의 자릿 수마다 뽑아내는 방법
        1.  나머지 연산을 사용
    2.  숫자 입력을 반복할 때 ``\n``을 받아오기 때문에 없애줘야 한다.