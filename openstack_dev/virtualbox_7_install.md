
#  Install virtualbox in ubuntu 22.04 and fixing kernel driver not installed, EFI secure boot enabled

https://www.youtube.com/watch?v=Xmz27Ldt3dc

#### 1. 설치
```sh
$ wget https://download.virtualbox.org/virtualbox/7.0.16/virtualbox-7.0_7.0.16-162802~Ubuntu~jammy_amd64.deb --no-check-certificate
--2024-04-27 08:50:50--  https://download.virtualbox.org/virtualbox/7.0.16/virtualbox-7.0_7.0.16-162802~Ubuntu~jammy_amd64.deb
Connecting to 70.10.15.10:8080... connected.
WARNING: cannot verify download.virtualbox.org's certificate, issued by ‘emailAddress=infosec@samsung.com,CN=SDS,O=SAMSUNG SDS,L=Gangman-gu,ST=Seoul,C=KR’:
  CA certificate key too weak

$ sudo apt remove virtualbox
$ sudo apt remove virtualbox-dkms
$ sudo apt autoremove 
$ sudo apt --fix-broken install

$ sudo dpkg -i  virtualbox-7.0_7.0.16-162802~Ubuntu~jammy_amd64.deb
==> password  12345678 입력
```
* 설치 로그
```log
There were problems setting up VirtualBox.  To re-start the set-up process, run
  /sbin/vboxconfig
as root.  If your system is using EFI Secure Boot you may need to sign the
kernel modules (vboxdrv, vboxnetflt, vboxnetadp, vboxpci) before you can load
them. Please see your Linux system's documentation for more information.
Processing triggers for libc-bin (2.35-0ubuntu3.7) ...
```

#### 2. reboot EFI Secure Boot



```sh
$ wget https://www.virtualbox.org/download/oracle_vbox_2016.asc --no-check-certificate
$ cat oracle_vbox_2016.asc | gpg --dearmor | sudo tee /usr/share/keyrings/virtualbox.gpg > /dev/null 2>&1
$ sudo nano /etc/apt/sources.list.d/virtualbox.list
  deb [arch=amd64 signed-by=/usr/share/keyrings/oracle-virtualbox-2016.gpg] https://download.virtualbox.org/virtualbox/debian jammy contrib
$ sudo apt update
$ sudo apt install virtualbox-7.0

```

```sh
deb [arch=amd64 signed-by=/usr/share/keyrings/oracle-virtualbox-2016.gpg] https://download.virtualbox.org/virtualbox/debian <mydist> contrib

sudo gpg --yes --output /usr/share/keyrings/oracle-virtualbox-2016.gpg --dearmor oracle_vbox_2016.asc
wget -O- https://www.virtualbox.org/download/oracle_vbox_2016.asc | sudo gpg --yes --output /usr/share/keyrings/oracle-virtualbox-2016.gpg --dearmor
sudo apt-get update
sudo apt-get install virtualbox-6.1
```

#### refreshing package 
What to do when experiencing The following signatures were invalid: BADSIG ... when refreshing the packages from the repository?
```sh
# sudo -s -H
# apt-get clean
# rm /var/lib/apt/lists/*
# rm /var/lib/apt/lists/partial/*
# apt-get clean
# apt-get update
```

### mokutil 사용방법

```sh
$ mokutil  --list-enrolled
$ mokutil --export
gpu@gpu-3:~$ ls -l  MOK*.der
-rw-r--r-- 1 gpu gpu 1080 11월 22 06:07 MOK-0001.der
-rw-r--r-- 1 gpu gpu  919 11월 22 06:07 MOK-0002.der
-rw-r--r-- 1 gpu gpu  769 11월 22 06:07 MOK-0003.der
$ mokutil --delete MOK-0002.der
$ mokutil --delete MOK-0003.der

$ mokutil --delete /root/module-signing/MOK.der 
```


#### dpkg remove 
```sh
gpu@gpu-3:~$ dpkg -l | grep  virtual

rc  virtualbox                                 6.1.50-dfsg-1~ubuntu1.22.04.1           amd64       
ii  virtualbox-7.0                             7.0.16-162802~Ubuntu~jammy              amd64       

```

