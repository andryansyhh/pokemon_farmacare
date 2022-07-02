package app

import (
	"fmt"
	"os"
	"pokemon_battle/infra"
	postgres "pokemon_battle/infra/postgre"
	"pokemon_battle/repository"
	"pokemon_battle/usecase"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	pokemonRepo := repository.NewPokeRepo(postgres.PSQL.DB.DB)
	app := usecase.NewPokeUsecase(pokemonRepo)
	infra.RegisterApi(router, app)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
