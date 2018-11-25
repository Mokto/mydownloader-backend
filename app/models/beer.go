package models

// import (
// 	_ "fmt"
// 	// "github.com/jinzhu/gorm"
// 	// "github.com/wangzitian0/golang-gin-starter-kit/common"
// 	// "github.com/wangzitian0/golang-gin-starter-kit/users"
// 	// "strconv"
// )

type Beer struct {
	// gorm.Model
	ID          int
	Name		string
	Type		string
	Brasserie		string
}
