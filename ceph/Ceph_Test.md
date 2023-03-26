# Ceph Test

### SCSI 디스크 제거 해다가 다시 붙이기

#### scsi Device 제거

```
# lsscsi  
[2:0:0:0]    cd/dvd  NECVMWar VMware SATA CD00 1.00  /dev/sr0
[32:0:0:0]   disk    VMware   Virtual disk     2.0   /dev/sda
[32:0:1:0]   disk    VMware   Virtual disk     2.0   /dev/sdb
[32:0:2:0]   disk    VMware   Virtual disk     2.0   /dev/sdc
[32:0:3:0]   disk    VMware   Virtual disk     2.0   /dev/sdd
[32:0:4:0]   disk    VMware   Virtual disk     2.0   /dev/sde


# echo 1 > /sys/block/sde/device/delete
root@c06:/sys/block/sde/device# lsscsi
[2:0:0:0]    cd/dvd  NECVMWar VMware SATA CD00 1.00  /dev/sr0
[32:0:0:0]   disk    VMware   Virtual disk     2.0   /dev/sda
[32:0:1:0]   disk    VMware   Virtual disk     2.0   /dev/sdb
[32:0:2:0]   disk    VMware   Virtual disk     2.0   /dev/sdc
[32:0:3:0]   disk    VMware   Virtual disk     2.0   /dev/sdd

```

* 디스크가 날라 갔는데 태평 스러운 로그만 나오네...

```
2023-03-22T12:25:16.249186+0000 mgr.c01.uuuowu [INF] Detected new or changed devices on c06
2023-03-22T12:25:24.394527+0000 mgr.c01.uuuowu [INF] Detected new or changed devices on c06
2023-03-22T12:25:32.711502+0000 mgr.c01.uuuowu [INF] Detected new or changed devices on c06
2023-03-22T12:25:40.721051+0000 mgr.c01.uuuowu [INF] Detected new or changed devices on c06
2023-03-22T12:26:50.800107+0000 mgr.c01.uuuowu [INF] Detected new or changed devices on c06
```



* osd tree 에서는 down 으로 표시

```
root@c01:/# ceph osd tree
ID   CLASS  WEIGHT   TYPE NAME       STATUS  REWEIGHT  PRI-AFF
 -1         0.23511  root default
 -5         0.03918      host c01
  0    hdd  0.00980          osd.0       up   1.00000  1.00000
  6    hdd  0.00980          osd.6       up   1.00000  1.00000
 13    hdd  0.00980          osd.13      up   1.00000  1.00000
 19    hdd  0.00980          osd.19      up   1.00000  1.00000
 -7         0.03918      host c02
  4    hdd  0.00980          osd.4       up   1.00000  1.00000
  9    hdd  0.00980          osd.9       up   1.00000  1.00000
 15    hdd  0.00980          osd.15      up   1.00000  1.00000
 18    hdd  0.00980          osd.18      up   1.00000  1.00000
-13         0.03918      host c03
  5    hdd  0.00980          osd.5       up   1.00000  1.00000
 11    hdd  0.00980          osd.11      up   1.00000  1.00000
 17    hdd  0.00980          osd.17      up   1.00000  1.00000
 20    hdd  0.00980          osd.20      up   1.00000  1.00000
 -9         0.03918      host c04
 21    hdd  0.00980          osd.21      up   1.00000  1.00000
 22    hdd  0.00980          osd.22      up   1.00000  1.00000
 23    hdd  0.00980          osd.23      up   1.00000  1.00000
 24    hdd  0.00980          osd.24      up   1.00000  1.00000
-11         0.03918      host c05
  7    hdd  0.00980          osd.7       up   1.00000  1.00000
 10    hdd  0.00980          osd.10      up   1.00000  1.00000
 12    hdd  0.00980          osd.12      up   1.00000  1.00000
 14    hdd  0.00980          osd.14      up   1.00000  1.00000
 -3         0.03918      host c06
  1    hdd  0.00980          osd.1       up   1.00000  1.00000
  2    hdd  0.00980          osd.2       up   1.00000  1.00000
  3    hdd  0.00980          osd.3       up   1.00000  1.00000
  8    hdd  0.00980          osd.8     down   1.00000  1.00000  <<----
```

* ceph -s

```
# ceph -s
  cluster:
    id:     0824744e-a65e-7c16-8bca-5317bd09676f
    health: HEALTH_ERR
            1 failed cephadm daemon(s)
            30/321 objects unfound (9.346%)
            1 osds down  <<----
            Possible data damage: 25 pgs recovery_unfound
            Degraded data redundancy: 142/1077 objects degraded (13.185%), 36 pgs degraded, 44 pgs undersized
            1 daemons have recently crashed
```



* cep orch device 에서는  device 목록에서 빠져 있게 된다. 

```
# ceph orch device ls
HOST  PATH      TYPE  DEVICE ID   SIZE  AVAILABLE  REFRESHED  REJECT REASONS
c01   /dev/sdb  hdd              10.7G             17m ago    Insufficient space 
c01   /dev/sdc  hdd              10.7G             17m ago    Insufficient space 
c01   /dev/sdd  hdd              10.7G             17m ago    Insufficient space 
c01   /dev/sde  hdd              10.7G             17m ago    Insufficient space 
c02   /dev/sdb  hdd              10.7G             17m ago    Insufficient space 
c02   /dev/sdc  hdd              10.7G             17m ago    Insufficient space 
c02   /dev/sdd  hdd              10.7G             17m ago    Insufficient space 
c02   /dev/sde  hdd              10.7G             17m ago    Insufficient space 
c03   /dev/sdb  hdd              10.7G             18m ago    Insufficient space 
c03   /dev/sdc  hdd              10.7G             18m ago    Insufficient space 
c03   /dev/sdd  hdd              10.7G             18m ago    Insufficient space 
c03   /dev/sde  hdd              10.7G             18m ago    Insufficient space 
c04   /dev/sdb  hdd              10.7G             27m ago    Insufficient space 
c04   /dev/sdc  hdd              10.7G             27m ago    Insufficient space 
c04   /dev/sdd  hdd              10.7G             27m ago    Insufficient space 
c04   /dev/sde  hdd              10.7G             27m ago    Insufficient space 
c05   /dev/sdb  hdd              10.7G             16m ago    Insufficient space 
c05   /dev/sdc  hdd              10.7G             16m ago    Insufficient space 
```



```
root@c01:/# ceph osd status
ID  HOST   USED  AVAIL  WR OPS  WR DATA  RD OPS  RD DATA  STATE
 0  c01    123M  9.87G      0        0       0        0   exists,up
 1  c06   83.3M  9.91G      0        0       0        0   exists,up
 2  c06   90.8M  9.90G      0        0       0        0   exists,up
 3  c06   67.4M  9.93G      0        0       0        0   exists,up
 4  c02    125M  9.87G      0        0       0        0   exists,up
 5  c03    124M  9.87G      0        0       0        0   exists,up
 6  c01    171M  9.82G      0        0       0        0   exists,up
 7  c05   71.7M  9.92G      0        0       0        0   exists,up
 8  c06   73.9M  9.92G      0        0       0        0   exists      <<<-------
 9  c02    123M  9.87G      0        0       0        0   exists,up
10  c05   61.2M  9.93G      0        0       0        0   exists,up
11  c03    100M  9.89G      0        0       0        0   exists,up
12  c05   64.6M  9.93G      0        0       0        0   exists,up
13  c01   91.5M  9.90G      0        0       0        0   exists,up
14  c05   65.1M  9.93G      0        0       0        0   exists,up
15  c02    128M  9.87G      0        0       0        0   exists,up
17  c03    115M  9.88G      0        0       0        0   exists,up
18  c02    130M  9.86G      0        0       0        0   exists,up
19  c01   98.9M  9.89G      0        0       0        0   exists,up
20  c03    128M  9.87G      0        0       0        0   exists,up
21  c04   74.8M  9.92G      0        0       0        0   exists,up
22  c04   79.1M  9.91G      0        0       0        0   exists,up
23  c04   73.4M  9.92G      0        0       0        0   exists,up
24  c04   62.0M  9.93G      0        0       0        0   exists,up

```

