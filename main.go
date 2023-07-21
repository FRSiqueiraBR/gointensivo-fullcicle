package main

import (
	"database/sql"
	"fmt"

	"github.com/FRSiqueiraBR/gointensivo-fullcycle/internal/infra/database"
	"github.com/FRSiqueiraBR/gointensivo-fullcycle/internal/usecase"
	"github.com/satori/go.uuid"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	myuuid := uuid.NewV4()

	input := usecase.OrderInput{
		ID: myuuid.String(),
		Price: 10.0,
		Tax: 1.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}