

#### can not connect  http://192.168.137.103/identity

```log
INFO keystone.cmd.bootstrap [None req-dd65f5a4-f5a2-4cba-9565-c9c5273f6041 None None] Created region RegionOne
INFO keystone.cmd.bootstrap [None req-dd65f5a4-f5a2-4cba-9565-c9c5273f6041 None None] Created public endpoint http://192.168.137.103/identity
+lib/keystone:bootstrap_keystone:615       '[' False == True ']'
+./stack.sh:main:1133                      create_keystone_accounts
+lib/keystone:create_keystone_accounts:314  local admin_project
++lib/keystone:create_keystone_accounts:315  oscwrap project show admin -f value -c id
++functions-common:oscwrap:2442             local xtrace
+++functions-common:oscwrap:2443             set +o
+++functions-common:oscwrap:2443             grep xtrace
++functions-common:oscwrap:2443             xtrace='set -o xtrace'
++functions-common:oscwrap:2444             set +o xtrace
Failed to discover available identity versions when contacting http://192.168.137.103/identity. Attempting to parse version from URL.
Could not find versioned identity endpoints when attempting to authenticate. Please check that your auth_url is correct. Bad Gateway (HTTP 502)
++functions-common:oscwrap:2461             return 1
+lib/keystone:create_keystone_accounts:315  admin_project=
+lib/keystone:create_keystone_accounts:1   exit_trap
+./stack.sh:exit_trap:549                  local r=1
++./stack.sh:exit_trap:550                  jobs -p
+./stack.sh:exit_trap:550                  jobs=263657
+./stack.sh:exit_trap:553                  [[ -n 263657 ]]
+./stack.sh:exit_trap:553                  [[ -n '' ]]
+./stack.sh:exit_trap:559                  '[' -f /tmp/tmp.yLt0jLhEXg ']'
+./stack.sh:exit_trap:560                  rm /tmp/tmp.yLt0jLhEXg
+./stack.sh:exit_trap:564                  kill_spinner
+./stack.sh:kill_spinner:459               '[' '!' -z '' ']'
+./stack.sh:exit_trap:566                  [[ 1 -ne 0 ]]
+./stack.sh:exit_trap:567                  echo 'Error on exit'
Error on exit
+./stack.sh:exit_trap:569                  type -p generate-subunit
+./stack.sh:exit_trap:570                  generate-subunit 1714643738 194 fail
+./stack.sh:exit_trap:572                  [[ -z /opt/stack/logs ]]
+./stack.sh:exit_trap:575                  /opt/stack/data/venv/bin/python3 /opt/stack/devstack/tools/worlddump.py -d /opt/stack/logs
World dumping... see /opt/stack/logs/worlddump-2024-05-02-095852.txt for details
+./stack.sh:exit_trap:584                  exit 1
stack@gpu-3:~/devstack$ 
```

==> localaddress  no proxy 

```sh
stack@gpu-3:~$ env | grep proxy
no_proxy=localhost,127.0.0.1,code.sdsdev.co.kr,sds.redii.net,70.150.192.213,70.150.192.231,70.150.192.119,70.60.31.37,70.60.31.51,70.70.202.215,70.60.31.41,70.60.31.47,70.60.31.38,70.60.31.36,70.60.31.72,70.60.31.104,70.60.31.106,70.60.192.231,70.60.31.103,70.150.192.201,70.150.192.202,192.168.137.3,192.168.137.101,192.168.137.102,192.168.137.103,192.168.137.104,192.168.137.105
ftp_proxy=http://70.10.15.10:8080
https_proxy=http://70.10.15.10:8080
http_proxy=http://70.10.15.10:8080

```