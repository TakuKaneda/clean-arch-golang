# clean-arch-golang

Implement this website <https://qiita.com/hirotakan/items/698c1f5773a3cca6193e>.

## How To Run

### Setting MySQL

```bash
$ mysql.server start
$ mysql -uroot
mysql> create database sample;
mysql> use sample;
mysql> create table users ( id int NOT NULL PRIMARY KEY auto_increment, first_name varchar(255), last_name varchar(255));
```

### Run server

```bash
$ go run server.go
```

### Sample Requests

```bash
# create
$ curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"FirstName": "David", "LastName": "Alaba"}' localhost:8080/users

# list users
$ curl -H 'Content-Type:application/json' localhost:8080/users

# get user by ID
$ curl -H 'Content-Type:application/json' localhost:8080/users/1
```
