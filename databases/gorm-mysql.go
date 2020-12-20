package databases

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type DbConfig struct {
	Host        string        `yaml:"host"`
	Port        int           `yaml:"port"`
	UserName    string        `yaml:"username"`
	Password    string        `yaml:"password"`
	DB          string        `yaml:"db"`
	MaxLifeTime time.Duration `yaml:"maxlifetime"`
	MaxIdleConn int           `yaml:"maxidleconn"`
	MaxOpenConn int           `yaml:"maxopenconn"`
}

//多数据源
func NewMySqlDb(dbConfig DbConfig) (*gorm.DB, error) {
	dsn :=  fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		dbConfig.UserName, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DB)
	gdb, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:  256,
	}), &gorm.Config{
		NamingStrategy:  schema.NamingStrategy{
			SingularTable: false,
		} ,
	})
	if err != nil {
		return  nil, err
	}
	db , err := gdb.DB()
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(dbConfig.MaxLifeTime)
	db.SetMaxIdleConns(dbConfig.MaxIdleConn)
	db.SetMaxOpenConns(dbConfig.MaxOpenConn)
	if err := db.Ping(); err != nil {
		return  nil, err
	}
	return gdb, nil
}
