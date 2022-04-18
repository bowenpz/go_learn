package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	DepartmentID int `gorm:"primary_key"`
	Name         string
	Employees    []Employee `gorm:"association_foreignkey:DepartmentID;foreignkey:DepartID"`
}

func (d Department) String() string {
	return fmt.Sprintf("【%d】%s %v", d.ID, d.Name, d.Employees)
}

func (d Department) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Printf("\n--------\ncreate department: %v\n--------\n", d)
	return nil
}

type Employee struct {
	gorm.Model
	UserID   int `gorm:"primary_key"`
	Name     string
	DepartID int
}

func (e Employee) String() string {
	return fmt.Sprintf("【%d】%s", e.ID, e.Name)
}

func init() {
	MigrateTable(&Department{})
	MigrateTable(&Employee{})
}
