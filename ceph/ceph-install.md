

## 개발환경

### VM 환경

* ESXi :  70.60.31.89  root/dlatl123#
* jhy-c01: 192.168.105.51 ~ 59  ceph/Good
* hostnamectl  set-hostname c01 ~ c09

```
# cat 00-installer-config.yaml
# This is the network config written by 'subiquity'
network:
  ethernets:
    ens160:
      addresses:
      - 192.168.105.51/24
      gateway4: 192.168.105.1
      nameservers:
        addresses: []
        search: []
  version: 2
```



```
# netplan apply
```



```
192.168.105.51 c01
192.168.105.52 c02
192.168.105.53 c03
192.168.105.54 c04
192.168.105.55 c05
192.168.105.56 c06
192.168.105.57 c07
192.168.105.58 c08
192.168.105.58 c09
```

#### 1. NAT 설정

```
vi /etc/sysctl.conf
net.ipv4.ip_forward=1
 
vi /usr/local/bin/nat.sh
#!/bin/bash
iptables -A FORWARD -o ens160 -j ACCEPT
iptables -t nat -A POSTROUTING -o ens160 -j MASQUERADE
 
vi /lib/systemd/system/nat.service
[Unit]
After=network.target
 
[Service]
ExecStart=/usr/local/bin/nat.sh
 
[Install]
WantedBy=default.target
 
chmod 744 /usr/local/bin/nat.sh
chmod 644 /lib/systemd/system/nat.service
```

#### 2. ssh ternneling 

```
c:/ ssh -L 7443:192.168.105.51:8443 ceph@70.60.31.95
"ceph" 입력
브라우저에서 https://localhost:7443 입력
```



#### 3. Master  node에서 cephadm

```
# cat /etc/ceph/ceph.conf
# minimal ceph.conf for 1373ff90-46ab-678d-2467-8030180fe5c0
[global]
        fsid = 1373ff90-46ab-678d-2467-8030180fe5c0
        mon_host = [v2:192.168.105.51:3300/0,v1:192.168.105.51:6789/0] [v2:192.168.105.52:3300/0,v1:192.168.105.52:6789/0] [v2:192.168.105.53:3300/0,v1:192.168.105.53:6789/0] [v2:192.168.105.54:3300/0,v1:192.168.105.54:6789/0] [v2:192.168.105.55:3300/0,v1:192.168.105.55:6789/0] [v2:192.168.105.56:3300/0,v1:192.168.105.56:6789/0]
[mon.c01]
public network = 192.168.105.0/24
mon_allow_pool_delete = true
```

#### 4. object gateway 실행\

```
# ceph orch apply rgw foo
```

* create user --> get key
* create bucket 



### Ceph  Cluster 재구성

#### 1. 기존 ceph cluster 제거

```
# for i in  {1..6}; do scp  reset.sh  c0${i}:~ ; done
```

* reset.sh

```
#!/bin/bash
 
docker container rm -f $(docker container ls -aq)
docker image rm -f $(docker image ls -q)
sleep 3
docker container rm -f $(docker container ls -aq)
docker image rm -f $(docker image ls -q)
sleep 3
docker container rm -f $(docker container ls -aq)
docker image rm -f $(docker image ls -q)
sleep 3
docker container rm -f $(docker container ls -aq)
docker image rm -f $(docker image ls -q)
 
apt-get purge -y docker-ce docker-ce-cli containerd.io docker-ce-rootless-extras docker-compose-plugin docker-scan-plugin docker-buildx-plugin
apt-get autoremove -y docker-ce docker-ce-cli docker-ce-rootless-extras docker-compose-plugin docker-scan-plugin docker-buildx-plugin
 
rm -rf /var/lib/docker /etc/docker
rm -rf /var/run/docker.sock
 
systemctl | grep 'ceph-' | awk '{print $2}' | grep -v loaded | xargs systemctl stop
systemctl | grep 'ceph-' | awk '{print $2}' | grep -v loaded | xargs systemctl disable
 
systemctl | grep 'system-ceph' | awk '{print $1}' | xargs systemctl stop
systemctl | grep 'system-ceph' | awk '{print $1}' | xargs systemctl disable
 
systemctl | grep 'ceph' | awk '{print $2}' | grep -v loaded | xargs systemctl stop
systemctl | grep 'ceph' | awk '{print $2}' | grep -v loaded | xargs systemctl disable
 
rm -rf  /etc/systemd/system/ceph-*
rm -rf /etc/systemd/system/ceph.*
systemctl daemon-reload
systemctl reset-failed
```

