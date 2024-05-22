## Wifi 

### command

```sh
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

$ sudo nmcli device
$ sudo lshw -businfo
$ sudo lshw -c network
$ sudo iwconfig 

```sh
$ sudo nmcli device
[sudo] password for jhyunlee: 
DEVICE           TYPE      STATE      CONNECTION         
enx00e04c3621af  ethernet  connected  Wired connection 1 
lo               loopback  unmanaged  -
```

### sudo lshw -businfo
```sh
$ sudo lshw -businfo
Bus info          Device           Class          Description
=============================================================
                                   system         Venue 11 Pro 7130 MS (0604)
                                   bus            Motherboard
                                   memory         64KiB BIOS
cpu@0                              processor      Intel(R) Core(TM) i5-4210Y CPU @ 1.50GHz
                                   memory         128KiB L1 cache
                                   memory         512KiB L2 cache
                                   memory         3MiB L3 cache
                                   memory         4GiB System Memory
                                   memory         2GiB SODIMM DDR3 Synchronous 1600 MHz (0.6 ns)
                                   memory         2GiB SODIMM DDR3 Synchronous 1600 MHz (0.6 ns)
pci@0000:00:00.0                   bridge         Haswell-ULT DRAM Controller
pci@0000:00:02.0  /dev/fb0         display        Intel Corporation
pci@0000:00:03.0  card0            multimedia     Haswell-ULT HD Audio Controller
                  input30          input          HDA Intel HDMI HDMI/DP,pcm=3
                  input31          input          HDA Intel HDMI HDMI/DP,pcm=7
                  input32          input          HDA Intel HDMI HDMI/DP,pcm=8
pci@0000:00:14.0                   bus            8 Series USB xHCI HC
usb@1             usb1             bus            xHCI Host Controller
usb@1:2                            bus            USB 2.0 Hub
usb@1:2.3                          input          USB Receiver
                  input33          input          Logitech Wireless Keyboard PID:4023
                  input34          input          Logitech Wireless Mouse PID:4022
usb@1:2.4                          generic        USB 10/100 LAN
usb@1:3                            multimedia     Integrated Webcam
usb@1:5                            multimedia     Video
usb@2             usb2             bus            xHCI Host Controller
pci@0000:00:16.0                   communication  8 Series HECI #0
pci@0000:00:1b.0  card1            multimedia     8 Series HD Audio Controller
                  input18          input          HDA Intel PCH Dock Headphone
                  input19          input          HDA Intel PCH Headphone
pci@0000:00:1c.0                   bridge         8 Series PCI Express Root Port 4
pci@0000:01:00.0  mmc1             bus            SD/MMC Card Reader Controller
pci@0000:00:1f.0                   bridge         8 Series LPC Controller
                                   system         PnP device PNP0c01
                                   system         PnP device PNP0c02
                                   system         PnP device PNP0b00
                                   generic        PnP device INT3f0d
                                   system         PnP device PNP0c02
                                   input          PnP device PNP0303
                                   system         PnP device PNP0c02
                                   system         PnP device PNP0c02
                  input6           input          Intel Virtual Buttons
                  input7           input          Intel Virtual Switches
pci@0000:00:1f.2  scsi0            storage        8 Series SATA Controller 1 [AHCI mode]
scsi@0:0.0.0      /dev/sda         disk           128GB LITEONIT LJT-128
scsi@0:0.0.0,1                     volume         511MiB Windows FAT volume
scsi@0:0.0.0,2    /dev/sda2        volume         118GiB EXT4 volume
pci@0000:00:1f.3                   bus            8 Series SMBus Controller
                                   power          DELL VJF0X46
                                   power          To Be Filled By O.E.M.
                  mmc0             bus            MMC Host
mmc@0:0001:1      mmc0:0001:1      generic        SDIO Device
                  input0           input          Lid Switch
                  input1           input          Power Button
                  input10          input          SYNA7500:00 06CB:3AF0 Stylus
                  input11          input          SYNA7500:00 06CB:3AF0
                  input2           input          AT Translated Set 2 keyboard
                  input20          input          Video Bus
                  input9           input          Dell WMI hotkeys
usb@1:2.4         enx00e04c3621af  network        Ethernet interface

```

### sudo lshw -c network
```sh
$ sudo lshw -c network
  *-network                 
       description: Ethernet interface
       physical id: b
       bus info: usb@1:2.4
       logical name: enx00e04c3621af
       serial: 00:e0:4c:36:21:af
       size: 100Mbit/s
       capacity: 100Mbit/s
       capabilities: ethernet physical tp mii 10bt 10bt-fd 100bt 100bt-fd autonegotiation
       configuration: autonegotiation=on broadcast=yes driver=r8152 driverversion=v1.12.13 duplex=full ip=192.168.0.29 link=yes multicast=yes port=MII speed=100Mbit/s
```
```sh
$ sudo iwconfig 
lo        no wireless extensions.

enx00e04c3621af  no wireless extensions.
```