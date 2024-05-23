## mariadb
### remove all purge

```sh
sudo systemctl stop mariadb
sudo service mysql stop
sudo apt-get remove --purge mariadb-server mariadb-client mariadb-common
sudo rm -rf / var / lib / mysql
sudo rm -rf / etc / mysql
sudo apt-get autoremove sudo deluser mysql
sudo delgroup mysql
sudo rm-rf / var / log / mysql *
sudo apt-get
```
### install  sudo apt - get
```sh
sudo apt-get install mariadb-server mariadb-client
sudo systemctl start mariadb
sudo mysql_secure_installation
sudo systemctl enable mariadb
sudo apt update
sudo apt install libmariadb-dev
```
### 가상환경 설정 
```sh
python3 -m venv venv
source venv/bin/activate
pip install pandas
pip install mariadb==1.1.7
pip install pyyaml
pip install sqlalchemy
pip install mysql-connector-python
deactivate
```
### 설치결과
```sh
libmariadb3/jammy-updates,jammy-security,now 1:10.6.12-0ubuntu0.22.04.1 amd64 [설치됨,자동]
mariadb-client-10.6/jammy-updates,jammy-security,now 1:10.6.12-0ubuntu0.22.04.1 amd64 [설치됨,자동]
mariadb-client-core-10.6/jammy-updates,jammy-security,now 1:10.6.12-0ubuntu0.22.04.1 amd64 [설치됨,자동]
mariadb-client/jammy-updates,jammy-updates,jammy-security,jammy-security,now 1:10.6.12-0ubuntu0.22.04.1 all [설치됨]
mariadb-common/jammy-updates,jammy-updates,jammy-security,jammy-security,now 1:10.6.12-0ubuntu0.22.04.1 all [설치됨,자동]
mariadb-server-10.6/jammy-updates,jammy-security,now 1:10.6.12-0ubuntu0.22.04.1 amd64 [설치됨,자동]
mariadb-server-core-10.6/jammy-updates,jammy-security,now 1:10.6.12-0ubuntu0.22.04.1 amd64 [설치됨,자동]
mariadb-server/jammy-updates,jammy-updates,jammy-security,jammy-security,now 1:10.6.12-0ubuntu0.22.04.1 all [설치됨]
libdbd-mysql-perl / jammy-updates,jammy-security,now 4.050 -5 ubuntu0.22.04.1 amd64 [설치됨,자동]
libmysqlclient21 / jammy-updates,jammy-security,now 8.0.35 -0 ubuntu0.22.04.1 amd64 [설치됨,자동]
mysql-common / jammy,jammy,now 5.8 + 1.0.8 all [설치됨,자동]
```
### DB 초기화
```sh
MariaDB [ (none)] > select user, host,authentication_string from mysql.user;
MariaDB [ (none)] > CREATE DATABASE TRADE;
MariaDB [ (none)] > SELECT User, Host, authentication_string FROM mysql.user;
MariaDB [ (none)] > create user 'trade' @'%' identified by 'tradepass';
MariaDB [ (none)] > FLUSH PRIVILEGES;
MariaDB [ (none)] > GRANT ALL PRIVILEGES ON TRADE.* TO trade @'%';
MariaDB [ (none)] > FLUSH PRIVILEGES;
MariaDB [ (none)] > show variables like " % version % ";
MariaDB [ (none)] > ALTER USER 'root' @'localhost' IDENTIFIED BY 'tradepass';
ALTER USER 'trade' IDENTIFIED BY 'tradepass';
MariaDB [ (none)] > use TRADE;
Database changed

$ mysql  -u trad -p
```
select user, host,authentication_string from mysql.user;
CREATE DATABASE TRADE;
SELECT User, Host, authentication_string FROM mysql.user;
create user 'trade' @'%' identified by 'tradepass';
FLUSH PRIVILEGES;
GRANT ALL PRIVILEGES ON TRADE.* TO trade @'%';
FLUSH PRIVILEGES;
show variables like " % version % ";
ALTER USER 'root' @'localhost' IDENTIFIED BY 'tradepass';