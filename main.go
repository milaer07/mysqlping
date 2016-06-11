package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
	 "time"
	 "log"
	 "os"
)  
 
var   logger *log.Logger   //log句柄全局变量
var logfile *os.File    //文件句柄全局变量

func logfileopen(){
	logfile,err := os.OpenFile("mysqlping.log",os.O_CREATE|os.O_RDONLY|os.O_APPEND,0)
	logger =log.New(logfile,"\r\n检查网络:",log.LstdFlags)
	CheckErr(err)
}
//log文件打开

func dbconn() {	
 //   db, err := sql.Open("mysql", "root:330327@/lcname?charset=utf8")
	db,err:= sql.Open("mysql","zhgb:123456@tcp(10.2.0.55:3306)/zhgbdb?charset=utf8")
	CheckErr(err)
     defer db.Close()
	err=db.Ping()
	CheckErr(err)
    rows, err := db.Query("select * from devices_list")  //查询表格
    CheckErr(err)
	defer rows.Close()  //退出后关闭rows
	updateMoney, err := db.Prepare("UPDATE devices_list SET flag_ip=?,flag_m_time=? WHERE device_id=?")          //Prepare创建一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。
    CheckErr(err) 
	//循环导出查询语句
		for rows.Next() {
        var line_id string
        var station string
		var device_id string
		var device_ip string
		var flag_ip int
		var flag_m_time []uint8
        err = rows.Scan(&line_id, &station,&device_id,&device_ip,&flag_ip,&flag_m_time)                      //复制查询语句
        CheckErr(err)
		errip := goping(device_ip)           //运行网络查找
		res, err := updateMoney.Exec(errip,time.Now(),device_id)  //返回值进行数据修改
		 CheckErr(err)
		_, err = res.RowsAffected()   //_原来名为affect,查询执行的次数
		CheckErr(err)
/*		fmt.Println(affect)     打印修改的次数
		println(errip)              打印GOPING返回值     
        fmt.Println(line_id,"   ",station,"   ",device_id,"      ",device_ip)       打印数据查询的列项   */
    
	}
	
}           //数据库链接并进行PING网络

func main() {
		logfileopen()               //开启日志
		t_start :=time.Now()    //执行开始时间
		dbconn()                 //连接数据库处理
		fmt.Printf("执行完毕,总共执行%dms",time.Now().Sub(t_start).Nanoseconds()/1e6)  //显示执行时间
		defer logfile.Close()
}

func CheckErr(err error) {
    if err != nil {
		logger.Fatal(err)
    //    panic(err)
    }
}

