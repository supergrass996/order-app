package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"

	"github.com/supergrass996/order-app/internal/notification"
	"github.com/supergrass996/order-app/internal/order"
	"github.com/supergrass996/order-app/internal/storage"
)

func main() {
	db, err := sql.Open("sqlite", "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := storage.NewSQLiteRepo(db)

	var initializer order.RepositoryInitializer = repo
	if err := initializer.Init(); err != nil {
		log.Fatal(err)
	}

	// Один сервис — разные отправители.
	emailService := order.NewOrderService(repo, &notification.EmailSender{})
	if err := emailService.CreateOrder("Иван", []string{"apple", "banana"}, 10.5); err != nil {
		log.Fatal(err)
	}

	smsService := order.NewOrderService(repo, &notification.SMSSender{})
	if err := smsService.CreateOrder("Мария", []string{"milk"}, 5.0); err != nil {
		log.Fatal(err)
	}
}
