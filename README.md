# gin-api-server

```
$ go get github.com/gin-gonic/gin
$ go get github.com/go-xorm/xorm
$ go get github.com/mattn/go-sqlite3
$ sqlite3 development.db
sqlite> CREATE TABLE user(id INTEGER PRIMARY KEY AUTOINCREMENT,  first_name VARCHAR(255) NOT NULL,  last_name VARCHAR(255) NOT NULL);
sqlite> .tables
sqlite> INSERT INTO user(first_name,  last_name) VALUES("Yamada",  "Taro");
sqlite> INSERT INTO user(first_name,  last_name) VALUES("Sato",  "Hajime");
sqlite> INSERT INTO user(first_name,  last_name) VALUES("Koike",  "Tatsuo");
sqlite> SELECT * FROM user;
1|Yamada|Taro
2|Sato|Hajime
3|Koike|Tatsuo
sqlite> .quit
```
