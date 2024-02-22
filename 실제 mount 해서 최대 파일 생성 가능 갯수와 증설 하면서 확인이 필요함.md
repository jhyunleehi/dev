## Netapp Inode Test

```
ontap-select-cluster::*> volume show  -volume V1  -fields  size,files,files-used,files-maximum-possible;                                                                                         vserver   volume size files files-used files-maximum-possible                                                                                                                                    --------- ------ ---- ----- ---------- ---------------------- 
SAM_ISCSI V1     20MB 566   96         4855               
```





```
ontap-select-cluster::*> volume create -volume V1 -state online -policy default -unix-permissions ---rwxr-xr-x -junction-active true -vsroot false -type RW -snapshot-policy default -foreground true -tiering-policy none -analytics-state off -activity-tracking-state off -anti-ransomware-state disabled -granular-data disabled -snapshot-locking-enabled false -is-preserve-unlink-enabled false -vserver SAM_ISCSI -aggregate N1_aggr1 
[Job 108] Job succeeded: Successful
```



## 실제 mount 해서 최대 파일 생성 가능 갯수 테스트 



```
FlexGroup::> volume show -vserver svm1
Vserver   Volume       Aggregate    State      Type       Size  Available Used%
--------- ------------ ------------ ---------- ---- ---------- ---------- -----
svm1      svm1_root    N1_aggr6     online     RW          1GB    972.5MB    0%
svm1      v1           N1_aggr6     online     RW      105.3MB    99.78MB    0%
svm1      v2           N1_aggr1     online     RW      31.58MB    29.79MB    0%
3 entries were displayed.

```

#### nfs mount 

```
root@j1:~# mount  | grep  mnt
nsfs on /run/snapd/ns/snap-store.mnt type nsfs (rw)
192.168.56.131:/v1 on /mnt/v1 type nfs (rw,relatime,vers=3,rsize=65536,wsize=65536,namlen=255,hard,proto=tcp,timeo=600,retrans=2,sec=sys,mountaddr=192.168.56.131,mountvers=3,mountport=635,mountproto=udp,local_lock=none,addr=192.168.56.131)
192.168.56.131:/v2 on /mnt/v2 type nfs (rw,relatime,vers=3,rsize=65536,wsize=65536,namlen=255,hard,proto=tcp,timeo=600,retrans=2,sec=sys,mountaddr=192.168.56.131,mountvers=3,mountport=635,mountproto=udp,local_lock=none,addr=192.168.56.131)
```



#### volume 상태 확인 

```
FlexGroup::> set d

Warning: These diagnostic commands are for use by NetApp personnel only.
Do you want to continue? {y|n}: y

FlexGroup::*> volume show  -volume v1  -fields  size,files,files-used,files-maximum-possible;  
vserver volume size    files files-used files-maximum-possible 
------- ------ ------- ----- ---------- ---------------------- 
svm1    v1     105.3MB 3192  96         25590                  

FlexGroup::*> volume show  -volume v2  -fields  size,files,files-used,files-maximum-possible; 
vserver volume size    files files-used files-maximum-possible 
------- ------ ------- ----- ---------- ---------------------- 
svm1    v2     31.58MB 944   96         7675                   
```



#### 생성해 보도록 하자 

```
for i in {0..999}
do
for j in {0..999}
do
for k in {0..999}
do 
mkdir  -p /mnt/v1/${i}/${j}/${k}
for n in {0..999}
do
echo ${n}> /mnt/v1/${i}/${j}/${k}/${n}
done
done
done
done

```

#### no space 

```
mkdir: cannot create directory ‘/mnt/v1/0/9’: No space left on device
./r.sh: line 10: /mnt/v1/0/9/811/{0...999}: No such file or directory
mkdir: cannot create directory ‘/mnt/v1/0/9’: No space left on device
```

#### disk 용량은 있으나 

```
root@j1:~/ftest# df -h | grep mnt
192.168.56.131:/v1  101M  7.4M   93M   8% /mnt/v1
192.168.56.131:/v2   31M  256K   30M   1% /mnt/v2
```



#### volume 에서  3098개 파일이 생성되고 더  이상 생성 안됨 

```
root@j1:~/ftest# find /mnt/v1 | wc -l
3098


FlexGroup::*> volume show  -volume v1  -fields  size,files,files-used,files-maximum-possible; 
vserver volume size    files files-used files-maximum-possible 
------- ------ ------- ----- ---------- ---------------------- 
svm1    v1     105.3MB 3192  3192       25590   
```

