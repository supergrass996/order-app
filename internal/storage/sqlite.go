package storage

import (
	"database/sql"
	"fmt"

	"github.com/supergrass996/order-app/internal/order"
)

// SQLiteRepo реализует RepositoryInitializer и RepositoryWriter для SQLite.
type SQLiteRepo struct {
	db *sql.DB
}

func NewSQLiteRepo(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{db: db}
}

func (r *SQLiteRepo) Init() error {
	_, err := r.db.Exec(`
        CREATE TABLE IF NOT EXISTS orders (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            customer TEXT NOT NULL,
            products TEXT NOT NULL,
            total REAL NOT NULL,
            status TEXT NOT NULL
        )`)
	return err
}

func (r *SQLiteRepo) SaveOrder(o order.Order) (order.Order, error) {
	result, err := r.db.Exec(
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		o.Customer, o.Products, o.Total, o.Status,
	)
	if err != nil {
		return order.Order{}, err
	}
	id, _ := result.LastInsertId()
	o.ID = int(id)
	fmt.Printf("Заказ #%d сохранён в БД\n", o.ID)
	return o, nil
}

// Проверка компиляции: SQLiteRepo обязан реализовывать оба интерфейса.
var (
	_ order.RepositoryInitializer = (*SQLiteRepo)(nil)
	_ order.RepositoryWriter      = (*SQLiteRepo)(nil)
)
