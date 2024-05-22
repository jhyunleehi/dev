## 실습 개요

AWS 솔루션스 아키텍트는 Amazon Web Services(AWS)의 전반적 기능과 AWS 네트워킹 구성 요소 간의 관계를 이해해야 합니다. 이 실습에서는 Amazon Virtual Private Cloud(Amazon VPC), 단일 가용 영역의 퍼블릭 서브넷과 프라이빗 서브넷, 퍼블릭 및 프라이빗 경로, NAT 게이트웨이 및 인터넷 게이트웨이를 생성합니다. 이러한 서비스는 AWS 내부의 네트워킹 아키텍처의 기반입니다. 이 설계에는 인프라, 설계, 라우팅, 보안 개념이 포함됩니다.

다음 이미지에는 이 실습 환경의 최종 아키텍처를 보여 줍니다.

![alt text](image.png)


### 과제1: 특정리전에 VPC 생성 

![alt text](image-1.png)

1. VPC  생성 
![alt text](image-2.png)

2. VPC 설정 편집 
* DNS 호스트 이름 활성화
* DNS 확인 활성화
![alt text](image-3.png)



### 과제2: 퍼블릭 서브넷 및 프라이빗 서브넷 생성
이 과제에서는 Lab VPC에서 퍼블릭 서브넷과 프라이빗 서브넷을 생성합니다. VPC에 새 서브넷을 추가하려면 VPC의 범위에서 서브넷의 IPv4 CIDR 블록을 지정해야 합니다. 서브넷을 위치할 가용 영역을 지정할 수 있습니다. 동일한 가용 영역에 여러 서브넷을 위치할 수 있습니다.

![alt text](image-4.png)


1. public subnet 생성

![alt text](image-5.png)

2. private subnet 생성

![alt text](image-6.png)
* 라우팅 테이블
![alt text](image-7.png)
* 네트워크 ACL
![alt text](image-8.png)

### 과제3: internet gateway

![alt text](image-9.png)

### 과제4: public subnet traffic을 ineternet gateway로 라우팅

![alt text](image-10.png)

![alt text](image-11.png)

### 과제5: public 보안 그룹 생성
1. 인바운드 규칙
![alt text](image-12.png)
2. 아웃바운드 규칙
![alt text](image-13.png)

### 과제6 : EC2 인스턴스 생성
![alt text](image-14.png)

이 과제에서는 퍼블릭 서브넷에서 Amazon EC2 인스턴스를 시작합니다. 
* 인터넷을 통한 IPv4 통신을 활성화하려면 인스턴스에 해당 인스턴스의 프라이빗 IPv4 주소와 연결된 퍼블릭 IPv4 주소가 있어야 합니다. 
* 기본적으로 인스턴스는 VPC와 서브넷 안에서 정의된 프라이빗(내부) IP 주소 공간만 인식합니다

![alt text](image-16.png)

![alt text](image-15.png)

* 사용자 프로파일 
```sh
#!/bin/bash
# To connect to your EC2 instance and install the Apache web server with PHP
yum update -y
yum install -y httpd php8.1
systemctl enable httpd.service
systemctl start httpd
cd /var/www/html
wget https://us-west-2-tcprod.s3.amazonaws.com/courses/ILT-TF-200-ARCHIT/v7.7.3.prod-f959cb1c/lab-2-VPC/scripts/instanceData.zip
unzip instanceData.zip
```
### 과제7: http public instance에 접속
* http를 이용한 public dns를 이용한 외부 접속

![alt text](image-17.png)




### 과제8: session manager 접속
* session manager 이용한 접속

![alt text](image-18.png)


### 과제9: NAT 구성한 private 접속

![alt text](image-19.png)

* 라우팅 테이블 등록

![alt text](image-20.png)

* private subnet 연결 

![alt text](image-21.png)

### 과제 10: 프라이빗 리소스용 보안 그룹 생성

![alt text](image-23.png)

![alt text](image-24.png)

### 과제 11: 프라이빗 서브넷에서 Amazon EC2 인스턴스 시작

![alt text](image-25.png)

![alt text](image-29.png)

* 보안 그룹

![alt text](image-26.png)

* 네트워크 
![alt text](image-27.png)

* 스토리지
![alt text](image-28.png)


