# Ceph

## Command

### 1. ceph

* 기본

```
# ceph orch ls 
# ceph orsch ps 
# ceph health 
# ceph status 
# ceph mon dump
# ceph mgr stat 
# ceph osd pool ls 
# ceph pg stat 
# ceph osd status
# ceph osd tree 
```



```
# ceph health 
# ceph status 
# cephadm shell -ceph orch ls
# cephadm shell
# ceph config get osd public_network
# ceph config get mon public_network
# ceph config show mon.serverc.lab.example.com mon_host
# ceph config dump

# ceph -w
# ceph config set mgr mgr/cephadm/log_to_cluster_level debug
# ceph -W cephadm --watch-debug

# ceph log last cephadm
# cephadm logs --name <name-of-daemon>
# cephadm logs --fsid <fsid> --name <name-of-daemon>

# ceph crash -l 
# ceph crash  info  <crash-id>


# cephadm shell
# ceph orch host ls 
# ceph orch host ls --detail 
# ceph orch host add c01

# ceph orch ls
# ceph orch ps 
# ceph orch ps  --daemon_type  osd

# ceph orch device ls
# ceph orch device ls —hostname=servere.lab.example.com
# ceph orch device zap <node> /dev/vda --force
# ceph orch apply osd --all-available-devices
# ceph orch apply osd --all-available-devices --unmanaged=true
# ceph orch daemon add osd c06:/dev/sdc 

# ceph orch apply mon c01,c02,c03,c04,c05,c06

# ceph orch daemon add osd node:/dev/vdb
# ceph orch daemon stop osd.12
# ceph orch daemon rm osd.12


# ceph osd status 
# ceph osd tree
# ceph mgr stat 
# ceph osd pool ls
# ceph osd df 
# ceph osd dump 
# ceph osd rm 12

# ceph osd crush tree 
# ceph orch device ls 

# ceph mon dump 
# ceph mon stat
# ceph pg stat 

# ceph-volume lvm create --bluestore --data /dev/vdc
# ceph-volume lvm prepare --bluestore --data /dev/vdc
# ceph-volume lvm activate <osd-fsid>
# ceph-volume lvm batch --bluestore /dev/vdc /dev/vdd /dev/nvme0n1
# ceph-volume inventory


# ceph osd set noout
# ceph osd set nobackfill
# ceph osd set norecover
# ceph osd set norebalance
# ceph osd set noscrub
# ceph osd set nodeep-scrub
```



```
# ceph osd pool create pool-name pg-num pgp-num erasure erasure-code-profile crush-rule-name 
```

* pool-name: 새 풀의 이름

* pg-num: 풀에 대한 총 PG 수
* pgp-num: 풀에 대해 유효한 PG 수 (일반적으로 총 PG 수와 같음)
* erasure: 이레이저 코딩 풀이라고 지정
* erasure-code-profile: 사용할 프로필 이름
* crush-rule-name: 풀에 사용할 CRUSH Rule의 이름

```
# ceph  orch apply rgw myorg ko-mid-1 --placement="2 c1 c2"
```

##### 상태 모니터링

```
# ceph orch ls 
# ceph orsch ps 
# ceph health 
# ceph status 
# ceph mon dump
# ceph mgr stat 
# ceph osd pool ls 
# ceph pg stat 
# ceph osd status
# ceph osd tree 
```





### 2. osd

```
# ceph osd lspools
# osd pool default pg num = 100
# osd pool default pgp num = 100
# ceph osd pool create {pool-name} {pg-num} [{pgp-num}] [replicated] [crush-ruleset-name] 
# ceph osd pool create {pool-name} {pg-num}  {pgp-num} erasure [erasure-code-profile] [crush-ruleset-name] [expected_num_objects]
# ceph osd pool rename {current-pool-name} {new-pool-name}
# ceph osd pool delete {pool-name} [{pool-name} --yes-i-really-really-mean-it]
# ceph osd pool set-quota data max_objects 10000
# rados df
# ceph osd pool set {poolname} size {num-replicas}
```

##### osd 추가

```
# ceph orch device ls
# ceph orch device ls —hostname=servere.lab.example.com
# ceph orch device zap c01 /dev/sdb  --force
# ceph orch apply osd --all-available-devices
# ceph orch apply osd --all-available-devices --unmanaged=true

# ceph orch daemon add osd c06:/dev/sdc 
# ceph osd tree
```

* osd 추가할때 아래 메세지 나오면 osd 정보가 완전히 정리되지 않아서 그런거다. 

```
root@c01:/# ceph orch daemon add osd c03:/dev/sdb
Created no osd(s) on host c03; already created?
```

##### osd stop -> osd  destroy -> osd  rm 

```
# cehp osd out osd.1
# ceph osd stop osd.1
# ceph osd rm osd.1
# ceph osd tree  <<=== 찌꺼기 있으면 Crush에서 제거
# cehp osd crush remove osd.1
# ceph osd destroy 1
```





##### osd disk 제거

* device만 zap 해서는 확실하게 repository에서 빠지지 않는다.
* daemon도 죽이고 crush에서 bucket 정보도 빼주고 auth도 모두 제거해줘야 다시 붙일 수 있다  (재활용)

```
# ceph orch daemon stop osd.0
# ceph orch device zap c01 /dev/sdb --force
# ceph osd crush remove osd.0
# ceph auth del osd.0
# ceph osd rm 0
 
### 만약 아래와 같이 에러가 발생한다면:
# ceph osd rm 9
# ceph osd down osd.9
# ceph osd rm 9
```

```
# ceph osd destroy 17 --yes-i-really-mean-it
```



##### osd 삭제

```
# ceph osd out osd.${num}
# ceph orch daemon stop osd.${num}
# ceph osd purge ${num} --yes-i-really-mean-it
# ceph orch daemon rm osd.${num} --force
# ceph osd crush remove osd.${num}
# ceph osd auth del osd.${num}
# ceph osd rm ${num}
```



#####  osd destroyed, exists 

```
root@c01:/# ceph osd status
ID  HOST   USED  AVAIL  WR OPS  WR DATA  RD OPS  RD DATA  STATE
19  c01      0      0       0        0       0        0   destroyed,exists  <<---
root@c01:/# ceph orch  device zap  c01 /dev/sdd --force
zap successful for /dev/sdd on c01

root@c01:/# ceph osd status
ID  HOST   USED  AVAIL  WR OPS  WR DATA  RD OPS  RD DATA  STATE
19  c01      0      0       0        0       0        0   exists,new <<-- 다시 올라옴.

```





### 3.erasure-code

```
# ceph osd erasure-code-profile ls
# ceph osd erasure-code-profile get default
# ceph osd erasure-code-profile set ecprofile-4k2m k=4 m=2 plugin=jerasure crush-failure-domain=host
# ceph osd erasure-code-profile set ecprofile-19k5m k=19 m=5 plugin=jerasure crush-failure-domain=host

# ceph osd erasure-code-profile ls
# ceph osd  erasure-code-profile get ecprofile-19k5m
# ceph osd  erasure-code-profile get ecprofile-4k2m

# ceph osd erasure-code-profile get ecprofile-4k2m
# ceph osd pool create ec-4k2m-pool  64 64 erasure ecprofile-4k2m
# ceph osd pool create ec-19k5m-pool 64 64 erasure ecprofile-19k5m

#  ceph osd lspools
3 ec-19k5m-pool
4 ec-4k2m-pool

# ceph osd erasure-code-profile rm  ecprofile-4k2m
```



```
# ceph osd erasure-code-profile get <name>  
# ceph osd erasure-code-profile ls
# ceph osd erasure-code-profile rm <name> 
# ceph osd erasure-code-profile set <name> [<profile>...] [--force]                          
```



### 4. pool

##### pool 생성

```
# ceph osd lspools
# ceph osd pool ls detail
# ceph osd pool get ec-19k5m-pool all
# ceph osd pool get ec-19k5m-pool all

# ceph osd pool delete ec-4k2m-pool ec-4k2m-pool --yes-i-really-really-mean-it
Error EPERM: pool deletion is disabled; you must first set the mon_allow_pool_delete config option to true before you can destroy a pool

# vi /etc/ceph/ceph.conf
mon_allow_pool_delete = true  << 라인추가
# ceph tell mon.\* injectargs --mon-allow-pool-delete=true
# ceph osd pool delete ec-4k2m-pool ec-4k2m-pool --yes-i-really-really-mean-it

```

##### pool 관리

```
# ceph osd lspools
# ceph osd pool ls detail

# ceph osd pool rename 
# ceph osd pool delete
# ceph osd pool set pool_name nodelete true
# ceph osd pool set
# ceph osd pool get
 
# ceph df
# ceph osd pool stats
# ceph osd pool application enable
# ceph osd pool set-quota
```