```sh
gpu@gpu-3:~$ sudo dpkg -r virtualbox
dpkg: 경고: ignoring request to remove virtualbox, only the config files of which are on the system; use --purge to remove them too

$ sudo dpkg  --purge virtualbox
(데이터베이스 읽는중 ...현재 239502개의 파일과 디렉터리가 설치되어 있습니다.)
Purging configuration files for virtualbox (6.1.50-dfsg-1~ubuntu1.22.04.1) ...

gpu@gpu-3:~$ sudo dpkg -r virtualbox-7.0
(데이터베이스 읽는중 ...현재 239500개의 파일과 디렉터리가 설치되어 있습니다.)
Removing virtualbox-7.0 (7.0.16-162802~Ubuntu~jammy) ...
debconf: DbDriver "config": /var/cache/debconf/config.dat is locked by another process: Resource temporarily unavailable
dpkg: error processing package virtualbox-7.0 (--remove):
 installed virtualbox-7.0 package pre-removal script subprocess returned error exit status 1


gpu@gpu-3:~$ sudo apt list | grep   virtualbox
virtualbox-7.0/now 7.0.16-162802~Ubuntu~jammy amd64 [installed,local]

==> 제거
gpu@gpu-3:~$ sudo apt remove virtualbox-7.0/now

```


#### gpu-4 virtual box 7 설치 이후
* 정상 설치된 경우 2 key가 생성된다. 
1. Subject: C=GB, ST=Isle of Man, L=Douglas, O=Canonical Ltd., CN=Canonical Ltd. Master Certificate Authority
2. Issuer: CN=gpu-4 Secure Boot Module Signature key

```sh
gpu@gpu-4:~$ mokutil  --list-enrolled
[key 1]
SHA1 Fingerprint: 76:a0:92:06:58:00:bf:37:69:01:c3:72:cd:55:a9:0e:1f:de:d2:e0
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            b9:41:24:a0:18:2c:92:67
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=GB, ST=Isle of Man, L=Douglas, O=Canonical Ltd., CN=Canonical Ltd. Master Certificate Authority
        Validity
            Not Before: Apr 12 11:12:51 2012 GMT
            Not After : Apr 11 11:12:51 2042 GMT
        Subject: C=GB, ST=Isle of Man, L=Douglas, O=Canonical Ltd., CN=Canonical Ltd. Master Certificate Authority
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
            X509v3 Authority Key Identifier:
                AD:91:99:0B:C2:2A:B1:F5:17:04:8C:23:B6:65:5A:26:8E:34:5A:63
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Key Usage:
                Digital Signature, Certificate Sign, CRL Sign
            X509v3 CRL Distribution Points:
                Full Name:
                  URI:http://www.canonical.com/secure-boot-master-ca.crl
    Signature Algorithm: sha256WithRSAEncryption
    Signature Value:    

[key 2]
SHA1 Fingerprint: 88:24:cb:24:c4:d6:c4:e6:69:db:1f:eb:d9:9b:48:d1:78:ac:d2:58
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            33:9c:20:c4:6f:3f:f0:d4:72:ce:33:d1:db:b8:f4:66:2e:cc:f0:56
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: CN=gpu-4 Secure Boot Module Signature key
        Validity
            Not Before: Nov 21 21:01:34 2023 GMT
            Not After : Oct 28 21:01:34 2123 GMT
        Subject: CN=gpu-4 Secure Boot Module Signature key
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)               
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Subject Key Identifier:
                8E:E3:BD:34:15:FC:A5:CF:9F:1D:52:2F:6C:F7:DB:90:32:45:79:D4
            X509v3 Authority Key Identifier:
                8E:E3:BD:34:15:FC:A5:CF:9F:1D:52:2F:6C:F7:DB:90:32:45:79:D4
            X509v3 Basic Constraints: critical
                CA:FALSE
            X509v3 Extended Key Usage:
                Code Signing, 1.3.6.1.4.1.2312.16.1.2
            Netscape Comment:
                OpenSSL Generated Certificate
    Signature Algorithm: sha256WithRSAEncryption    
gpu@gpu-4:~$
```



