package main

import (
	"fmt"
	"gorm.io/gorm/clause"
)

func main() {
	AutoMigrate()

	var department Department
	res := db.Preload(clause.Associations).Where("department_id = ?", 1).Find(&department)
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Printf("%v", department)
}
