# URL SHORTENER
System to shorten URL
## CLONE

```bash
$ git clone git@github.com:DiegoSantosWS/urlshortener.git
$ cd urlshortener
```
## PACKAGES
if your descktop does not have the packages it will have download them
```bash
$ go get -u go get -u github.com/gorilla/...
```

## CONFIGURATION

create a MySQL `database` and enter user data and password
## CONNECTIONS

```bash
$ go get https://github.com/go-sql-driver/... OR
$ go get https://github.com/go-sql-driver/mysql
```
AFTER MAKE A CONNECTION WITH DATABASE
```go
Db, err = sqlx.Open("mysql", "user:pass@tcp(host:port)/dbname")
```
## URL EXAMPLE

`URL ORIGINAL: http://www.example.com.br`

`URL SHORTENER: http://localhost:3000/ADqF`