* 6번 host에서 보면 docker가 없음

```
root@c06:/sys/block/sde/device# systemctl | grep ceph
  ceph-0824744e-a65e-7c16-8bca-5317bd09676f@crash.c06.service                                           loaded active running   Ceph crash.c06 for 0824744e-a65e-7c16-8bca-5317bd09676f
  ceph-0824744e-a65e-7c16-8bca-5317bd09676f@node-exporter.c06.service                                   loaded active running   Ceph node-exporter.c06 for 0824744e-a65e-7c16-8bca-5317bd09676f
  ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.1.service                                               loaded active running   Ceph osd.1 for 0824744e-a65e-7c16-8bca-5317bd09676f
  ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.2.service                                               loaded active running   Ceph osd.2 for 0824744e-a65e-7c16-8bca-5317bd09676f
  ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.3.service                                               loaded active running   Ceph osd.3 for 0824744e-a65e-7c16-8bca-5317bd09676f
● ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service                                               loaded failed failed    Ceph osd.8 for 0824744e-a65e-7c16-8bca-5317bd09676f   <<<=====
  ceph-0824744e-a65e-7c16-8bca-5317bd09676f@promtail.c06.service                                        loaded active running   Ceph promtail.c06 for 0824744e-a65e-7c16-8bca-5317bd09676f
  system-ceph\x2d0824744e\x2da65e\x2d7c16\x2d8bca\x2d5317bd09676f.slice                                 loaded active active    system-ceph\x2d0824744e\x2da65e\x2d7c16\x2d8bca\x2d5317bd09676f.slice
  ceph-0824744e-a65e-7c16-8bca-5317bd09676f.target                                                      loaded active active    Ceph cluster 0824744e-a65e-7c16-8bca-5317bd09676f
  ceph.target                                                                                           loaded active active    All Ceph clusters and services
```



* osd의 로그를 보자

```
# journalctl -u  ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8


-- Logs begin at Fri 2023-03-17 12:41:24 UTC, end at Wed 2023-03-22 12:38:13 UTC. --
Mar 22 12:27:34 c06 systemd[1]: Failed to start Ceph osd.8 for 0824744e-a65e-7c16-8bca-5317bd09676f.
Mar 22 12:27:34 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Failed with result 'exit-code'.
Mar 22 12:27:34 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Start request repeated too quickly.
Mar 22 12:27:34 c06 systemd[1]: Stopped Ceph osd.8 for 0824744e-a65e-7c16-8bca-5317bd09676f.
Mar 22 12:27:34 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Scheduled restart job, restart counter is at 6.
Mar 22 12:27:24 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Failed with result 'exit-code'.
Mar 22 12:27:23 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Main process exited, code=exited, status=1/FAILURE
Mar 22 12:27:23 c06 bash[155648]: debug 2023-03-22T12:27:23.494+0000 7f0ec3d343c0 -1  ** ERROR: unable to open OSD superblock on /var/lib/ceph/osd/ceph-8: (2) No such file or directory
Mar 22 12:27:23 c06 bash[155648]: debug 2023-03-22T12:27:23.494+0000 7f0ec3d343c0 -1 bluestore(/var/lib/ceph/osd/ceph-8/block) _read_bdev_label failed to read from /var/lib/ceph/osd/ceph-8/block: (5) Input/output error
Mar 22 12:27:23 c06 bash[155648]: debug 2023-03-22T12:27:23.490+0000 7f0ec3d343c0  0 pidfile_write: ignore empty --pid-file
Mar 22 12:27:23 c06 bash[155648]: debug 2023-03-22T12:27:23.490+0000 7f0ec3d343c0  0 ceph version 17.2.3 (dff484dfc9e19a9819f375586300b3b79d80034d) quincy (stable), process ceph-osd, pid 7
Mar 22 12:27:23 c06 bash[155648]: debug 2023-03-22T12:27:23.490+0000 7f0ec3d343c0  0 set uid:gid to 167:167 (ceph:ceph)
Mar 22 12:27:22 c06 bash[155430]: --> Failed to activate any OSD(s)
Mar 22 12:27:22 c06 bash[155430]: --> Failed to activate via simple: 'Namespace' object has no attribute 'json_config'
Mar 22 12:27:22 c06 bash[155430]: --> Failed to activate via lvm: could not find osd.8 with osd_fsid 640c49da-554c-49b3-9cf0-7f74b2d2009c
Mar 22 12:27:22 c06 bash[155430]: --> Failed to activate via raw: did not find any matching OSD to activate
Mar 22 12:27:21 c06 systemd[1]: Started Ceph osd.8 for 0824744e-a65e-7c16-8bca-5317bd09676f.
Mar 22 12:27:21 c06 systemd[1]: Stopped Ceph osd.8 for 0824744e-a65e-7c16-8bca-5317bd09676f.
Mar 22 12:27:21 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Scheduled restart job, restart counter is at 5.
Mar 22 12:27:11 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Failed with result 'exit-code'.
Mar 22 12:27:10 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Main process exited, code=exited, status=1/FAILURE
Mar 22 12:27:10 c06 bash[155104]: debug 2023-03-22T12:27:10.739+0000 7f4dbe5a23c0 -1  ** ERROR: unable to open OSD superblock on /var/lib/ceph/osd/ceph-8: (2) No such file or directory
Mar 22 12:27:10 c06 bash[155104]: debug 2023-03-22T12:27:10.739+0000 7f4dbe5a23c0 -1 bluestore(/var/lib/ceph/osd/ceph-8/block) _read_bdev_label failed to read from /var/lib/ceph/osd/ceph-8/block: (5) Input/output error
Mar 22 12:27:10 c06 bash[155104]: debug 2023-03-22T12:27:10.739+0000 7f4dbe5a23c0  0 pidfile_write: ignore empty --pid-file
Mar 22 12:27:10 c06 bash[155104]: debug 2023-03-22T12:27:10.739+0000 7f4dbe5a23c0  0 ceph version 17.2.3 (dff484dfc9e19a9819f375586300b3b79d80034d) quincy (stable), process ceph-osd, pid 6
Mar 22 12:27:10 c06 bash[155104]: debug 2023-03-22T12:27:10.739+0000 7f4dbe5a23c0  0 set uid:gid to 167:167 (ceph:ceph)
Mar 22 12:27:10 c06 bash[154891]: --> Failed to activate any OSD(s)
Mar 22 12:27:10 c06 bash[154891]: --> Failed to activate via simple: 'Namespace' object has no attribute 'json_config'
Mar 22 12:27:10 c06 bash[154891]: --> Failed to activate via lvm: could not find osd.8 with osd_fsid 640c49da-554c-49b3-9cf0-7f74b2d2009c
Mar 22 12:27:10 c06 bash[154891]: --> Failed to activate via raw: did not find any matching OSD to activate
Mar 22 12:27:08 c06 systemd[1]: Started Ceph osd.8 for 0824744e-a65e-7c16-8bca-5317bd09676f.
Mar 22 12:27:08 c06 systemd[1]: Stopped Ceph osd.8 for 0824744e-a65e-7c16-8bca-5317bd09676f.
Mar 22 12:27:08 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Scheduled restart job, restart counter is at 4.
Mar 22 12:26:58 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Failed with result 'exit-code'.
Mar 22 12:26:57 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Main process exited, code=exited, status=1/FAILURE
Mar 22 12:26:57 c06 bash[154565]: debug 2023-03-22T12:26:57.931+0000 7f4045c963c0 -1  ** ERROR: unable to open OSD superblock on /var/lib/ceph/osd/ceph-8: (2) No such file or directory
Mar 22 12:26:57 c06 bash[154565]: debug 2023-03-22T12:26:57.931+0000 7f4045c963c0 -1 bluestore(/var/lib/ceph/osd/ceph-8/block) _read_bdev_label failed to read from /var/lib/ceph/osd/ceph-8/block: (5) Input/output error
Mar 22 12:26:57 c06 bash[154565]: debug 2023-03-22T12:26:57.915+0000 7f4045c963c0  0 pidfile_write: ignore empty --pid-file
Mar 22 12:26:57 c06 bash[154565]: debug 2023-03-22T12:26:57.915+0000 7f4045c963c0  0 ceph version 17.2.3 (dff484dfc9e19a9819f375586300b3b79d80034d) quincy (stable), process ceph-osd, pid 6
Mar 22 12:26:57 c06 bash[154565]: debug 2023-03-22T12:26:57.915+0000 7f4045c963c0  0 set uid:gid to 167:167 (ceph:ceph)
Mar 22 12:26:57 c06 bash[154352]: --> Failed to activate any OSD(s)
Mar 22 12:26:57 c06 bash[154352]: --> Failed to activate via simple: 'Namespace' object has no attribute 'json_config'
Mar 22 12:26:57 c06 bash[154352]: --> Failed to activate via lvm: could not find osd.8 with osd_fsid 640c49da-554c-49b3-9cf0-7f74b2d2009c
Mar 22 12:26:57 c06 bash[154352]: --> Failed to activate via raw: did not find any matching OSD to activate
Mar 22 12:26:55 c06 systemd[1]: Started Ceph osd.8 for 0824744e-a65e-7c16-8bca-5317bd09676f.
Mar 22 12:26:55 c06 systemd[1]: Stopped Ceph osd.8 for 0824744e-a65e-7c16-8bca-5317bd09676f.
Mar 22 12:26:55 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Scheduled restart job, restart counter is at 3.
Mar 22 12:26:45 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Failed with result 'exit-code'.
Mar 22 12:26:45 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Main process exited, code=exited, status=1/FAILURE
Mar 22 12:26:44 c06 bash[153697]: debug 2023-03-22T12:26:44.976+0000 7f096e5363c0  0 pidfile_write: ignore empty --pid-file
Mar 22 12:26:44 c06 bash[153697]: debug 2023-03-22T12:26:44.976+0000 7f096e5363c0  0 ceph version 17.2.3 (dff484dfc9e19a9819f375586300b3b79d80034d) quincy (stable), process ceph-osd, pid 8
Mar 22 12:26:44 c06 bash[153697]: debug 2023-03-22T12:26:44.976+0000 7f096e5363c0  0 set uid:gid to 167:167 (ceph:ceph)
Mar 22 12:26:44 c06 bash[153485]: --> Failed to activate any OSD(s)
Mar 22 12:26:44 c06 bash[153485]: --> Failed to activate via simple: 'Namespace' object has no attribute 'json_config'
Mar 22 12:26:44 c06 bash[153485]: --> Failed to activate via lvm: could not find osd.8 with osd_fsid 640c49da-554c-49b3-9cf0-7f74b2d2009c
Mar 22 12:26:44 c06 bash[153485]: --> Failed to activate via raw: did not find any matching OSD to activate
Mar 22 12:26:43 c06 systemd[1]: Started Ceph osd.8 for 0824744e-a65e-7c16-8bca-5317bd09676f.
Mar 22 12:26:43 c06 systemd[1]: Stopped Ceph osd.8 for 0824744e-a65e-7c16-8bca-5317bd09676f.
Mar 22 12:26:43 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Scheduled restart job, restart counter is at 2.
Mar 22 12:26:33 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Failed with result 'exit-code'.
Mar 22 12:26:32 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Main process exited, code=exited, status=1/FAILURE
Mar 22 12:26:32 c06 bash[153093]: debug 2023-03-22T12:26:32.133+0000 7f069a9823c0 -1  ** ERROR: unable to open OSD superblock on /var/lib/ceph/osd/ceph-8: (2) No such file or directory
Mar 22 12:26:32 c06 bash[153093]: debug 2023-03-22T12:26:32.133+0000 7f069a9823c0 -1 bluestore(/var/lib/ceph/osd/ceph-8/block) _read_bdev_label failed to read from /var/lib/ceph/osd/ceph-8/block: (5) Input/output error
Mar 22 12:26:32 c06 bash[153093]: debug 2023-03-22T12:26:32.129+0000 7f069a9823c0  0 pidfile_write: ignore empty --pid-file
Mar 22 12:26:32 c06 bash[153093]: debug 2023-03-22T12:26:32.129+0000 7f069a9823c0  0 ceph version 17.2.3 (dff484dfc9e19a9819f375586300b3b79d80034d) quincy (stable), process ceph-osd, pid 7
Mar 22 12:26:32 c06 bash[153093]: debug 2023-03-22T12:26:32.129+0000 7f069a9823c0  0 set uid:gid to 167:167 (ceph:ceph)
Mar 22 12:26:31 c06 bash[152880]: --> Failed to activate any OSD(s)
Mar 22 12:26:31 c06 bash[152880]: --> Failed to activate via simple: 'Namespace' object has no attribute 'json_config'
Mar 22 12:26:31 c06 bash[152880]: --> Failed to activate via lvm: could not find osd.8 with osd_fsid 640c49da-554c-49b3-9cf0-7f74b2d2009c
Mar 22 12:26:31 c06 bash[152880]: --> Failed to activate via raw: did not find any matching OSD to activate
Mar 22 12:26:30 c06 systemd[1]: Started Ceph osd.8 for 0824744e-a65e-7c16-8bca-5317bd09676f.
Mar 22 12:26:30 c06 systemd[1]: Stopped Ceph osd.8 for 0824744e-a65e-7c16-8bca-5317bd09676f.
Mar 22 12:26:30 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Scheduled restart job, restart counter is at 1.
Mar 22 12:26:19 c06 systemd[1]: ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service: Failed with result 'exit-code'.

```