##### erasure-code pool

```
# ceph osd erasure-code-profile ls
# ceph osd erasure-code-profile get default
# ceph osd erasure-code-profile set ecprofile-4k2m k=4 m=2 plugin=jerasure crush-failure-domain=host
# ceph osd erasure-code-profile set ecprofile-19k5m k=19 m=5 plugin=jerasure crush-failure-domain=host

# ceph osd erasure-code-profile ls
# ceph osd erasure-code-profile get ecprofile-19k5m
# ceph osd erasure-code-profile get ecprofile-4k2m

# ceph osd erasure-code-profile get ecprofile-4k2m
# ceph osd pool create ec-4k2m-pool  64 64 erasure ecprofile-4k2m
# ceph osd pool create ec-19k5m-pool 64 64 erasure ecprofile-19k5m

#  ceph osd lspools
3 ec-19k5m-pool
4 ec-4k2m-pool

# ceph osd erasure-code-profile rm  ecprofile-4k2m
```



```
# ceph osd erasure-code-profile get <name>  
# ceph osd erasure-code-profile ls
# ceph osd erasure-code-profile rm <name> 
# ceph osd erasure-code-profile set <name> [<profile>...] [--force]                          
```





##### pool 제거

```
# ceph osd pool delete ec-4k2m-pool ec-4k2m-pool --yes-i-really-really-mean-it
Error EPERM: pool deletion is disabled; you must first set the mon_allow_pool_delete config option to true before you can destroy a pool

# vi /etc/ceph/ceph.conf
mon_allow_pool_delete = true  << 라인추가
# ceph tell mon.\* injectargs --mon-allow-pool-delete=true
# ceph osd pool delete ec-4k2m-pool ec-4k2m-pool --yes-i-really-really-mean-it


# ceph osd pool delete  mypool1 mypool1  --yes-i-really-really-mean-it
# ceph osd pool delete  mypool2 mypool2  --yes-i-really-really-mean-it
```



### 5. PG

```
# ceph osd lspools
# ceph pg ls 
# ceph pg ls-by-pool ec-4k2m-pool
# ceph pg ls-by-pool default.rgw.buckets.data
# ceph pg ls-by-osd osd.1
# ceph pg ls-by-primary osd.1
# ceph pg map 3.1c
```

- PG: Placement Group의 ID를 나타냅니다.
- OBJECTS: 해당 PG에 속한 오브젝트(object)의 수를 나타냅니다.
- DEGRADED: 해당 PG에서 복제(replication)된 오브젝트 중에서 몇 개가 손상되었는지를 나타냅니다.
- MISPLACED: 해당 PG에서 복제(replication)된 오브젝트 중에서 몇 개가 다른 OSD(Object Storage Daemon)에 저장되어야 하는데 잘못 저장된 상태인지를 나타냅니다.
- UNFOUND: 해당 PG에서 복제(replication)된 오브젝트 중에서 몇 개가 찾을 수 없는 상태인지를 나타냅니다.
- BYTES: 해당 PG에 속한 모든 오브젝트의 크기 합계를 바이트 단위로 나타냅니다.
- OMAP_BYTES*: 해당 PG에 저장된 모든 오브젝트의 omap 데이터 크기를 바이트 단위로 나타냅니다.
- OMAP_KEYS*: 해당 PG에 저장된 모든 오브젝트의 omap 데이터 키(key)의 수를 나타냅니다.
- LOG: 해당 PG의 로그(log) 크기를 나타냅니다.
- STATE: 해당 PG의 상태(state)를 나타냅니다.
- SINCE: 해당 PG의 상태가 언제부터 지속되고 있는지를 나타냅니다.
- VERSION: 해당 PG의 버전(version)을 나타냅니다.
- REPORTED: 해당 PG의 상태가 언제 마지막으로 보고되었는지를 나타냅니다.
- UP: 해당 PG에서 복제(replication)된 오브젝트가 어느 OSD에서 작동하고 있는지를 나타냅니다.
- ACTING: 해당 PG에서 복제(replication)된 오브젝트를 어느 OSD에서 읽고 쓸 수 있는지를 나타냅니다.
- SCRUB_STAMP: 해당 PG의 스크럽(scrub)이 마지막으로 수행된 시간을 나타냅니다.
- DEEP_SCRUB_STAMP: 해당 PG의 딥 스크럽(deep scrub)이 마지막으로 수행된 시간을 나타냅니다.
- AST_SCRUB_DURATION: 해당 PG의 자동 스크럽(scrub) 기간을 나타냅니다. (AST는 Automatic Scrubbing and Repairing Technology의 약어입니다.)
- SCRUB_SCHEDULING: 해당 PG의 스크럽(scrub)이 예정된 시간을 나타냅니다.



#### File -> Object-> Pool -> PG

```
# ceph osd map ec-4k2m-pool n01484850_8845.JPEG
osdmap e234 pool 'ec-4k2m-pool' (12) object 'n01484850_8845.JPEG' -> pg 12.e026cb5f (12.1f) -> up ([4,0,5,3,1,2], p4) acting ([4,0,5,3,1,2], p4)
```

- 특정 pool과 object의 OSD (Object Storage Device) 매핑 정보를 조회
- 데이터의 분산 및 복제 상태를 확인하고 문제가 발생한 OSD를 식별할 수 있다
- "up"은 복제된 데이터의 저장소 위치
- "acting"은 현재 데이터의 저장소 위치
- Crush 맵은 데이터를 저장하기 위해 사용 가능한 OSD들을 나타내는 맵



##### 10MB 파일 f1 업로드 

```
# ceph  osd  map  default.rgw.buckets.data  f1
osdmap e979 pool 'default.rgw.buckets.data' (24) object 'f1' -> pg 24.67b909e9 (24.9) -> up ([10,6,3], p10) acting ([10,6,3], p10)
```

* 10MB 파일은 3개 objet로 나눠지고 각각의 PG에 할당되었다.
* 10MB file -> { 24.0, 24.5, 24.6} 3개 pg에 할당되었다. 이 PG는 자신의 OSD할당 모듬을 가지고 있으니 group 이라는 이름을 사용한것 같다. 
* pg 앞에 붙는 24는 pool id 이다. 그리고 거기에 하나씩 pg 일련 번호가 붙는다. 32개이다. 
* 이렇게 덩어리로 관리한다. 
* PG라고 하는 이유가 결국은 그냥 

```
root@c01:/# ceph pg ls-by-pool default.rgw.buckets.data
PG     OBJECTS  DEGRADED  MISPLACED  UNFOUND  BYTES    OMAP_BYTES*  OMAP_KEYS*  LOG  STATE         SINCE  VERSION  REPORTED  UP             ACTING         
24.0         1         0          0        0  4194304            0           0    6  active+clean    28m    955'6    978:61    [4,15,20]p4    [4,15,20]p4  
...
24.5         1         0          0        0  1611392            0           0    6  active+clean    28m    955'6    978:42    [18,4,9]p18    [18,4,9]p18  
24.6         2         0          0        0  4194304            0           0    6  active+clean    28m    955'6    978:51    [5,11,14]p5    [5,11,14]p5  
24.9         0         0          0        0        0            0           0    6  active+clean    28m    955'6    978:36    [10,6,3]p10    [10,6,3]p10  
...
24.1f        0         0          0        0        0            0           0    6  active+clean    28m    955'6    978:36   [16,11,0]p16   [16,11,0]p16 
```



##### pg 조회 

*  pool에서 OSD로 찾아 가는 길인데.. 덩어리로 찾아 간다는 것

```
root@c01:/# ceph pg map 24.0
osdmap e979 pg 24.0 (24.0) -> up [4,15,20] acting [4,15,20]
root@c01:/# ceph pg map 24.5
osdmap e979 pg 24.5 (24.5) -> up [18,4,9] acting [18,4,9]
root@c01:/# ceph pg map 24.6
osdmap e979 pg 24.6 (24.6) -> up [5,11,14] acting [5,11,14]
```





### 6. rados

```
# rados lspools
# rados df
```

##### get

```
# rados --pool default.rgw.buckets.data ls
>> 조회된 결과 중에서 선택
# rados -p default.rgw.buckets.data get edb79dac-6966-41fd-8a17-84f2c06323d8.34467.1_test_data/n01484850/n01484850_8845.JPEG n01484850_8845.JPEG
```

- pool name : default.rgw.buckets.data
- 특정 object 경로 : test_data/n01484850/
- 미리 생성해 둔 local 폴더 이름 : n01484850

