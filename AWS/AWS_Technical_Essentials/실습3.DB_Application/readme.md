# 실습 3: Amazon S3 버킷과 Amazon DynamoDB 테이블을 사용하도록 웹 애플리케이션 구성

목표
이 실습을 완료하면 다음을 수행할 수 있습니다.

* Amazon Simple Storage Service(Amazon S3) 버킷 생성
* S3 버킷 정책 생성
* S3 버킷을 사용하도록 애플리케이션 수정
*S3 버킷에 객체 업로드
* Amazon DynamoDB 테이블 생성
* 애플리케이션 웹 인터페이스를 사용하여 애플리케이션 테스트
* AWS Management Console을 사용하여 기존 DynamoDB 항목 관리
* AWS Management Console을 사용하여 DynamoDB 테이블에서 항목 생성

## 과제 1: Amazon Simple Storage Service(Amazon S3) 버킷 생성

실습 지원을 위해 필수 리소스가 미리 생성되어 있습니다. 이러한 리소스에는 퍼블릭 서브넷 2개가 가용 영역 2개에 각각 포함되어 있는 VPC, 인터넷 게이트웨이, 인터넷으로의 경로가 포함된 라우팅 테이블, 그리고 Employee Directory 애플리케이션을 호스팅하는 EC2 인스턴스가 포함됩니다. 아래 이미지에 리소스가 대략적으로 나와 있습니

![alt text](image.png)


![alt text](image-1.png)
![alt text](image-2.png)
![alt text](image-3.png)
![alt text](image-4.png)

## 과제 2: S3 버킷 정책 생성

![alt text](image-5.png)

![alt text](image-6.png)

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "AllowS3ReadAccess",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::737833037758:role/EmployeeDirectoryAppRole"
      },
      "Action": "s3:*",
      "Resource": [
        "arn:aws:s3:::employee-photo-bucket-jhyunlee-7747",
        "arn:aws:s3:::employee-photo-bucket-jhyunlee-7747/*"
      ]
    }
  ]
}
```
![alt text](image-8.png)

![alt text](image-9.png)
![alt text](image-10.png)

## 과제 3: S3 버킷을 사용하도록 애플리케이션 수정

![alt text](image-11.png)


## 과제 4: S3 버킷에 객체 업로드

![alt text](image-12.png)
![alt text](image-13.png)

## 과제 5: Amazon DynamoDB 테이블 생성

![alt text](image-14.png)


## 과제 6: 애플리케이션 웹 인터페이스를 사용하여 애플리케이션 테스트
![alt text](image-15.png)
![alt text](image-16.png)
![alt text](image-17.png)

## 과제 7: AWS Management Console을 사용하여 기존 DynamoDB 항목 관리
![alt text](image-18.png)
![alt text](image-19.png)

## 과제 8: AWS Management Console을 사용하여 DynamoDB 테이블에서 항목 생성

![alt text](image-20.png)

