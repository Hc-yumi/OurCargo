package main

import (
	"log"

	"example.com/m/dao"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Person struct {
// 	gorm.Model
// 	Name string
// 	Age  int
// }

var db *gorm.DB
var err error

func main() {
	dsn := "host=myrds.c8eoe8ahfumy.ap-northeast-1.rds.amazonaws.com user=postgres password=Hach8686 dbname=test port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// エラーでたらプロセス終了
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// db, err := gorm.Open("sqlite3", "mydb.sqlite3")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// db.LogMode(true)

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// create persons table
		{
			ID: "00004",
			Migrate: func(tx *gorm.DB) error {
				tx.AutoMigrate(
					&dao.User{},
					&dao.Company{},
					&dao.TruckType{},
					&dao.TruckSize{},
					&dao.Order{},
				)

				tx.Migrator().CreateConstraint(&dao.User{}, "CompanyID")
				tx.Migrator().CreateConstraint(&dao.User{}, "fk_users_company")

				tx.Migrator().CreateConstraint(&dao.TruckType{}, "UserID")
				tx.Migrator().CreateConstraint(&dao.TruckType{}, "fk_trucktypes_user")

				tx.Migrator().CreateConstraint(&dao.TruckType{}, "TruckSizeID")
				tx.Migrator().CreateConstraint(&dao.TruckType{}, "fk_trucktypes_trucksize")

				tx.Migrator().CreateConstraint(&dao.Order{}, "UserID")
				tx.Migrator().CreateConstraint(&dao.Order{}, "fk_orders_user")

				tx.Migrator().CreateConstraint(&dao.Order{}, "TruckTypeID")
				tx.Migrator().CreateConstraint(&dao.Order{}, "fk_orders_trucktype")

				tx.Migrator().CreateConstraint(&dao.Order{}, "TruckSizeID")
				tx.Migrator().CreateConstraint(&dao.Order{}, "fk_orders_trucksize")

				return nil
			},

			Rollback: func(tx *gorm.DB) error {
				// return tx.Migrator().DropTable("people")
				return nil
			},
		},

		{
			ID: "00005",
			Migrate: func(tx *gorm.DB) error {
				tx.AutoMigrate(
					&dao.User{},
				)

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				// return tx.Migrator().DropTable("people")
				return nil
			},
		},

		{
			ID: "00006",
			Migrate: func(tx *gorm.DB) error {

				// User table
				tx.Create(&dao.User{LineID: "111"})
				tx.Create(&dao.User{Name: "aaa"})
				tx.Create(&dao.User{Role: "発注先"})
				tx.Create(&dao.User{Role: "発注元"})
				tx.Create(&dao.User{Email: "example@example.com"})
				// tx.Create(&dao.User{Tel: "09000001111"})

				// Company table
				tx.Create(&dao.Company{CompanyName: "transport株式会社"})
				tx.Create(&dao.Company{CompanyLocation: "千葉県市川市"})

				// TruckType table
				// tx.Create(&dao.TruckType{Price: "10000"})
				// tx.Create(&dao.TruckType{Price: "20000"})
				// tx.Create(&dao.TruckType{Price: "35000"})
				// tx.Create(&dao.TruckType{Price: "50000"})

				// TruckSize table
				tx.Create(&dao.TruckSize{TruckSize: "軽"})
				tx.Create(&dao.TruckSize{TruckSize: "2t"})
				tx.Create(&dao.TruckSize{TruckSize: "4t"})
				tx.Create(&dao.TruckSize{TruckSize: "10t"})

				// Order table
				tx.Create(&dao.Order{PickupLocation: "千葉県市川市"})
				tx.Create(&dao.Order{ArrivalLocation: "東京都墨田区"})
				// tx.Create(&dao.Order{PickupDatetime: "2006-01-02 00:00"})
				// tx.Create(&dao.Order{ArrivalDatetime: "2006-01-02 00:00"})
				// tx.Create(&dao.Order{Mileage: "20km"})
				// tx.Create(&dao.Order{OrderDatetime: "111"})
			
				return nil
			},

			Rollback: func(tx *gorm.DB) error {
				// return tx.Migrator().DropTable("people")
				return nil
			},
		},
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
