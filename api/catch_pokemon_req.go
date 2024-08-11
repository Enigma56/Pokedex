
package api

import (
    "fmt"
    "encoding/json"
    "net/http"
    "io"
    "log"

    "github.com/Enigma56/pokedex/internal/pokedict"
)


func (c *Client) FetchPokemonInfo(pokemonName string) (pokedict.Pokemon, error) {
    url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
    dat, ok := c.cache.Get(url)
    if ok {
        fmt.Println("cache hit!")
        var pokemonDetails pokedict.Pokemon
        err := json.Unmarshal(dat, &pokemonDetails)
        if err != nil {
            fmt.Println(err)
            return pokedict.Pokemon{}, err
        }
        
        return pokemonDetails, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal(err)
        return pokedict.Pokemon{}, err
    }

    res, err := c.httpClient.Do(req)
    
	if err != nil {
		log.Fatal(err)
        return pokedict.Pokemon{}, err
	}

    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if res.StatusCode > 399 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
        return pokedict.Pokemon{}, err
	}

    if err != nil {
        return pokedict.Pokemon{}, err
    }

    err = c.cache.Add(url, body)
    fmt.Printf("Added url to cache: %s\n", url)
    if err != nil {
        return pokedict.Pokemon{}, err
    }

    var pokemonDetails pokedict.Pokemon
    err = json.Unmarshal([]byte(body), &pokemonDetails)
    if err != nil {
        fmt.Println(err)
        return pokedict.Pokemon{}, err
    }

    return pokemonDetails, nil
}

func catchSuccessful(rand_float float64, base_exp int) bool {
    chance := rand_float / (1 / float64(base_exp))

    if chance > .5 {
        return true
    } else {
        return false
    }
} 

