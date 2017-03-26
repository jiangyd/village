

系统环境准备:

apt-get update
apt-get install git
apt-get install golang
apt-get install mysql-server

安装第三方库:
go get github.com/astaxie/beego
go get github.com/beego/bee
go get -u qiniupkg.com/api.v7
go get github.com/go-sql-driver/mysql


创建数据库:
create database village CHARACTER SET utf8 COLLATE utf8_general_ci;