https://joonghyunlee.github.io/remote-debugging-nova-1

Visual Studio Code의 디버깅 프로파일을 추가해야 한다. launch.json 파일을 열어 아래 그림과 같이 디버깅 프로파일을 추가한다. Visual Studio Code의 원격 디버깅 설정은 크게 두 가지로 나뉜다. 첫 번째는 원격지 호스트에서 직접 Python 인터프리터를 실행하는 launch 방식이고, 두 번째는 소스 코드에 디버깅용 서버를 띄우는 코드를 삽입하여, 이 서버에 붙는 attach 방식이다. launch 방식이 보다 직관적이고 간편하지만 WSGI application과 같이 별도 worker 프로세스를 띄우는 경우는 attach 방식을 써야 한다.

