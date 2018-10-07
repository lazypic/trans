# trans
언어인식 라이브러리 + 카카오API를 이용해서 제작된 터미널 번역기입니다.

#### 사용법
- 영어단어를 번역할 수 있습니다.
```bash
$ trans test
$ trans -text test # 이렇게도 작성할 수 있습니다.
시험
```

- 한글을 입력하면 자동으로 언어를 인식해서 영어로 바꾸어줍니다.
```bash
$ trans 테스트
test
```

- 중국어도 가능합니다.
```bash
$ trans 早上好
좋은 아침
```

- 파일을 trans 명령으로 보낼수 있습니다. 좋아하는 시인데 번역이 좋지 않습니다.
```bash
$ cat woohwa.txt | trans
```

- echo 명령을 이용해서 번역을 할 수 있습니다.
```bash
$ echo "test" | trans
시험
```

- 모든 Args를 사용한 문법은 아래와 같습니다. -src, -target에는 kr(한국어), en(영어), jp(일본어), cn(중국어), vi(베트남어), id(인도네시아어)를 설정할 수 있습니다.
```bash
$ trans -src en -target kr -text test
```

- man 형태의 유닉스/리눅스 명령어도 trans 명령어로 번역할 수 있습니다.
```bash
$ man -help 2>&1 | trans
```

#### library
- https://github.com/abadojack/whatlanggo
- https://developers.kakao.com/features/platform#번역
