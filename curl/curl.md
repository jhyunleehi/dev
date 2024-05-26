# curl  


```sh
curl -V
curl  -help all
```

-k --insecure https 프로토콜에서 SSL 인증서에 대한 검증없이 연결
-i --head HTTP 헤더만 보여주고 컨텐츠는 표시하지 않음
-D --dump-header <file> HTTP 헤더를 file에 기록 (덤프)
-L --location HTTP 301, 302 응답을 받은 경우 리디렉션 URL로 따라간다.
-d --data HTTP POST 요청 데이터 입력
-J --remote-header-name  헤더에 있는 파일 이름으로 다운로드 파일을 저장
-o --output FILE curl로 받아온 내용을 FILE 이라는 이름의 파일로 저장
-O --remote-name 파일 저장시 리모트에 저장되어 있던 이름을 그대로 가져와서 로컬에 저장
-s --silent 진행 내용이나 메시지들을 출력하지 않음HTTP response code 만 가져오거나 할 경우 유리
-X --request 요청시 사용할 메소드의 종류 (GET, POST, PUT, PATCH, DELETE)
-A --user-agent 서버에 User-Agent <name> 보내기
-u --user 서버 사용자 및 비밀번호
-T --upload-file 로컬 FILE 을 대상으로 전송
-H 전송할 헤더를 지정
-v --verbose 동작하면서 자세한 헤더 통신 옵션을 출력한다.



####  GET
```sh
$ curl www.example.com
$ curl -X GET www.example.com
```
#### POST 
```sh 
$ curl -d "key1=value1&key2=value2" \ 
-H "Content-Type: application/x-www-form-urlencoded" \ 
-X POST http://localhost:8000/data 

$ curl -d '{"key1":"value1", "key2":"value2"}' \
-H "Content-Type: application/json" \
-X POST http://localhost:8000/data
```
### PUT
```sh
$ curl -X PUT -d 'name=inpa&email=inpa@gmail.com' http://localhost:8080/user/100
$ curl -X PUT -H "Content-Type: application/json" -d '{"name":"inpa","email":"inpa@gmail.com"}' http://localhost:8080/user/100

# 파일명으로 PUT
$ curl -T filename.txt http://www.example.com/dir/
```
### option 
```sh 
$ curl -I google.com
$ curl -v www.example.com
```

### proxy, 인증  
```sh 
$ curl -u username:password https://example.com/protected_page
``
$ curl -x 192.168.44.1:8888 http://linux.com/
```