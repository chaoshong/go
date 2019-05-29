package config

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/viper"
)

func init() {

}

type DbInfo struct {
	Dbname   string
	Host     string
	Port     string
	User     string
	Password string
	Sslmode  string
}

type Db struct {
	DbType string
	Db     DbInfo
}

func GetDBInfo() (string, string) {
	Os := runtime.GOOS
	// 获取系统类型，win 系统用 %，linux 用 $
	if Os == "windows" {
		//viper.AddConfigPath("d:/golang/src/dbconfig")
		viper.AddConfigPath(os.Getenv("GOPATH") + "/src/dbconfig")
		fmt.Println("windows")

	} else {
		viper.AddConfigPath("$GOPATH/src/config")
		fmt.Println("Linux")
	}
	viper.SetConfigName("db")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read DB error: %v", err)
	}

	var dbs Db
	err = viper.Unmarshal(&dbs)
	if err != nil {
		fmt.Println("read info error: %v", err)
	}

	return dbs.DbType, fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s password=%s",
		dbs.Db.Host, dbs.Db.User, dbs.Db.Dbname, dbs.Db.Port, dbs.Db.Sslmode, dbs.Db.Password)

}
