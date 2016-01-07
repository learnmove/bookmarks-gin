# Bookmarks Gin (Golang)

### This is part of a project: [Bookmarks](https://github.com/stefanjarina/bookmarks)

* Gin Framework
* GORM
* MySQL

#### TODO:
- [ ] Change Mongoose to GORM
- [ ] Fix code formatting

### Installation

I presume you have MySQL installed and running.

In *MySQL CLI* create user, database and give user privileges for database.
e.g.:
```sql
CREATE USER 'bookmarks_gin'@'localhost' IDENTIFIED BY 'password';
CREATE DATABASE bookmarks_gin;
GRANT ALL ON bookmarks_gin.* TO 'bookmarks_gin'@'localhost';
```

**You need to edit mysql connection URI**
```bash
export GO_MYSQL_URI="mysql://bookmarks_gin:password@localhost/bookmarks_gin"
```

Download source, fetch dependencies, compile and run
```bash
git clone https://github.com/stefanjarina/bookmarks-gin
cd $GOPATH/src/github.com/stefanjarina/bookmarks-gin

go get
go build
./bookmarks-gin
```

You can specify different IP/PORT by creating/changing these env variables
- $IP
- $PORT