## ubuntu 환경에서 mouse가 멈추는 현상에 대한 조치 방법


```bash
sudo nano /etc/default/grub

GRUB_CMDLINE_LINUX="quiet splash usbcore.autosuspend=-1"
```

```bash
sudo update-grub
sudo reboot
```