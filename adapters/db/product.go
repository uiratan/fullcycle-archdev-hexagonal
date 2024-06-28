package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uiratan/fullcycle-archdev-hexagonal/application"
)

type ProductDB struct {
	db *sql.DB
}

func (p *ProductDB) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

//Save(product ProductInterface) (ProductInterface, error)
