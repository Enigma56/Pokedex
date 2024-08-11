package api
import (
    "fmt"
    "os"
    "net/http"
    "time"
    "errors"

    "github.com/Enigma56/pokedex/internal/cache"
)

type Client struct {
    cache cache.PokeCache
    httpClient http.Client
}

func NewClient() Client {
    return Client{
        cache: cache.NewCache(5 * time.Minute),
        httpClient: http.Client{
            Timeout: time.Minute,
        },
    }
}

type Config struct {
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

func CmdExploreArea(cfg *Config, args ...string) error {
    if len(args) < 1 {
        return errors.New("No argument passed to function!")
    }

    area, err := cfg.ApiClient.ListAreaDetails(args[0])

    if err != nil {
        return err
    }

    fmt.Printf("Exploring %s...", args[0])
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
