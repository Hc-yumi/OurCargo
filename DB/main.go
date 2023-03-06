package main

import (
	"log"

	"github.com/ourcargo/constant"
	"github.com/ourcargo/dao"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	dsn := "host=myrds.c8eoe8ahfumy.ap-northeast-1.rds.amazonaws.com user=postgres password=Hach8686 dbname=test port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// エラーでたらプロセス終了
		log.Fatalf("Some error occured. Err: %s", err)
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// create persons table
		{
			ID: "2023030011",
			Migrate: func(tx *gorm.DB) error {
				tx.AutoMigrate(
					&dao.User{},
					&dao.Company{},
					&dao.TruckType{},
					&dao.TruckSize{},
					&dao.Order{},
				)

				// tx.Migrator().CreateConstraint(&dao.User{}, "Company")
				// tx.Migrator().CreateConstraint(&dao.User{}, "fk_company_id")

				// tx.Migrator().CreateConstraint(&dao.TruckType{}, "User")
				// tx.Migrator().CreateConstraint(&dao.TruckType{}, "fk_user_id")

				// tx.Migrator().CreateConstraint(&dao.TruckType{}, "TruckSize")
				// tx.Migrator().CreateConstraint(&dao.TruckType{}, "fk_trucksize_id")

				// tx.Migrator().CreateConstraint(&dao.Order{}, "User")
				// tx.Migrator().CreateConstraint(&dao.Order{}, "fk_user_id")

				// tx.Migrator().CreateConstraint(&dao.Order{}, "TruckType")
				// tx.Migrator().CreateConstraint(&dao.Order{}, "fk_trucktype_id")

				// tx.Migrator().CreateConstraint(&dao.Order{}, "TruckSize")
				// tx.Migrator().CreateConstraint(&dao.Order{}, "fk_trucksize_id")

				return nil
			},

			Rollback: func(tx *gorm.DB) error {
				return nil
			},
		},

		{
			ID: "2023030012",
			Migrate: func(tx *gorm.DB) error {

				// Company table
				tx.Create(&dao.Company{CompanyName: "transport株式会社", CompanyLocation: "千葉県市川市"})

				// User table
				tx.Create(&dao.User{Name: "aaa",
					LineID:  "yyy",
					Role:    "発注先",
					Email:   "example@example.com",
					Company: dao.Company{ID: 1}})

				// TruckSize table
				tx.Create(&dao.TruckSize{TruckSize: constant.TruckSizeSmall})
				tx.Create(&dao.TruckSize{TruckSize: constant.TruckSizeTowTon})
				tx.Create(&dao.TruckSize{TruckSize: constant.TruckSizeFourTon})
				tx.Create(&dao.TruckSize{TruckSize: constant.TruckSizeTenTon})

				// TruckType table
				tx.Create(&dao.TruckType{User: dao.User{ID: 1},
					TruckSize: dao.TruckSize{ID: 1},
					Price:     "￥10,000"})

				// Order table
				tx.Create(&dao.Order{User: dao.User{ID: 1},
					TruckType:       dao.TruckType{ID: 1},
					TruckSize:       dao.TruckSize{ID: 1},
					PickupLocation:  "千葉県市川市",
					ArrivalLocation: "東京都墨田区",
					Mileage:         "20km"})

				return nil
			},

			Rollback: func(tx *gorm.DB) error {
				return nil
			},
		},
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