* install docker.sh

```
#!/bin/bash 
curl -fsSL http://nexus.sdsdev.co.kr:8081/nexus/content/repositories/docker-ubuntu/gpg | sudo apt-key add 
apt update 
apt install -y docker-ce
apt install -y docker-ce-cli
apt install -y containerd.io 
docker version
docker compose version
```



```
# chmod  755 *.sh
# for i in  {1..6}; do scp  *.sh  c0${i}:~ ; done

# for i in  {1..6}; do rsh c0${i}  /root/reset.sh ; done
# for i in  {1..6}; do rsh c0${i}  docker ps  ; done
```



#### 2.so-ceph-controller로 설치 

https://code.sdsdev.co.kr/carl/so-ceph-controller

```
# git clone https://code.sdsdev.co.kr/carl/so-ceph-controller.git
 최신 release Tag: v1.1.0.20230310.0
# vi /etc/apt/sources.list
# apt install -y ntp
# vi /etc/ntp.conf
  server ntp.carl.io
# apt install -y make
# cp deploy/CARL.crt /usr/local/share/ca-certificates/
# update-ca-certificates

# curl -fsSL http://nexus.sdsdev.co.kr:8081/nexus/content/repositories/docker-ubuntu/gpg | sudo apt-key add
# vi /etc/apt/sources.list.d/docker-ce.list
deb [arch=amd64] http://nexus.sdsdev.co.kr:8081/nexus/content/repositories/docker-ubuntu/ focal stable
# apt update

# apt-get install -y docker-ce
# apt-get install -y docker-ce-cli
# apt-get install -y containerd.io

# vi so-ceph-controller/.env
# vi so-ceph-controller/deploy/host.yml
# make prepare
# make create-spec
# make deploy
# docker logs -f so-adapter-ceph
# make destroy <<---- destroy 
```

#### 3. osd 정리 등록

```
# for i in {1..6}; do ceph orch device zap c0${i} /dev/sdb  --force ; done
# for i in {1..6}; do ceph orch device zap c0${i} /dev/sdc  --force ; done
# for i in {1..6}; do ceph orch device zap c0${i} /dev/sdd  --force ; done
# for i in {1..6}; do ceph orch device zap c0${i} /dev/sde  --force ; done
# ceph orch apply osd --all-available-devices
```

* 자동으로 등록됨.

```
# cephadm shell  --fsid  0824744e-a65e-7c16-8bca-5317bd09676f
```



#### 4. fsid 여러개 일때  rm-cluster 

```
c# cephadm shell  -- ceph -v
ERROR: Cannot infer an fsid, one must be specified (using --fsid): ['0824744e-a65e-7c16-8bca-5317bd09676f', '1373ff90-46ab-678d-2467-8030180fe5c0']

# cephadm ls

# cephadm rm-cluster --fsid 1373ff90-46ab-678d-2467-8030180fe5c0 --force
```



#### 5. Ceph Cluster Purging 

```
# ceph mgr moudle disable cephadm
# ceph fsid
# cephadm rm-cluster --force --zap-osds --fsid <fsid>
```





## Ceph install : Ceph Guide

#### 1. host setup

```
# cat /etc/hosts
192.168.57.11 c1
192.168.57.12 c2
192.168.57.13 c3

# hostnamectl set-hostname c1
# root@c1:~# cat /etc/netplan/00-installer-config.yaml
# This is the network config written by 'subiquity'
network:
  ethernets:
    enp0s3:
      dhcp4: true
    enp0s8:
      addresses:
      - 192.168.57.11/24
      gateway4: 192.168.57.1
      nameservers:
        addresses: []
        search: []
  version: 2
root@c1:~# netplan apply

```

* docker 설치, package 설치

```
# sudo apt-get install ca-certificates curl gnupg lsb-release

# curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

#  echo \
  "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# cat /etc/apt/sources.list.d/docker.list
deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu   focal stable

# apt update 
# apt-get install -y docker-ce  docker-ce-cli containerd.io

# cp deploy/CARL.crt /usr/local/share/ca-certificates/
# update-ca-certificates
# vi /etc/apt/sources.list
# apt install -y ntp
# vi /etc/ntp.conf
# apt install -y make
```



#### 2. install ceph adm  

