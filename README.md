# village


## 系统环境准备,以ubuntu为例子:

1. apt-get update
2. apt-get install git
3. apt-get install golang
4. apt-get install mysql-server
5. 配置golang环境

## 安装第三方库:
1. go get github.com/astaxie/beego
2. go get github.com/beego/bee
3. go get -u qiniupkg.com/api.v7
4. go get github.com/go-sql-driver/mysql
5. go get -u github.com/astaxie/beego/session/mysql
6. go get github.com/golang-commonmark/markdown
7. go get github.com/golang-commonmark/mdtool

## 使用mysql数据库:
1. 创建village数据库
```
    create database village CHARACTER SET utf8 COLLATE utf8_general_ci;
```

2. 创建session mysql存储引擎表

```
CREATE TABLE `session` (
    `session_key` char(64) NOT NULL,
	`session_data` blob,
	`session_expiry` int(11) unsigned NOT NULL,
	PRIMARY KEY (`session_key`)
	) ENGINE=MyISAM DEFAULT CHARSET=utf8;
```

3. 在main.go里面配置数据库地址