### GPU-2 재설치
#### 1. 현재상황파악
* 잘못 설치되어 있음 
* rc : remove and configfile 
* iU : instll and Unpacked 
```
$ apt list | grep virtual | grep  re
WARNING: apt does not have a stable CLI interface. Use with caution in scripts.
virtualbox/jammy-updates,now 6.1.50-dfsg-1~ubuntu1.22.04.1 amd64 [residual-config]

==> 여기도 꼬여 있음 
$ dpkg -l | grep virtual
rc  virtualbox                                 6.1.50-dfsg-1~ubuntu1.22.04.1           amd64        x86 virtualization solution - base binaries
iU  virtualbox-7.0                             7.0.16-162802~Ubuntu~jammy              amd64        Oracle VM VirtualBox

$ dpkg -l | grep virtual 
$ sudo apt remove  virtualbox-7.0
$ sudo dpkg  --purge virtualbox
$ sudo dpkg  --purge virtualbox-7.0
$ sudo apt --fix-broken install
$ sudo  dpkg -i virtualbox-7.0_7.0.16-162802~Ubuntu~jammy_amd64.deb
$ sudo apt --fix-broken install
==> mock password 12345678 입력
==> reboot
==> During the boot when prompted choose Enroll MOK
    You will see the keys that were created and signed and choose Continue
    Keyinput : 12345678
    reboot
$ mokutil  --list-enrolled
==> chek key list

$ dpkg -l | grep virtual
ii  virtualbox-7.0                             7.0.16-162802~Ubuntu~jammy              amd64        Oracle VM VirtualBox
==> check ii install statues

$ virtualbox 
==> ok
```





### 시행착오
```
gpu@gpu-2:~$ sudo  dpkg -i virtualbox-7.0_7.0.16-162802~Ubuntu~jammy_amd64.deb
Selecting previously unselected package virtualbox-7.0.
(데이터베이스 읽는중 ...현재 238810개의 파일과 디렉터리가 설치되어 있습니다.)
Preparing to unpack virtualbox-7.0_7.0.16-162802~Ubuntu~jammy_amd64.deb ...
Unpacking virtualbox-7.0 (7.0.16-162802~Ubuntu~jammy) ...
dpkg: dependency problems prevent configuration of virtualbox-7.0:
 virtualbox-7.0 패키지는 다음 패키지에 의존: libqt5help5 (>= 5.15.1): 하지만:
  libqt5help5 패키지는 설치하지 않았습니다.
 virtualbox-7.0 패키지는 다음 패키지에 의존: libqt5xml5 (>= 5.0.2): 하지만:
  libqt5xml5 패키지는 설치하지 않았습니다.

dpkg: error processing package virtualbox-7.0 (--install):
 의존성 문제 - 설정하지 않고 남겨둠
Processing triggers for mailcap (3.70+nmu1ubuntu1) ...
Processing triggers for gnome-menus (3.36.0-1ubuntu3) ...
Processing triggers for desktop-file-utils (0.26-1ubuntu3) ...
Processing triggers for hicolor-icon-theme (0.17-2) ...
Processing triggers for shared-mime-info (2.1-2) ...
처리하는데 오류가 발생했습니다:
 virtualbox-7.0
gpu@gpu-2:~$ sudo apt --fix-broken install
패키지 목록을 읽는 중입니다... 완료
의존성 트리를 만드는 중입니다... 완료
상태 정보를 읽는 중입니다... 완료
의존성을 바로잡는 중입니다... 완료
다음 패키지가 자동으로 설치되었지만 더 이상 필요하지 않습니다:
  libgsoap-2.8.117 liblzf1 libsdl1.2debian
Use 'sudo apt autoremove' to remove them.
The following additional packages will be installed:
  libqt5help5 libqt5sql5 libqt5sql5-sqlite libqt5xml5
다음 새 패키지를 설치할 것입니다:
  libqt5help5 libqt5sql5 libqt5sql5-sqlite libqt5xml5
0개 업그레이드, 4개 새로 설치, 0개 제거 및 3개 업그레이드 안 함.
1개를 완전히 설치하지 못했거나 지움.
462 k바이트 아카이브를 받아야 합니다.
이 작업 후 1,793 k바이트의 디스크 공간을 더 사용하게 됩니다.
계속 하시겠습니까? [Y/n] y
```

