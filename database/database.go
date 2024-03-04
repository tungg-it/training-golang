package database

import (
	"fmt"

	"gin-api/app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(config *config.ConfigAppEnv) *gorm.DB {
	// "user:password@tcp(host:port)/database?charset=utf8&parseTime=True&loc=Local"
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.MySqlUser, config.MySqlPassword, config.MySqlHost, config.MySqlPort, config.MySqlDatabase)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       url,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect database!!")
	}

	fmt.Println("Connect database successfully!!")

	// sqlDb, _ := db.DB()
	// driver, _ := migrateSql.WithInstance(sqlDb, &migrateSql.Config{})

	// // Create Migrate instance
	// m, err := migrate.NewWithDatabaseInstance(
	// 	"migrations",
	// 	"mysql",
	// 	driver,
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// // Run migrations unfulfilleda
	// err = m.Up()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Migrations completed!")

	return db
}
