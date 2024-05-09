
### host 
```sh
jh@jh:~$ cat /etc/hosts
192.168.0.28 client 
192.168.0.26 server
```
#### server 
```sh
jh@jh:~$ sudo hostnamectl set-hostname server

```
#### client 
```sh
jh@jh:~$ sudo hostnamectl set-hostname client
```
## ssh-keygen 

```sh
jh@client:~/.ssh$ ll
total 16
drwx------  2 jh jh 4096  5월  9 21:12 ./
drwxr-x--- 22 jh jh 4096  5월  9 21:12 ../
-rw-------  1 jh jh 2590  5월  9 21:12 id_rsa
-rw-r--r--  1 jh jh  563  5월  9 21:12 id_rsa.pub
jh@client:~/.ssh$ cat id_rsa
-----BEGIN OPENSSH PRIVATE KEY-----
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
v08CgCH+bel58tAAAACWpoQGNsaWVudAE=
-----END OPENSSH PRIVATE KEY-----

jh@client:~/.ssh$ cat id_rsa.pub
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCWQGV0FJ3sYUDFsRZQGuJNt6ATGQ4P4RlTfZZmNY+H7TxEzC2z1vu+NQ5cgMEJIhUCU0Uth+vmhnJFit9nI1hBqlq2OZuEm51/7iBgprv5ASt+MOViivMQTCMJuCODpxrzstmgReTWigKIKqFYBE4LFQWHgPm1mJafDjwVkOwXs1/a6FlFc+X5oLeESwN2eQpl6Nai1d67hCFvFRL4Topbdq228z7p0okRRdPTCDvXmd4v5j+eroXz8iFOWgah8Kh/oEZs7J/IHh74swqQV5d3xWUhc5s+WzLrj5/sPgdefnVloniXhvPgPLnkRhXwXMFEioMrjdY2Bsl3PlCfuhiEioeaoeHwzjo/o9iFvlWgm2vmwcJ8aLkHiSekMU7BKy99RAWJMGOHyF/JRgz8/Cg+NNXSajD8QAyICAHG9DPyh2kgZSXG3kwz4GQUCVy3lmS47yHc5eXxrwxBjJOHF8Mu6RSaeN9qTL/1QJ2DD4Gtiy5oUkeEEA+ySKW7ya0uvBU= jh@client
```
### ssh-copy-id server
* client의 public key가 remote의 user의 .ssh 아래 authorized_keys 복사된다.  
```sh
jh@client:~/.ssh$ ssh-copy-id server
The authenticity of host 'server (192.168.0.26)' can't be established.
ED25519 key fingerprint is SHA256:lENF9ksfGHQaFqFuajFt/egVGU2D8VyLUP1x5WjAOO4.
This key is not known by any other names
Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
/usr/bin/ssh-copy-id: INFO: attempting to log in with the new key(s), to filter out any that are already installed
/usr/bin/ssh-copy-id: INFO: 1 key(s) remain to be installed -- if you are prompted now it is to install the new keys
jh@server's password: 

Number of key(s) added: 1

Now try logging into the machine, with:   "ssh 'server'"
and check to make sure that only the key(s) you wanted were added.
```
### server  authorized_keys 
```sh
jh@server:~/.ssh$ ls
authorized_keys
jh@server:~/.ssh$ cat authorized_keys 
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCWQGV0FJ3sYUDFsRZQGuJNt6ATGQ4P4RlTfZZmNY+H7TxEzC2z1vu+NQ5cgMEJIhUCU0Uth+vmhnJFit9nI1hBqlq2OZuEm51/7iBgprv5ASt+MOViivMQTCMJuCODpxrzstmgReTWigKIKqFYBE4LFQWHgPm1mJafDjwVkOwXs1/a6FlFc+X5oLeESwN2eQpl6Nai1d67hCFvFRL4Topbdq228z7p0okRRdPTCDvXmd4v5j+eroXz8iFOWgah8Kh/oEZs7J/IHh74swqQV5d3xWUhc5s+WzLrj5/sPgdefnVloniXhvPgPLnkRhXwXMFEioMrjdY2Bsl3PlCfuhiEioeaoeHwzjo/o9iFvlWgm2vmwcJ8aLkHiSekMU7BKy99RAWJMGOHyF/JRgz8/Cg+NNXSajD8QAyICAHG9DPyh2kgZSXG3kwz4GQUCVy3lmS47yHc5eXxrwxBjJOHF8Mu6RSaeN9qTL/1QJ2DD4Gtiy5oUkeEEA+ySKW7ya0uvBU= jh@client
```

