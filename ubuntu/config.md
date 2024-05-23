
### lsblk 
```sh
root@V11:~# lsblk -a  -o name,maj:min,rm,size,ro,type,mountpoint
NAME        MAJ:MIN RM   SIZE RO TYPE MOUNTPOINT
loop0         7:0    0    62M  1 loop /snap/core20/1587
loop1         7:1    0 400.8M  1 loop /snap/gnome-3-38-2004/112
loop2         7:2    0     4K  1 loop /snap/bare/5
loop3         7:3    0  91.7M  1 loop /snap/gtk-common-themes/1535
loop4         7:4    0 163.3M  1 loop /snap/firefox/1635
loop5         7:5    0  45.9M  1 loop /snap/snap-store/582
loop6         7:6    0    47M  1 loop /snap/snapd/16292
loop7         7:7    0   284K  1 loop /snap/snapd-desktop-integration/14
loop8         7:8    0  63.9M  1 loop /snap/core20/2318
loop9         7:9    0     0B  0 loop 
sda           8:0    0 119.2G  0 disk 
├─sda1        8:1    0   512M  0 part /boot/efi
└─sda2        8:2    0 118.7G  0 part /
mmcblk1     179:0    0 119.7G  0 disk 
└─mmcblk1p1 179:1    0 119.7G  0 part 
root@V11:~# lsblk -S
NAME HCTL       TYPE VENDOR   MODEL                                  REV SERIAL               TRAN
sda  0:0:0:0    disk ATA      LITEONIT LJT-128L6G-11 M.2 2260 128GB 10C  TW0Y48CM5508545D9026 sata
```

### 노트북 덮게 닫아도 전원 유지

systemd 전원 설정방법
Ubuntu 20.04 설치된 노트북에서 덮개가 닫힌 상태에서도 전원을 유지해도록 설정한다. 터미널 창에서 /etc/systemd/logind.conf 파일을 sudo 권한으로 수정한다.

/etc/systemd/logind.conf
```
[Login]
HandleLidSwitch=ignore
```
systemctl restart systemd-logind.service
```
jhyunlee@snote:~/code/dev$ systemctl | grep login
  systemd-logind.service                                                                                     loaded active running   User Login Management
```