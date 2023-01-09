package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/ssengalanto/biscuit/cmd/account/app"
)

func main() {
	app.Run()
}
