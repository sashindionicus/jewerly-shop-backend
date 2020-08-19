package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	userTable          = "users"
	categoriesTable    = "categories"
	titlesTable        = "titles"
	descriptionsTable  = "descriptions"
	materialsTable     = "materials"
	imagesTable        = "images"
	productsTable      = "products"
	productImagesTable = "product_images"
	ordersTable        = "orders"
	orderItemsTable    = "order_items"
	transactionsTable  = "transactions"
	adminUsersTable    = "admin_users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	DBName   string
	SSLMode  string
	Password string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	logrus.Infof("db config: %+v\n", cfg)

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.SSLMode, cfg.Password,
	))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
