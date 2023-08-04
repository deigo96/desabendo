package bootrstrap

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(env *Env) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	env.Db.DbUser,
	env.Db.DbPass,
	env.Db.DbHost,
	env.Db.DbPort,
		env.Db.Dbname,
)

db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func CloseConnection(db *gorm.DB) {
	if db == nil {
		return
	}

	if db != nil {
		db, _ := db.DB()
		db.Close()
	}
}