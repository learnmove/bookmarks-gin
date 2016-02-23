# Bookmarks Gin (Golang)

### This is part of a project: [Bookmarks](https://github.com/stefanjarina/bookmarks)

* Gin Framework
* GORM
* MySQL

#### TODO:
- [x] Change MongoDB to GORM and MySQL
- [x] Fix code formatting
- [ ] Proper testing + TDD onwards
- [x] Users + authentification - probably leverage jwt tokens
- [ ] Public/Private Bookmarks
- [ ] Backend filters support

### Installation

I presume you have MySQL installed and running.

In *MySQL CLI* create user, database and give user privileges for database.
e.g.:
```sql
CREATE USER 'bookmarks_gin'@'localhost' IDENTIFIED BY 'password';
CREATE DATABASE bookmarks_gin;
GRANT ALL ON bookmarks_gin.* TO 'bookmarks_gin'@'localhost';
```

**You need to specify few variables**
```bash
export GO_MYSQL_URI="bookmarks_gin:password@/bookmarks_gin"
export IP=0.0.0.0
export PORT=4000
```

Download source, fetch dependencies, compile and run
```bash
go get github.com/stefanjarina/bookmarks-gin
cd $GOPATH/src/github.com/stefanjarina/bookmarks-gin

go get
go build
./bookmarks-gin initdb --password "your_new_admin_pass"
./bookmarks-gin
```