package memory

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/jibaru/ctx-logging/internal/domain"
)

// RawPokemonRepository retrieves Pokémon data from a text file.
type RawPokemonRepository struct {
	pokemons map[int]*domain.Pokemon
}

// NewRawPokemonRepository creates a new RawPokemonRepository.
func NewRawPokemonRepository(pokemons map[int]*domain.Pokemon) *RawPokemonRepository {
	return &RawPokemonRepository{
		pokemons: pokemons,
	}
}

// GetByName retrieves a Pokémon by its name.
func (r *RawPokemonRepository) GetByID(ctx context.Context, id int) (*domain.Pokemon, error) {
	time.Sleep(1 * time.Second)

	pokemon, exists := r.pokemons[id]
	if !exists {
		slog.InfoContext(ctx, "pokemon not found", "id", id)
		return nil, errors.New("pokemon not found")
	}

	slog.InfoContext(ctx, "pokemon found")

	return pokemon, nil
}
