package main

import (
	"fmt"
	"log"

	"github.com/go-faker/faker/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hamedhaghi/goforum/config"
	"github.com/hamedhaghi/goforum/models"
	"github.com/hamedhaghi/goforum/services"

	"gorm.io/gorm"
)

func main() {

	storage := config.GetConfig().Storage
	storage.Migrator().DropTable(
		&models.Thread{},
		&models.Post{},
		&models.Comment{},
	)
	storage.AutoMigrate(
		&models.Thread{},
		&models.Post{},
		&models.Comment{},
	)
	threadSeeder(storage)
}

func threadSeeder(storage *gorm.DB) {
	thread := services.NewThreadService(storage)
	for i := 1; i <= 100; i++ {
		tm := &models.Thread{
			Title:       faker.Sentence(),
			Description: faker.Paragraph(),
			Posts: []models.Post{
				{
					Title:   faker.Sentence(),
					Content: faker.Paragraph(),
					Votes:   i,
					Comments: []models.Comment{
						{
							Content: faker.Paragraph(),
							Votes:   i,
						},
					},
				},
			},
		}
		if err := thread.Create(tm); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Migrated Successfully!")
}
