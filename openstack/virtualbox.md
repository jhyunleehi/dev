### 1. node
1. controller 
2. compute
3. storge

#### user: stack/ok

### configurageion 
#### network 
1. host only network #1 10.10.0.1/24
2. host only network #2 203.0.11.3.1/24

### NIC
1. hostonly network#1
2. internel network eth1
3. hostonly network#2 
4. NAT network

### controller node
1. cpu :1
2. mem :1G
3. disk: 40G, 4G
4. nw : NIC1,NIC2,NIC3,NIC4

### compute node 
1. cput 1
2. mem 1G
3. disk 40G 4G 4G 4G
4. nw : NIC1,NIC2,NIC3,NIC4

### storage node 
1. cput 1
2. mem 1G
3. disk 40G 4G
4. nw : NIC1,NIC2,NIC3,NIC4


## setup controller node

1. connection ready 
```sh
$ sudo hostnamectl set-hostname controller
$ sudo date
$ sudo date -s "2024-05-01 12:00:00"
$ sudo hwclk -w 

$ sudo ipaddr add 192.168.56.11/24 dev enp0s3
$ sudo apt update --fix-missing

$ sudo ip link set dev enps10 down
$ sudo ip link set dev enps10 up
==> /etc/environment
==> SDS.crt install 

$ sudo apt install openssh-server 
$ sudo systemctl stop ufw
$ sudo ip  route add default via 10.0.5.2
$ sudo apt install git
```
2. setting host from ssh clinet
```sh
$ ssh stack@192.168.56.11
$ sudo vi  /etc/hosts
192.168.56.11 controller
192.168.56.12 compute
192.168.56.13 storage

stack@controller:/etc/apt$ grep  "^deb" /etc/apt/sources.list
deb http://kr.archive.ubuntu.com/ubuntu/ jammy main restricted
deb http://kr.archive.ubuntu.com/ubuntu/ jammy-updates main restricted
deb http://kr.archive.ubuntu.com/ubuntu/ jammy universe
deb http://kr.archive.ubuntu.com/ubuntu/ jammy-updates universe
deb http://kr.archive.ubuntu.com/ubuntu/ jammy multiverse
deb http://kr.archive.ubuntu.com/ubuntu/ jammy-updates multiverse
deb http://kr.archive.ubuntu.com/ubuntu/ jammy-backports main restricted universe multiverse
deb http://security.ubuntu.com/ubuntu jammy-security main restricted
deb http://security.ubuntu.com/ubuntu jammy-security universe
deb http://security.ubuntu.com/ubuntu jammy-security multiverse

$ sudo resolvectl status | grep "DNS Server"
Current DNS Server: 168.126.63.2
       DNS Servers: 168.126.63.1 168.126.63.2
Current DNS Server: 168.126.63.2
       DNS Servers: 168.126.63.1 168.126.63.2

$ dig www.naver.com

$ sudo  vi /etc/netplan/01-network-manager-all.yaml
network:
  version: 2
  renderer: networkd
  ethernets:
    enp0s3:
      dhcp4: no
      dhcp6: no
      addresses: [192.168.56.11/24]
      gateway4: 192.168.56.1
      nameservers:
        addresses: [168.126.63.1,168.126.63.2]
      routes:
        - to: 192.168.56.0/24
          via: 192.168.56.1
    enp0s8:
      dhcp4: no
      dhcp6: no      
    enp0s9:
      dhcp4: no
      dhcp6: no
      addresses: [192.168.57.11/24]
      gateway4: 192.168.57.1  
      routes:
        - to: 192.168.57.0/24
          via: 192.168.57.1
    enp0s10:
      dhcp4: yes
      dhcp6: no 
      routes:
        - to: default
          via:10.0.5.2     
```

### make clone compute node, storage node

1. make vm clone  compute node in virtual box
2. make vm clone  storage node in virtual box 

* 완전한 복제로 생성한다. 
* MAC 주소 생성하는 옵션을 사용한다. 


## setup compute node
* add 4G disk
* add 4G disk 
```sh
$ sudo hostnamectl set-hostname compute 
$ sudo  vi /etc/netplan/01-network-manager-all.yaml
network:
  version: 2
  renderer: networkd
  ethernets:
    enp0s3:
      dhcp4: no
      dhcp6: no
      addresses: [192.168.56.31/24]
      gateway4: 192.168.56.1
      nameservers:
        addresses: [168.126.63.1,168.126.63.2]
      routes:
        - to: 192.168.56.0/24
          via: 192.168.56.1
    enp0s8:
      dhcp4: no
      dhcp6: no      
    enp0s9:
      dhcp4: no
      dhcp6: no
      addresses: [192.168.57.31/24]
      gateway4: 192.168.57.1  
      routes:
        - to: 192.168.57.0/24
          via: 192.168.57.1
    enp0s10:
      dhcp4: yes
      dhcp6: no 
      routes:
        - to: default
          via:10.0.5.2     
```

## setup storage node

```sh
$ sudo hostnamectl set-hostname storage 
$ sudo  vi /etc/netplan/01-network-manager-all.yaml
network:
  version: 2
  renderer: networkd
  ethernets:
    enp0s3:
      dhcp4: no
      dhcp6: no
      addresses: [192.168.56.41/24]
      gateway4: 192.168.56.1
      nameservers:
        addresses: [168.126.63.1,168.126.63.2]
      routes:
        - to: 192.168.56.0/24
          via: 192.168.56.1
    enp0s8:
      dhcp4: no
      dhcp6: no      
    enp0s9:
      dhcp4: no
      dhcp6: no
      addresses: [192.168.57.41/24]
      gateway4: 192.168.57.1  
      routes:
        - to: 192.168.57.0/24
          via: 192.168.57.1
    enp0s10:
      dhcp4: yes
      dhcp6: no 
      routes:
        - to: default
          via:10.0.5.2     
```

## hosts, ssh-keygen, ssh-copy-id

```sh
jhyunlee@Good:~$ cat /etc/hosts
192.168.56.11 controller
192.168.56.31 compute
192.168.56.41 storage

ssh-keygen -P ""
ssh-copy-id controller
ssh-copy-id compute
ssh-copy-id storage

```

