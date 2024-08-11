package pokedict

import (
    "fmt"
    "errors" 
)

type Pokemon struct {
    Name    string  `json:"name"`
    Height  int `json:"height"`
    Weight  int `json:"weight"`
    BaseExperience int `json:"base_experience"`
    Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type Pokedex struct {
    pokedex map[string]Pokemon
}

func NewPokedex() Pokedex {
    return Pokedex{
        pokedex: map[string]Pokemon{},
    }
}

func (pd *Pokedex) Add(pokemon Pokemon) error {
    _, ok := pd.pokedex[pokemon.Name]
    if ok {
        return errors.New("This pokemon has already been caught!")
    }

   pd.pokedex[pokemon.Name] = pokemon
   return nil
}


func (pd *Pokedex) GetInfo(name string) error {
    p, ok := pd.pokedex[name]
    if !ok {
        return errors.New("This pokemon is not in your pokedex")
    }

    fmt.Printf("Name: %s\n", p.Name)
    fmt.Printf("Height: %v\n", p.Height)
    fmt.Printf("Weight: %v\n", p.Weight)

    fmt.Println("Stats:")
    for _, stat := range p.Stats {
        fmt.Printf("\t-%s: %v\n", stat.Stat.Name, stat.BaseStat)
    }

    fmt.Println("Types:")
    for _, t := range p.Types {
        fmt.Printf("\t- %s\n", t.Type.Name)
    }
    return nil
}

func (pd *Pokedex) GetAllPokemon() {
    fmt.Println("Your Pokedex:")
    for name := range pd.pokedex {
        fmt.Printf("\t- %s", name)
    }
}