##### 일단 로그를 좀 보면..

* 여기의 block Device 가 없다고 나온다.  그런데 실제 그 device 가 없어서 그런다.

```
root@c06:/var/lib/ceph/0824744e-a65e-7c16-8bca-5317bd09676f/osd.8# ls -l
total 64
lrwxrwxrwx 1 167 167  111 Mar 22 07:23 block -> /dev/mapper/ceph--6a42d416--d50b--45e4--ad58--7b73221b61b7-osd--block--640c49da--554c--49b3--9cf0--7f74b2d2009c
-rw------- 1 167 167   37 Mar 22 07:23 ceph_fsid
-rw------- 1 167 167  387 Mar 22 07:23 config
-rw------- 1 167 167   37 Mar 22 07:23 fsid
-rw------- 1 167 167   55 Mar 22 07:23 keyring
-rw------- 1 167 167    6 Mar 22 07:23 ready
-rw------- 1 167 167    3 Mar 22 07:23 require_osd_release
-rw------- 1 167 167   10 Mar 22 07:23 type
-rw------- 1 167 167   38 Mar 22 07:23 unit.configured
-rw------- 1 167 167   48 Mar 22 07:23 unit.created
-rw------- 1 167 167   99 Mar 22 07:23 unit.image
-rw------- 1 167 167  357 Mar 22 07:23 unit.meta
-rw------- 1 167 167 1301 Mar 22 07:23 unit.poststop
-rw------- 1 167 167 2786 Mar 22 07:23 unit.run
-rw------- 1 167 167  318 Mar 22 07:23 unit.stop
-rw------- 1 167 167    2 Mar 22 07:23 whoami
```



#### scsi device 다시 등록

```
# lsblk
# lsscsi 
root@c06:/sys/class/scsi_host# rescan-scsi-bus.sh -a
Scanning SCSI subsystem for new devices
Scanning host 0 for  SCSI target IDs  0 1 2 3 4 5 6 7, all LUNs
Scanning host 1 for  SCSI target IDs  0 1 2 3 4 5 6 7, all LUNs
Scanning host 2 for  SCSI target IDs  0 1 2 3 4 5 6 7, all LUNs
... Scanning for device 32 0 4 0 ...
NEW: Host: scsi32 Channel: 00 Id: 04 Lun: 00
      Vendor: VMware   Model: Virtual disk     Rev: 2.0
      Type:   Direct-Access                    ANSI SCSI revision: 06
1 new or changed device(s) found.
        [32:0:4:0]
0 remapped or resized device(s) found.
0 device(s) removed.

root@c06:/sys/class/scsi_host# lsscsi
[2:0:0:0]    cd/dvd  NECVMWar VMware SATA CD00 1.00  /dev/sr0
[32:0:0:0]   disk    VMware   Virtual disk     2.0   /dev/sda
[32:0:1:0]   disk    VMware   Virtual disk     2.0   /dev/sdb
[32:0:2:0]   disk    VMware   Virtual disk     2.0   /dev/sdc
[32:0:3:0]   disk    VMware   Virtual disk     2.0   /dev/sdd
[32:0:4:0]   disk    VMware   Virtual disk     2.0   /dev/sdf    <<<==== sdf로 

```

