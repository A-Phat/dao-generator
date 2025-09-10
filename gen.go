package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// 🔁 โหลด .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_LOCAL")
	if dsn == "" {
		log.Fatal("Missing DB_LOCAL in .env")
	}

	// 🔗 เชื่อมต่อ DB
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	outPath, _ := filepath.Abs("D:/platform-service/auth-service/pkg/entity")
	fmt.Println("OutPath:", outPath)
	g := gen.NewGenerator(gen.Config{
		OutPath:        outPath,
		ModelPkgPath:   "entity", 
		Mode:           gen.WithoutContext,
		FieldNullable:  true,
		FieldCoverable: true,
		FieldSignable:  true,
		FieldWithIndexTag: true,             
		FieldWithTypeTag:  true,  
		
	})

	g.UseDB(db)

	// เลือก table ที่ต้องการ generate
	g.GenerateModel("apps")


	g.Execute()
}