```
# for obj in $(rados -p default.rgw.buckets.data ls | grep test_data/n01484850/)
do
rados -p default.rgw.buckets.data get $obj ./n01484850/$(basename $obj)
done
```

##### put

```
# rados -p ec-4k2m-pool-``01` `put n01484850_8845.JPEG /home/testdata/n01484850_8845.JPEG
```

##### delete

```
# rados -p ec19k5m-test-01-pool rm n01484850_8845.JPEG
```



```
# radosgw-admin realm create --rgw-realm=<realm-name> --default
# radosgw-admin zonegroup create --rgw-zonegroup=<zonegroup-name>  --master --default
# radosgw-admin zone create --rgw-zonegroup=<zonegroup-name> --rgw-zone=<zone-name> --master --default
# radosgw-admin period update --rgw-realm=<realm-name> --commit
```



### Object Gateway

#### 1. rgw application 생성

* `cephadm` 을 사용하면 Ceph Object Gateway 데몬은 `ceph.conf` 파일 또는 명령줄 옵션 대신 Ceph Monitor 구성 데이터베이스를 사용하여 구성됩니다. 구성이 `client.rgw` 섹션에 없는 경우 Ceph Object Gateway 데몬은 기본 설정으로 시작하고 포트 `80` 에 바인드됩니다.

##### 방법1

```
# cephadm shell
# radosgw-admin realm create --rgw-realm=test_realm --default
# radosgw-admin zonegroup create --rgw-zonegroup=default  --master --default
# radosgw-admin zone create --rgw-zonegroup=default --rgw-zone=test_zone --master --default
# radosgw-admin period update --rgw-realm=test_realm --commit
# ceph orch apply rgw test --realm=test_realm --zone=test_zone --placement="2 host01 host02"
```

##### 방법2

```
# ceph orch apply rgw foo
```

##### 방법3

```
# ceph orch host label add host01 rgw  # the 'rgw' label can be anything
# ceph orch host label add host02 rgw
# ceph orch apply rgw foo '--placement=label:rgw count-per-host:2' --port=8000

# ceph orch host label add ceph-01 rgw
# ceph orch host label add ceph-02 rgw
# ceph orch host label add ceph-03 rgw
```

#### 2. rgw appplication 제거

```
# cephadm shell
# ceph orch ls
# ceph orch rm rgw.foo
```



```
# radosgw-admin zone get
# radosgw-admin zone rm --rgw-zone=default
# ceph osd pool rm default.rgw.log default.rgw.log --yes-i-really-really-mean-it
# ceph osd pool rm default.rgw.meta default.rgw.meta --yes-i-really-really-mean-it
# ceph osd pool rm default.rgw.control default.rgw.control --yes-i-really-really-mean-it
# ceph osd pool rm default.rgw.data.root default.rgw.data.root --yes-i-really-really-mean-it
# ceph osd pool rm default.rgw.gc default.rgw.gc --yes-i-really-really-mean-it
```





### 7. rbd

##### 1. ceph

```
# rbd create mypool/myimage --size 102400
# rbd create mypool/myimage --size 102400 --object-size 8M
# rbd rm mypool/myimage
# rbd map mypool/myimage --id admin --keyfile secretfile
# rbd unmap /dev/rbd0
```

##### 2. client

```
# rbd ls -l
# rbd  map vol01 --pool <pool-name>
# rbd  map vol02 --pool <pool-name>
# rbd  map vol03 --pool <pool-name>
# lsblk
```

```
[client]
# yum -y install ceph-fuse
# ceph-authtool -p /etc/ceph/ceph.client.admin.keyring > admin.key
# cat admin.key
# mount -t ceph h6:6789:/  /cephfile  -o name=admin,secretfile=/root/admin.key
```



CRUSH – Contraolled, Scalable, Decentralized Placement of Replicated Data



## Ceph tool

