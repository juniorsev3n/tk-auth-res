package structs

import "github.com/jinzhu/gorm"

type Reseller struct {
	gorm.Model
	First_Name string
	Last_Name  string
	Password   string
}
