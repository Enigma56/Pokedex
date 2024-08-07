package api
import (
    "fmt"
    "os"
    "net/http"
    "time"
    "errors"
)

type Client struct {
    httpClient http.Client
}

func NewClient() Client {
    return Client{
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

type cmd func(*Config) error

var CommandMap = map[string]cmd{
    "help": CmdHelp,
    "exit": CmdExit,
    "map": CmdMap,
    "mapb": CmdMapb,
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

func CmdMap(cfg *Config) error {
    areas, err := cfg.ApiClient.ListLocationAreas(&cfg.CurrLocationAreaURL)

    if err != nil {
        return err
    }
    
    cfg.NextLocationAreaURL = areas.Next
    cfg.PrevLocationAreaURL = areas.Previous
    cfg.CurrLocationAreaURL = *cfg.NextLocationAreaURL



    for _, loc := range areas.Results {
        fmt.Println(loc.Name)
    }
    return nil
}



func CmdMapb(cfg *Config) error {
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

func CmdHelp(cfg *Config) error {
    fmt.Println("\nWelcome to the Pokedex!")
    fmt.Println("Usage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")

    return nil
}

func CmdExit(cfg *Config) error {
    os.Exit(0)

    return nil
}