- [ceph-volume – Ceph OSD deployment and inspection tool](https://docs.ceph.com/en/quincy/man/8/ceph-volume/)
- [ceph-volume-systemd – systemd ceph-volume helper tool](https://docs.ceph.com/en/quincy/man/8/ceph-volume-systemd/)
- [ceph – ceph administration tool](https://docs.ceph.com/en/quincy/man/8/ceph/)
- [ceph-authtool – ceph keyring manipulation tool](https://docs.ceph.com/en/quincy/man/8/ceph-authtool/)
- [ceph-clsinfo – show class object information](https://docs.ceph.com/en/quincy/man/8/ceph-clsinfo/)
- [ceph-conf – ceph conf file tool](https://docs.ceph.com/en/quincy/man/8/ceph-conf/)
- [ceph-debugpack – ceph debug packer utility](https://docs.ceph.com/en/quincy/man/8/ceph-debugpack/)
- [ceph-dencoder – ceph encoder/decoder utility](https://docs.ceph.com/en/quincy/man/8/ceph-dencoder/)
- [ceph-mon – ceph monitor daemon](https://docs.ceph.com/en/quincy/man/8/ceph-mon/)
- [ceph-osd – ceph object storage daemon](https://docs.ceph.com/en/quincy/man/8/ceph-osd/)
- [ceph-kvstore-tool – ceph kvstore manipulation tool](https://docs.ceph.com/en/quincy/man/8/ceph-kvstore-tool/)
- [ceph-run – restart daemon on core dump](https://docs.ceph.com/en/quincy/man/8/ceph-run/)
- [ceph-syn – ceph synthetic workload generator](https://docs.ceph.com/en/quincy/man/8/ceph-syn/)
- [crushdiff – ceph crush map test tool](https://docs.ceph.com/en/quincy/man/8/crushdiff/)
- [crushtool – CRUSH map manipulation tool](https://docs.ceph.com/en/quincy/man/8/crushtool/)
- [librados-config – display information about librados](https://docs.ceph.com/en/quincy/man/8/librados-config/)
- [monmaptool – ceph monitor cluster map manipulation tool](https://docs.ceph.com/en/quincy/man/8/monmaptool/)
- [osdmaptool – ceph osd cluster map manipulation tool](https://docs.ceph.com/en/quincy/man/8/osdmaptool/)
- [rados – rados object storage utility](https://docs.ceph.com/en/quincy/man/8/rados/)

## 



### AWS Client 설치

#### Clinet 설치

##### 1. S3 client 설치

```
# curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
# unzip awscliv2.zip
# sudo ./aws/install
```

##### 2.  credentials

```
# cat ~/.aws/credentials
[default]
aws_access_key_id = 2VEVI5DW1GOO2QUX14BL
aws_secret_access_key = GefWYyvAlQhszFNZEwDjeDyMa8fS0vCUJ5jSvUgP
```

##### 3. upload file 

```
# dd if=/dev/urandom of=file1 bs=1MB count=10
# aws s3 cp ./file1 s3://mybucket/file1 --endpoint-url http://192.168.105.51:80
upload: ./file1 to s3://mybucket/file1
```

* 하위 폴더 포함해서 many file 

```
# aws s3 sync ./test_data/ s3://mybucket/test_data/ --endpoint-url http://192.168.105.11:80
```

* 삭제

```
# aws s3  rm    s3://mybucket/f1  --endpoint-url http://192.168.105.51:80
delete: s3://mybucket/f1
```



##### 4. get list

```
# aws s3  ls mybucket --endpoint-url http://192.168.105.51:80
2023-03-19 05:32:00   10000000 file1
```

##### 5. download

```
# aws s3 cp s3://mybucket/file1 ~/f2 --endpoint-url http://192.168.105.51:80
```

##### 6. pg map

```
# ceph  osd  map  default.rgw.buckets.data  f1
osdmap e979 pool 'default.rgw.buckets.data' (24) object 'f1' -> pg 24.67b909e9 (24.9) -> up ([10,6,3], p10) acting ([10,6,3], p10)
```

* pg 목록

```
root@c01:/# ceph pg ls-by-pool default.rgw.buckets.data
PG    OBJECTS  DEGRADED  MISPLACED  UNFOUND  BYTES   LOG  STATE         SINCE  VERSION  REPORTED  UP         
24.0  1  0   0  0  4194304  6  active+clean    28m    955'6    978:61    [4,15,20]p4    [4,15,20]p4  
...
24.5  1  0   0  0  1611392   6  active+clean    28m    955'6    978:42    [18,4,9]p18    [18,4,9]p18  
24.6  2  0   0  0  4194304   6  active+clean    28m    955'6    978:51    [5,11,14]p5    [5,11,14]p5  
24.9  0  0   0  0        0   6  active+clean    28m    955'6    978:36    [10,6,3]p10    [10,6,3]p10  
...
24.1f 0  0   0  0        0   6  active+clean    28m    955'6    978:36   [16,11,0]p16   [16,11,0]p16 
```



##### 7. 저장 용량 분석

* 10MB 파일 1개가 4개 object 되고 이것을 PG에서 14MB되고 이것이 pool 기준으로는 29MiB로 표시되는 군

```
root@c01:/# rados df
POOL_NAME                      USED  OBJECTS  CLONES  COPIES  RD_OPS       RD  WR_OPS       WR
.mgr                        3.1 MiB        2       0       6     116  100 KiB     141  1.9 MiB
.rgw.root                    72 KiB        6       0      18       0      0 B       0      0 B
default.rgw.buckets.data     29 MiB        4       0      12       9  9.5 MiB       0      0 B
default.rgw.buckets.index       0 B       11       0      33     115  115 KiB      17    9 KiB
default.rgw.buckets.non-ec      0 B        0       0       0       0      0 B       0      0 B
default.rgw.control             0 B        8       0      24       0      0 B       0      0 B
default.rgw.log             408 KiB      209       0     627    4350  4.2 MiB    2892      0 B
default.rgw.meta             72 KiB        7       0      21      20   16 KiB      12    6 KiB
ec-4k2m-pool                    0 B        0       0       0       0      0 B       0      0 B
```



##### 8. 파일 쪼개진것 object 확인

```
# rados --pool default.rgw.buckets.data ls
e52d5020-4ee3-448f-b912-ac6356d4d5f4.105479.1__multipart_file1.2~gPUErBGcA0XFYXVXl-dKxuziKX-wVwx.1
e52d5020-4ee3-448f-b912-ac6356d4d5f4.105479.1__shadow_file1.2~gPUErBGcA0XFYXVXl-dKxuziKX-wVwx.1_1
e52d5020-4ee3-448f-b912-ac6356d4d5f4.105479.1_file1
e52d5020-4ee3-448f-b912-ac6356d4d5f4.105479.1__multipart_file1.2~gPUErBGcA0XFYXVXl-dKxuziKX-wVwx.2
```

##### 9.  osd 10 제거하면 pg는 어떻게 할당되는가?

* 기존 사용하던 PG 24.9가  ([10,6,3], p10) ==> up ([6,9,3], p6) 로 변경 되었다.  6과3은 유지되고 10이 9로 변경됨.

```
# ceph osd rm 10
# ceph osd  map  default.rgw.buckets.data  f1
osdmap e993 pool 'default.rgw.buckets.data' (24) object 'f1' -> pg 24.67b909e9 (24.9) -> up ([6,9,3], p6) acting ([6,9,3], p6)

```

* osd 디스크를 다시 붙이기 위해서 zap 실행해서 cleanzing : lv,vg,pv 정보 제거되면서 다시 osd 디스크가 붙는다. (다행인지는 모르겠지만 osd 번호는 그대로 가져간다.  )

```
# ceph orch device zap c06 /dev/sdd --force
zap successful for /dev/sdd on c06
```

##### 10. osd disk 를 dd 해버리면...

```
# ceph  osd  map  default.rgw.buckets.data  f1
osdmap e1005 pool 'default.rgw.buckets.data' (24) object 'f1' -> pg 24.67b909e9 (24.9) -> up ([10,6,3], p10) acting ([10,6,3], p10)
```

* 여기서 나온 osd 3개를 모두 dd로 밀어 버리고 hexedit를 통해서 정확히 밀어 버렸는지 확인한다.  
* osd 10, 6, 3 번 모두 밀어 버림.

```
# dd if=/dev/zero of=/dev/ceph-dd50d221-fdad-4049-a1f0-9591fc5878a1/osd-block-c441467b-5937-4b00-a44b-d2b37698af0e bs=1M
10237+0 records in
10236+0 records out
10733223936 bytes (11 GB, 10 GiB) copied, 21.7701 s, 493 MB/s
# apt install hexedit
# hexedit /dev/ceph-9d62c648-791f-4204-b117-fa074ecc77a5/osd-block-f4001470-3fc9-449d-8cbd-b42f7bdabea1
# disk 손상 
```

* 목록도 잘 보이고, 다운로드도 잘 됨 ?? 이상하네...

```
root@c01:~# aws s3  ls mybucket --endpoint-url http://192.168.105.51:80
2023-03-19 05:32:00   10000000 file1
root@c01:~# aws s3 cp s3://mybucket/file1 ~/f2 --endpoint-url http://192.168.105.51:80
download: s3://mybucket/file1 to ./f2
```

* 일단 여기 까지만 뭔가...  다른 osd에 저장되는 듯



### Ceph Erasure Coding pool 

https://devops.sdsdev.co.kr/confluence/pages/viewpage.action?pageId=656939555

##### 1. gateway 설치

```
# ceph orch apply rgw foo
# ceph orch ls
rgw.foo                    ?:80             2/2  2m ago     3m   count:2
```



##### 2. radosgw  상태  확인 : zone, zonegroup, placement

```
root@c01:/# radosgw-admin zonegroup  list
{
    "default_info": "c83359d5-dc52-462d-b1d6-1032a2315ef5",
    "zonegroups": ["default"]
}

root@c01:/# radosgw-admin zone list
{
    "default_info": "b100b5a2-6e09-44e0-a9be-a0174deaf432",
    "zones": ["default" ]
}

root@c01:/# radosgw-admin zone placement  list
[
    {
        "key": "default-placement",
        "val": {
            "index_pool": "default.rgw.buckets.index",
            "storage_classes": {
                "STANDARD": {"data_pool": "default.rgw.buckets.data"}
            },
            "data_extra_pool": "default.rgw.buckets.non-ec",
            "index_type": 0
        }
    }
]
```

* zonegroup, zone, placement

```
# radosgw-admin zonegroup get
# radosgw-admin zone get
# radosgw-admin zone placement get --placement-id default-placement
```



##### 3. EC profile 생성, Pool 생성

```
# ceph osd  erasure-code-profile ls
# ceph osd erasure-code-profile get default
# ceph config set global mon_max_pg_per_osd 500 
```

* mypool => PG status: 64 creating + incomplete 상태.... 

```
# ceph osd erasure-code-profile set myprofile k=19 m=5 plugin=jerasure crush-failure-domain=host
# ceph osd erasure-code-profile get myprofile
# ceph osd pool create mypool 64 64 erasure myprofile
```

* ecpool =>PG status : 12 active+clean 상태

```
# ceph osd erasure-code-profile set myecprof k=3 m=2 crush-failure-domain=host
# ceph osd pool create ecpool 12 12 erasure myecprof
# ceph osd pool application enable ecpool rgw
```

* ecpool195 -> PG status : 32 Unknown 상태 

```
# ceph osd erasure-code-profile set ecprofile k=19 m=5 crush-failure-domain=rack
# ceph osd pool create ecpool195 32 32 erasure ecprofile
# ceph osd pool application enable ecpool195 rgw
```

* ec-4k2m-pool

```
# ceph osd erasure-code-profile ls
# ceph osd erasure-code-profile set ecprofile-4k2m k=4 m=2 plugin=jerasure crush-failure-domain=host
# ceph osd  erasure-code-profile get ecprofile-4k2m
# ceph osd pool create ec-4k2m-pool  64 64 erasure ecprofile-4k2m
#  ceph osd lspools
4 ec-4k2m-pool
```



##### 4. Placement Target 등록

```
# radosgw-admin zonegroup placement add \
    --rgw-zonegroup default \
    --placement-id myplacement
    
# radosgw-admin zone placement add  \
   --rgw-zone default  \
   --placement-id myplacement \
   --data-pool ec-4k2m-pool \
   --index-pool default.rgw.buckets.index \
   --data-extra-pool default.rgw.temporary.non-ec \
   --compression lz4
   
# radosgw-admin zonegroup get
# radosgw-admin zone placement list
```

##### 5. Placement 사용하는 S3 user 생성

```
# radosgw-admin user create \
       --uid s3-user \
       --display-name s3-user.fullname \
       --placement-id myplacement \
       --storage-class STANDARD \
       --tags GLACIER
```

* 그냥 web에서 제공하는 방식으로 user 생성하면  오류 발생한다

* ```
  root@c01:~# aws s3 cp ./f2 s3://mybucket/f2 --endpoint-url http://192.168.105.51:80
  upload failed: ./f2 to s3://mybucket/f2 An error occurred (AccessDenied) when calling the UploadPart operation: Unknown
  ```

```
#  radosgw-admin user list
[
    "s3-user",
    "dashboard",
    "myuser"
]
# radosgw-admin user info --uid  myuser
```





##### 5. rgw 서비스를 재기동

```
# ceph orch ps  ==> daemon 이름 확인
# ceph orch daemon restart rgw.foo.c01.mrawxe
# ceph orch daemon restart rgw.foo.c02.udcxyw
Scheduled to restart rgw.foo.c01.mrawxe on host 'c01'
```



##### 6. Test

```
# dd if=/dev/urandom of=f1 bs=1MB count=5000
5000000000 bytes (5.0 GB, 4.7 GiB) copied, 27.3202 s, 183 MB/s
# aws s3 cp ./f1 s3://mybucket/f1 --endpoint-url http://192.168.105.51:80

# raods lspools
# rados df 
# rados -p ec-4k2m-pool put f2 ./f2
# rados -p ec-4k2m-pool put f1 ./f1   
# rados -p ec-4k2m-pool ls

# ceph pg ls-by-pool ec-4k2m-pool

# ceph osd map ec-4k2m-pool f1
osdmap e288 pool 'ec-4k2m-pool' (9) object 'f1' -> pg 9.67b909e9 (9.29) -> up ([18,14,21,6,20,16], p18) acting ([18,14,21,6,20,16], p18)
root@c01:/rootfs/root# ceph osd map ec-4k2m-pool f2
osdmap e288 pool 'ec-4k2m-pool' (9) object 'f2' -> pg 9.54e93ef4 (9.34) -> up ([3,5,6,7,15,2], p3) acting ([3,5,6,7,15,2], p3)

```

* rados를 통해서 object가 4MB 단위로 쪼개 지지 않고 등록되고, AWS S3 같은 api를 통해서 입력하면 4MB 단위로 객체가 쪼개 져서 등록 되는 것을 확인했다. 



* EC 코드 pool 에 upload  할때 오류 발생한다.  : CreateMultiparUpload operation.

  ```
  # aws s3 cp ./f3 s3://mybucket/f3 --endpoint-url http://192.168.105.51:80
  upload failed: ./f3 to s3://mybucket/f3 An error occurred (AccessDenied) when calling the CreateMultipartUpload operation: Unknown
  
  # aws s3 cp ./f4 s3://mybucket/f4 --endpoint-url http://192.168.105.51:80
  upload failed: ./f4 to s3://mybucket/f4 An error occurred (AccessDenied) when calling the PutObject operation: Unknown
  ```

* 내일은 그냥 복제 pool에다 upload를 해봐야 겠다.  

* 로그를 보면서 하려면  뭔가 log 분석 방법이 필요하다.

* 



### Ceph placement/Pool 추가

https://docs.ceph.com/en/latest/radosgw/placement/#adding-a-placement-target

##### 1. zonegroup placement add 



##### 1. erasure-code profile 이용한 Pool 생성 

```
# ceph osd erasure-code-profile ls
# ceph osd erasure-code-profile get default
# ceph osd erasure-code-profile set ecprofile-4k2m k=4 m=2 plugin=jerasure crush-failure-domain=host
# ceph osd erasure-code-profile set ecprofile-19k5m k=19 m=5 plugin=jerasure crush-failure-domain=host

# ceph osd erasure-code-profile ls
# ceph osd  erasure-code-profile get ecprofile-19k5m
# ceph osd  erasure-code-profile get ecprofile-4k2m

# ceph osd erasure-code-profile get ecprofile-4k2m
# ceph osd pool create ec-4k2m-pool  64 64 erasure ecprofile-4k2m
# ceph osd pool create ec-19k5m-pool 64 64 erasure ecprofile-19k5m

#  ceph osd lspools
3 ec-19k5m-pool
4 ec-4k2m-pool
# ceph osd erasure-code-profile rm  ecprofile-4k2m
```



##### 2. date pool  :  erasure

```
# ceph osd pool create kr-west1.rgw.archive.data erasure default --autoscale-mode=on
# ceph osd pool application enable kr-west1.rgw.archive.data rgw
//# ceph osd pool set kr-west1.rgw.archive.data compression_mode force
```

##### 3. index pool 

```
# ceph osd pool create kr-west1.rgw.archive.index replicated --autoscale-mode=on
# ceph osd pool application enable kr-west1.rgw.archive.index rgw
# ceph osd pool create kr-west1.rgw.archive.index.non-ec replicated --autoscale-mode=on
# ceph osd pool application enable kr-west1.rgw.archive.index.non-ec rgw
```

##### 4. get pool list

```
# ceph osd pool ls
# ceph osd pool ls detail
# ceph osd pool get  kr-west1.rgw.archive.data all
```

##### 5. default zonegroup, zone에  archive-placement  생성한 pool 등록

```
# radosgw-admin zonegroup placement add \
            --rgw-zonegroup default \
            --placement-id archive-placement \
            --storage-class GLACIER
            
# radosgw-admin zone placement add \
      --rgw-zone default \
      --placement-id archive-placement \
      --storage-class GLACIER \
      --data-pool kr-west1.rgw.archive.data \
      --index-pool kr-west1.rgw.archive.index \
      --data-extra-pool kr-west1.rgw.archive.index.non-ec \
      --compression lz4
      
# radosgw-admin zone placement list
```

##### 6.  zonegroup, zone, placement  확인 

```
# radosgw-admin zonegroup get
# radosgw-admin zone get
# radosgw-admin zone placement get --placement-id default-placement
```

##### 7. radosgw-admin  user

````
# radosgw-admin user create \
       --uid s3-user-1 \
       --display-name s3-user-1.fullname \
       --placement-id archive-placement \
       --storage-class GLACIER \
       --tags GLACIER
#  radosgw-admin user info --uid s3-user-1
````

##### 9. rgw 서비스를 재기동

```
# ceph orch ps  ==> daemon 이름 확인
# ceph orch daemon restart rgw.foo.c01.mrawxe
# ceph orch daemon restart rgw.foo.c02.udcxyw
Scheduled to restart rgw.foo.c01.mrawxe on host 'c01'
```



#### rgw realm으로 생성해보기.

```
# radosgw-admin realm create --rgw-realm=<realm-name> --default
# radosgw-admin zonegroup create --rgw-zonegroup=<zonegroup-name>  --master --default
# radosgw-admin zone create --rgw-zonegroup=<zonegroup-name> --rgw-zone=<zone-name> --master --default
# radosgw-admin period update --rgw-realm=<realm-name> --commit
```





### EC 19+5 placement 생성

##### 1. EC profile 생성

```
# ceph osd erasure-code-profile ls
# ceph osd erasure-code-profile get default
# ceph osd erasure-code-profile set ecprofile-4k2m k=4 m=2 plugin=jerasure crush-failure-domain=host
# ceph osd erasure-code-profile set ecprofile-19k5m k=19 m=5 plugin=jerasure crush-failure-domain=host

# ceph osd erasure-code-profile ls
# ceph osd erasure-code-profile get ecprofile-19k5m
# ceph osd erasure-code-profile get ecprofile-4k2m

# ceph osd erasure-code-profile get ecprofile-4k2m
# ceph osd pool create ec-4k2m-pool  64 64 erasure ecprofile-4k2m
# ceph osd pool create ec-19k5m-pool 64 64 erasure ecprofile-19k5m

#  ceph osd lspools
3 ec-19k5m-pool
4 ec-4k2m-pool

# 삭제
# ceph osd erasure-code-profile  rm  ecprofile-4k2m
```



##### 2. crushmap 추출하여 rule 추가

```
# ceph osd getcrushmap -o map.bin
# crushtool -d map.bin -o map.txt
```

##### 3. crushmap rule 추가

```
# vi map.txt
...
# rule 추가
rule ec-rule-19k5m {
        id 8
        type erasure
        step set_chooseleaf_tries 5
        step set_choose_tries 100
        step take default class hdd
        step choose indep 0 type osd
        step emit
}
```

##### 4. crushmap rule 적용

```
# crushtool -c map.txt -o map.bin
# crushtool -i map.bin --test --show-mappings --rule=5 --num-rep 48
# ceph osd setcrushmap -i map.bin

# ceph osd crush rule ls
replicated_rule
ec-4k2m-pool-01
ec-19k5m-pool
ec-4k2m-pool
erasure-code
ec-rule-19k5m
```

##### 5. osd eraure-code profile 

```
# ceph osd erasure-code-profile set ec-rule-19k5m-profile erasure k=19 m=5 plugin=jerasure crush-failure-domain=host replicated_rule
```

##### 6. erasure pool 생성

```
# ceph osd pool create ec-19k5m-pool erasure ec-rule-19k5m-profile ec-rule-19k5m
```

```
# ceph osd erasure-code-profile ls
# ceph osd erasure-code-profile get default
# ceph osd erasure-code-profile set ecprofile-19k5m k=19 m=5 plugin=jerasure crush-failure-domain=host
# ceph osd erasure-code-profile get ecprofile-19k5m

# ceph osd lspools
# ceph osd pool delete ec-19k5m-pool ec-19k5m-pool --yes-i-really-really-mean-it
```



##### 7. zone group에 placement 추가

```
# radosgw-admin zonegroup placement add \
    --rgw-zonegroup default \
    --placement-id myplacement
    
# radosgw-admin zone placement add  \
   --rgw-zone default  \
   --placement-id myplacement \
   --data-pool ec-19k5m-pool \
   --index-pool default.rgw.buckets.index \
   --data-extra-pool default.rgw.temporary.non-ec \
   --compression lz4
   
# radosgw-admin zonegroup get
# radosgw-admin zone placement list
```



##### 8. rgw 서비스를 재기동

```
# ceph orch ps  ==> daemon 이름 확인
# ceph orch daemon restart rgw.foo.c01.mrawxe
# ceph orch daemon restart rgw.foo.c02.udcxyw
```







#### 문제발생

#### 현상 : bucket 만드는 화면에서 오류 발생

##### 1. /var/log/syslog 현상

```
Mar 20 06:21:51 c01 bash[17810]: debug 2023-03-20T06:21:51.408+0000 7f69c63ee700  0 [dashboard INFO rgw_client] Found RGW daemon with configuration: host=192.168.105.53, port=9443, ssl=False
Mar 20 06:21:51 c01 bash[17810]: debug 2023-03-20T06:21:51.412+0000 7f69c3be9700  0 [dashboard INFO rgw_client] Found RGW daemon with configuration: host=192.168.105.53, port=9443, ssl=False
Mar 20 06:21:51 c01 bash[17810]: debug 2023-03-20T06:21:51.412+0000 7f69c3be9700  0 [dashboard ERROR exception] Internal Server Error
Mar 20 06:21:51 c01 bash[17810]: Traceback (most recent call last):
Mar 20 06:21:51 c01 bash[17810]:   File "/usr/share/ceph/mgr/dashboard/services/exception.py", line 47, in dashboard_exception_handler
Mar 20 06:21:51 c01 bash[17810]:     return handler(*args, **kwargs)
Mar 20 06:21:51 c01 bash[17810]:   File "/lib/python3.6/site-packages/cherrypy/_cpdispatch.py", line 54, in __call__
Mar 20 06:21:51 c01 bash[17810]:     return self.callable(*self.args, **self.kwargs)
Mar 20 06:21:51 c01 bash[17810]:   File "/usr/share/ceph/mgr/dashboard/controllers/_base_controller.py", line 258, in inner
Mar 20 06:21:51 c01 bash[17810]:     ret = func(*args, **kwargs)
Mar 20 06:21:51 c01 bash[17810]:   File "/usr/share/ceph/mgr/dashboard/controllers/_rest_controller.py", line 191, in wrapper
Mar 20 06:21:51 c01 bash[17810]:     return func(*vpath, **params)
Mar 20 06:21:51 c01 bash[17810]:   File "/usr/share/ceph/mgr/dashboard/controllers/rgw.py", line 159, in list
Mar 20 06:21:51 c01 bash[17810]:     return RgwClient.admin_instance(daemon_name=daemon_name).get_placement_targets()
Mar 20 06:21:51 c01 bash[17810]:   File "/usr/share/ceph/mgr/dashboard/services/rgw_client.py", line 571, in get_placement_targets
Mar 20 06:21:51 c01 bash[17810]:     'data_pool': placement_pool['val']['storage_classes']['STANDARD']['data_pool']
Mar 20 06:21:51 c01 bash[17810]: KeyError: 'data_pool'
Mar 20 06:21:51 c01 bash[17810]: debug 2023-03-20T06:21:51.412+0000 7f69c3be9700  0 [dashboard ERROR request] [::ffff:192.168.105.1:34356] [GET] [500] [0.006s] [admin] [513.0B] /api/rgw/si
```

* [이강규] 2023-03-20 15:32  여기 코드가 문제라고 함. 
  https://github.com/ceph/ceph/blob/main/src/pybind/mgr/dashboard/services/rgw_client.py

```
    def get_placement_targets(self):  # type: () -> dict
        zone = self._get_daemon_zone_info()
        placement_targets = []  # type: List[Dict]
        for placement_pool in zone['placement_pools']:
            placement_targets.append(
                {
                    'name': placement_pool['key'],
                    'data_pool': placement_pool['val']['storage_classes']['STANDARD']['data_pool']
                }
            )

        return {'zonegroup': self.daemon.zonegroup_name,
                'placement_targets': placement_targets}
```



* 이강규 프로: `/usr/share/ceph/mgr/dashboard/services/rgw_client.py` 여기 파일에서 STANDARD 클래스가 없어서 그런다고 함. 

```
# radosgw-admin zone list
# radosgw-admin zone get 
# radosgw-admin zone placement list
# radosgw-admin zone placement rm --placement-id archive-placement
# radosgw-admin zone placement rm --placement-id default-placement
# radosgw-admin zone placement rm --placement-id myplacement
```



```
radosgw-admin zonegroup delete --rgw-zonegroup=default --rgw-zone=default
radosgw-admin period update --commit
radosgw-admin zone delete --rgw-zone=default
radosgw-admin period update --commit
radosgw-admin zonegroup delete --rgw-zonegroup=default
radosgw-admin period update --commit
```



```
ceph osd pool rm default.rgw.control default.rgw.control --yes-i-really-really-mean-it
ceph osd pool rm default.rgw.data.root default.rgw.data.root --yes-i-really-really-mean-it
ceph osd pool rm default.rgw.gc default.rgw.gc --yes-i-really-really-mean-it
ceph osd pool rm default.rgw.log default.rgw.log --yes-i-really-really-mean-it
ceph osd pool rm default.rgw.users.uid default.rgw.users.uid --yes-i-really-really-mean-it
```





```
# radosgw-admin zonegroup placement add \
       --rgw-zonegroup default \
       --placement-id archive-placement \
       --storage-class STANDARD
	   
# radosgw-admin zone placement add \
      --rgw-zone default \
      --placement-id archive-placement \
      --storage-class STANDARD \
      --data-pool kr-west1.rgw.archive.data \
      --index-pool kr-west1.rgw.archive.index \
      --data-extra-pool kr-west1.rgw.archive.index.non-ec \
      --compression lz4		   

# radosgw-admin period update --commit

이렇게 실행 한번 해보셔요.
```

* 이 단계에서 오류 발생 

```
root@c01:/# radosgw-admin period update --commit
Error initializing realm (2) No such file or directory

```





####  PG osd dd  밀기

##### 1. 현상

```
# ceph health detail
HEALTH_WARN 3 failed cephadm daemon(s); 4 daemons have recently crashed
[WRN] CEPHADM_FAILED_DAEMON: 3 failed cephadm daemon(s)
    daemon osd.3 on c04 is in error state
    daemon osd.5 on c05 is in error state
    daemon osd.8 on c06 is in error state
[WRN] RECENT_CRASH: 4 daemons have recently crashed
    mon.c01 crashed on host c01 at 2023-03-18T09:05:34.288781Z
    mon.c01 crashed on host c01 at 2023-03-18T09:07:41.259705Z
    osd.8 crashed on host c06 at 2023-03-19T08:34:23.412027Z
    osd.5 crashed on host c05 at 2023-03-19T08:34:22.210687Z
```



#### Test Object repair

#### Host: remove, add 

##### 1. drain all daemon from host

```
# ceph orch host drain c06
# ceph orch osd rm status 
# ceph orch ps c06
# ceph orch host rm c06
Removed  host 'c06'
```

##### 2. offline host remove  force

```
# ceph orch host rm c06 --offline --force
```

##### 3. host maintenance mode

```
# ceph orch host maintenance enter c05
# ceph orch host maintenance exit c05
```

##### 4. add host

```
# ceph orch host add c06  192.168.105.56 --labels _admin
```



#### OSD : remove add

##### 1. list device 

```
# ceph orch device ls
# ceph orch apply osd --all-available-devices
# ceph orch daemon add osd c05:/dev/sdb
```

* 현상

```
root@c01:/# ceph orch daemon add osd c05:/dev/sdb
Created no osd(s) on host c05; already created?
```

##### 2. osd 제거

* 클러스터에서 모든 PG 제거
* 클러스터에서 PG 없는 OSD 제거

```
# ceph orch osd rm 0
# ceph orch osd rm status
No OSD remove/replace operations reported
```

* lvremove, vgremove, pgremove 하고 available device로 되면 다시 osd로 등록 할 수 있다. 

##### 3.제거 중지

```
# cep orch osd rm stop 0.3
```

##### 4. lv, vg, pg 제거

```
# lvremove
# vgremove
# pvremove
```

* 이렇게 하면 해당 device가 다시 available 상태로 된다. 

##### 5. osd 등록

```
# ceph orch apply osd --all-available-devices
```





#### ceph node os 재설치

##### 설치전 백업.

```
root@c06:/etc/ceph# cat ceph.conf
# minimal ceph.conf for 1373ff90-46ab-678d-2467-8030180fe5c0
[global]
        fsid = 1373ff90-46ab-678d-2467-8030180fe5c0
        mon_host = [v2:192.168.105.51:3300/0,v1:192.168.105.51:6789/0] [v2:192.168.105.52:3300/0,v1:192.168.105.52:6789/0] [v2:192.168.105.53:3300/0,v1:192.168.105.53:6789/0] [v2:192.168.105.54:3300/0,v1:192.168.105.54:6789/0] [v2:192.168.105.55:3300/0,v1:192.168.105.55:6789/0]

root@c06:/etc/ceph# cat ceph.client.admin.keyring
[client.admin]
        key = AQARgRRkPCOXHhAAMcXUf2ue3GdWQaiJWWcGUA==
        caps mds = "allow *"
        caps mgr = "allow *"
        caps mon = "allow *"
        caps osd = "allow *"
```

* 원래 위치



### Ceph Map

#### 1.  Monitor map

* 클러스터 fsid, 위치, 이름 주소 및 각 모니터의 포트를 포함합니다.
* 현재 에포크, 맵이 생성 된 시간 및 마지막으로 변경된 시간을 나타냅니다

```
[ceph: root@c1 /]# ceph mon dump
dumped monmap epoch 5
epoch 5
fsid fdc6ff5a-4fae-11eb-9165-000c2938f508
last_changed 2021-01-06T00:29:46.083864+0000
created 2021-01-05T23:37:38.199123+0000
min_mon_release 15 (octopus)
0: [v2:192.168.56.21:3300/0,v1:192.168.56.21:6789/0] mon.c1
1: [v2:192.168.56.22:3300/0,v1:192.168.56.22:6789/0] mon.c2
2: [v2:192.168.56.23:3300/0,v1:192.168.56.23:6789/0] mon.c3
3: [v2:192.168.56.25:3300/0,v1:192.168.56.25:6789/0] mon.c5
4: [v2:192.168.56.24:3300/0,v1:192.168.56.24:6789/0] mon.c4
```



#### 2. OSD map

* fsid 맵이 생성되고 마지막으로 수정 된 클러스터 , 풀 목록, 복제본 크기, PG 번호, OSD 목록 및 해당 상태 (예 : up, in)가 포함됩니다

```
[ceph: root@c1 /]# ceph osd dump
epoch 10279
fsid fdc6ff5a-4fae-11eb-9165-000c2938f508
created 2021-01-05T23:37:39.119640+0000
modified 2023-03-13T12:13:24.488369+0000
flags sortbitwise,recovery_deletes,purged_snapdirs,pglog_hardlimit
crush_version 17
full_ratio 0.95
backfillfull_ratio 0.9
nearfull_ratio 0.85
require_min_compat_client jewel
min_compat_client jewel
require_osd_release octopus
pool 1 'device_health_metrics' replicated size 3 min_size 2 crush_rule 0 object_hash rjenkins pg_num 1 pgp_num 1 autoscale_mode on last_change 14 flags hashpspool stripe_width 0 pg_num_min 1 application mgr_devicehealth
pool 4 'myfs' replicated size 3 min_size 2 crush_rule 0 object_hash rjenkins pg_num 32 pgp_num 32 autoscale_mode on last_change 72 flags hashpspool stripe_width 0 pg_autoscale_bias 4 pg_num_min 16 recovery_priority 5 application cephfs
pool 5 'myfsmds' replicated size 3 min_size 2 crush_rule 0 object_hash rjenkins pg_num 32 pgp_num 32 autoscale_mode on last_change 72 flags hashpspool stripe_width 0 application cephfs

max_osd 8
osd.0 up   in  weight 1 up_from 10187 up_thru 10256 down_at 10179 last_clean_interval [9166,10178) [v2:192.168.56.22:6802/3312945622,v1:192.168.56.22:6803/3312945622] [v2:192.168.56.22:6804/3312945622,v1:192.168.56.22:6805/3312945622] exists,up 68bc5d6f-d2e0-4b7e-9f6a-0df8f1084147
osd.1 up   in  weight 1 up_from 10185 up_thru 10234 down_at 10184 last_clean_interval [9159,10181) [v2:192.168.56.23:6800/498223093,v1:192.168.56.23:6801/498223093] [v2:192.168.56.23:6802/498223093,v1:192.168.56.23:6803/498223093] exists,up 1018123a-93d4-482c-8f0d-a4ef007d4aae
osd.2 up   in  weight 1 up_from 10185 up_thru 10257 down_at 10184 last_clean_interval [9168,10181) [v2:192.168.56.24:6808/2335915194,v1:192.168.56.24:6809/2335915194] [v2:192.168.56.24:6810/2335915194,v1:192.168.56.24:6811/2335915194] exists,up dc88e368-0ad9-4363-bfe7-eabfb5f8d133
...
```



#### 3. PG map

* ceph pg dump -h 
* PG 버전, 타임 스탬프, 마지막 OSD 맵 에포크, 전체 비율 및 PG ID, Up Set , Acting Set , PG 상태 (예 : PG 상태) 와 같은 각 배치 그룹의 세부 정보가 포함됩니다. , )  각 풀의 데이터 사용량 통계active + clean

```
[ceph: root@c1 /]# ceph pg dump summary
dumped summary
version 129636
stamp 2023-03-13T12:16:39.591591+0000
last_osdmap_epoch 0
last_pg_scan 0
PG_STAT  OBJECTS  MISSING_ON_PRIMARY  DEGRADED  MISPLACED  UNFOUND  BYTES        OMAP_BYTES*  OMAP_KEYS*  sum        25009                   0         0          0        0  55473487490      1279458       98929  
OSD_STAT  USED     AVAIL    USED_RAW  TOTAL  
sum       155 GiB  189 GiB   161 GiB  350 GiB


[ceph: root@c1 /]# ceph pg dump pools
dumped pools
POOLID  OBJECTS    DEGRADED  MISPLACED  UNFOUND  BYTES        OMAP_BYTES*  OMAP_KEYS*  LOG     DISK_LOG
1277          0    0         0          0        0            0            0           0       0         0
32        12038    0         0          0        0  50237747772         1175         103  230054    230054
1272         38    0         0          0        0       103610            0           0     325       325
31            0    0         0          0        0            0            0           0       0         0
30           22    0         0          0        0         6939         4710          10     154       154
1278          6    0         0          0        0          231          214          12   34618     34618
20           29    0         0          0        0           19         3274         217     641       641
17        11261    0         0          0        0   1040892215      1200004       98492   44878     44878
10          286    0         0          0        0            0         8041          35   10332     10332
9            70    0         0          0        0        13996         4794          26     728       728
4            23    0         0          0        0       405364         4710          10     470       470
1             1    0         0          0        0            0        52536          24      28        28
8             8    0         0          0        0            0            0           0    9952      9952
5          1002    0         0          0        0   4194306003            0           0    9702      9702
6            16    0         0          0        0         3713            0           0    1281      1281
7           209    0         0          0        0         7628            0           0   39325     39325


[ceph: root@c1 /]# ceph pg dump osds 
dumped osds
OSD_STAT  USED     AVAIL    USED_RAW  TOTAL    HB_PEERS     PG_SUM  PRIMARY_PG_SUM
6          25 GiB   24 GiB    26 GiB   50 GiB  [0,1,2,3,4]     287              81
1          28 GiB   21 GiB    29 GiB   50 GiB  [0,2,3,4,6]     253              61
0          30 GiB   19 GiB    31 GiB   50 GiB  [1,2,3,4,6]     288              85
2          18 GiB   31 GiB    19 GiB   50 GiB  [0,1,3,4,6]     166              83
3          25 GiB   24 GiB    26 GiB   50 GiB  [0,1,2,4,6]     281              68
4          30 GiB   69 GiB    31 GiB  100 GiB  [0,1,2,3,6]     312             151
5             0 B      0 B       0 B      0 B           []       0               0
sum       155 GiB  189 GiB   161 GiB  350 GiB                                     

```



#### 4. CRUSH map

```
# ceph osd getcrushmap -o {comp-crushmap-filename}
# crushtool -d {comp-crushmap-filename} -o {decomp-crushmap-filename} cat
```



```
# begin crush map
tunable choose_local_tries 0
tunable choose_local_fallback_tries 0
tunable choose_total_tries 50
tunable chooseleaf_descend_once 1
tunable chooseleaf_vary_r 1
tunable chooseleaf_stable 1
tunable straw_calc_version 1
tunable allowed_bucket_algs 54

# devices
device 0 osd.0 class ssd
device 1 osd.1 class ssd
device 2 osd.2 class ssd
device 3 osd.3 class ssd
device 4 osd.4 class ssd
device 5 osd.5 class ssd
device 6 osd.6 class ssd

# types
type 0 osd
type 1 host
type 2 chassis
type 3 rack
type 4 row
type 5 pdu
type 6 pod
type 7 room
type 8 datacenter
type 9 zone
type 10 region
type 11 root

# buckets  <<<<<====== bucket 
host c2 {
	id -3		# do not change unnecessarily
	id -4 class ssd		# do not change unnecessarily
	# weight 0.049
	alg straw2
	hash 0	# rjenkins1
	item osd.0 weight 0.049
}
host c3 {
	id -5		# do not change unnecessarily
	id -6 class ssd		# do not change unnecessarily
	# weight 0.049
	alg straw2
	hash 0	# rjenkins1
	item osd.1 weight 0.049
}
host c5 {
	id -7		# do not change unnecessarily
	id -8 class ssd		# do not change unnecessarily
	# weight 0.098
	alg straw2
	hash 0	# rjenkins1
	item osd.3 weight 0.049
	item osd.5 weight 0.049
}
host c4 {
	id -9		# do not change unnecessarily
	id -10 class ssd		# do not change unnecessarily
	# weight 0.146
	alg straw2
	hash 0	# rjenkins1
	item osd.2 weight 0.049
	item osd.4 weight 0.098
}
host c1 {
	id -11		# do not change unnecessarily
	id -12 class ssd		# do not change unnecessarily
	# weight 0.049
	alg straw2
	hash 0	# rjenkins1
	item osd.6 weight 0.049
}
root default {
	id -1		# do not change unnecessarily
	id -2 class ssd		# do not change unnecessarily
	# weight 0.390
	alg straw2
	hash 0	# rjenkins1
	item c2 weight 0.049
	item c3 weight 0.049
	item c5 weight 0.098
	item c4 weight 0.146
	item c1 weight 0.049
}

# rules
rule replicated_rule {
	id 0
	type replicated
	min_size 1
	max_size 10
	step take default
	step chooseleaf firstn 0 type host
	step emit
}

# end crush map
```

##### osd tree

```
[ceph: root@c1 /]# ceph osd tree
ID   CLASS  WEIGHT   TYPE NAME      STATUS  REWEIGHT  PRI-AFF
 -1         0.39047  root default                            
-11         0.04880      host c1                             
  6    ssd  0.04880          osd.6      up   1.00000  1.00000
 -3         0.04880      host c2                             
  0    ssd  0.04880          osd.0      up   1.00000  1.00000
 -5         0.04880      host c3                             
  1    ssd  0.04880          osd.1      up   1.00000  1.00000
 -9         0.14648      host c4                             
  2    ssd  0.04880          osd.2      up   1.00000  1.00000
  4    ssd  0.09769          osd.4      up   1.00000  1.00000
 -7         0.09760      host c5                             
  3    ssd  0.04880          osd.3      up   1.00000  1.00000
  5    ssd  0.04880          osd.5    down         0  1.00000
```

* CRUSH는 CRUSH Map이라는 Storage Topology를 이용한다. 
* CRUSH Map은 Bucket이라는 논리적 단위의 계층으로 구성된다. Bucket은 root, region, datacentor, room, pod, pdu, row, rack, chassis, host, osd 11가지 type으로 구성되어 있다. CRUSH Map의 Leaf는 반드시 osd bucket이어야 한다.
* 각 Bucket은 Weight값을 가지고 있는데 Weight는 각 Bucket이 갖고있는 Object의 비율을 나타낸다. 만약 Bucket A의 Weight가 100이고 Bucket B의 Weight가 200이라면 Bucket B가 Bucket A보다 2배많은 Object를 갖는다는걸 의미한다. 따라서 일반적으로 osd Bucket Type의 Weight값은 OSD가 관리하는 Disk의 용량에 비례하여 설정한다. 나머지 Bucket Type의 weight는 모든 하위 Bucket의 Weight의 합이다. 
* CRUSH는 CRUSH Map의 root Bucket부터 시작하여 하위 Bucket을 Replica 개수 만큼 선택하고, 선택한 Bucket에서 동일한 작업을 반복하여 Leaf에 있는 osd Bucket을 찾는다. Object의 Replica 개수는 Bucket Type에 설정한 Replica에 따라 정해진다. Rack Bucket Type에는 3개의 Replica를 설정하고 Row Bucket Type에는 2개의 Replica를 설정하였다면, CRUSH는 3개의 Rack Bucket을 선택하고 선택한 Rack Bucket의 하위 Bucket인 Row Bucket을 각 Rack Bucket당 2개씩 선택하기 때문에 Object의 Replica는 6이 된다.
*  하위 Bucket을 선택하는 기준은 각 Bucket Type에 설정한 Bucket 알고리즘에 따라 결정된다.
* Bucket 알고리즘에는 Uniform, List, Tree, Straw, Straw2가 있다.



#### 5. MDS map

```
[ceph: root@c1 /]# ceph fs dump
dumped fsmap epoch 721
e721
enable_multiple, ever_enabled_multiple: 1,0
compat: compat={},rocompat={},incompat={1=base v0.20,2=client writeable ranges,3=default file layouts on dirs,4=dir inode in separate object,5=mds uses versioned encoding,6=dirfrag is stored in omap,8=no anchor table,9=file layout v2,10=snaprealm v2}
legacy client fscid: 2
 
Filesystem 'myfs' (2)
fs_name	myfs
epoch	719
flags	12
created	2021-01-06T02:31:20.479051+0000
modified	2023-03-10T12:13:37.147064+0000
tableserver	0
root	0
session_timeout	60
session_autoclose	300
max_file_size	1099511627776
min_compat_client	0 (unknown)
last_failure	0
last_failure_osd_epoch	10183
compat	compat={},rocompat={},incompat={1=base v0.20,2=client writeable ranges,3=default file layouts on dirs,4=dir inode in separate object,5=mds uses versioned encoding,6=dirfrag is stored in omap,8=no anchor table,9=file layout v2,10=snaprealm v2}
max_mds	1
in	0
up	{0=86504139}
failed	
damaged	
stopped	
data_pools	[5]
metadata_pool	4
inline_data	disabled
balancer	
standby_count_wanted	1
[mds.myfs.c3.itxgwq{0:86504139} state up:active seq 12 join_fscid=2 addr [v2:192.168.56.23:6808/1674461965,v1:192.168.56.23:6809/1674461965]]
 
 
Filesystem 'cephfs' (3)
fs_name	cephfs
epoch	716
flags	12
created	2021-01-08T02:51:41.916579+0000
modified	2023-03-10T12:13:29.084277+0000
tableserver	0
root	0
session_timeout	60
session_autoclose	300
max_file_size	1099511627776
min_compat_client	0 (unknown)
last_failure	0
last_failure_osd_epoch	10188
compat	compat={},rocompat={},incompat={1=base v0.20,2=client writeable ranges,3=default file layouts on dirs,4=dir inode in separate object,5=mds uses versioned encoding,6=dirfrag is stored in omap,8=no anchor table,9=file layout v2,10=snaprealm v2}
max_mds	1
in	0
up	{0=86504109}
failed	
damaged	
stopped	
data_pools	[31]
metadata_pool	30
inline_data	disabled
balancer	
standby_count_wanted	1
[mds.cephfs.c2.mhueee{0:86504109} state up:active seq 9 join_fscid=3 addr [v2:192.168.56.22:6800/1601949077,v1:192.168.56.22:6801/1601949077]]
 
 
Standby daemons:
 
[mds.myfs.c5.jukhiq{-1:86504172} state up:standby seq 4 join_fscid=2 addr [v2:192.168.56.25:6800/1272481591,v1:192.168.56.25:6801/1272481591]]
[mds.cephfs.c1.lswbph{-1:86504221} state up:standby seq 3 join_fscid=3 addr [v2:192.168.56.21:6800/496946002,v1:192.168.56.21:6801/496946002]]
```

