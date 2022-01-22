# go_mysql_backup
使用 Golang 写的一个 Mysql 备份工具，支持单机 Docker Mysql 容器和 Mysql 持续更新中...（练手项)


## docker 备份请看 docker_mysql_bak.go
## 单机mysql备份请看 mysql.go

打包二进制文件
set GOARCH=amd64
set GOOS=linux
go build -o jstgs-image-bak-mysql docker_mysql_bak.go

go build -o test-jstgs-image-bak-mysql docker_mysql_bak.go