### 1. node
1. controller 
2. compute
3. storge

가이드: https://www.youtube.com/watch?v=uEYxBaFq_V4&t=553s

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

## virtual box host-only network 
### /etc/vbox/networks.conf 설정 
==> 이것 지정해야만 network 주소를 설정할 수 있음
```sh
$ sudo vi /etc/vbox/networks.conf
#* 10.10.0.0/24 10.10.1.0/24 10.10.0.0/24 10.10.1.0/24
* 0.0.0.0/0 ::/0
```

## setup controller node

1. connection ready 
```sh
$ sudo hostnamectl set-hostname controller
$ sudo date
$ sudo date -s "2024-05-01 12:00:00"
$ sudo hwclk -w 

$ sudo ipaddr add 10.10.0.11/24 dev enp0s3
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
$ ssh stack@10.10.0.11
$ sudo vi  /etc/hosts
10.10.0.11 controller
10.10.0.31 compute
10.10.0.41 storage


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
      addresses: [10.10.0.11/24]
      gateway4: 10.10.0.1
      nameservers:
        addresses: [168.126.63.1,168.126.63.2]
      routes:
        - to: 10.10.0.0/24
          via: 10.10.0.1
    enp0s8:
      dhcp4: no
      dhcp6: no      
    enp0s9:
      dhcp4: no
      dhcp6: no
      addresses: [10.10.1.11/24]
      gateway4: 10.10.1.1  
      routes:
        - to: 10.10.1.0/24
          via: 10.10.1.1
    enp0s10:
      dhcp4: yes
      dhcp6: no 
      routes:
        - to: default
          via: 10.0.5.2     
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
      addresses: [10.10.0.31/24]
      gateway4: 10.10.0.1
      nameservers:
        addresses: [168.126.63.1,168.126.63.2]
      routes:
        - to: 10.10.0.0/24
          via: 10.10.0.1
    enp0s8:
      dhcp4: no
      dhcp6: no      
    enp0s9:
      dhcp4: no
      dhcp6: no
      addresses: [10.10.1.31/24]
      gateway4: 10.10.1.1  
      routes:
        - to: 10.10.1.0/24
          via: 10.10.1.1
    enp0s10:
      dhcp4: yes
      dhcp6: no 
      routes:
        - to: default
          via: 10.0.5.2     
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
      addresses: [10.10.0.41/24]
      gateway4: 10.10.0.1
      nameservers:
        addresses: [168.126.63.1,168.126.63.2]
      routes:
        - to: 10.10.0.0/24
          via: 10.10.0.1
    enp0s8:
      dhcp4: no
      dhcp6: no      
    enp0s9:
      dhcp4: no
      dhcp6: no
      addresses: [10.10.1.41/24]
      gateway4: 10.10.1.1  
      routes:
        - to: 10.10.1.0/24
          via: 10.10.1.1
    enp0s10:
      dhcp4: yes
      dhcp6: no 
      routes:
        - to: default
          via: 10.0.5.2     
```

### hosts, ssh-keygen, ssh-copy-id

```sh
jhyunlee@Good:~$ cat /etc/hosts
10.10.0.11 controller
10.10.0.31 compute
10.10.0.41 storage

ssh-keygen -P ""
ssh-copy-id controller
ssh-copy-id compute
ssh-copy-id storage

```
### git clone 
```sh
$ git clone https://github.com/Sangwan70/openstack-zed.git
$ scp -r  openstack-zed  compute:~
$ scp -r  openstack-zed  storage:~
```


### /etc/vbox/networks.conf 설정 
```sh
$ sudo vi /etc/vbox/networks.conf
#* 10.10.0.0/24 10.10.1.0/24 10.10.0.0/24 10.10.1.0/24
* 0.0.0.0/0 ::/0
```
-> 주의할 점은 별표로 시작한다는 점이다. 
-> host only network 에 뭔가 취약점이 있어서 이것의 시작 주소를 제한 한 것 같다. 


==> host-only network 지정할 때 오류 발생하는 현상 관련

https://forums.virtualbox.org/viewtopic.php?t=104357