* 설치는 되었으나 실행하면 
==> reboot 
```
gpu@gpu-1:~$ virtualbox
WARNING: The vboxdrv kernel module is not loaded. Either there is no module
         available for the current kernel (6.5.0-18-generic) or it failed to
         load. Please recompile the kernel module and install it by

           sudo /sbin/vboxconfig

         You will not be able to start VMs until this problem is fixed.
qgpu@gpu-1:~$ 
```

### virtualbox 설치 중에 gcc-12가 필요하네 ㅎㅎ
==> 설치 과정에서 cat /var/log/vbox-setup.log 로그를 참조하라고 함.
==> 그것을 보면... /bin/sh: 1: gcc-12: not found    
==> 그래서 gcc-12를 설치해줌 

```sh
gpu@gpu-3:~$ sudo apt install gdebi

gpu@gpu-3:~$ sudo gdebi virtualbox-7.0_7.0.16-162802~Ubuntu~jammy_amd64.deb
Reading package lists... Done
Building dependency tree... Done
Reading state information... Done
Reading state information... Done

Oracle VM VirtualBox
 VirtualBox is a powerful PC virtualization solution allowing you to run a
 wide range of PC operating systems on your Linux system. This includes
 Windows, Linux, FreeBSD, DOS, OpenBSD and others. VirtualBox comes with a broad
 feature set and excellent performance, making it the premier virtualization
 software solution on the market.
소프트웨어 패키지를 설치하시겠습니까? [y/N]:y
/usr/bin/gdebi:113: FutureWarning: Possible nested set at position 1
  c = findall("[[(](\S+)/\S+[])]", msg)[0].lower()
Selecting previously unselected package virtualbox-7.0.
(데이터베이스 읽는중 ...현재 219444개의 파일과 디렉터리가 설치되어 있습니다.)
Preparing to unpack virtualbox-7.0_7.0.16-162802~Ubuntu~jammy_amd64.deb ...
Unpacking virtualbox-7.0 (7.0.16-162802~Ubuntu~jammy) ...
virtualbox-7.0 (7.0.16-162802~Ubuntu~jammy) 설정하는 중입니다 ...
addgroup: The group `vboxusers' already exists as a system group. Exiting.
vboxdrv.sh: failed: Look at /var/log/vbox-setup.log to find out what went wrong.

There were problems setting up VirtualBox.  To re-start the set-up process, run
  /sbin/vboxconfig
as root.  If your system is using EFI Secure Boot you may need to sign the
kernel modules (vboxdrv, vboxnetflt, vboxnetadp, vboxpci) before you can load
them. Please see your Linux system's documentation for more information.
Processing triggers for mailcap (3.70+nmu1ubuntu1) ...
Processing triggers for gnome-menus (3.36.0-1ubuntu3) ...
Processing triggers for desktop-file-utils (0.26-1ubuntu3) ...
Processing triggers for hicolor-icon-theme (0.17-2) ...
Processing triggers for shared-mime-info (2.1-2) ...



gpu@gpu-3:~$ cat /var/log/vbox-setup.log
Building the main VirtualBox module.
Error building the module:
make V=1 CONFIG_MODULE_SIG= CONFIG_MODULE_SIG_ALL= -C /lib/modules/6.5.0-28-generic/build M=/tmp/vbox.0 SRCROOT=/tmp/vbox.0 -j12 modules
make[1]: warning: -j12 forced in submake: resetting jobserver mode.
warning: the compiler differs from the one used to build the kernel
  The kernel was built by: x86_64-linux-gnu-gcc-12 (Ubuntu 12.3.0-1ubuntu1~22.04) 12.3.0
  You are using:
make -f ./scripts/Makefile.build obj=/tmp/vbox.0 need-builtin=1 need-modorder=1
# cmd_mod /tmp/vbox.0/vboxdrv.mod
  printf '%s
