# How to Install OpenStack on Ubuntu with DevStack [ Easy method ]

https://opendev.org/openstack/devstack/src/branch/stable/2024.1#


### Step 1: Preparing the system

```sh
$ sudo apt-get update && sudo apt-get upgrade -y
```

### Step 2: Creating stack user with Sudo privileges
```sh
$ sudo useradd -s /bin/bash -d /opt/stack -m dstack
$ echo "dstack ALL=(ALL) NOPASSWD: ALL" | sudo tee /etc/sudoers.d/dstack
$ sudo su - dstack
```


### Step 3: Downloading Devstack
```sh
$ echo "check_certificate = off" >> ~/.wgetrc
$ sudo apt install git 
$ git clone https://opendev.org/openstack/devstack
```

### Step 4: Creating configuration (.conf) file for Devstack
```sh
$ cd devstack
$ vim local.conf

[[local|localr]]
 
ADMIN_PASSWORD=1234qwer
DATABASE_PASSWORD=$ADMIN_PASSWOCinder
RABBIT_PASSWORD=$ADMIN_PASSWORD
SERVICE_PASSWORD=$ADMIN_PASSWORD
ENABLED_SERVICES=placement-api
```

Note:
1. StrongAdminSecret is the password we used here, you can change it with your choice.
2. You can find a sample configuration file for local.conf in the Samples directory under the Devstack repository.

### Step 5: Installing Openstack with Devstack
```sh
$ ./stack.sh
```

The script will install the listed features for your OpenStack environment –

Horizon – OpenStack Dashboard
Keystone – Identity Service
Nova – Compute Service
Glance – Image Service
Neutron – Network Service
Placement – Placement API
Cinder – Block Storage Service
The setup will take around 10 to 20 minutes, based on your system performance and internet speed, as many git trees and packages are installed during the process.

After your installation successfully finishes, your terminal will look like the image below.


### Step 6: Accessing OpenStack using a web browser
```sh 
https://server-ip/dashboard
https://localhost/dashboard

Openstack Login
Openstack Login
Now, enter the credentials. You can also log in as admin here, by having User Name as admin & for Password using the one we added to local.conf file.
```

## ==> reinstall reset all 
$  ./unstack.sh
$  ./clean.sh
$ ps -ef | grep  "^stack"
$ kill -9 `ps -ef | grep  "^stack" | awk '{print $2}'`
$ sudo  deluser stack
$ sudo  rm  -rf /opt/stack

```sh
stack@gpu-3:~$ env | grep proxy
no_proxy=localhost,127.0.0.1,code.sdsdev.co.kr,sds.redii.net,70.150.192.213,70.150.192.231,70.150.192.119,70.60.31.37,70.60.31.51,70.70.202.215,70.60.31.41,70.60.31.47,70.60.31.38,70.60.31.36,70.60.31.72,70.60.31.104,70.60.31.106,70.60.192.231,70.60.31.103,70.150.192.201,70.150.192.202,192.168.137.3,192.168.137.101,192.168.137.102,192.168.137.103,192.168.137.104,192.168.137.105
ftp_proxy=http://70.10.15.10:8080
https_proxy=http://70.10.15.10:8080
http_proxy=http://70.10.15.10:8080
```



==> Didn't find service registered by hostname after 60 seconds - Openstack/devstack deployment fails

```sh
stack@gpu-4:~/devstack$ head /etc/hosts
127.0.0.1	localhost gpu-4
```


### install package 
```sh
$ sudo apt install -y mariadb-server mariadb rabbitmq-server python3-openstackclient memcached python3-memcached  \
	openstack-keystone httpd python3-mod_wsgi python3-heatclient openstack-glance openstack-neutron openstack-neutron-ml2 \
	openstack-nova-api openstack-nova-conductor openstack-nova-novncproxy openstack-nova-scheduler \
	openstack-neutron-linuxbridge openstack-cinder openstack-placement-api openstack-dashboard openstack-heat-common \
	openstack-heat-api openstack-heat-api-cfn openstack-heat-engine

$ sudo apt install -y openstack-swift-proxy python3-swift python3-swiftclient rdo-release 
        python3-keystoneclient python3-keystonemiddleware

$ sudo apt install -y openstack-selinux python3-openstacksdk
```

