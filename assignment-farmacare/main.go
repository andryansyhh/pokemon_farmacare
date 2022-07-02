package main

import (
	"pokemon_battle/app"
	infra "pokemon_battle/infra/postgre"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	if err := infra.InitPostgre(); err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
