package model

import (
	"fmt"

	"github.com/xilepeng/Blog/utils"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(utils.Db,fmt.Sprintf("%s:%s@tcp(%s:%s/)"))
}
