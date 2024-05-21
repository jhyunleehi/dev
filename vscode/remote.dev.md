

### .ssh  authorized_keys

* client에서 ssh key 생성하고 서버에 pub key ssh-copy-id로 복사
* client에서 ssh-copy-id  root@dev_server
* client에서  ~/.ssh/config 파일 설정
* config 파일에서 설정한 이름으로 ssh 접속

```sh 
$ ssh-keygen -q -f ~/.ssh/jhyunlee.pem -N ""
$ ls -l .ssh
$ ssh-copy-id -i ~/.ssh/jhyunlee.pem.pub root@dev_server
$ ssh root@dev_server
$ cat ~/.ssh/config

Host my-server
    hostName  192.168.0.29
    User root
    IdentityFile ~/.ssh/jhyunlee.pem

$ ssh my-server
```