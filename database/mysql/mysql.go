package mysql

import (
	"fmt"

	"github.com/berrylradianh/ecowave-go/config"
	"github.com/berrylradianh/ecowave-go/database/seed"
	rt "github.com/berrylradianh/ecowave-go/modules/entity/role"
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	InitDB()
	InitialMigration()
	seed.DBSeed(DB)
}

func InitDB() {
	var err error

	configurations := config.GetConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configurations.DB_USERNAME,
		configurations.DB_PASSWORD,
		configurations.DB_HOST,
		configurations.DB_PORT,
		configurations.DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
}

func InitialMigration() {
	DB.AutoMigrate(
		rt.Role{},
		ut.User{},
		ut.UserDetail{},
	)
	DB.Migrator().HasConstraint(&ut.User{}, "UserDetail")
	DB.Migrator().HasConstraint(&rt.Role{}, "Users")
}
