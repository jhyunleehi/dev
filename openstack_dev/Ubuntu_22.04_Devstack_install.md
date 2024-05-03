# DevStack 
DevStack은 빠르고 쉽게 오픈스택의 핵심 컴포넌트를 구축하기 위한 스크립트
* Nova: 가상의 서버를 생성 및 관리
* Glance: 서버의 이미지 생성 및 관리
* Cinder: 블록 스토리지를 생성 및 관리
* Neutron: 가상의 네트워크 생성 및 관리
* Keystone: 사용자에 대한 인증
* Swift: 오브젝트 스토리지를 구성
* Horizon: Horizon 대시보드를 통해 GUI 환경에서 오픈스택의 구성요소들을 컨트롤 가능

### 1. swap 
```sh
sudo fallocate -l 8G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
```

### 2. stack 계정 생성
```sh
sudo useradd -s /bin/bash -d /opt/stack -m stack
sudo chown stack:stack -R /opt/stack/
echo "stack ALL=(ALL) NOPASSWD: ALL" | sudo tee /etc/sudoers.d/stack
sudo chmod 755 /opt/stack
sudo su - stack
```

### 3. devstack 
* 설치버젼 선택: https://opendev.org/openstack/devstack 
```sh
sudo apt-get update
sudo apt install git -y
git clone https://opendev.org/openstack/devstack
cd devstack
```

### 4. 가상 Floating(Public) ip 대역 생성
```sh
sudo apt install bridge-utils net-tools
sudo brctl addbr mybr0
sudo ifconfig mybr0 192.168.100.1 netmask 255.255.255.0 up
sudo iptables -I FORWARD -j ACCEPT
sudo iptables -t nat -I POSTROUTING -s 192.168.100.0/24 -j MASQUERADE
sudo ip addr add $할당받은_공인_IP/32 dev lo
```

### 5. Openstack 설치

* devstack에 대한 설정은 local.conf를 통해 이루어진다.
* local.conf의 샘플은 devstack/samples에 위치
* devstack 설치를 진행할 때는 /devstack에 local.conf 파일을 설치

* local.conf 파일 
```sh
[[local|localrc]]
HOST_IP=$할당받은_공인IP
FORCE=yes
ADMIN_PASSWORD=secret
DATABASE_PASSWORD=$ADMIN_PASSWORD
RABBIT_PASSWORD=$ADMIN_PASSWORD
SERVICE_PASSWORD=$ADMIN_PASSWORD

# Enable Swift
enable_service s-proxy s-object s-container s-account
SWIFT_REPLICAS=1
SWIFT_HASH=66a3d6b56c1f479c8b4e70ab5c2000f5
SWIFT_DATA_DIR=/opt/stack/data/swift

# Disable services
disable_service etcd3
disable_service ovn
disable_service ovn-controller
disable_service ovn-northd
disable_service q-ovn-metadata-agent

# Use openvswitch as the ml2 plugin driver
Q_AGENT=openvswitch

# Enable Neutron services neutron-server, neutron-openvswitch-agent,
# neutron-dhcp-agent, neutron-l3-agent and neutron-metadata-agent
enable_service q-svc
enable_service q-agt
enable_service q-dhcp
enable_service q-l3
enable_service q-meta

## Neutron options
#Q_USE_SECGROUP=True
FLOATING_RANGE="192.168.100.0/24"    
IPV4_ADDRS_SAFE_TO_USE="10.0.0.0/22"    
Q_FLOATING_ALLOCATION_POOL=start=192.168.100.50,end=192.168.100.250
PUBLIC_NETWORK_GATEWAY="192.168.100.1"
PUBLIC_INTERFACE=mybr0

# Open vSwitch provider networking configuration
Q_USE_PROVIDERNET_FOR_PUBLIC=True
OVS_PHYSICAL_BRIDGE=br-ex
PUBLIC_BRIDGE=br-ex
OVS_BRIDGE_MAPPINGS=public:br-ex

[[post-config|/$Q_PLUGIN_CONF_FILE]]
[ml2]
type_drivers=flat,gre,vlan,vxlan
tenant_network_types=vxlan
mechanism_drivers=openvswitch,l2population

[agent]
tunnel_types=vxlan,gre
```

### 6. 재설치
```sh
./clean.sh
./stack.sh

sudo dpkg -l | grep  -v  ii
sudo dpkg -r pkg_name
sudo dpkg -P pkg_name
```

### 7. 실행
```sh

ll ~/devstack/openrc
source openrc $username $project_name
source openrc admin admin
evn | grep OS_

https://10.10.50/dashboard
https://localhost/dashboard

Openstack Login
Now, enter the credentials. You can also log in as admin here, by having User Name as admin & for Password using the one we added to local.conf file.
```

