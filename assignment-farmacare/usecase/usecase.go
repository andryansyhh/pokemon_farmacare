package usecase

import "pokemon_battle/repository"

type PokeUsecase struct {
	PokeRepository repository.PokeRepoInterface
}

type PokeUsecaseInterface interface {
	PokemonUsecase
}

func NewPokeUsecase(pokeRepo repository.PokeRepoInterface) PokeUsecaseInterface {
	return &PokeUsecase{
		PokeRepository: pokeRepo,
	}
}
