package main

import (
	"log"
	"os"

	"banking_system/config"
	"banking_system/models"
	"banking_system/routes"
)

func main() {
	config.InitDB()

	config.DB.Migrator().DropTable(
		&models.Transaction{},
		&models.Repayment{},
		&models.Loan{},
		&models.AccountCustomer{},
		&models.Account{},
		&models.Customer{},
		&models.Branch{},
		&models.Bank{},
	)

	err := config.DB.AutoMigrate(
		&models.Bank{},
		&models.Branch{},
		&models.Customer{},
		&models.Account{},
		&models.AccountCustomer{},
		&models.Loan{},
		&models.Repayment{},
		&models.Transaction{},
	)
	if err != nil {
		log.Println("Migration failed, dropping and recreating tables...")
		config.DB.Migrator().DropTable(
			&models.Transaction{},
			&models.Repayment{},
			&models.Loan{},
			&models.AccountCustomer{},
			&models.Account{},
			&models.Customer{},
			&models.Branch{},
			&models.Bank{},
		)
		err = config.DB.AutoMigrate(
			&models.Bank{},
			&models.Branch{},
			&models.Customer{},
			&models.Account{},
			&models.AccountCustomer{},
			&models.Loan{},
			&models.Repayment{},
			&models.Transaction{},
		)
		if err != nil {
			log.Fatalf("failed to run migrations after dropping tables: %v", err)
		}
	}

	router := routes.SetupRouter(config.DB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

