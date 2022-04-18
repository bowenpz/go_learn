package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db                *gorm.DB
	autoMigrateTables []interface{}
)

func init() {
	var (
		dbUser   = "root"
		dbPasswd = "123456"
		dbHost   = "127.0.0.1"
		dbPort   = "3306"
		dbName   = "mysql"
		err      error
	)

	mysqlConfig := mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPasswd, dbHost, dbPort, dbName),
	})
	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	}
	if db, err = gorm.Open(mysqlConfig, gormConfig); err != nil {
		panic(err)
	}
}

func AutoMigrate() {
	autoDB := db.Set("gorm:table_options", "ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = Dynamic")
	if err := autoDB.AutoMigrate(autoMigrateTables...); err != nil {
		panic(err)
	}
}

func MigrateTable(table interface{}) {
	// table: optional pointer
	autoMigrateTables = append(autoMigrateTables, table)
}
