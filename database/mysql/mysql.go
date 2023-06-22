package mysql

import (
	"fmt"

	"github.com/berrylradianh/ecowave-go/config"
	"github.com/berrylradianh/ecowave-go/database/seed"
	pe "github.com/berrylradianh/ecowave-go/modules/entity/product"
	er "github.com/berrylradianh/ecowave-go/modules/entity/review"
	re "github.com/berrylradianh/ecowave-go/modules/entity/role"
	et "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"

	ie "github.com/berrylradianh/ecowave-go/modules/entity/information"

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
		re.Role{},
		ue.User{},
		ue.UserDetail{},
		ue.UserAddress{},
		ie.Information{},
		ve.Voucher{},
		pe.ProductCategory{},
		pe.Product{},
		pe.ProductImage{},
		et.Transaction{},
		et.TransactionDetail{},
		ve.VoucherType{},
		er.RatingProduct{},
	)
	DB.Migrator().HasConstraint(&ue.User{}, "UserDetail")
	DB.Migrator().HasConstraint(&re.Role{}, "Users")
	// DB.Migrator().HasConstraint(&pe.Product{}, "TransactionDetails")
}
