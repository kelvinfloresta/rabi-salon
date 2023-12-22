package fixtures

import (
	"rabi-salon/config"
	"rabi-salon/frameworks/database/gorm_adapter"
	"rabi-salon/frameworks/database/gorm_adapter/models"
)

var TestDatabase = gorm_adapter.New(config.TestDatabase)

var tables = []string{
	models.User{}.TableName(),
}

func CleanDatabase() {
	gormDatabase, ok := TestDatabase.(*gorm_adapter.GormAdapter)
	if !ok {
		panic(gormDatabase)
	}

	if gormDatabase.Conn == nil {
		if err := gormDatabase.Connect(); err != nil {
			panic(err)
		}
	}

	for _, table := range tables {
		if err := gormDatabase.Conn.Exec("TRUNCATE " + table + " CASCADE").Error; err != nil {
			panic(err)
		}
	}
}