* 잘 안잡히네..
* reboot으로 해결해보자.
* 재부팅 안하고 ...# rescan-scsi-bus.sh -a 옵션으로 수행해서 ...



##### osd start1

```
root@c01:/# ceph orch daemon start osd.8
Scheduled to start osd.8 on host 'c06'
```

* 하지만 여전히 osd tree에서 보면 down 상태로  나온다. 



* osd destroy는 osd id와 crush map을 유지한체  파괴된 디스크를 교체한다.
* osd delete는 osd id와 crush map에서 정보를 제거하기 때문에 완전히 새로 디스크로 등록하는 것이다. 
* osd 제거 : 인증키를 제거하며, osd 맵에서 osd를 제거하고 파일에서 osd를 제거한다. 



#### osd start2

```
root@c01:/# ceph cephadm osd activate c06
Created no osd(s) on host c06; already created?
```

* 이렇게 해도 osd가 올라오지 않는 구나...



#### 재부팅..

```
root@c06:~# lsscsi
[2:0:0:0]    cd/dvd  NECVMWar VMware SATA CD00 1.00  /dev/sr0
[32:0:0:0]   disk    VMware   Virtual disk     2.0   /dev/sda
[32:0:1:0]   disk    VMware   Virtual disk     2.0   /dev/sdb
[32:0:2:0]   disk    VMware   Virtual disk     2.0   /dev/sdc
[32:0:3:0]   disk    VMware   Virtual disk     2.0   /dev/sdd
[32:0:4:0]   disk    VMware   Virtual disk     2.0   /dev/sde
```

* 디스크 위치가 잘 올라옴

```
root@c01:/# ceph osd tree
ID   CLASS  WEIGHT   TYPE NAME       STATUS  REWEIGHT  PRI-AFF
 -1         0.23511  root default
 -5         0.03918      host c01
  0    hdd  0.00980          osd.0       up   1.00000  1.00000
  6    hdd  0.00980          osd.6       up   1.00000  1.00000
 13    hdd  0.00980          osd.13      up   1.00000  1.00000
 19    hdd  0.00980          osd.19      up   1.00000  1.00000
 -7         0.03918      host c02
  4    hdd  0.00980          osd.4       up   1.00000  1.00000
  9    hdd  0.00980          osd.9       up   1.00000  1.00000
 15    hdd  0.00980          osd.15      up   1.00000  1.00000
 18    hdd  0.00980          osd.18      up   1.00000  1.00000
-13         0.03918      host c03
  5    hdd  0.00980          osd.5       up   1.00000  1.00000
 11    hdd  0.00980          osd.11      up   1.00000  1.00000
 17    hdd  0.00980          osd.17      up   1.00000  1.00000
 20    hdd  0.00980          osd.20      up   1.00000  1.00000
 -9         0.03918      host c04
 21    hdd  0.00980          osd.21      up   1.00000  1.00000
 22    hdd  0.00980          osd.22      up   1.00000  1.00000
 23    hdd  0.00980          osd.23      up   1.00000  1.00000
 24    hdd  0.00980          osd.24      up   1.00000  1.00000
-11         0.03918      host c05
  7    hdd  0.00980          osd.7       up   1.00000  1.00000
 10    hdd  0.00980          osd.10      up   1.00000  1.00000
 12    hdd  0.00980          osd.12      up   1.00000  1.00000
 14    hdd  0.00980          osd.14      up   1.00000  1.00000
 -3         0.03918      host c06
  1    hdd  0.00980          osd.1       up   1.00000  1.00000
  2    hdd  0.00980          osd.2       up   1.00000  1.00000
  3    hdd  0.00980          osd.3       up   1.00000  1.00000
  8    hdd  0.00980          osd.8       up   1.00000  1.00000
```



#### 결국.

디스크를 빼고 다시 인식할때 /dev/sd* 정보가 정확하게 만  올라고면 /pv/vg/lv 정보가 매핑되고 거기에 따라서  잘 올라오게 된다.

그런데 /dev/sde ==> /dev/sdf 이렇게 올라오면  LVM에서 뭔가 잡았어도 osd  start가 안된다.   



### SCSI disk 강제로 format 해보기





```
root@c06:~# lsscsi
[2:0:0:0]    cd/dvd  NECVMWar VMware SATA CD00 1.00  /dev/sr0
[32:0:0:0]   disk    VMware   Virtual disk     2.0   /dev/sda
[32:0:1:0]   disk    VMware   Virtual disk     2.0   /dev/sdb
[32:0:2:0]   disk    VMware   Virtual disk     2.0   /dev/sdc
[32:0:3:0]   disk    VMware   Virtual disk     2.0   /dev/sdd
[32:0:4:0]   disk    VMware   Virtual disk     2.0   /dev/sde
root@c06:~# lvs
  LV                                             VG                                        Attr       LSize   Pool Origin Data%  Meta%  Move Log Cpy%Sync Convert
  osd-block-77e02e02-96d2-4de1-b6ec-57da44a92dc5 ceph-34b8eb36-7896-46b3-aa17-a7d74620020e -wi-ao---- <10.00g
  osd-block-640c49da-554c-49b3-9cf0-7f74b2d2009c ceph-6a42d416-d50b-45e4-ad58-7b73221b61b7 -wi-ao---- <10.00g
  osd-block-1c3107f0-4f0b-4996-9d67-aa3510c4aeed ceph-af7bb618-81a6-4bdf-bb06-0b06b9e30699 -wi-ao---- <10.00g
  osd-block-2800e7b4-1deb-4172-adb2-cef325aee5e0 ceph-fe47a268-6524-4ac2-9301-99944df3dfb6 -wi-ao---- <10.00g
  lv-0                                           ubuntu-vg                                 -wi-ao---- <28.00g

root@c06:~# mkfs.ext4  /dev/ceph-6a42d416-d50b-45e4-ad58-7b73221b61b7/osd-block-640c49da-554c-49b3-9cf0-7f74b2d2009c
mke2fs 1.45.5 (07-Jan-2020)
/dev/ceph-6a42d416-d50b-45e4-ad58-7b73221b61b7/osd-block-640c49da-554c-49b3-9cf0-7f74b2d2009c contains a ceph_bluestore file system
Proceed anyway? (y,N) y
Discarding device blocks: done
Creating filesystem with 2620416 4k blocks and 655360 inodes
Filesystem UUID: 7cbcc024-0891-455e-b52b-dc8086126dc5
Superblock backups stored on blocks:
        32768, 98304, 163840, 229376, 294912, 819200, 884736, 1605632

Allocating group tables: done
Writing inode tables: done
Creating journal (16384 blocks): done
Writing superblocks and filesystem accounting information: done

```



#### cephadm log 에러

