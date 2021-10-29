package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
type mysqlConfig struct {
	UserName string
	UserPass string
	DBName string
}
var saveDirName string = "/data/back_up/mysql/"
var filePix string = "gin_vue_blog_"

func main() {
	nowT := time.Now().Format("15:04:05")
	fileName := filePix+time.Now().Format("2006:01:02") + "-" + nowT
	CreateFile(fileName)
	shellFileName := saveDirName + fileName+ ".sh"

	c1 := "chmod 777 " +  shellFileName +" && " + shellFileName +" && rm -rf "+shellFileName
	Command(c1)
}
// 创建文件
func CreateFile(fileName string) {
	sqlName := fileName + ".sql"
	shellFileName := saveDirName + fileName + ".sh"
	mysqlContainerName := "mysql"//容器名称
	var DbConfig = mysqlConfig{
		UserName: "root",
		UserPass: "root",
		DBName:   "tcsd",
	}
	s1 := "#/bin/bash \n"
	mysqlLogin := fmt.Sprintf(" -u%s -p%s %s", DbConfig.UserName, DbConfig.UserPass, DbConfig.DBName)
	s2 := "docker exec -i " + mysqlContainerName + " /bin/bash -c 'mysqldump "+mysqlLogin+" > /" + sqlName + "' \n"
	s3 := "docker cp " + mysqlContainerName + ":/" + sqlName + " " +  saveDirName + sqlName  + " \n"
	s := fmt.Sprintf("%s %s %s", s1, s2, s3)
	err := ioutil.WriteFile(shellFileName, []byte(s), 0666) //直接覆盖原来的内容
	if err != nil {
		fmt.Println(err)
	}
}

// 执行命令
func Command(cmd string) {
	c := exec.Command("bash", "-c", cmd)
	output, _ := c.CombinedOutput()
	fmt.Println(string(output))
	fmt.Print("命令执行完成")
}


func main1() {
	//who := "World!"
	//有参数通过空格隔开
	if len(os.Args) > 1 {
		l := len(os.Args)
		for i := 0; i < l; i++ {
			fmt.Println(os.Args[i])
			fmt.Println("\n")
		}
		//who = strings.Join(os.Args[1:], " ")
	}
	//fmt.Println("Hello", who)
}
