

## linux Host ip config
```
: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000    
    inet 127.0.0.1/8 scope host lo       
    inet6 ::1/128 scope host        
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000    
    inet 10.10.0.51/24 brd 10.10.0.255 scope global noprefixroute enp0s3               
3: enp0s8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:f5:ab:d5 brd ff:ff:ff:ff:ff:ff
4: enp0s9: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000    
    inet 192.168.0.24/24 brd 192.168.0.255 scope global dynamic noprefixroute enp0s9    
5: enp0s10: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000    
    inet 192.168.0.23/24 brd 192.168.0.255 scope global dynamic noprefixroute enp0s10    
6: ovs-system: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 66:aa:cf:fd:52:dc brd ff:ff:ff:ff:ff:ff
7: br-ex: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 62:f3:4f:28:a8:4d brd ff:ff:ff:ff:ff:ff
8: br-tun: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether a6:9e:93:2a:7b:4d brd ff:ff:ff:ff:ff:ff
12: br-int: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether ce:3e:78:17:3d:4b brd ff:ff:ff:ff:ff:ff
17: virbr0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
    link/ether 52:54:00:00:6e:cc brd ff:ff:ff:ff:ff:ff
    inet 192.168.122.1/24 brd 192.168.122.255 scope global virbr0
       valid_lft forever preferred_lft forever
```

```sh
    Overview Subnets Ports DHCP Agents
Name : public
ID : 1169aaf8-75f2-45cb-81af-799a22ac86a8
Project ID : cf778a2ca4f845dbaa35e7a4a9c12487
Status : Active
Admin State UP
Shared No
External Network Yes
MTU 1500
Provider Network
    Network Type: flat
    Physical Network: public
    Segmentation ID: -
```