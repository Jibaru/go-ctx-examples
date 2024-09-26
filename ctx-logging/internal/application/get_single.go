package application

import (
	"context"
	"log/slog"

	"github.com/jibaru/ctx-logging/internal/domain"
)

// GetSingleService handles getting a single Pokémon by name.
type GetSingleService struct {
	repository domain.PokemonRepository
}

// NewGetSingleService creates a new GetSingleService.
func NewGetSingleService(repository domain.PokemonRepository) GetSingleService {
	return GetSingleService{repository: repository}
}

// GetSingle retrieves a single Pokémon by name.
func (s GetSingleService) GetSingle(ctx context.Context, id int) (*domain.Pokemon, error) {
	slog.InfoContext(ctx, "get single pokemon called", "id", id)
	return s.repository.GetByID(ctx, id)
}
