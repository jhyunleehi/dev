network UNCLAIMED with Intel Wi-Fi 6 AX200 after fresh Ubuntu 20.04 installation
Asked 1 year ago
Modified 11 months ago
Viewed 2k times
2

Some useful information to help with diagnostics:

$ lspci -knn | grep Net -A3; rfkill list
25:00.0 Network controller [0280]: Intel Corporation Wi-Fi 6 AX200 [8086:2723] (rev 1a)
    Subsystem: Intel Corporation Wi-Fi 6 AX200 [8086:0084]
    Kernel modules: iwlwifi

lshw

$ sudo lshw -C network
  *-network UNCLAIMED
       description: Network controller
       product: Wi-Fi 6 AX200
       vendor: Intel Corporation
       physical id: 0
       bus info: pci@0000:25:00.0
       version: 1a
       width: 64 bits
       clock: 33MHz
       capabilities: pm msi pciexpress msix cap_list
       configuration: latency=0

Kernel

$ uname -r
5.15.0-41-generic

dmesg

$ sudo dmesg | grep iwl
[  133.187211] iwlwifi 0000:25:00.0: pcim_iomap_regions_request_all failed
[  133.187224] iwlwifi: probe of 0000:25:00.0 failed with error -22

System Info

$ uname -a
Linux a25f126-lcedt 5.15.0-41-generic #44~20.04.1-Ubuntu SMP Fri Jun 24 13:27:29 UTC 2022 x86_64 x86_64 x86_64 GNU/Linux

p.s. No dual boot. Ubuntu only, i.e., disabling Fast Startup in BIOS didn't help.

Edit

$ sudo dmesg | grep 25:00
[sudo] password for startup: 
[    2.214779] pci 0000:25:00.0: [8086:2723] type 00 class 0x028000
[    2.214811] pci 0000:25:00.0: reg 0x10: [mem 0xf0500000-0xf0503fff 64bit]
[    2.214979] pci 0000:25:00.0: PME# supported from D0 D3hot D3cold
[    2.251872] pci 0000:25:00.0: BAR 0: no space for [mem size 0x00004000 64bit]
[    2.251873] pci 0000:25:00.0: BAR 0: failed to assign [mem size 0x00004000 64bit]
[    2.252073] pci 0000:25:00.0: BAR 0: no space for [mem size 0x00004000 64bit]
[    2.252074] pci 0000:25:00.0: BAR 0: failed to assign [mem size 0x00004000 64bit]
[    2.253941] pci 0000:25:00.0: Adding to iommu group 43
[   53.193450] iwlwifi 0000:25:00.0: pcim_iomap_regions_request_all failed
[   53.193473] iwlwifi: probe of 0000:25:00.0 failed with error -22

After upgrade

$ cat /proc/version_signature
Ubuntu 5.15.0-48.54~20.04.1-generic 5.15.53

$ uname -r
5.15.0-48-generic

$ sudo dmesg | grep iwl
[   47.972918] iwlwifi 0000:25:00.0: pcim_iomap_regions_request_all failed
[   47.972929] iwlwifi: probe of 0000:25:00.0 failed with error -22

Contents of /etc/default/grub

$ more /etc/default/grub
# If you change this file, run 'update-grub' afterwards to update
# /boot/grub/grub.cfg.
# For full documentation of the options in this file, see:
#   info -f grub -n 'Simple configuration'

GRUB_DEFAULT=0
GRUB_TIMEOUT_STYLE=hidden
GRUB_TIMEOUT=0
GRUB_DISTRIBUTOR=`lsb_release -i -s 2> /dev/null || echo Debian`
GRUB_CMDLINE_LINUX_DEFAULT="quiet splash pci=assign-busses"
GRUB_CMDLINE_LINUX="quiet splash"
...

Then, I ran $ sudo update-grub and $ sudo reboot. Unfortunately, the dmesg output didn't change.

$ sudo dmesg | grep iwl
[   28.553306] iwlwifi 0000:26:00.0: pcim_iomap_regions_request_all failed
[   28.553320] iwlwifi: probe of 0000:26:00.0 failed with error -22

    networkingdrivers20.04iwlwifi


https://community.intel.com/t5/Wireless/Wifi-not-working-in-Linux-using-AX200-on-Beelink-mini-PC/td-p/1498596



```
Hi Deivid,

I got the wifi to work today. I wish I would understand why it is working now, what made the difference. Here is what i did:

Used a wired network connection with an ethernet cable
installed Ubuntu 22.04.02 LTS from a USB stick
During installation, chose to update packages and install proprietary drivers
After the installation completed I was able to connect to my wifi network and remove the wired connection
Thanks for sending the additional troubleshooting steps. I ran them and will try to understand their output.

Now that wifi is working, we can close this conversation. Thanks again for the pointers.

Paul
```