'   linux/SUPDrv-linux.o SUPDrv.o SUPDrvGip.o SUPDrvSem.o SUPDrvTracer.o SUPLibAll.o common/string/strformatrt.o combined-agnostic1.o combined-agnostic2.o combined-os-specific.o | awk '!x[$0]++ { print("/tmp/vbox.0/"$0) }' > /tmp/vbox.0/vboxdrv.mod
# CC [M]  /tmp/vbox.0/linux/SUPDrv-linux.o
  gcc-12 -Wp,-MMD,/tmp/vbox.0/linux/.SUPDrv-linux.o.d -nostdinc -I./arch/x86/include -I./arch/x86/include/generated  -I./include -I./arch/x86/include/uapi -I./arch/x86/include/generated/uapi -I./include/uapi -I./include/generated/uapi -include ./include/linux/compiler-version.h -include ./include/linux/kconfig.h -I./ubuntu/include -include ./include/linux/compiler_types.h -D__KERNEL__ -std=gnu11 -fshort-wchar -funsigned-char -fno-common -fno-PIE -fno-strict-aliasing -Wall -Wundef -Werror=implicit-function-declaration -Werror=implicit-int -Werror=return-type -Werror=strict-prototypes -Wno-format-security -Wno-trigraphs -mno-sse -mno-mmx -mno-sse2 -mno-3dnow -mno-avx -m64 -mno-80387 -mtune=generic -mno-red-zone -mcmodel=kernel -Wno-sign-compare -fno-asynchronous-unwind-tables -mfunction-return=thunk-extern -fno-jump-tables -mharden-sls=all -fpatchable-function-entry=16,16 -fno-delete-null-pointer-checks -O2 -Wframe-larger-than=1024 -fstack-protector-strong -Wno-main -fno-omit-frame-pointer -fno-optimize-sibling-calls -ftrivial-auto-var-init=zero -fzero-call-used-regs=used-gpr -pg -mrecord-mcount -falign-functions=16 -Wvla -Wno-pointer-sign -Wno-maybe-uninitialized -Wno-array-bounds -Wno-alloc-size-larger-than -Wimplicit-fallthrough=5 -fno-strict-overflow -fno-stack-check -fconserve-stack -Werror=date-time -g -gdwarf-5 -include /tmp/vbox.0/include/VBox/SUPDrvMangling.h -fno-omit-frame-pointer -fno-pie -Wno-declaration-after-statement -I./include -I/tmp/vbox.0/ -I/tmp/vbox.0/include -I/tmp/vbox.0/r0drv/linux -D__KERNEL__ -DMODULE -DRT_WITHOUT_PRAGMA_ONCE -DRT_OS_LINUX -DIN_RING0 -DIN_RT_R0 -DIN_SUP_R0 -DVBOX -DRT_WITH_VBOX -DVBOX_WITH_HARDENING -DSUPDRV_WITH_RELEASE_LOGGER -DVBOX_WITHOUT_EFLAGS_AC_SET_IN_VBOXDRV -DIPRT_WITHOUT_EFLAGS_AC_PRESERVING -DVBOX_WITH_64_BITS_GUESTS -DCONFIG_VBOXDRV_AS_MISC -DRT_ARCH_AMD64  -fsanitize=bounds-strict -fsanitize=shift -fsanitize=bool -fsanitize=enum  -DMODULE  -DKBUILD_BASENAME='"SUPDrv_linux"' -DKBUILD_MODNAME='"vboxdrv"' -D__KBUILD_MODNAME=kmod_vboxdrv -c -o /tmp/vbox.0/linux/SUPDrv-linux.o /tmp/vbox.0/linux/SUPDrv-linux.c   ; ./tools/objtool/objtool --hacks=jump_label --hacks=noinstr --hacks=skylake --retpoline --rethunk --sls --stackval --static-call --uaccess --prefix=16   --module /tmp/vbox.0/linux/SUPDrv-linux.o
/bin/sh: 1: gcc-12: not found
make[2]: *** [scripts/Makefile.build:251: /tmp/vbox.0/linux/SUPDrv-linux.o] Error 127
make[1]: *** [Makefile:2039: /tmp/vbox.0] Error 2
make: *** [/tmp/vbox.0/Makefile-footer.gmk:133: vboxdrv] Error 2
gpu@gpu-3:~$ gcc
gcc: fatal error: no input files
compilation terminated.

