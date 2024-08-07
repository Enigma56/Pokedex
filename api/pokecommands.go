package api
import (
    "fmt"
    "os"
    "net/http"
    "time"
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
}

type cmd func(*Config) error

var CommandMap = map[string]cmd{
    "help": CmdHelp,
    "exit": CmdExit,
    "map": CmdMap,
}

type LocationArea struct {
    Count int `json:"count"`
    Next string `json:"next"`
    Previous *string `json:"previous"`
    Results []struct {
        Name string `json:"name"`
        Url string `json:"url"`
    } `json:"results"`
}

func CmdMap(cfg *Config) error {
    //apiClient := NewClient()
    areas, err := cfg.ApiClient.ListLocationAreas()

    if err != nil {
        return err
    }
    
    cfg.NextLocationAreaURL = areas.Next
    cfg.PrevLocationAreaURL = areas.Prev
    printLocations(areas)
    return nil
}



func CmdMapb(cfg *Config) error {
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



func printLocations(location LocationArea) {
    maps := location.Results

    for _, loc := range maps {
        fmt.Println(loc.Name)
    }
}
