package domain

import "context"

// Pokemon represents a Pokémon entity.
type Pokemon struct {
	ID   int `json:"id"`
	Name struct {
		English  string `json:"english"`
		Japanese string `json:"japanese"`
		Chinese  string `json:"chinese"`
		French   string `json:"french"`
	} `json:"name"`
	Type []string `json:"type"`
	Base struct {
		HP        int `json:"HP"`
		Attack    int `json:"Attack"`
		Defense   int `json:"Defense"`
		SpAttack  int `json:"Sp. Attack"`
		SpDefense int `json:"Sp. Defense"`
		Speed     int `json:"Speed"`
	} `json:"base"`
	Species     string `json:"species"`
	Description string `json:"description"`
	Evolution   struct {
		Prev []string   `json:"prev,omitempty"`
		Next [][]string `json:"next,omitempty"`
	} `json:"evolution"`
	Profile struct {
		Height  string     `json:"height"`
		Weight  string     `json:"weight"`
		Egg     []string   `json:"egg"`
		Ability [][]string `json:"ability"`
		Gender  string     `json:"gender"`
	} `json:"profile"`
	Image struct {
		Sprite    string `json:"sprite"`
		Thumbnail string `json:"thumbnail"`
		Hires     string `json:"hires"`
	} `json:"image"`
}

// PokemonRepository defines the methods a repository must implement.
type PokemonRepository interface {
	GetByID(ctx context.Context, id int) (*Pokemon, error)
}
