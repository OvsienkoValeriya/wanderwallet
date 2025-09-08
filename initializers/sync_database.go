package initializers

import (
	"log"
	"wanderwallet/internal/models"
)

func SyncDatabase() {
	if err := DB.AutoMigrate(
		&models.User{},
		&models.Travel{},
		&models.Expense{},
		&models.Category{},
	); err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}

	seedCategories()
}
func seedCategories() {
	categories := []models.Category{
		{Name: "Транспорт", Builtin: true},
		{Name: "Питание", Builtin: true},
		{Name: "Жильё", Builtin: true},
		{Name: "Продукты", Builtin: true},
		{Name: "Шоппинг", Builtin: true},
		{Name: "Достопримечательности", Builtin: true},
		{Name: "Развлечения", Builtin: true},
		{Name: "Общие", Builtin: true},
		{Name: "Дополнительные сборы и платежи", Builtin: true},
	}

	for _, c := range categories {
		var existing models.Category
		if err := DB.Where("name = ? AND user_id IS NULL", c.Name).First(&existing).Error; err != nil {
			if err := DB.Create(&c).Error; err != nil {
				log.Printf("не удалось создать категорию %s: %v", c.Name, err)
			}
		}
	}
}
