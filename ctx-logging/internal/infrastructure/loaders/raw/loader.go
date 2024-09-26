package raw

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/jibaru/ctx-logging/internal/domain"
)

type Loader struct {
	filePath string
}

func NewLoader(filePath string) *Loader {
	return &Loader{filePath: filePath}
}

func (l *Loader) Load() (map[int]*domain.Pokemon, error) {
	// Open the JSON file
	absPath, err := filepath.Abs(l.filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open JSON file: %w", err)
	}

	file, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("could not open JSON file: %w", err)
	}
	defer file.Close()

	// Read the content of the file
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("could not read JSON file: %w", err)
	}

	// Parse the JSON data
	var pokemons []domain.Pokemon
	err = json.Unmarshal(data, &pokemons)
	if err != nil {
		return nil, fmt.Errorf("could not parse JSON file: %w", err)
	}

	pokemonsMap := make(map[int]*domain.Pokemon, 0)
	for _, pokemon := range pokemons {
		pokemonsMap[pokemon.ID] = &pokemon
	}

	return pokemonsMap, nil
}
