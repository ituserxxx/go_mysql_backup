package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"reflect"
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
		UserName:    "",
		UserPass:    "",
		DBName:      "",
		SaveDirName: "",
	}
	for{
		i := 0
		for i<reflect.ValueOf(mysql{}).NumField(){
			if cof.SaveDirName == ""{
				fmt.Println(">>> 请输入保存的路径：")
				var flag string
				_, _ = fmt.Scanln(&flag)
				cof.SaveDirName = flag
			}
			if cof.UserName == ""{
				fmt.Println(">>> 请输入保存的登录Mysql用户名：")
				var flag string
				_, _ = fmt.Scanln(&flag)
				cof.UserName = flag
			}
			if cof.UserPass == ""{
				fmt.Println(">>> 请输入保存的登录Mysql密码：")
				var flag string
				_, _ = fmt.Scanln(&flag)
				cof.UserPass = flag
			}
			if cof.DBName == ""{
				fmt.Println(">>> 请输入需要同步的库：")
				var flag string
				_, _ = fmt.Scanln(&flag)
				cof.DBName = flag
			}
			i++

		}
		break
	}

	goRun(cof)
}
func goRun(cof *mysql)  {

	nowT := time.Now().Format("15:04:05")
	fileName := cof.DBName+time.Now().Format("2006:01:02") + "-" + nowT

	sqlName := fileName + ".sql"
	shellFileName := cof.SaveDirName + fileName + ".sh"

	s1 := "#/bin/bash \n"
	mysqlLogin := fmt.Sprintf(" -u%s -p%s %s", cof.UserName, cof.UserPass, cof.DBName)
	s2 := "mysqldump "+mysqlLogin+" > /" + sqlName + "' \n"
	s3 :=  sqlName + " " +  cof.SaveDirName+ sqlName  + " \n"
	s := fmt.Sprintf("%s %s %s", s1, s2, s3)
	err := ioutil.WriteFile(shellFileName, []byte(s), 0666) //直接覆盖原来的内容
	if err != nil {
		fmt.Println(err)
	}

	shellFileName = cof.SaveDirName + fileName + ".sh"
	c1 := "chmod 777 " +  shellFileName +" && " + shellFileName +" && rm -rf "+shellFileName


	c := exec.Command("bash", "-c", c1)
	output, _ := c.CombinedOutput()
	fmt.Println(string(output))
	fmt.Println(cof,"命令执行完成")
}
