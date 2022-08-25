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

var (
	dbUser     = "root"
	dbPasswd   = "123456"
	dbHost     = "127.0.0.1"
	dbPort     = "3306"
	dpPort5738 = "33061"
	dpPort800  = "33062"
	dpPort8029 = "33063"
	dbName     = "mysql"
)

func MySQLVersion800() {
	if err := startGorm(dbUser, dbPasswd, dbHost, dpPort800, dbName); err != nil {
		panic(err)
	}
}

func MySQLVersion8029() {
	if err := startGorm(dbUser, dbPasswd, dbHost, dpPort8029, dbName); err != nil {
		panic(err)
	}
}

func MySQLVersion5738() {
	if err := startGorm(dbUser, dbPasswd, dbHost, dpPort5738, dbName); err != nil {
		panic(err)
	}
}

func startGorm(dbUser, dbPasswd, dbHost, dbPort, dbName string) error {
	mysqlConfig := mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPasswd, dbHost, dbPort, dbName),
	})
	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	}
	var err error
	db, err = gorm.Open(mysqlConfig, gormConfig)
	return err
}

func MigrateTable(tables ...interface{}) {
	// table: optional pointer
	if len(tables) == 0 {
		return
	}
	autoMigrateTables = append(autoMigrateTables, tables...)
	autoMigrate()
}

func autoMigrate() {
	autoDB := db.Set("gorm:table_options", "ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 ROW_FORMAT = Dynamic")
	if err := autoDB.AutoMigrate(autoMigrateTables...); err != nil {
		panic(err)
	}
}