* sshd config 
```sh
jh@server:/etc/ssh$ sudo vi /etc/ssh/sshd_config

PermitRootLogin yes
```


### ssh-copy-id root@server 
```sh
jh@client:~$ ssh-copy-id root@server
/usr/bin/ssh-copy-id: INFO: attempting to log in with the new key(s), to filter out any that are already installed
/usr/bin/ssh-copy-id: INFO: 1 key(s) remain to be installed -- if you are prompted now it is to install the new keys
root@server's password: 

Number of key(s) added: 1

Now try logging into the machine, with:   "ssh 'root@server'"
and check to make sure that only the key(s) you wanted were added.
```
### root@server의 .ssh
* client의 pub key가 root/.ssh/authorized_keys에 복사된다.  
```sh
root@server:~/.ssh# ls -l
total 4
-rw------- 1 root root 563  5월  9 21:22 authorized_keys
root@server:~/.ssh# cat authorized_keys 
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCWQGV0FJ3sYUDFsRZQGuJNt6ATGQ4P4RlTfZZmNY+H7TxEzC2z1vu+NQ5cgMEJIhUCU0Uth+vmhnJFit9nI1hBqlq2OZuEm51/7iBgprv5ASt+MOViivMQTCMJuCODpxrzstmgReTWigKIKqFYBE4LFQWHgPm1mJafDjwVkOwXs1/a6FlFc+X5oLeESwN2eQpl6Nai1d67hCFvFRL4Topbdq228z7p0okRRdPTCDvXmd4v5j+eroXz8iFOWgah8Kh/oEZs7J/IHh74swqQV5d3xWUhc5s+WzLrj5/sPgdefnVloniXhvPgPLnkRhXwXMFEioMrjdY2Bsl3PlCfuhiEioeaoeHwzjo/o9iFvlWgm2vmwcJ8aLkHiSekMU7BKy99RAWJMGOHyF/JRgz8/Cg+NNXSajD8QAyICAHG9DPyh2kgZSXG3kwz4GQUCVy3lmS47yHc5eXxrwxBjJOHF8Mu6RSaeN9qTL/1QJ2DD4Gtiy5oUkeEEA+ySKW7ya0uvBU= jh@client
```


## ssh pem format
### 1. Generate SSH Key in PEM Format
```sh
ssh-keygen -m PEM -t rsa -b 4096  -f ~/.ssh/id_rsa.pem
```
* -m PEM specifies that the key should be generated in PEM format.
* -t rsa specifies the type of key to create, in this case, RSA.
* -b 4096 specifies the number of bits in the key, in this case, 4096 bits for added security.
* -f ~/.ssh/id_rsa.pem specified the key file name

```sh
ssh-keygen  -q -f  ~/.ssh/root.pem -N  “”
```

### 2. Copy the Public Key to Your Server

```sh
ssh-copy-id -i ~/.ssh/id_rsa.pem.pub your_username@hostname
```

### 3. ssh login with pem key

```sh
ssh -i  ~/.ssh/root.pem root@server
```
### 4. ssh config 

```sh
$ cd ~/.ssh
$ vi config

Host my-website.com
    HostName my-website.com
    User my-user
    IdentityFile ~/.ssh/id_rsa

$ chmod 600 ~/.ssh/config

$ ssh my-website.com
```

### 5. example
```sh
jh@client:~$ ssh-keygen -q -f ~/.ssh/jhyunlee.pem -N ""
jh@client:~$ ls -l .ssh
total 8
-rw------- 1 jh jh 2590  5월  9 21:55 jhyunlee.pem
-rw-r--r-- 1 jh jh  563  5월  9 21:55 jhyunlee.pem.pub

jh@client:~$ ssh-copy-id -i ~/.ssh/jhyunlee.pem.pub root@server

jh@client:~$ ssh root@server
root@server:~# ls -l .ssh
-rw------- 1 root root 563  5월  9 21:56 authorized_keys

jh@client:~/.ssh$ cat config 
Host my-server
    HostName 192.168.0.26
    User root
    IdentityFile ~/.ssh/jhyunlee.pem

jh@client:~/.ssh$ ssh my-server
```