```
2023-03-22T14:10:00.001353+0000 mon.c01 [ERR]     osd.8 crashed on host c06 at 2023-03-22T12:26:14.074640Z
2023-03-22T14:11:24.860274+0000 osd.2 [WRN] 9 slow requests (by type [ 'delayed' : 9 ] most affected pool [ 'default.rgw.control' : 9 ])
2023-03-22T14:11:25.328158+0000 osd.3 [WRN] 1 slow requests (by type [ 'delayed' : 1 ] most affected pool [ 'default.rgw.control' : 1 ])
2023-03-22T14:11:25.823984+0000 osd.2 [WRN] 9 slow requests (by type [ 'delayed' : 9 ] most affected pool [ 'default.rgw.control' : 9 ])
2023-03-22T14:11:26.354881+0000 osd.3 [WRN] 1 slow requests (by type [ 'delayed' : 1 ] most affected pool [ 'default.rgw.control' : 1 ])
2023-03-22T14:11:26.860827+0000 osd.2 [WRN] 9 slow requests (by type [ 'delayed' : 9 ] most affected pool [ 'default.rgw.control' : 9 ])
2023-03-22T14:11:27.358559+0000 osd.3 [WRN] 3 slow requests (by type [ 'delayed' : 3 ] most affected pool [ 'default.rgw.control' : 3 ])
2023-03-22T14:11:27.867085+0000 osd.2 [WRN] 11 slow requests (by type [ 'delayed' : 11 ] most affected pool [ 'default.rgw.control' : 11 ])
2023-03-22T14:11:28.182573+0000 mon.c01 [WRN] Health check failed: 12 slow ops, oldest one blocked for 174 sec, daemons [osd.2,osd.3] have slow ops. (SLOW_OPS)
2023-03-22T14:11:28.378332+0000 osd.3 [WRN] 3 slow requests (by type [ 'delayed' : 3 ] most affected pool [ 'default.rgw.control' : 3 ])
2023-03-22T14:11:28.919211+0000 osd.2 [WRN] 13 slow requests (by type [ 'delayed' : 13 ] most affected pool [ 'default.rgw.control' : 13 ])
2023-03-22T14:11:29.358186+0000 osd.3 [WRN] 5 slow requests (by type [ 'delayed' : 5 ] most affected pool [ 'default.rgw.control' : 5 ])
2023-03-22T14:11:29.894818+0000 osd.2 [WRN] 13 slow requests (by type [ 'delayed' : 13 ] most affected pool [ 'default.rgw.control' : 13 ])
2023-03-22T14:11:30.361070+0000 osd.3 [WRN] 5 slow requests (by type [ 'delayed' : 5 ] most affected pool [ 'default.rgw.control' : 5 ])
2023-03-22T14:11:30.893286+0000 osd.2 [WRN] 13 slow requests (by type [ 'delayed' : 13 ] most affected pool [ 'default.rgw.control' : 13 ])
2023-03-22T14:11:31.364618+0000 osd.3 [WRN] 5 slow requests (by type [ 'delayed' : 5 ] most affected pool [ 'default.rgw.control' : 5 ])
2023-03-22T14:11:31.890725+0000 osd.2 [WRN] 15 slow requests (by type [ 'delayed' : 15 ] most affected pool [ 'default.rgw.control' : 15 ])
2023-03-22T14:11:32.336753+0000 osd.3 [WRN] 7 slow requests (by type [ 'delayed' : 7 ] most affected pool [ 'default.rgw.control' : 7 ])
2023-03-22T14:11:32.877229+0000 osd.2 [WRN] 15 slow requests (by type [ 'delayed' : 15 ] most affected pool [ 'default.rgw.control' : 15 ])
2023-03-22T14:11:33.291759+0000 osd.3 [WRN] 7 slow requests (by type [ 'delayed' : 7 ] most affected pool [ 'default.rgw.control' : 7 ])
2023-03-22T14:11:33.847212+0000 osd.2 [WRN] 17 slow requests (by type [ 'delayed' : 17 ] most affected pool [ 'default.rgw.control' : 17 ])
2023-03-22T14:11:34.274299+0000 osd.3 [WRN] 9 slow requests (by type [ 'delayed' : 9 ] most affected pool [ 'default.rgw.control' : 9 ])
2023-03-22T14:11:34.478251+0000 mon.c01 [WRN] Health check update: 22 slow ops, oldest one blocked for 179 sec, daemons [osd.2,osd.3] have slow ops. (SLOW_OPS)
2023-03-22T14:11:34.861738+0000 osd.2 [WRN] 17 slow requests (by type [ 'delayed' : 17 ] most affected pool [ 'default.rgw.control' : 17 ])
2023-03-22T14:11:35.813675+0000 osd.2 [WRN] 17 slow requests (by type [ 'delayed' : 17 ] most affected pool [ 'default.rgw.control' : 17 ])
2023-03-22T14:11:37.970654+0000 mon.c01 [INF] Health check cleared: SLOW_OPS (was: 22 slow ops, oldest one blocked for 179 sec, daemons [osd.2,osd.3] have slow ops.)
```



#### osd tree는 아무것도 모르는 상태

```
root@c01:/# ceph osd tree
ID   CLASS  WEIGHT   TYPE NAME       STATUS  REWEIGHT  PRI-AFF
 -1         0.23511  root default
 -5         0.03918      host c01
  0    hdd  0.00980          osd.0       up   1.00000  1.00000
  6    hdd  0.00980          osd.6       up   1.00000  1.00000
 13    hdd  0.00980          osd.13      up   1.00000  1.00000
 19    hdd  0.00980          osd.19      up   1.00000  1.00000
 -7         0.03918      host c02
  4    hdd  0.00980          osd.4       up   1.00000  1.00000
  9    hdd  0.00980          osd.9       up   1.00000  1.00000
 15    hdd  0.00980          osd.15      up   1.00000  1.00000
 18    hdd  0.00980          osd.18      up   1.00000  1.00000
-13         0.03918      host c03
  5    hdd  0.00980          osd.5       up   1.00000  1.00000
 11    hdd  0.00980          osd.11      up   1.00000  1.00000
 17    hdd  0.00980          osd.17      up   1.00000  1.00000
 20    hdd  0.00980          osd.20      up   1.00000  1.00000
 -9         0.03918      host c04
 21    hdd  0.00980          osd.21      up   1.00000  1.00000
 22    hdd  0.00980          osd.22      up   1.00000  1.00000
 23    hdd  0.00980          osd.23      up   1.00000  1.00000
 24    hdd  0.00980          osd.24      up   1.00000  1.00000
-11         0.03918      host c05
  7    hdd  0.00980          osd.7       up   1.00000  1.00000
 10    hdd  0.00980          osd.10      up   1.00000  1.00000
 12    hdd  0.00980          osd.12      up   1.00000  1.00000
 14    hdd  0.00980          osd.14      up   1.00000  1.00000
 -3         0.03918      host c06
  1    hdd  0.00980          osd.1       up   1.00000  1.00000
  2    hdd  0.00980          osd.2       up   1.00000  1.00000
  3    hdd  0.00980          osd.3       up   1.00000  1.00000
  8    hdd  0.00980          osd.8       up   1.00000  1.00000 <<<---

```



#### osd.8 로그

* 멀쩡하게 살아 있음. 

```
root@c06:~# systemctl status  ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8
● ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service - Ceph osd.8 for 0824744e-a65e-7c16-8bca-5317bd09676f
     Loaded: loaded (/etc/systemd/system/ceph-0824744e-a65e-7c16-8bca-5317bd09676f@.service; enabled; vendor preset: enabled)
     Active: active (running) since Wed 2023-03-22 14:05:56 UTC; 11min ago
   Main PID: 1420 (bash)
      Tasks: 10 (limit: 19102)
     Memory: 23.9M
     CGroup: /system.slice/system-ceph\x2d0824744e\x2da65e\x2d7c16\x2d8bca\x2d5317bd09676f.slice/ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8.service
             ├─1420 /bin/bash /var/lib/ceph/0824744e-a65e-7c16-8bca-5317bd09676f/osd.8/unit.run
             └─3018 /usr/bin/docker run --rm --ipc=host --stop-signal=SIGTERM --net=host --entrypoint /usr/bin/ceph-osd --privileged --group-add=disk --init --name ceph-0824744e-a65e-7c16-8bca-5317bd09676f-osd-8 -e CONTAINER_IMAGE=ha>

Mar 22 14:16:10 c06 bash[3018]: Uptime(secs): 600.2 total, 0.0 interval
Mar 22 14:16:10 c06 bash[3018]: Flush(GB): cumulative 0.002, interval 0.000
Mar 22 14:16:10 c06 bash[3018]: AddFile(GB): cumulative 0.000, interval 0.000
Mar 22 14:16:10 c06 bash[3018]: AddFile(Total Files): cumulative 0, interval 0
Mar 22 14:16:10 c06 bash[3018]: AddFile(L0 Files): cumulative 0, interval 0
Mar 22 14:16:10 c06 bash[3018]: AddFile(Keys): cumulative 0, interval 0
Mar 22 14:16:10 c06 bash[3018]: Cumulative compaction: 0.00 GB write, 0.00 MB/s write, 0.00 GB read, 0.00 MB/s read, 0.0 seconds
Mar 22 14:16:10 c06 bash[3018]: Interval compaction: 0.00 GB write, 0.00 MB/s write, 0.00 GB read, 0.00 MB/s read, 0.0 seconds
Mar 22 14:16:10 c06 bash[3018]: Stalls(count): 0 level0_slowdown, 0 level0_slowdown_with_compaction, 0 level0_numfiles, 0 level0_numfiles_with_compaction, 0 stop for pending_compaction_bytes, 0 slowdown for pending_compaction_bytes, >
Mar 22 14:16:10 c06 bash[3018]: ** File Read Latency Histogram By Level [P] **

```