```
# curl --silent --remote-name --location https://github.com/ceph/ceph/raw/octopus/src/cephadm/cephadm

# chmod +x cephadm
# ./cephadm add-repo --release octopus
# ./cephadm install
# which cephadm
/usr/sbin/cephadm

# which cephadm
/usr/sbin/cephadm

# mkdir -p /etc/ceph
# cephadm bootstrap --mon-ip  192.168.57.11
Ceph Dashboard is now available at:

URL: https://c1:8443/
User: admin
Password: 9hpblkbcnu
sudo /usr/sbin/cephadm shell --fsid 420b849a-cb10-11ed-91b1-5d6354a9bda7 -c /etc/ceph/ceph.conf -k /etc/ceph/ceph.client.admin.keyring

# cephadm shell
# cephadm add-repo --release octopus
# cephadm install ceph-common

# ceph  -v
ceph version 15.2.17 (8a82819d84cf884bd39c17e3236e0632ac146dc4) octopus (stable)

```

#### 3. add host to cluster

```
# ssh-copy-id -f -i /etc/ceph/ceph.pub root@c1
# ssh-copy-id -f -i /etc/ceph/ceph.pub root@c2
# ssh-copy-id -f -i /etc/ceph/ceph.pub root@c3
```

```
# cephadm shell
# ceph orch host add c2
# ceph orch host add c3


# ceph config set mon public_network 192.168.57./24
# ceph  orch apply mon 3
# ceph orch host label add c1 _admin
# ceph orch apply mon *<host1,host2,host3,...>*
# ceph orch host ls
# ceph orch apply mon label:mon
```

* yaml로 설치하기.

```
# ceph orch apply -i file.yaml

service_type: mon
placement:
  hosts:
   - host1
   - host2
   - host3
```

* osd 정리

```
# ceph orch device zap c1 /dev/sdb --force
# ceph orch device zap c2 /dev/sdb --force
# ceph orch device zap c3 /dev/sdb --force
# ceph orch apply osd --all-available-devices
# ceph orch daemon add osd host1:/dev/sdb
```



* MDS

```
# ceph orch apply mds *<fs-name>* --placement="*<num-daemons>* [*<host1>* ...]"
```



* RGW : realm, zone 설치

```
# ceph orch apply rgw *<realm-name>* *<zone-name>* --placement="*<num-daemons>* [*<host1>* ...]"
# ceph orch apply rgw myorg us-east-1 --placement="2 myhost1 myhost2"
```

```
# radosgw-admin realm create --rgw-realm=myorg --default
# radosgw-admin zonegroup create --rgw-zonegroup=myzonegroup --master --default
# radosgw-admin zone create --rgw-zonegroup=myzonegroup --rgw-zone=myzone --master --default
# radosgw-admin period update --rgw-realm=myorg --commit

# ceph orch apply rgw myorg myzone --placement="2 c1 c2"

# radosgw-admin zone get
```

* user add : dashboard

```
# radosgw-admin user create --uid=dashboard --display-name=dashboard --system
# radosgw-admin user info --uid=dashboard
```

* RGW 화면이 안 보일 때 : 뭔가 user 를 생성해줘야 하는 듯

```
# radosgw-admin user create --uid=dashboard --display-name=dashboard --system
root@c1:/# radosgw-admin user info --uid=dashboard
{
    "user_id": "dashboard",
    "display_name": "dashboard",
     "keys": [
        {
            "user": "dashboard",
            "access_key": "B1S1DYWLG52YNPXFJ7F0",
            "secret_key": "noi7877VTggsLmNDqQj0uCpfeaUcH4UTW93j6t7u"
        }
    ],
}

# echo  "B1S1DYWLG52YNPXFJ7F0" > akey
# echo  "noi7877VTggsLmNDqQj0uCpfeaUcH4UTW93j6t7u" > skey
# ceph dashboard set-rgw-api-access-key -i akey
# ceph dashboard set-rgw-api-secret-key -i skey
```





#### 4. fsid 여러개 일때  rm-cluster 

```
c# cephadm shell  -- ceph -v
ERROR: Cannot infer an fsid, one must be specified (using --fsid): ['0824744e-a65e-7c16-8bca-5317bd09676f', '1373ff90-46ab-678d-2467-8030180fe5c0']

# cephadm ls

# cephadm rm-cluster --fsid 1373ff90-46ab-678d-2467-8030180fe5c0 --force
```





