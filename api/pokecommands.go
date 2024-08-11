package api
import (
    "fmt"
    "os"
    "net/http"
    "time"
    "errors"
    "math/rand"

    "github.com/Enigma56/pokedex/internal/cache"
    "github.com/Enigma56/pokedex/internal/pokedict"
)

type Client struct {
    rand *rand.Rand
    cache cache.PokeCache
    httpClient http.Client
}

func NewClient() Client {
    randSrc := rand.NewSource(1)
    return Client{
        rand: rand.New(randSrc),
        cache: cache.NewCache(5 * time.Minute),
        httpClient: http.Client{
            Timeout: time.Minute,
        },
    }
}

type Config struct {
    Pokedex pokedict.Pokedex
    ApiClient Client
    NextLocationAreaURL *string
    PrevLocationAreaURL *string
    CurrLocationAreaURL string
}

type cmd func(*Config, ...string) error

var CommandMap = map[string]cmd{
    "help": CmdHelp,
    "exit": CmdExit,
    "map": CmdMap,
    "mapb": CmdMapb,
    "explore": CmdExploreArea,
    "catch": CmdCatch,
    "inspect": CmdInspect,
    "pokedex": CmdPokedex,
}

type LocationArea struct {
    Count int `json:"count"`
    Next *string `json:"next"`
    Previous *string `json:"previous"`
    Results []struct {
        Name string `json:"name"`
        Url string `json:"url"`
    } `json:"results"`
}

func CmdPokedex(cfg *Config, args ...string) error {
    cfg.Pokedex.GetAllPokemon()
    return nil
}

func CmdInspect(cfg *Config, args ...string) error { 
    if len(args) < 1 {
        return errors.New("No argument passed to function!")
    }

    err := cfg.Pokedex.GetInfo(args[0])
    if err != nil {
        return err
    }   
    return nil
}

func CmdCatch(cfg *Config, args ...string) error {
    if len(args) < 1 {
        return errors.New("No argument passed to function!")
    }

    float := cfg.ApiClient.rand.Float64()
    pokemon, err := cfg.ApiClient.FetchPokemonInfo(args[0])

    if err != nil {
        return err
    }

    fmt.Printf("Throwing a Pokeball at %s!\n", pokemon.Name)
    catchSuccessful := catchSuccessful(float, pokemon.BaseExperience) 
    if catchSuccessful {
        err := cfg.Pokedex.Add(pokemon) //no possible way to return an error
        if err != nil {
            return err
        }
        fmt.Printf("%s was caught!\n", pokemon.Name)
        return nil
    } else {
        fmt.Printf("%s escaped!\n", pokemon.Name)
        return nil
    } 
}

func CmdExploreArea(cfg *Config, args ...string) error {
    if len(args) < 1 {
        return errors.New("No argument passed to function!")
    }

    area, err := cfg.ApiClient.ListAreaDetails(args[0])

    if err != nil {
        return err
    }

    fmt.Printf("Exploring %s...\n", args[0])
    fmt.Println("Pokemon:")
    for _, encounter := range area.PokemonEncounters {
        fmt.Println(encounter.Pokemon.Name)
    }

    return nil
}
func CmdMap(cfg *Config, args ...string) error {
    areas, err := cfg.ApiClient.ListLocationAreas(&cfg.CurrLocationAreaURL)

    if err != nil {
        return err
    }
    
    cfg.NextLocationAreaURL = areas.Next
    cfg.PrevLocationAreaURL = areas.Previous
    cfg.CurrLocationAreaURL = *cfg.NextLocationAreaURL

    //Cast results to a []byte --> add to PokeCache
    for _, loc := range areas.Results {
        fmt.Println(loc.Name)
    }
    return nil
}



func CmdMapb(cfg *Config, args ...string) error {
    if cfg.PrevLocationAreaURL == nil {
        return errors.New("No prev location URL has been set, try calling map cmd?")
    }

    cfg.CurrLocationAreaURL = *cfg.PrevLocationAreaURL

    areas, err := cfg.ApiClient.ListLocationAreas(&cfg.CurrLocationAreaURL)
    if err != nil {
        return err
    }
 
    cfg.NextLocationAreaURL = areas.Next
    cfg.PrevLocationAreaURL = areas.Previous
    

    for _, loc := range areas.Results {
        fmt.Println(loc.Name)
    }

    return nil
}

func CmdHelp(cfg *Config, args ...string) error {
    fmt.Println("\nWelcome to the Pokedex!")
    fmt.Println("Usage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")

    return nil
}

func CmdExit(cfg *Config, args ...string) error {
    os.Exit(0)

    return nil
}
