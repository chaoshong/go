package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("dbconfig")
	viper.AddConfigPath("$GOPATH/src/dbconfig")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read DB error: %v", err)
	}

}

func GetDBInfo() string {
	var (
		dbName   string
		port     string
		user     string
		sslMode  string
		password string
		host     string
	)

	dbName = viper.GetString("PostgreSql.db.dbname")
	port = viper.GetString("PostgreSql.db.port")
	user = viper.GetString("PostgreSql.db.user")
	sslMode = viper.GetString("PostgreSql.db.sslmode")
	password = viper.GetString("PostgreSql.db.password")
	host = viper.GetString("PostgreSql.db.host")

	return fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s password=%s", host, user, dbName, port, sslMode, password)

}
