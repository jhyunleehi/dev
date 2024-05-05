## NFS

### server 

```sh
$ lsmod | grep nfs
$ sudo apt install nfs-kernel-server
$ lsmod | grep nfs
nfsd                  839680  3
auth_rpcgss           184320  1 nfsd
nfs_acl                12288  1 nfsd
lockd                 143360  1 nfsd
grace                  12288  2 nfsd,lockd
sunrpc                811008  16 nfsd,auth_rpcgss,lockd,nfs_acl

$ sudo  mkdir -p /mnt/nfs_share
$ sudo  chown -R nobody:nogroup  /mnt/nfs_share/
$ sudo chmod  777 /mnt/nfs_share/
```

```sh
$ sudo vi  /etc/exports
/mnt/nfs_share  192.168.0.0/24(rw,sync,no_subtree_check)
```
* rw: Stands for Read/Write.
* sync: Requires changes to be written to the disk before they are applied.
* No_subtree_check: Eliminates subtree checkin

```sh
$ sudo exportfs -a
$ sudo systemctl restart nfs-kernel-server
$ sudo ufw allow from 192.168.0.0/24 to any port nfs

$ sudo ufw enable
Firewall is active and enabled on system startup
$ sudo ufw status
Status: active

To                         Action      From
--                         ------      ----
2049                       ALLOW       192.168.0.0/24
```

### client

```sh
$ sudo apt update
$ lsmod | grep nfs

$ sudo apt install nfs-common
$ sudo mkdir -p /mnt/nfs_clientshare
$ sudo mount 192.168.0.25:/mnt/nfs_share /mnt/nfs_clientshare

192.168.0.25:/mnt/nfs_share  50770432 14734848  33423872  31% /mnt/nfs_clientshare

$ modprobe nfs
$ lsmod  | grep nfs
nfsv4                1114112  1
nfs                   581632  2 nfsv4
lockd                 143360  1 nfs
fscache               389120  1 nfs
netfs                  61440  2 fscache,nfs
sunrpc                811008  9 nfsv4,auth_rpcgss,lockd,rpcsec_gss_krb5,nfs
```
