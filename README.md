# Bookmarks Gin (Golang)

### This is part of a project: [Bookmarks](https://github.com/stefanjarina/bookmarks)

* Gin Framework
* GORM
* MySQL

#### TODO:
- [x] Change MongoDB to GORM and MySQL
  - [x] Add support for other databases (PostgreSQL, SQLite3, MSSQL)
- [x] Add TOML config file support
- [x] Fix code formatting
- [ ] Proper testing + TDD onwards
- [x] Users + authentification - probably leverage jwt tokens
- [ ] Public/Private Bookmarks
- [ ] Backend filters support

### Installation

#### Database Setup

##### MySQL
I presume you have MySQL installed and running.

In *MySQL CLI* create user, database and give user privileges for database.
e.g.:
```sql
CREATE USER 'bookmarks_gin'@'localhost' IDENTIFIED BY 'password';
CREATE DATABASE bookmarks_gin;
GRANT ALL ON bookmarks_gin.* TO 'bookmarks_gin'@'localhost';
```

Update `app.toml` and application will build up connection string automatically.

##### PostgreSQL
I presume you have PostgreSQL installed and running.

```bash
$ sudo su - postgres

postgres@servername:~$ createdb bookmarks_gin

postgres@servername:~$ createuser -P bookmarks_gin
Enter password for new role: 
Enter it again: 

postgres@servername:~$ psql
psql (9.1.9)
Type "help" for help.

postgres=# GRANT ALL PRIVILEGES ON DATABASE bookmarks_gin TO bookmarks_gin;
GRANT
postgres=# \q
postgres@servername:~$ logout
```

Update `app.toml` and application will build up connection string automatically.

##### SqLite 3
I presume you have SQLite 3 installed.

```bash
$ sqlite3 /tmp/bookmarks_gin.db
```

Update `app.toml` and application will build up connection string automatically.

#### Download/Build App and initialize the Database
Download source, fetch dependencies, compile and run
```bash
go get github.com/stefanjarina/bookmarks-gin
cd $GOPATH/src/github.com/stefanjarina/bookmarks-gin

go get
go build
./bookmarks-gin initdb --password "your_new_admin_pass"
./bookmarks-gin
```

Application defaults to localhost:4000. You may change this behaviour in `app.toml` or set the following environment variables:
```bash
export IP=0.0.0.0
export PORT=4000
```