Linux에서 DNS 서버를 등록하는 방법은 사용하는 배포판에 따라 조금씩 다를 수 있습니다. 아래는 일반적인 방법들입니다.

1. /etc/resolv.conf 파일 수정
이 방법은 대부분의 Linux 배포판에서 사용할 수 있습니다.

터미널을 열고 root 권한으로 편집기를 실행합니다:

bash
코드 복사
sudo nano /etc/resolv.conf
또는

bash
코드 복사
sudo vi /etc/resolv.conf
파일에 다음 내용을 추가합니다:

plaintext
코드 복사
nameserver 8.8.8.8
nameserver 8.8.4.4
여기서 8.8.8.8과 8.8.4.4는 구글의 공용 DNS 서버입니다. 다른 DNS 서버를 사용하고 싶다면 해당 IP 주소로 변경하면 됩니다.

파일을 저장하고 종료합니다.

2. NetworkManager 사용 (Ubuntu, Fedora 등)
NetworkManager를 사용하는 배포판에서는 GUI 또는 명령어를 통해 DNS 서버를 설정할 수 있습니다.

GUI를 통한 설정 (Ubuntu 예시)
네트워크 설정으로 이동합니다.
사용하는 네트워크를 선택하고, 설정(기어 아이콘)을 클릭합니다.
IPv4 또는 IPv6 탭으로 이동합니다.
DNS 서버를 수동으로 설정하고, 원하는 DNS 서버 주소를 입력합니다.
설정을 저장하고 네트워크 연결을 재시작합니다.
CLI를 통한 설정
nmcli 명령어를 사용하여 DNS 서버를 설정합니다:
bash
코드 복사
sudo nmcli dev show | grep DNS
sudo nmcli con mod <connection_name> ipv4.dns "8.8.8.8 8.8.4.4"
sudo nmcli con up <connection_name>
여기서 <connection_name>은 네트워크 연결 이름입니다. 이를 찾으려면 nmcli con show 명령어를 사용할 수 있습니다.
3. Systemd-resolved 사용 (Ubuntu 18.04 이상)
systemd-resolved를 사용하는 경우 /etc/systemd/resolved.conf 파일을 편집합니다:

bash
코드 복사
sudo nano /etc/systemd/resolved.conf
DNS=와 FallbackDNS= 라인을 찾아 주석을 제거하고 원하는 DNS 서버 주소를 입력합니다:

plaintext
코드 복사
[Resolve]
DNS=8.8.8.8 8.8.4.4
FallbackDNS=1.1.1.1 1.0.0.1
파일을 저장하고 종료한 다음, systemd-resolved 서비스를 재시작합니다:

bash
코드 복사
sudo systemctl restart systemd-resolved
/etc/resolv.conf 파일이 systemd-resolved에 의해 관리되고 있는지 확인합니다:

bash
코드 복사
sudo ln -sf /run/systemd/resolve/resolv.conf /etc/resolv.conf
4. Netplan 사용 (Ubuntu 20.04 이상)
Netplan 설정 파일을 편집합니다. 보통 /etc/netplan/01-netcfg.yaml 또는 /etc/netplan/50-cloud-init.yaml 파일이 사용됩니다:

bash
코드 복사
sudo nano /etc/netplan/01-netcfg.yaml
다음과 같이 DNS 서버 설정을 추가합니다:

yaml
코드 복사
network:
  version: 2
  ethernets:
    eth0:
      dhcp4: yes
      nameservers:
        addresses:
          - 8.8.8.8
          - 8.8.4.4
파일을 저장하고 Netplan을 적용합니다:

bash
코드 복사
sudo netplan apply
이러한 방법들을 통해 Linux 시스템에서 DNS 서버를 설정할 수 있습니다. 사용 중인 배포판과 네트워크 관리 도구에 따라 적절한 방법을 선택하세요.