gpu@gpu-3:~$ apt install gcc
E: 잠금 파일 /var/lib/dpkg/lock-frontend 파일을 열 수 없습니다 - open (13: Permission denied)
E: Unable to acquire the dpkg frontend lock (/var/lib/dpkg/lock-frontend), are you root?
gpu@gpu-3:~$ sudo  apt install gcc

패키지 목록을 읽는 중입니다... 완료
의존성 트리를 만드는 중입니다... 완료
상태 정보를 읽는 중입니다... 완료
gcc is already the newest version (4:11.2.0-1ubuntu1).
gcc 패키지는 수동설치로 지정합니다.
0개 업그레이드, 0개 새로 설치, 0개 제거 및 3개 업그레이드 안 함.

gpu@gpu-3:~$ sudo  apt install gcc-12
패키지 목록을 읽는 중입니다... 완료
의존성 트리를 만드는 중입니다... 완료
상태 정보를 읽는 중입니다... 완료
The following additional packages will be installed:
  cpp-12 libasan8 libgcc-12-dev libtsan2
제안하는 패키지:
  gcc-12-locales cpp-12-doc gcc-12-multilib gcc-12-doc
다음 새 패키지를 설치할 것입니다:
  cpp-12 gcc-12 libasan8 libgcc-12-dev libtsan2
0개 업그레이드, 5개 새로 설치, 0개 제거 및 3개 업그레이드 안 함.
40.1 M바이트 아카이브를 받아야 합니다.
이 작업 후 138 M바이트의 디스크 공간을 더 사용하게 됩니다.
계속 하시겠습니까? [Y/n] y
받기:1 http://nexus.sdsdev.co.kr:8081/repository/mirror-proxy/ubuntu jammy-updates/main amd64 cpp-12 amd64 12.3.0-1ubuntu1~22.04 [10.8 MB]
받기:2 http://nexus.sdsdev.co.kr:8081/repository/mirror-proxy/ubuntu jammy-updates/main amd64 libasan8 amd64 12.3.0-1ubuntu1~22.04 [2,442 kB]
받기:3 http://nexus.sdsdev.co.kr:8081/repository/mirror-proxy/ubuntu jammy-updates/main amd64 libtsan2 amd64 12.3.0-1ubuntu1~22.04 [2,477 kB]
받기:4 http://nexus.sdsdev.co.kr:8081/repository/mirror-proxy/ubuntu jammy-updates/main amd64 libgcc-12-dev amd64 12.3.0-1ubuntu1~22.04 [2,618 kB]
받기:5 http://nexus.sdsdev.co.kr:8081/repository/mirror-proxy/ubuntu jammy-updates/main amd64 gcc-12 amd64 12.3.0-1ubuntu1~22.04 [21.7 MB]
내려받기 40.1 M바이트, 소요시간 1초 (31.1 M바이트/초)
Selecting previously unselected package cpp-12.
(데이터베이스 읽는중 ...현재 220183개의 파일과 디렉터리가 설치되어 있습니다.)
Preparing to unpack .../cpp-12_12.3.0-1ubuntu1~22.04_amd64.deb ...
Unpacking cpp-12 (12.3.0-1ubuntu1~22.04) ...
Selecting previously unselected package libasan8:amd64.
Preparing to unpack .../libasan8_12.3.0-1ubuntu1~22.04_amd64.deb ...
Unpacking libasan8:amd64 (12.3.0-1ubuntu1~22.04) ...
Selecting previously unselected package libtsan2:amd64.
Preparing to unpack .../libtsan2_12.3.0-1ubuntu1~22.04_amd64.deb ...
Unpacking libtsan2:amd64 (12.3.0-1ubuntu1~22.04) ...
Selecting previously unselected package libgcc-12-dev:amd64.
Preparing to unpack .../libgcc-12-dev_12.3.0-1ubuntu1~22.04_amd64.deb ...
Unpacking libgcc-12-dev:amd64 (12.3.0-1ubuntu1~22.04) ...
Selecting previously unselected package gcc-12.
Preparing to unpack .../gcc-12_12.3.0-1ubuntu1~22.04_amd64.deb ...
Unpacking gcc-12 (12.3.0-1ubuntu1~22.04) ...
cpp-12 (12.3.0-1ubuntu1~22.04) 설정하는 중입니다 ...
libasan8:amd64 (12.3.0-1ubuntu1~22.04) 설정하는 중입니다 ...
libtsan2:amd64 (12.3.0-1ubuntu1~22.04) 설정하는 중입니다 ...
libgcc-12-dev:amd64 (12.3.0-1ubuntu1~22.04) 설정하는 중입니다 ...
gcc-12 (12.3.0-1ubuntu1~22.04) 설정하는 중입니다 ...
Processing triggers for man-db (2.10.2-1) ...
Processing triggers for libc-bin (2.35-0ubuntu3.7) ...

