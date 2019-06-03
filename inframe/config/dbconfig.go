package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
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

	viper.AddConfigPath(os.Getenv("GOPATH") + "/src/config")
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
