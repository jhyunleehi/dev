### wifi install 

#### chip : 02:00.0 Network controller: Intel Corporation Wi-Fi 6 AX200 (rev 1a)

```
jhyunlee@good:~/build/rtl8821CU$ lspci
00:00.0 Host bridge: Advanced Micro Devices, Inc. [AMD] Renoir/Cezanne Root Complex
00:00.2 IOMMU: Advanced Micro Devices, Inc. [AMD] Renoir/Cezanne IOMMU
00:01.0 Host bridge: Advanced Micro Devices, Inc. [AMD] Renoir PCIe Dummy Host Bridge
00:01.2 PCI bridge: Advanced Micro Devices, Inc. [AMD] Renoir/Cezanne PCIe GPP Bridge
00:01.3 PCI bridge: Advanced Micro Devices, Inc. [AMD] Renoir/Cezanne PCIe GPP Bridge
00:02.0 Host bridge: Advanced Micro Devices, Inc. [AMD] Renoir PCIe Dummy Host Bridge
00:02.1 PCI bridge: Advanced Micro Devices, Inc. [AMD] Renoir/Cezanne PCIe GPP Bridge
00:08.0 Host bridge: Advanced Micro Devices, Inc. [AMD] Renoir PCIe Dummy Host Bridge
00:08.1 PCI bridge: Advanced Micro Devices, Inc. [AMD] Renoir Internal PCIe GPP Bridge to Bus
00:14.0 SMBus: Advanced Micro Devices, Inc. [AMD] FCH SMBus Controller (rev 51)
00:14.3 ISA bridge: Advanced Micro Devices, Inc. [AMD] FCH LPC Bridge (rev 51)
00:18.0 Host bridge: Advanced Micro Devices, Inc. [AMD] Cezanne Data Fabric; Function 0
00:18.1 Host bridge: Advanced Micro Devices, Inc. [AMD] Cezanne Data Fabric; Function 1
00:18.2 Host bridge: Advanced Micro Devices, Inc. [AMD] Cezanne Data Fabric; Function 2
00:18.3 Host bridge: Advanced Micro Devices, Inc. [AMD] Cezanne Data Fabric; Function 3
00:18.4 Host bridge: Advanced Micro Devices, Inc. [AMD] Cezanne Data Fabric; Function 4
00:18.5 Host bridge: Advanced Micro Devices, Inc. [AMD] Cezanne Data Fabric; Function 5
00:18.6 Host bridge: Advanced Micro Devices, Inc. [AMD] Cezanne Data Fabric; Function 6
00:18.7 Host bridge: Advanced Micro Devices, Inc. [AMD] Cezanne Data Fabric; Function 7
01:00.0 Ethernet controller: Realtek Semiconductor Co., Ltd. RTL8111/8168/8411 PCI Express Gigabit Ethernet Controller (rev 15)
02:00.0 Network controller: Intel Corporation Wi-Fi 6 AX200 (rev 1a)
03:00.0 Non-Volatile memory controller: Micron/Crucial Technology P2 NVMe PCIe SSD (rev 01)
04:00.0 VGA compatible controller: Advanced Micro Devices, Inc. [AMD/ATI] Cezanne [Radeon Vega Series / Radeon Vega Mobile Series] (rev c5)
04:00.1 Audio device: Advanced Micro Devices, Inc. [AMD/ATI] Renoir Radeon High Definition Audio Controller
04:00.2 Encryption controller: Advanced Micro Devices, Inc. [AMD] Family 17h (Models 10h-1fh) Platform Security Processor
04:00.3 USB controller: Advanced Micro Devices, Inc. [AMD] Renoir/Cezanne USB 3.1
04:00.4 USB controller: Advanced Micro Devices, Inc. [AMD] Renoir/Cezanne USB 3.1
04:00.5 Multimedia controller: Advanced Micro Devices, Inc. [AMD] ACP/ACP3X/ACP6x Audio Coprocessor (rev 01)
04:00.6 Audio device: Advanced Micro Devices, Inc. [AMD] Family 17h/19h HD Audio Controller
```

#### intel 시스템 지원 utility 설치가 필요한듯  SSU

https://community.intel.com/t5/Wireless/Wifi-not-working-in-Linux-using-AX200-on-Beelink-mini-PC/td-p/1498596


https://www.intel.com/content/www/us/en/support/articles/000005511/wireless.html


2.5기가비트 이더넷은 Realtek RTL8125 네트워크 인터페이스 컨트롤러를 사용하고 WiFi 6은 이론적으로 최대 2.4Gb/s의 처리량을 제공할 수 있는 2×2 WiFi 6 기술을 지원하는 Intel Cyclone Peak AX200 M.2 2230 카드를 사용합니다. Bluetooth 5.2 지원 제공:


#### intel AX200 driver 

https://gist.github.com/mixxen/339846df6f316416336d038090a3c848
```
# run this to determine which firmware the kernal is looking for
modinfo iwlwifi | grep iwlwifi-cc

# example output:
# firmware:       iwlwifi-cc-a0-50.ucode

# go https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tree/ and download the firmware

# copy file to /lib/firmware
sudo cp ~/Download/iwlwifi-cc-a0-50.ucode /lib/firmware

# reboot
sudo reboot

# check that firmware was loaded
sudo dmesg | grep iwlwifi | grep version

# example output
# [    4.566709] iwlwifi 0000:05:00.0: loaded firmware version 50.3e391d3e.0 op_mode iwlmvm
```


```
ACupofAir가 댓글을 달았습니다. on 2021년 8월 6일
아니요, 귀하의 방법은 구식입니다. Ax200은 커널 5.1+부터 지원되었습니다. 우분투 20.04를 사용하면 확실히 ax200을 지원합니다. Wi-Fi를 사용할 수 없는 가장 큰 문제는 win10의 빠른 시작일 수 있습니다. 비활성화한 다음 문제를 해결합니다. 대체로 귀하의 아이디어에 감사드립니다.

@진왈린
zinwalin이 댓글을 달았습니다. on 2022년 2월 14일
BIOS에서 빠른 시작을 비활성화하고 Windows 10을 부팅한 다음 다시 시작하고 Ubuntu 20.04를 입력하면 해결됩니다.

@jason-shen
Jason-Shen이 댓글을 달았습니다. 2022년 5월 28일
여기 이 문제에 대한 2센트가 있습니다. 지금까지 이것은 여전히 ​​문제입니다.@zinwalin맞습니다. 팁도 추가하세요. 새로 설치했는데 Wi-Fi가 표시되지 않으면 스트레스 받지 마세요. 1x 미만의 속도로 창으로 이동한 다음 재부팅하면 Wi-Fi가 나타날 것입니다. 누군가 나에게 말했으면 좋았을 텐데요. 그거, 3일 뒤에야 알겠는데 ㅋㅋㅋ
```