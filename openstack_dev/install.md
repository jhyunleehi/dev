# openstack install in Virtual Box 7

## Setup virtual box host-only network 
### /etc/vbox/networks.conf 설정 
==> 이것 지정해야만 network 주소를 설정할 수 있음
```sh
$ sudo vi /etc/vbox/networks.conf
#* 10.10.0.0/24 10.10.1.0/24 10.10.0.0/24 10.10.1.0/24
* 0.0.0.0/0 ::/0
```

1. host only network #1 10.10.0.1/24
2. host only network #2 203.0.113.1/24


## controller node 
configuration 
1. cpu :1
2. mem :1G
3. disk: 40G, 4G
4. nic1 hostonly network#1
5. nic2 hostonly network#2 
6. nic3 not config
7. nic4 NAT 203.0.113

install ubunut 22.04
1. hostname : controller
2. user : stack/ok

ubunut setting 
```sh
$ sudo hostnamectl set-hostname controller
$ sudo date
$ sudo date -s "2024-05-01 12:00:00"
$ sudo hwclock -w 

$ sudo ipaddr add 10.10.0.11/24 dev enp0s3


$ sudo ip link set dev enps10 down
$ sudo ip link set dev enps10 up

$ scp  gpu@10.10.0.1:/etc/hosts .
$ scp  gpu@10.10.0.1:/etc/apt/sources.list .
$ scp  gpu@10.10.0.1:/etc/netplan/01* .
$ scp  gpu@10.10.0.1:SDS.crt .
$ scp  gpu@10.10.0.1:/etc/environment .

$ sudo  cp hosts /etc/hosts
$ sudo  cp sources.list /etc/apt/sources.list
$ sudo  cp  01*  /etc/netplan/
$ sudo  cp  SDS.crt /usr/local/share/ca-certificates/ 
$ sudo  update-ca-certificates
$ sudo  cp environment /etc/environment
$ sudo apt update --fix-missing

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

$ sudo resolvectl status | grep  "DNS Server"
Current DNS Server: 70.10.98.4
       DNS Servers: 70.10.98.4 203.241.132.34

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
        addresses: [70.10.98.4,203.241.132.34]
      routes:
        - to: 10.10.0.0/24
          via: 10.10.0.1
    enp0s8:
      dhcp4: no
      dhcp6: no                
    enp0s10:
      dhcp4: yes
      dhcp6: no 
      routes:
        - to: default
          via: 10.0.5.2     
```

sudo ip route delete default via 10.10.0.1 dev enp0s3


## shutdown and clone compute, storage 

### compute node 
1. cpu 1
2. mem 1G
3. disk 40G 4G 4G 4G
4. nw : NIC1,NIC2,NIC3,NIC4
5. set hostname
6. ip addr set
7. ssh-keygen
8. ssh-copy-id 

### storage node 
1. cpu 1
2. mem 1G
3. disk 40G 4G
4. nw : NIC1,NIC2,NIC3,NIC4
5. set hostname
6. ip addr set
7. ssh-keygen
8. ssh-copy-id 


## openstack installer 

```sh
$ git clone https://github.com/Sangwan70/openstack-zed.git
Cloning into 'openstack-zed'...

$ scp  -r openstack-zed/ compute:~
$ scp  -r openstack-zed/ storage:~
$ ./pre-download.sh 
[sudo] password for stack: 
--2024-05-01 22:15:22--  http://download.cirros-cloud.net/0.4.0/MD5SUMS
Connecting to 70.10.15.10:8080... 
--no-check-certificate

stack@controller:~$ sudo ip route delete default via 10.10.0.1 dev enp0s3
stack@controller:~$ sudo ip route delete  default via 10.0.5.2 dev enp0s10

stack@controller:~$ ip route 
default via 10.0.5.2 dev enp0s10 proto dhcp src 10.0.5.15 metric 100  <<== only one 

==> cheeck wget 
stack@controller:~$ wget www.naver.com
--2024-05-02 00:12:48--  http://www.naver.com/
Connecting to 70.10.15.10:8080... connected.
Proxy request sent, awaiting response... 302 Moved Temporarily
Location: https://www.naver.com/ [following]
--2024-05-02 00:12:48--  https://www.naver.com/
Connecting to 70.10.15.10:8080... connected.
ERROR: cannot verify www.naver.com's certificate, issued by ‘emailAddress=infosec@samsung.com,CN=SDS,O=SAMSUNG SDS,L=Gangman-gu,ST=Seoul,C=KR’:
  CA certificate key too weak
To connect to www.naver.com insecurely, use `--no-check-certificate'.
```

```sh
stack@controller:~/openstack-zed/controller/scripts$ ./pre-download.sh

2024-05-02 00:17:34 (175 KB/s) - ‘/home/stack/openstack-zed/controller/img/cirros-0.4.0-x86_64-disk.img’ saved [12716032/12716032]

cirros-0.4.0-x86_64-disk.img: OK
```

```sh
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
```
