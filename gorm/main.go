package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Genre struct {
	ID         int
	Genre_name string
}

func connectionDB() (*gorm.DB, error) {
	// user=postgres.ciotjxaqztwqcxlzhvis password=[YOUR-PASSWORD] host=aws-0-ap-southeast-1.pooler.supabase.com port=5432 dbname=postgres
	var consStr = "host=aws-0-ap-southeast-1.pooler.supabase.com user=postgres.ciotjxaqztwqcxlzhvis password=cPb6LdnTd7lxokCF dbname=postgres port=5432"

	db, err := gorm.Open(postgres.Open(consStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetAllGenre(db *gorm.DB) ([]Genre, error) {
	var rv []Genre
	err := db.Find(&rv).Error
	if err != nil {
		return nil, err
	}

	return rv, nil
}

func main() {

	dbCon, err := connectionDB()
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	fmt.Println(dbCon)

	db, err := GetAllGenre(dbCon)

	if err != nil {
		fmt.Println("error Get:", err)
		return
	}

	fmt.Println(db)

	// dbCon.AutoMigrate(&Genre{})
}
