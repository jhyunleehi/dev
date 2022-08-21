# SQL



### 1. SQL 패키지

Go에서 SQL 데이타베이스를 사용하기 위해서는 표준패키지 database/sql을 사용한다. database/sql 패키지는 관계형 데이타베이스들에게 공통적으로 사용되는 인터페이스들을 제공하고 있다. 즉, sql 패키지는 데이타베이스를 연결하고, 쿼리를 실행하고, DML 명령을 수행하는 등의 SQL 관련 함수들을 제공한다.

database/sql 패키지는 여러 종류의 SQL 데이타베이스를 지원하는데, 각각의 데이타베이스 Driver와 함께 사용된다. 흔히 사용되는 데이타베이스에 대해 다음과 같은 드라이버들이 현재 지원되는데, Full List는 [여기](https://github.com/golang/go/wiki/SQLDrivers)에서 참조할 수 있다.

- MySQL: https://github.com/go-sql-driver/mysql
- MSSQL: https://github.com/denisenkom/go-mssqldb
- Oracle: https://github.com/rana/ora
- Postgres: https://github.com/lib/pq
- SQLite: https://github.com/mattn/go-sqlite3
- DB2: https://bitbucket.org/phiggins/db2cli



### 2. sql.DB 타입

database/sql 패키지에서 가장 중요한 Type은 sql.DB 인데, 일반적으로 sql.Open() 함수를 사용하여 sql.DB 객체를 얻는다. 즉, sql.Open(드라이버, Connection) 함수에서 어떤 DB 드라이버를 사용할 것인지 그리고 해당 DB의 연결 정보를 제공하면, 결과로 sql.DB 객체를 얻게 된다. 일단 이 sql.DB 객체를 얻은 후, sql.DB의 여러 메서드들을 사용하여 쿼리를 하고, SQL문을 실행한다. 

예를 들어, 자주 사용되는 sql.DB 메서드로 쿼리를 하는 Query(), QueryRow(), 그리고 INSERT, UPDATE, DELETE등을 실행하는 Exec()를 들 수 있다.



```go
// sql.DB 객체 db 생성
db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
 
// db 차후에 닫기
defer db.Close()
 
// SELECT 쿼리
rows, err := db.Query("SELECT id, name FROM test")
 
// INSERT 실행
db.Exec("INSERT INTO test(id, name) VALUES (1, 'Alex')")
```



### 3. mysql Query

Go에서 MySql 사용하기 위해서는 표준패키지 database/sql과 MySql 드라이버를 import해야 한다. MySql Driver를 처음 사용하는 경우는 아래와 같이 go get 명령을 통해 패키지를 먼저 다운받아야 한다. *(주: 몇가지 다른 MySql 드라이버들이 있는데, 여기서는 가장 많이 쓰는 github.com/go-sql-driver/mysql 드라이버를 사용)*
패키지가 정상적으로 다운로드되면 아무런 에러가 표시되지 않고 새 프롬프트만 표시된다.

```go
$ go get github.com/go-sql-driver/mysql
```

MySql 드라이버가 설치된 후, 아래와 같이 database/sql과 MySql 드라이버를 import하는데, MySql 드라이버 패키지는 _ 로 alias를 주어 개발자가 드라이버 패키지를 직접 사용하지 않게 한다. 이 경우 드라이버 패키지는 database/sql 패키지가 내부적으로 사용하게 되며, 개발자는 database/sql를 통해서 모든 SQL 프로세싱을 진행하게 된다.

```go
package main
 
import (
    "database/sql" 
    _ "github.com/go-sql-driver/mysql"
    "log"
)
 
func main() {
    db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
     
    //...(db 사용)....
}
```



패키지 database/sql에서 가장 먼저 사용하는 것은 sql.Open()으로 이는 sql.DB 객체를 리턴한다.  

sql.Open()의 첫번째 파라미터는 드라이버명으로 여기서는 mysql로 적으며, 두번째 파라미터는 각 드라이버의 Connection String으로서 MySql Connection을 적으면 된다. 

위의 Connection은 로컬서버에 TCP 3306 포트로 root와 그 암호를 사용하여 접속하며, testdb라는 Database에 접속할 것을 나타낸다. 

여기서 한가지 주목할 것은 sql.Open()은 실제 DB Connection을 Open하지 않는다는 점이다. 즉, sql.DB는 드라이버종류와 Connection 정보를 가지고는 있지만, 실제 DB를 연결하지 않으며, 많은 경우 Connection 정보조차 체크하지도 않는다. 실제 DB Connection은 Query 등과 같이 실제 DB 연결이 필요한 싯점에 이루어지게 된다.



#### 1. single row query

MySql에서 쿼리를 위해 2종류의 메서드를 사용한다. 즉, 하나의 Row만을 리턴할 경우 QueryRow() 메서드를, 복수개의 Row를 리턴할 경우 Query() 메서드를 사용한다. 

하나의 Row에서 실제 데이타를 읽어 로컬 변수에 할당하기 위해 Scan() 메서드를 사용하며, 복수 Row에서 다음 Row로 이동하기 위해 Next() 메서드를 사용한다. 아래 예제에서 QueryRow()는 SQL 문을 실행해서 리턴된 하나의 ROW의 데이타를 Scan() 안의 파라미터(name)에 넣고 있음을 볼 수 있다.

```go
package main
 
import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
)
 
func main() {
    // sql.DB 객체 생성
    db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
 
    // 하나의 Row를 갖는 SQL 쿼리
    var name string
    err = db.QueryRow("SELECT name FROM test1 WHERE id = 1").Scan(&name)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(name)
}
```



#### 2. multi row query

MySql에서 복수개의 Row를 리턴할 경우 Query() 메서드를 사용한다. 물론 Row가 0개 리턴되더라도 사용할 수 있다. 

복수 Row들에서 다음 Row로 하나씩 이동하기 위해 Next() 메서드를 사용하는데, 흔히 for 루프에서 계속 Next()를 호출하며 다음 Row가 있는지 체크한다. 아래 예제는 Query()를 사용하여 복수 Row를 리턴받고, 각 Row를 순서대로 하나씩 읽어 (Scan 메서드) 그 데이타를 출력하는 코드이다.

```go
func main() {
    // sql.DB 객체 생성
    db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
 
    // 복수 Row를 갖는 SQL 쿼리
    var id int
    var name string
    rows, err := db.Query("SELECT id, name FROM test1 where id >= ?", 1)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close() //반드시 닫는다 (지연하여 닫기)
 
    for rows.Next() {
        err := rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(id, name)
    }
}
```



여기서 한가지 주목할 것은 SQL 쿼리에서 ? (Placeholder)를 사용하여 Parameterized Query를 사용하고 있다는 점이다. 

이는 SQL Injection과 같은 문제를 방지하기 위해 파라미터를 문자열 결합이 아닌 별도의 파라미터로 대입시키는 방식이다. 

위의 예제에서 Placeholder ? 에는 1이 대입된다. Placeholder는 데이타베이스의 종류에 따라 다르게 사용하는데, 예를 들어 MySql은 ? 를 사용하고, Oracle은 :val1, :val2 등을 사용하고, PostgreSQL은 $1, $2 등을 사용한다.



### 4. mysql  DML

#### 1. insert, update, delete

데이타를 INSERT, UPDATE, DELETE (DML Operation)하기 위해서 sql.DB 객체의 Exec() 메서드를 사용한다. 

* Query/QueryRow 메서드는 데이타를 리턴할 때 사용하는 반면, DML과 같이 리턴되는 데이타가 없는 경우는 Exec 메서드를 사용해야 한다.

* Exec 메서드의 첫번째 파라미터에는 SQL문을 적고, 그 SQL문 안에 ? 이 있는 경우 계속해서 상응하는 파라미터를 넣어 준다. 
* Exec 메서드는 sql.Result와 error 객체를 리턴하며, sql.Result 객체로부터 갱신된 레코드수(RowsAffected())와 새로 추가된 Id (LastInsertId())를 구할 수 있다.

```go
package main
 
import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
)
 
func main() {
    // sql.DB 객체 생성
    db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
 
    // INSERT 문 실행
    result, err := db.Exec("INSERT INTO test1 VALUES (?, ?)", 11, "Jack")
    if err != nil {
        log.Fatal(err)
    }
 
    // sql.Result.RowsAffected() 체크
    n, err := result.RowsAffected()
    if n == 1 {
        fmt.Println("1 row inserted.")
    }
}
```



#### 2. prepared Statement 

Prepared Statement는 데이타베이스 서버에 Placeholder를 가진 SQL문을 미리 준비시키는 것으로, 차후 해당 Statement를 호출할 때 준비된 SQL문을 빠르게 실행하도록 하는 기법이다. Go에서 Prepared Statement를 사용하기 위해서는 sql.DB의 Prepare() 메서드를 써서 Placeholder를 가진 SQL문을 미리 준비시키고, sql.Stmt 객체를 리턴받는다. 차후 이 sql.Stmt 객체의 Exec (혹은 Query/QueryRow) 메서드를 사용하여 준비된 SQL문을 실행한다.

```go
package main
 
import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)
 
func main() {
    // sql.DB 객체 생성
    db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
 
    // Prepared Statement 생성
    stmt, err := db.Prepare("UPDATE test1 SET name=? WHERE id=?")
    checkError(err)
    defer stmt.Close()
 
    // Prepared Statement 실행
    _, err = stmt.Exec("Tom", 1) //Placeholder 파라미터 순서대로 전달
    checkError(err)
    _, err = stmt.Exec("Jack", 2)
    checkError(err)
    _, err = stmt.Exec("Shawn", 3)
    checkError(err)
}
 
func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
```