gpu@gpu-3:~$ sudo gdebi virtualbox-7.0_7.0.16-162802~Ubuntu~jammy_amd64.deb
Reading package lists... Done
Building dependency tree... Done
Reading state information... Done
Reading state information... Done

Oracle VM VirtualBox
 VirtualBox is a powerful PC virtualization solution allowing you to run a
 wide range of PC operating systems on your Linux system. This includes
 Windows, Linux, FreeBSD, DOS, OpenBSD and others. VirtualBox comes with a broad
 feature set and excellent performance, making it the premier virtualization
 software solution on the market.
소프트웨어 패키지를 설치하시겠습니까? [y/N]:y
/usr/bin/gdebi:113: FutureWarning: Possible nested set at position 1
  c = findall("[[(](\S+)/\S+[])]", msg)[0].lower()
(데이터베이스 읽는중 ...현재 220425개의 파일과 디렉터리가 설치되어 있습니다.)
Preparing to unpack virtualbox-7.0_7.0.16-162802~Ubuntu~jammy_amd64.deb ...
Unpacking virtualbox-7.0 (7.0.16-162802~Ubuntu~jammy) over (7.0.16-162802~Ubuntu~jammy) ...
virtualbox-7.0 (7.0.16-162802~Ubuntu~jammy) 설정하는 중입니다 ...
addgroup: The group `vboxusers' already exists as a system group. Exiting.
Can't load /var/lib/shim-signed/mok/.rnd into RNG
40A7ABBD65710000:error:12000079:random number generator:RAND_load_file:Cannot open file:../crypto/rand/randfile.c:106:Filename=/var/lib/shim-signed/mok/.rnd
.+.+.....+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++*...+.................+.+...+......+.....+.........+.+...+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++*..+..........+.....+.........+.+......+........+....+...+.........+.....+...+......+.+.....+.+.........+...+..+..................+................+.........+.....+...+.......+...........+.+..............+....+..+......+.......+......+.....+....+......+.........+.....+......+.........+.+..+....+......+...+..............+...+...+.+.....+.+.....+....+...+...+.................+.+.................+..................+.........+.+.....+...................+.....+...+....+......+...+.........+..+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
......+....+..+.+........+.+......+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++*.....+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++*........+...............+.+..+.........+.+...+............+..+..........+........+....+...+...+............+.......................+...+.........+.+......+..............+.......+.....+.........+.+............+......+..+...+.......+.....+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
-----
vboxdrv.sh: failed: modprobe vboxdrv failed. Please use 'dmesg' to find out why.

There were problems setting up VirtualBox.  To re-start the set-up process, run
  /sbin/vboxconfig
as root.  If your system is using EFI Secure Boot you may need to sign the
kernel modules (vboxdrv, vboxnetflt, vboxnetadp, vboxpci) before you can load
them. Please see your Linux system's documentation for more information.
Processing triggers for mailcap (3.70+nmu1ubuntu1) ...
Processing triggers for gnome-menus (3.36.0-1ubuntu3) ...
Processing triggers for desktop-file-utils (0.26-1ubuntu3) ...
Processing triggers for hicolor-icon-theme (0.17-2) ...
Processing triggers for shared-mime-info (2.1-2) ...
```

### gdebi 방식으로 install
```sh
$ sudo gdebi virtualbox-7.0_7.0.16-162802~Ubuntu~jammy_amd64.deb
```
