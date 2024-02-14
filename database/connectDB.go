package database

import (
	"fmt"
	"log"
	"os"

	"github.com/github.com/chinnnawat/Practice_Golang_RestApi-Docker-Postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	//  dsn ย่อมาจาก Data source Name คือรูปแบบที่เป็น string ใช้เก็บข้อมูลวิธีการเชื่อมต่อ
	// ดยปกติจะใช้ในการเชื่อมต่อกับฐานข้อมูล โดย DSN จะรวมถึงข้อมูลสำคัญที่ใช้ในการเชื่อมต่อ เช่นชื่อผู้ใช้ (user), รหัสผ่าน (password), ชื่อฐานข้อมูล (dbname),
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Fact{})

	DB = Dbinstance{
		Db: db,
	}
}
