package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/codeedu/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProducts(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
		);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProducts(db *sql.DB) {
	insert := `insert into products values ("abc", "cadeira", 0, "disabled")`
	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDB := db.NewProductDB(Db)

	product, err := productDB.Get("abc")

	require.Nil(t, err)
	require.NotNil(t, product)

}

func TestProductDB_Create(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDB(Db)

	productToSave := application.NewProduct()

	productToSave.Price = 10
	productToSave.Name = "cadeira"

	savedProduct, err := productDb.Save(productToSave)

	require.Nil(t, err)
	require.NotNil(t, savedProduct)

	recoveredProduct, err := productDb.Get(savedProduct.GetId())

	require.Nil(t, err)
	require.NotNil(t, recoveredProduct)
	require.Equal(t, savedProduct.GetId(), recoveredProduct.GetId())

}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDB(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "enabled"

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

}
