package main

import (
	sql2 "database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-jet/jet/v2/postgres"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"io/ioutil"
	"log"
	"test-jet-go/.gen/postgres/public/model"
	. "test-jet-go/.gen/postgres/public/table"
)

func main() {
	stmt := postgres.
		SELECT(
			Products.AllColumns,
			Properties.AllColumns,
			Values.AllColumns,
		).
		FROM(
			Products.
				LEFT_JOIN(Properties, Properties.ProductID.EQ(Products.ID)).
				LEFT_JOIN(Values, Values.PropertyID.EQ(Properties.ID)),
		).
		ORDER_BY(Products.ID)

	var connectString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "", "postgres")

	db, err := sql2.Open("pgx", connectString)

	if err != nil {
		log.Fatal(err)
	}

	var result []struct {
		model.Products

		Properties []struct {
			model.Properties

			Value []model.Values
		}
	}

	err = stmt.Query(db, &result)

	if err != nil {
		log.Fatal(err)
	}

	text, _ := json.MarshalIndent(result, "", "    ")
	ioutil.WriteFile("./test.json", text, 0644)
}