* 최대 생성 갯수는 : 3192개
* 초기 생성: 96개
* 추가 생성: 3098     ==> 3194  (2개 더 있는 것은 .snapshot )



```
root@j1:/mnt/v2# df -h | grep mnt
192.168.56.131:/v1  101M  1.7M   99M   2% /mnt/v1
192.168.56.131:/v2   31M  640K   30M   3% /mnt/v2
root@j1:/mnt/v2# 
root@j1:/mnt/v2# 
root@j1:/mnt/v2# mount | grep mnt
nsfs on /run/snapd/ns/snap-store.mnt type nsfs (rw)
192.168.56.131:/v1 on /mnt/v1 type nfs (rw,relatime,vers=3,rsize=65536,wsize=65536,namlen=255,hard,proto=tcp,timeo=600,retrans=2,sec=sys,mountaddr=192.168.56.131,mountvers=3,mountport=635,mountproto=udp,local_lock=none,addr=192.168.56.131)
192.168.56.131:/v2 on /mnt/v2 type nfs (rw,relatime,vers=3,rsize=65536,wsize=65536,namlen=255,hard,proto=tcp,timeo=600,retrans=2,sec=sys,mountaddr=192.168.56.131,mountvers=3,mountport=635,mountproto=udp,local_lock=none,addr=192.168.56.131)


root@j1:/mnt/v2# find /mnt/v2 | wc -l
850

```

#### 96 + 850 => 946개

```
FlexGroup::*> volume show  -volume v2  -fields  size,files,files-used,files-maximum-possible;
vserver volume size    files files-used files-maximum-possible 
------- ------ ------- ----- ---------- ---------------------- 
svm1    v2     31.58MB 944   944        7675 
```



#### 30MB / 944 => 33KB

#### 100MB/3192 => 32KB 

대략 1 inode 당 32kb 데이터 블럭을 사용한다고 가정하고 inode 갯수를 예약한듯하다.



#### file갯수를  2000으로 변경하고 추가 파일 생성 

```
FlexGroup::*> volume modify  -vserver svm1 -volume v2 -files 2000   
Volume modify successful on volume v2 of Vserver svm1. 

FlexGroup::*> volume show  -volume v2  -fields  size,files,files-used,files-maximum-possible;
vserver volume size    files files-used files-maximum-possible 
------- ------ ------- ----- ---------- ---------------------- 
svm1    v2     31.58MB 1995  944        7675                   


root@j1:~/ftest# find /mnt/v2 | wc -l
1901

FlexGroup::*> 
FlexGroup::*> volume show  -volume v2  -fields  size,files,files-used,files-maximum-possible;
vserver volume size    files files-used files-maximum-possible 
------- ------ ------- ----- ---------- ---------------------- 
svm1    v2     31.58MB 1995  1995       7675    
```



* snapshot의 파일 갯수는 inode 갯수에서 계산하지 않는듯

```
root@j1:/mnt/v2# find /mnt/v2 | grep  snapshot | wc -l
1902
root@j1:/mnt/v2# find /mnt/v2 | grep -v  snapshot | wc -l
1900
root@j1:/mnt/v2# find /mnt/v2  | wc -l
3802

```



#### 1GB 볼륨으로 테스트

```
root@j1:~/ftest# df -h | grep mnt
192.168.56.131:/v1                                   101M  1.7M   99M   2% /mnt/v1
192.168.56.131:/v2                                    31M  1.3M   29M   5% /mnt/v2
192.168.56.131:/v3                                   1.1G   13M 1012M   2% /mnt/v3
192.168.56.131:/v2/.snapshot/snap.2023-08-16_051212  1.7M  1.4M  256K  85% /mnt/v2/.snapshot/snap.2023-08-16_051212
192.168.56.131:/v2/.snapshot/snap.2023-08-16_051301  1.7M  1.4M  256K  85% /mnt/v2/.snapshot/snap.2023-08-16_051301

root@j1:~/ftest# find /mnt/v3 | wc -l
32664

FlexGroup::*> volume show  -volume v3  -fields  size,files,files-used,files-maximum-possible;
vserver volume size   files files-used files-maximum-possible 
------- ------ ------ ----- ---------- ---------------------- 
svm1    v3     1.05GB 32758 32758      262143  

```

#### inode 1개당 사용하는 Bytes 추정 

* 12.05 MB used  / 32758개 = 384bytes