* journal log도 별 문제는 없어 보이고..

```
root@c06:~# journalctl  -r -u  ceph-0824744e-a65e-7c16-8bca-5317bd09676f@osd.8

-- Logs begin at Fri 2023-03-17 12:41:24 UTC, end at Wed 2023-03-22 14:14:46 UTC. --
Mar 22 14:06:24 c06 bash[3018]: debug 2023-03-22T14:06:24.136+0000 7fca8cd0b700 -1 log_channel(cluster) log [ERR] : 9.1c has 2 objects unfound and apparently lost
Mar 22 14:06:23 c06 bash[3018]: debug 2023-03-22T14:06:23.444+0000 7fca8c50a700  0 log_channel(cluster) log [DBG] : 7.4 scrub ok
Mar 22 14:06:23 c06 bash[3018]: debug 2023-03-22T14:06:23.037+0000 7fca8c50a700  1 osd.8 pg_epoch: 742 pg[6.18( v 733'4562 (0'0,733'4562] local-lis/les=629/630 n=5 ec=112/103 lis/c=740/731 les/c/f=741/732/0 sis=742) [17,10,8] r=2 lpr>
Mar 22 14:06:23 c06 bash[3018]: debug 2023-03-22T14:06:23.037+0000 7fca8c50a700  1 osd.8 pg_epoch: 742 pg[6.18( v 733'4562 (0'0,733'4562] local-lis/les=629/630 n=5 ec=112/103 lis/c=740/731 les/c/f=741/732/0 sis=742) [17,10,8] r=2 lpr>
Mar 22 14:06:23 c06 bash[3018]: debug 2023-03-22T14:06:23.037+0000 7fca8cd0b700 -1 log_channel(cluster) log [ERR] : 9.1c has 2 objects unfound and apparently lost

```





### ceph host 제거

#### 1. osd tree 확인

```
# ceph osd tree  <<=== 찌꺼기 있으면 Crush에서 제거

# cehp osd out osd.1
# ceph osd stop osd.1
# ceph osd rm osd.1
# ceph osd tree  <<=== 찌꺼기 있으면 Crush에서 제거
# cehp osd crush remove osd.1
# ceph osd destroy 
```

#### 2. host drain: host rm 

```
# ceph orch ps | grep c06
# ceph orch ps | grep c06
osd.1              c06                stopped           9s ago   5h    
osd.2              c06                stopped           9s ago   5h    
osd.3              c06                stopped           9s ago   5h    
osd.8              c06                error             9s ago   5h    

# ceph  orch daemon rm osd.1 --force
# ceph orch host drain c06

# ceph orch host rm c06
Error EINVAL: Not allowed to remove c06 from cluster. The following daemons are running in the host:
type                 id
-------------------- ---------------
osd                  8
# ceph orch daemon rm osd.8 --force
Removed osd.8 from host 'c06'

root@c01:/# ceph orch host rm c06
Removed  host 'c06'
root@c01:/# ceph orch host ls
HOST  ADDR            LABELS  STATUS
c01   192.168.105.51  _admin
c02   192.168.105.52  _admin
c03   192.168.105.53  _admin
c04   192.168.105.54  _admin
c05   192.168.105.55  _admin
5 hosts in cluster

# ceph  osd  crush remove c06
removed item id -3 name 'c06' from crush map
```





#### 3. host c05 제거

```
# ceph orch ps  | grep  c05
# ceph orch daemon rm osd.2 --force

# ceph osd out osd.2
# ceph osd down osd.2
# ceph osd tree
# ceph osd rm osd.2
# ceph osd tree  # 혹시 DNE do not exists 상태
# ceph osd crush remove osd.2
# ceph osd crush remove c05


# ceph orch host drain c05
# ceph orch ps | grep c05
# ceph orch host rm c05
```



* host add 

```
# ssh-copy-id -f -i /etc/ceph/ceph.pub root@c05
# ceph orch host add c05 192.168.105.55
# ceph orch host ls
# ceph orch device ls
# ceph osd unset noout

# ceph osd unset nobackfill
# ceph osd unset norecover
# ceph osd unset norebalance
# ceph osd unset noscrub
# ceph osd unset nodeep-scrub

# ceph orch daemon add osd c05:/dev/sdb
# ceph orch daemon add osd c05:/dev/sdc
# ceph orch daemon add osd c05:/dev/sdd
# ceph orch daemon add osd c05:/dev/sde
```



```
root@c01:/# ceph orch daemon add  osd c02:/dev/sdb
Error EINVAL: Traceback (most recent call last):
  File "/usr/share/ceph/mgr/mgr_module.py", line 1730, in _handle_command
    return self.handle_command(inbuf, cmd)
  File "/usr/share/ceph/mgr/orchestrator/_interface.py", line 171, in handle_command
    return dispatch[cmd['prefix']].call(self, cmd, inbuf)
  File "/usr/share/ceph/mgr/mgr_module.py", line 462, in call
    return self.func(mgr, **kwargs)
  File "/usr/share/ceph/mgr/orchestrator/_interface.py", line 107, in <lambda>
    wrapper_copy = lambda *l_args, **l_kwargs: wrapper(*l_args, **l_kwargs)  # noqa: E731
  File "/usr/share/ceph/mgr/orchestrator/_interface.py", line 96, in wrapper
    return func(*args, **kwargs)
  File "/usr/share/ceph/mgr/orchestrator/module.py", line 830, in _daemon_add_osd
    raise_if_exception(completion)
  File "/usr/share/ceph/mgr/orchestrator/_interface.py", line 228, in raise_if_exception
    raise e
mgr_module.MonCommandFailed: auth get failed: failed to find osd.4 in keyring retval: -2

```



```
# ceph orch client-keyring set client.admin "jerryp-ceph-01 jerryp-ceph-02 jerryp-ceph-03 jerryp-ceph-04 jerryp-ceph-05 jerryp-ceph-06"


```



#### 4. host c04 제거

```
# ceph orch ps  | grep  c04
# ceph orch daemon rm osd.2 --force

# ceph osd out osd.2
# ceph osd down osd.2
# ceph osd tree
# ceph osd rm osd.2
# ceph osd tree  # 혹시 DNE do not exists 상태
# ceph osd crush remove osd.2
# ceph osd crush remove c05


# ceph orch host drain c05
# ceph orch ps | grep c05
# ceph orch host rm c05
```



## Object 위치를 찾아서 손상해 보자

#### 1. text 파일 생성

```
root@c01:~# cat AAA.txt
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
```

#### 2. S3 upload

```
# aws s3 cp  ./AAA.txt s3://mybucket/AAA  --endpoint-url http://192.168.105.51:80
upload: ./AAA.txt to s3://mybucket/AAA
```



#### 3. File 저장된 object  식별한다. 

```
root@c01:/# rados --pool default.rgw.buckets.data ls

d1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA

```

#### 4. PG 식별 : 8.1f

```
# ceph  osd  map  default.rgw.buckets.data d1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA
osdmap e276 pool 'default.rgw.buckets.data' (8) object 'd1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA' -> pg 8.4c785a7f (8.1f) -> up ([4,13,5], p4) actin
```

* osd.4에 할당된것 확인.



#### 4. host의 osd 확인

```
root@c02:~# lsscsi
[2:0:0:0]    cd/dvd  NECVMWar VMware SATA CD00 1.00  /dev/sr0
[32:0:0:0]   disk    VMware   Virtual disk     2.0   /dev/sda
[32:0:1:0]   disk    VMware   Virtual disk     2.0   /dev/sdb  <<---- osd.4
[32:0:2:0]   disk    VMware   Virtual disk     2.0   /dev/sdc
[32:0:3:0]   disk    VMware   Virtual disk     2.0   /dev/sdd
[32:0:4:0]   disk    VMware   Virtual disk     2.0   /dev/sde
```



