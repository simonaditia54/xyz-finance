package main

import (
	"log"
	"os"
	"time"

	"xyz-finance/internal/auth"
	"xyz-finance/internal/consumer"
	"xyz-finance/internal/limit"
	"xyz-finance/internal/model"
	"xyz-finance/internal/transaction"
	"xyz-finance/pkg/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := mysql.Connect()
	// repo := transaction.NewInMemoryRepo()

	// // Tambahkan dummy limit
	// repo.UpdateLimit(&model.Limit{
	// 	ConsumerID: "uuid-123",
	// 	TenorMonth: 3,
	// 	TotalLimit: 2000000,
	// 	UsedLimit:  0,
	// })

	// service := transaction.NewService(repo)
	// handler := transaction.NewHandler(service)
	// consumerRepo := consumer.NewInMemoryRepo()
	consumerRepo := consumer.NewMySQLRepo(db)
	consumerService := consumer.NewService(consumerRepo)
	consumerHandler := consumer.NewHandler(consumerService)

	consumerRepo.Save(model.Consumer{
		ID:            "uuid-123",
		NIK:           "1234567890",
		FullName:      "Budi Santoso",
		LegalName:     "BUDI SANTOSO",
		TempatLahir:   "Jakarta",
		TanggalLahir:  time.Now(),
		Gaji:          7000000,
		FotoKTPURL:    "https://dummy.com/ktp.jpg",
		FotoSelfieURL: "https://dummy.com/selfie.jpg",
	})

	// transactionRepo := transaction.NewInMemoryRepo()
	transactionRepo := transaction.NewMySQLRepo(db)
	transactionService := transaction.NewService(transactionRepo)
	transactionHandler := transaction.NewHandler(transactionService)

	// transactionRepo.UpdateLimit(&model.Limit{
	// 	ConsumerID: "uuid-123",
	// 	TenorMonth: 3,
	// 	TotalLimit: 2000000,
	// 	UsedLimit:  0,
	// })

	limitRepo := limit.NewMySQLRepo(db)
	limitService := limit.NewService(limitRepo)
	limitHandler := limit.NewHandler(limitService)

	authService := auth.NewService(consumerRepo, os.Getenv("secret_key"))
	authHandler := auth.NewHandler(authService)
	authMiddleware := auth.JWTAuthMiddleware(os.Getenv("secret_key"))

	// handler.RegisterRoutes(r)
	// consumerHandler.RegisterRoutes(r)
	authHandler.RegisterRoutes(r)

	authorized := r.Group("/", authMiddleware)
	consumerHandler.RegisterRoutes(authorized)
	transactionHandler.RegisterRoutes(authorized)
	limitHandler.RegisterRoutes(authorized)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
