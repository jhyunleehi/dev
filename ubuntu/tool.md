
### host
```
jhyunlee@Good:~/code/eBPF/bcc/src/cc/libbpf$ host naver
Host naver not found: 2(SERVFAIL)
jhyunlee@Good:~/code/eBPF/bcc/src/cc/libbpf$ host www.naver.com
www.naver.com is an alias for www.naver.com.nheos.com.
www.naver.com.nheos.com has address 223.130.195.200
www.naver.com.nheos.com has address 223.130.200.107
```

### dig 
```
jhyunlee@Good:~/code/eBPF/bcc/src/cc/libbpf$ dig  www.naver.com

; <<>> DiG 9.18.18-0ubuntu0.22.04.2-Ubuntu <<>> www.naver.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 40977
;; flags: qr rd ra; QUERY: 1, ANSWER: 3, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 65494
;; QUESTION SECTION:
;www.naver.com.			IN	A

;; ANSWER SECTION:
www.naver.com.		12	IN	CNAME	www.naver.com.nheos.com.
www.naver.com.nheos.com. 12	IN	A	223.130.200.107
www.naver.com.nheos.com. 12	IN	A	223.130.195.200

;; Query time: 0 msec
;; SERVER: 127.0.0.53#53(127.0.0.53) (UDP)
;; WHEN: Thu Feb 22 14:18:25 KST 2024
;; MSG SIZE  rcvd: 108
```