```
sudo  dpkg -r  apache2                        
sudo  dpkg -r  bridge-utils                   
sudo  dpkg -r  docutils-common                
sudo  dpkg -r  erlang-base                    
sudo  dpkg -r  ieee-data                      
sudo  dpkg -r  ipvsadm                        
sudo  dpkg -r  keepalived                     
sudo  dpkg -r  libapache2-mod-proxy-uwsgi     
sudo  dpkg -r  libapache2-mod-wsgi-py3        
sudo  dpkg -r  libvirt-clients                
sudo  dpkg -r  libvirt-daemon-config-network  
sudo  dpkg -r  libvirt-daemon-config-nwfilter 
sudo  dpkg -r  libvirt-daemon-system          
sudo  dpkg -r  nvme-cli                       
sudo  dpkg -r  openstack-dashboard-common     
sudo  dpkg -r  os-brick-common                
sudo  dpkg -r  pycadf-common                  
sudo  dpkg -r  python3-django                 
sudo  dpkg -r  python3-greenlet               
sudo  dpkg -r  python3-json-pointer           
sudo  dpkg -r  python3-jsonpatch              
sudo  dpkg -r  python3-jsonschema             
sudo  dpkg -r  python3-ldap:amd64             
sudo  dpkg -r  python3-migrate                
sudo  dpkg -r  python3-oslo.messaging         
sudo  dpkg -r  python3-oslo.policy            
sudo  dpkg -r  python3-oslo.privsep           
sudo  dpkg -r  python3-oslo.rootwrap          
sudo  dpkg -r  python3-rtslib-fb              
sudo  dpkg -r  python3-yaql                   
sudo  dpkg -r  radvd                          
sudo  dpkg -r  swtpm                          
sudo  dpkg -r  swtpm-tools                    
sudo  dpkg -r  zvmcloudconnector-common       

sudo  dpkg --purge  apache2                        
sudo  dpkg --purge  bridge-utils                   
sudo  dpkg --purge  docutils-common                
sudo  dpkg --purge  erlang-base                    
sudo  dpkg --purge  ieee-data                      
sudo  dpkg --purge  ipvsadm                        
sudo  dpkg --purge  keepalived                     
sudo  dpkg --purge  libapache2-mod-proxy-uwsgi     
sudo  dpkg --purge  libapache2-mod-wsgi-py3        
sudo  dpkg --purge  libvirt-clients                
sudo  dpkg --purge  libvirt-daemon-config-network  
sudo  dpkg --purge  libvirt-daemon-config-nwfilter 
sudo  dpkg --purge  libvirt-daemon-system          
sudo  dpkg --purge  nvme-cli                       
sudo  dpkg --purge  openstack-dashboard-common     
sudo  dpkg --purge  os-brick-common                
sudo  dpkg --purge  pycadf-common                  
sudo  dpkg --purge  python3-django                 
sudo  dpkg --purge  python3-greenlet               
sudo  dpkg --purge  python3-json-pointer           
sudo  dpkg --purge  python3-jsonpatch              
sudo  dpkg --purge  python3-jsonschema             
sudo  dpkg --purge  python3-ldap:amd64             
sudo  dpkg --purge  python3-migrate                
sudo  dpkg --purge  python3-oslo.messaging         
sudo  dpkg --purge  python3-oslo.policy            
sudo  dpkg --purge  python3-oslo.privsep           
sudo  dpkg --purge  python3-oslo.rootwrap          
sudo  dpkg --purge  python3-rtslib-fb              
sudo  dpkg --purge  python3-yaql                   
sudo  dpkg --purge  radvd                          
sudo  dpkg --purge  swtpm                          
sudo  dpkg --purge  swtpm-tools                    
sudo  dpkg --purge  zvmcloudconnector-common       
```