```
root@c02:~# lvs
  LV                                             VG                                        Attr      
  osd-block-2b1c291f-56be-4912-8fcf-233ca1072627 ceph-1386b6b8-1b67-4ac2-83af-2715def0a81b 
  osd-block-d5af417d-f345-416d-a31a-0acf7397d9cc ceph-14c977f8-80e0-4c1f-9894-707118e99fb0 
  osd-block-5113df5b-dcc8-413f-b305-9beb13154218 ceph-5d737799-c7c8-43d9-bdb9-e40a854ff63d  <<----
  osd-block-8c81470f-87ac-4fbc-a7d4-8bdeeecf7519 ceph-a39c81b7-bd17-4768-a54d-13ec677bc2b2 
```



#### 5. hexedit

```
#hexedit /dev/ceph-5d737799-c7c8-43d9-bdb9-e40a854ff63d/osd-block-5113df5b-dcc8-413f-b305-9beb13154218

210F6FD8   00 00 00 00  00 00 00 00  00 00 00 00  00 00 00 00  ................
210F6FE8   00 00 00 00  00 00 00 00  00 00 00 00  00 00 00 00  ................
210F6FF8   00 00 00 00  00 00 00 00  41 41 41 41  41 41 41 41  ........AAAAAAAA
210F7008   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAA
210F7018   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAA
210F7028   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAA
210F7038   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAA
210F7048   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAA
210F7058   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAA
210F7068   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAA
210F7078   41 41 41 41  41 41 0A 0A  00 00 00 00  00 00 00 00  AAAAAA..........
210F7088   00 00 00 00  00 00 00 00  00 00 00 00  00 00 00 00  ................
210F7098   00 00 00 00  00 00 00 00  00 00 00 00  00 00 00 00  ................
```

* 시작위치는 0x210F7000 
* AAAAA => ABCDEFG로 수정후 저장

```
#hexedit /dev/ceph-5d737799-c7c8-43d9-bdb9-e40a854ff63d/osd-block-5113df5b-dcc8-413f-b305-9beb13154218
F4 눌러서 해당 위치로 다시 가서 정상적으로 기록 되어 있는지 확인
```





로그를 걸고.. : 아무런 오류가 없네...

```
root@c02:~# journalctl -f -u ceph-44f7f3b7-04b4-2af4-97ca-10275fad3d87@osd.4.service
-- Logs begin at Fri 2023-03-17 12:41:24 UTC. --
Mar 24 09:14:39 c02 bash[347283]: Uptime(secs): 8401.3 total, 0.0 interval
Mar 24 09:14:39 c02 bash[347283]: Flush(GB): cumulative 0.000, interval 0.000
Mar 24 09:14:39 c02 bash[347283]: AddFile(GB): cumulative 0.000, interval 0.000
Mar 24 09:14:39 c02 bash[347283]: AddFile(Total Files): cumulative 0, interval 0
Mar 24 09:14:39 c02 bash[347283]: AddFile(L0 Files): cumulative 0, interval 0
Mar 24 09:14:39 c02 bash[347283]: AddFile(Keys): cumulative 0, interval 0
Mar 24 09:14:39 c02 bash[347283]: Cumulative compaction: 0.00 GB write, 0.00 MB/s write, 0.00 GB read, 0.00 MB/s read, 0.0 seconds
Mar 24 09:14:39 c02 bash[347283]: Interval compaction: 0.00 GB write, 0.00 MB/s write, 0.00 GB read, 0.00 MB/s read, 0.0 seconds
Mar 24 09:14:39 c02 bash[347283]: Stalls(count): 0 level0_slowdown, 0 level0_slowdown_with_compaction, 0 level0_numfiles, 0 level0_numfiles_with_compaction, 0 stop for pending_compaction_bytes, 0 slowdown for pending_compaction_bytes, 0 memtable_compaction, 0 memtable_slowdown, interval 0 total count
Mar 24 09:14:39 c02 bash[347283]: ** File Read Latency Histogram By Level [P] **

```

#### 6. OSD scrub 실행 

```
Mar 24 09:21:11 c02 bash[347283]: debug 2023-03-24T09:21:11.232+0000 7f156bd4e700  0 log_channel(cluster) log [DBG] : 4.8 scrub ok
Mar 24 09:21:14 c02 bash[347283]: debug 2023-03-24T09:21:14.256+0000 7f156cd50700  0 log_channel(cluster) log [DBG] : 4.1f scrub ok
Mar 24 09:21:17 c02 bash[347283]: debug 2023-03-24T09:21:17.228+0000 7f156bd4e700  0 log_channel(cluster) log [DBG] : 4.1c scrub ok
Mar 24 09:21:21 c02 bash[347283]: debug 2023-03-24T09:21:21.336+0000 7f156cd50700  0 log_channel(cluster) log [DBG] : 3.1a scrub ok
Mar 24 09:21:27 c02 bash[347283]: debug 2023-03-24T09:21:27.391+0000 7f156c54f700  0 log_channel(cluster) log [DBG] : 7.16 scrub ok
Mar 24 09:21:29 c02 bash[347283]: debug 2023-03-24T09:21:29.467+0000 7f156cd50700  0 log_channel(cluster) log [DBG] : 8.1f scrub ok   <<==== 여기에는 문제가 있다고 나와야 할 것 같은데... 아직도 정상으로 보임.
```

* 이것은 쓸모가 없군.

#### 7. Deep Scrub 실행 : (fsck 처럼 발견)

* ceph -w 로그에는 아래와 같은 오류 발생

```
2023-03-24T09:22:45.437919+0000 osd.4 [ERR] 8.1f shard 4 soid 8:fe5a1e32:::d1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA:head : candidate had a read error
2023-03-24T09:22:45.438481+0000 osd.4 [ERR] 8.1f deep-scrub 0 missing, 1 inconsistent objects
2023-03-24T09:22:45.438484+0000 osd.4 [ERR] 8.1f deep-scrub 1 errors
2023-03-24T09:22:48.994574+0000 mon.c01 [ERR] Health check failed: 1 scrub errors (OSD_SCRUB_ERRORS)
2023-03-24T09:22:48.994598+0000 mon.c01 [ERR] Health check failed: Possible data damage: 1 pg inconsistent (PG_DAMAGED)
```

* osd.8 로그에는  checksum이 틀리다고 오류 나옴. device location [0x210f7000~1000], <<= 수정한 위치 정보
* bluestore(/var/lib/ceph/osd/ceph-4) _verify_csum bad crc32c/0x1000 
* checksum at blob offset 0x0, got 0x2e31c8b9, expected 0xa483cd1a,   <<-- bad checksum 
* device location [0x210f7000~1000],   <<-- 문제가 있는 위치
* logical extent 0x0~1000, 
* object #8:fe5a1e32:::d1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA:head# <<-- object 정보

