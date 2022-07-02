package infra

import (
	"pokemon_battle/handler"
	"pokemon_battle/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app usecase.PokeUsecaseInterface) {
	pokeSrv := handler.NewPokemonHttpServer(app)
	api := r.Group("/pokemon")
	{
		api.POST("/battle", pokeSrv.PostBattle)
		api.GET("/", pokeSrv.GetAllPokemons)
		api.GET("/battles", pokeSrv.GetBattle)
		api.GET("/scores", pokeSrv.GetPokemonScore)
	}
}