```
오류: 호스트 전용 네트워크 - 호스트 네트워크 인터페이스 매개변수를 저장하지 못했습니다. - E_ACCESSDENIED
우편 ~에 의해마르셀_»3. 2021년 11월 18:57

내 호스트 전용 네트워크의 IP 주소나 서브넷 마스크를 변경하려고 하면 다음 오류 코드가 나타납니다.
호스트 네트워크 인터페이스 매개변수를 저장하지 못했습니다.
수신자 RC:
E_ACCESSDENIED (0x80070005)
해결 방법을 아는 사람이 있나요? 나는 아치 기반의 리눅스 시스템을 가지고 있습니다
맨 위
스코거스1
사이트 중재자
게시물: 20945
가입일: 2009년 12월 30일, 20:14
기본 OS: MS Windows 10
VBox 버전: PUEL
게스트 OS: Windows, Linux
Re: 오류: 호스트 전용 네트워크 - 호스트 네트워크 인터페이스 매개변수를 저장하지 못했습니다. - E_ACCESSDENIED
우편 작성자: scottgus1 »3. 2021년 11월 19:00

Linux 호스트에 대한 새로운 설정이 6.1.28에 도입되었습니다. 보다https://www.virtualbox.org/manual/ch06. ... k_hostonly
Linux, Mac OS X 및 Solaris Oracle VM VirtualBox에서는 192.68.56.0/21 범위의 IP 주소만 호스트 전용 어댑터에 할당되도록 허용합니다. IPv6의 경우 링크 로컬 주소만 허용됩니다. 다른 범위가 필요한 경우 /etc/vbox/networks.conf를 만들고 거기에서 허용되는 범위를 지정하여 활성화할 수 있습니다. 예를 들어, 10.0.0.0/8 및 192.168.0.0/16 IPv4 범위와 2001::/64 범위를 허용하려면 /etc/vbox/networks.conf에 다음 줄을 입력합니다.

* 10.0.0.0/8 192.168.0.0 /16
* 2001::/64

해시 #으로 시작하는 줄은 무시됩니다. 다음 예에서는 모든 주소를 허용하여 범위 제어를 효과적으로 비활성화합니다.

* 0.0.0.0/0 ::/0

파일이 있지만 파일에 범위가 지정되지 않은 경우 호스트 전용 어댑터에는 주소가 할당되지 않습니다. 다음 예에서는 모든 범위를 효과적으로 비활성화합니다.

# 호스트 전용 어댑터에는 주소가 허용되지 않습니다 .
```

## install

### 1. ./pre-download.sh
stack@controller:~/scripts$ ./pre-download.sh

### 2. ~/openstack-zed/controller/scripts/ubuntu


./1_apt_init.sh
./2_apt_upgrade.sh
./3_install_mysql.sh
./4_install_rabbitmq.sh
./5_install_memcached.sh
./6_setup_keystone_1.sh
./7_setup_keystone_2.sh
./8_setup_glance_1.sh
./9_setup_glance_2.sh
./10_setup_placement_1.sh
./11_setup_placement_2.s
./12_setup_nova_1.sh
./13_setup_nova_2.sh
./14_setup_nova_3.sh
./15_setup_nova_4.sh
./16_setup_neutron_1.sh
./17_setup_neutron_2.sh
./18_setup_neutron_3.sh
./19_setup_neutron_4.sh
./20_setup_horizon.sh
./21_setup_cinder_1.sh
./22_setup_cinder_2.sh
./23_setup_cinder_3.sh
./24_setup_heat_1.sh
./25_setup_heat_2.sh
./26_setup_swift_1.sh
./27_setup_swift_2.sh
./28_setup_barbican_1.sh
./29_setup_barbican_2.sh
./30_setup_barbican_3.sh
./31_setup_swift_3.sh
./32_setup_trove_1.sh
./33_setup_trove_2.sh
./34_setup_trove_3.sh
./35_setup_trove_4.sh
./setup_etcd.sh
./tacker



./1_apt_init.sh
./2_apt_upgrade.sh
./3_setup_nova_1.sh
./4_setup_nova_2.sh
./5_setup_neutron_1.sh
./6_setup_neutron_2.sh
./7_setup_neutron_3.sh
./8_setup_neutron_4.sh
./9_setup_swift_1.sh
./10_setup_swift_2.sh
./11_setup_swift_3.sh
./tacker



#------------------------------------------------------------------------------
# Verify operation
#------------------------------------------------------------------------------

echo "Verifying keystone installation."

# From this point on, we are going to use keystone for authentication
unset OS_AUTH_URL OS_PASSWORD

echo "Requesting an authentication token as an admin user."
openstack \
    --os-auth-url http://controller:5000/v3 \
    --os-project-domain-name Default \
    --os-user-domain-name Default \
    --os-project-name "$ADMIN_PROJECT_NAME" \
    --os-username "$ADMIN_USER_NAME" \
    --os-auth-type password \
    --os-password "$ADMIN_PASS" \
    token issue

echo "Requesting an authentication token for the demo user."
openstack \
    --os-auth-url http://controller:5000/v3 \
    --os-project-domain-name Default \
    --os-user-domain-name Default \
    --os-project-name "$DEMO_PROJECT_NAME" \
    --os-username "$DEMO_USER_NAME" \
    --os-auth-type password \
    --os-password "$DEMO_PASS" \
    token issue