```
Mar 24 09:22:45 c02 bash[347283]: debug 2023-03-24T09:22:45.433+0000 7f156cd50700 -1 bluestore(/var/lib/ceph/osd/ceph-4) _verify_csum bad crc32c/0x1000 checksum at blob offset 0x0, got 0x2e31c8b9, expected 0xa483cd1a, device location [0x210f7000~1000], logical extent 0x0~1000, object #8:fe5a1e32:::d1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA:head#

Mar 24 09:22:45 c02 bash[347283]: debug 2023-03-24T09:22:45.433+0000 7f156cd50700 -1 
bluestore(/var/lib/ceph/osd/ceph-4) _verify_csum bad crc32c/0x1000 checksum at blob offset 0x0, got 0x2e31c8b9, expected 0xa483cd1a, device location [0x210f7000~1000], logical extent 0x0~1000, object #8:fe5a1e32:::d1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA:head#

Mar 24 09:22:45 c02 bash[347283]: debug 2023-03-24T09:22:45.433+0000 7f156cd50700 -1 bluestore(/var/lib/ceph/osd/ceph-4) _verify_csum bad crc32c/0x1000 checksum at blob offset 0x0, got 0x2e31c8b9, expected 0xa483cd1a, device location [0x210f7000~1000], logical extent 0x0~1000, object #8:fe5a1e32:::d1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA:head#

Mar 24 09:22:45 c02 bash[347283]: debug 2023-03-24T09:22:45.433+0000 7f156cd50700 -1 bluestore(/var/lib/ceph/osd/ceph-4) _verify_csum bad crc32c/0x1000 checksum at blob offset 0x0, got 0x2e31c8b9, expected 0xa483cd1a, device location [0x210f7000~1000], logical extent 0x0~1000, object #8:fe5a1e32:::d1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA:head#

Mar 24 09:22:45 c02 bash[347283]: debug 2023-03-24T09:22:45.433+0000 7f156cd50700 -1 log_channel(cluster) log [ERR] : 8.1f shard 4 soid 8:fe5a1e32:::d1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA:head : candidate had a read error

Mar 24 09:22:45 c02 bash[347283]: debug 2023-03-24T09:22:45.433+0000 7f156cd50700 -1 log_channel(cluster) log [ERR] : 8.1f deep-scrub 0 missing, 1 inconsistent objects
Mar 24 09:22:45 c02 bash[347283]: debug 2023-03-24T09:22:45.433+0000 7f156cd50700 -1 log_channel(cluster) log [ERR] : 8.1f deep-scrub 1 errors

```



#### 8. OSD 가서 복구 되었는지 확인

```
아직 복구는 안되었네.
```



````
2023-03-24T09:30:00.000180+0000 mon.c01 [ERR] [ERR] OSD_SCRUB_ERRORS: 1 scrub errors
2023-03-24T09:30:00.000184+0000 mon.c01 [ERR] [ERR] PG_DAMAGED: Possible data damage: 1 pg inconsistent
2023-03-24T09:30:00.000189+0000 mon.c01 [ERR]     pg 8.1f is active+clean+inconsistent, acting [4,13,5]
2023-03-24T09:40:00.000125+0000 mon.c01 [ERR] overall HEALTH_ERR 1 failed cephadm daemon(s); 1 scrub errors; Possible data damage: 1 pg inconsistent
2023-03-24T09:50:00.000113+0000 mon.c01 [ERR] overall HEALTH_ERR 1 failed cephadm daemon(s); 1 scrub errors; Possible data damage: 1 pg inconsistent
````



#### 9. 이때 해당 object를 다운로드 받아 보면 

* ceph -w 로그. 

```
2023-03-24T09:52:09.491203+0000 osd.4 [ERR] 8.1f missing primary copy of 8:fe5a1e32:::d1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA:head, will try copies on 5,13
```

* 로그를 찍은 소소코드

```
bool PrimaryLogPG::primary_error(
  const hobject_t& soid, eversion_t v)
{
  recovery_state.force_object_missing(pg_whoami, soid, v);
  bool uhoh = recovery_state.get_missing_loc().is_unfound(soid);
  if (uhoh)
    osd->clog->error() << info.pgid << " missing primary copy of " << soid << ", unfound";
  else
    osd->clog->error() << info.pgid << " missing primary copy of " << soid << ", will try copies on "
		       << recovery_state.get_missing_loc().get_locations(soid);
  return uhoh;
}
```

* osd.4 로그

```
Mar 24 09:52:09 c02 bash[347283]: debug 2023-03-24T09:52:09.487+0000 7f156cd50700 -1 bluestore(/var/lib/ceph/osd/ceph-4) _verify_csum bad crc32c/0x1000 checksum at blob offset 0x0, got 0x2e31c8b9, expected 0xa483cd1a, device location [0x210f7000~1000], logical extent 0x0~1000, object #8:fe5a1e32:::d1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA:head#
Mar 24 09:52:09 c02 bash[347283]: debug 2023-03-24T09:52:09.487+0000 7f156cd50700 -1 bluestore(/var/lib/ceph/osd/ceph-4) _verify_csum bad crc32c/0x1000 checksum at blob offset 0x0, got 0x2e31c8b9, expected 0xa483cd1a, device location [0x210f7000~1000], logical extent 0x0~1000, object #8:fe5a1e32:::d1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA:head#
Mar 24 09:52:09 c02 bash[347283]: debug 2023-03-24T09:52:09.487+0000 7f156cd50700 -1 bluestore(/var/lib/ceph/osd/ceph-4) _verify_csum bad crc32c/0x1000 checksum at blob offset 0x0, got 0x2e31c8b9, expected 0xa483cd1a, device location [0x210f7000~1000], logical extent 0x0~1000, object #8:fe5a1e32:::d1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA:head#
Mar 24 09:52:09 c02 bash[347283]: debug 2023-03-24T09:52:09.487+0000 7f156cd50700 -1 bluestore(/var/lib/ceph/osd/ceph-4) _verify_csum bad crc32c/0x1000 checksum at blob offset 0x0, got 0x2e31c8b9, expected 0xa483cd1a, device location [0x210f7000~1000], logical extent 0x0~1000, object #8:fe5a1e32:::d1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA:head#
Mar 24 09:52:09 c02 bash[347283]: debug 2023-03-24T09:52:09.487+0000 7f156cd50700 -1 log_channel(cluster) log [ERR] : 8.1f missing primary copy of 8:fe5a1e32:::d1e8ec50-d0ea-4b5f-a12d-04903859e2a3.14815.1_AAA:head, will try copies on 5,13
```

* 다운로드 받은 내용은 정상적으로 다운로드 됨

```
# cat AAA.txt
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
```



#### 9. pg repair

```
root@c01:/# ceph pg repair 8.1f
instructing pg 8.1f on osd.4 to repair
```

* osd 로그에는 수리 되었다고 나옴. 

```
Mar 24 10:03:37 c02 bash[347283]: debug 2023-03-24T10:03:37.967+0000 7f156cd50700  0 log_channel(cluster) log [DBG] : 8.1f repair ok, 0 fixed

```

* 그러나 hexedit 해보면 아직도 예전의 내용임.

```
210F7000   41 42 43 44  45 46 47 41  41 41 41 41  41 41 41 41  ABCDEFGAAAAAAAAA
210F7010   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAA
210F7020   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAA
210F7030   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAA
210F7040   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAA
210F7050   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAA
210F7060   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAA
210F7070   41 41 41 41  41 41 41 41  41 41 41 41  41 41 0A 0A  AAAAAAAAAAAAAA..
210F7080   00 00 00 00  00 00 00 00  00 00 00 00  00 00 00 00  ................
210F7090   00 00 00 00  00 00 00 00  00 00 00 00  00 00 00 00  ................
```

* repair 되었다고 했는데 아직 그대로 인 이유는???



```
root@c01:/# ceph pg map 8.1f
osdmap e276 pg 8.1f (8.1f) -> up [4,13,5] acting [4,13,5]
root@c01:/# ceph pg repair 8.1f
instructing pg 8.1f on osd.4 to repair
root@c01:/# ceph pg repair 8.1f
instructing pg 8.1f on osd.4 to repair
root@c01:/# ceph pg force-recovery  8.1f
pg 8.1f doesn't require recovery;
```



* 아 다른 위치로 복구를 했구나. 기존 데이터는 그대로 두고...

````/ㄱ/
210F7FD0   00 00 00 00  00 00 00 00  00 00 00 00  00 00 00 00  00 00 00 00  ....................
210F7FE4   00 00 00 00  00 00 00 00  00 00 00 00  00 00 00 00  00 00 00 00  ....................
210F7FF8   00 00 00 00  00 00 00 00  41 41 41 41  41 41 41 41  41 41 41 41  ........AAAAAAAAAAAA
210F800C   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAAAAAA
210F8020   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAAAAAA
210F8034   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAAAAAA
210F8048   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAAAAAA
210F805C   41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  41 41 41 41  AAAAAAAAAAAAAAAAAAAA
210F8070   41 41 41 41  41 41 41 41  41 41 41 41  41 41 0A 0A  00 00 00 00  AAAAAAAAAAAAAA......
210F8084   00 00 00 00  00 00 00 00  00 00 00 00  00 00 00 00  00 00 00 00  ....................
````

* 0x210F7000  ==>> 0x210F8000 으로 위치를 변경 했음  : 4MB 점프 해서 기록 했군.
* 
