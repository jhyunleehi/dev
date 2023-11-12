#### command

```
$ lspci
$ lshw -C network
$ modinfo iwlwifi | grep iwlwifi-cc
$ dmesg | grep iwlwifi | grep version
$ modinfo iwlwifi | grep iwlwifi-cc
firmware:       iwlwifi-cc-a0-77.ucode
$ ip link 
$ rfkill list
$ apt install ethtool
$ ip a
$ ethtool -i enp1s0
```

[    2.218075] iwlwifi: probe of 0000:02:00.0 failed with error -110 <<<------ 이렇게 올라오지 않음 

#### 성공의 경우

```
$lsmod | grep iwlwifi
iwlwifi 331776 1 iwlmvm
cfg80211 708608 3 iwlmvm,iwlwifi,mac80211
```