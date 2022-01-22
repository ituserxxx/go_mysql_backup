package main

import (
	"fmt"
	"os/exec"
	"time"
)

// 流程逻辑
// docker 容器中
// docker exec -it mysql /bin/bash -c 'mysqldump -uroot -proot databaseName > ./bak.sql'
// docker cp mysql:/bak.sql ./bak.sql
// 非容器中
// mysqldump -uroot -proot databaseName > ./bak.sql
// 持续更新中
//
type mysql struct {
	UserName string
	UserPass string
	DBName string
	SaveDirName string
}


func main() {
	var cof = &mysql{
		UserName:    "root",
		UserPass:    "root",
		DBName:      "gin_vue_blog",
		SaveDirName: "/data/back_up/mysql/",
	}
	goRun(cof)
}
func goRun(cof *mysql)  {
	nowT := time.Now().Format("2006_01_02_15_04_05")
	fileName := cof.DBName + "-" + nowT

	sqlName := fileName + ".sql"

	mysqlLogin := fmt.Sprintf("mysqldump  -u%s -p%s %s > %s", cof.UserName, cof.UserPass, cof.DBName, cof.SaveDirName +sqlName)

	fmt.Println("执行命令："+mysqlLogin)

	c := exec.Command("bash", "-c", mysqlLogin)
	output, _ := c.CombinedOutput()
	fmt.Println(string(output))
	fmt.Println(cof,"命令执行完成")
}
