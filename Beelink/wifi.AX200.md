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

Share
Improve this question
Follow
edited Oct 3, 2022 at 14:34
asked Sep 28, 2022 at 23:53
NV-CSG's user avatar
NV-CSG
2133 bronze badges

    @chili555 is this something that you can help, please? – 
    NV-CSG
    Sep 28, 2022 at 23:54
    I'm working on it. I'll post back as soon as/if I find something. – 
    chili555
    Sep 29, 2022 at 2:13
    May we please see the result of the terminal command: sudo dmesg | grep 25:00 in an edit to your question? Welcome to Ask Ubuntu. – 
    chili555
    Sep 29, 2022 at 12:31
    Also, the latest kernel version is -48 and you are running -41. If you update and reboot, is the problem resolved? sudo dmesg | grep iwl – 
    chili555
    Sep 29, 2022 at 12:34
    @chili555 I added the output of sudo dmesg | grep 25:00, and I ran apt get update and apt get upgrade. Now I am working on upgrading the kernel/ – 
    NV-CSG
    Sep 29, 2022 at 16:13

Show 6 more comments
2 Answers
Sorted by:
0

Please try the technique mentioned in the link I gave above. Open a terminal and do:

sudo nano /etc/default/grub

In the editor window, use the arrow keys to move the cursor to the line beginning with "GRUB_CMDLINE_LINUX_DEFAULT". Change that line to read:

GRUB_CMDLINE_LINUX_DEFAULT="quiet splash pci=assign-busses"

Proofread carefully twice. Save (Ctrl+o followed by Enter) and exit (Ctrl+x) the text editor nano. Follow with:

sudo update-grub

Reboot and tell us if there is any improvement.


