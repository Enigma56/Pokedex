package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "log"

    "github.com/Enigma56/pokedex/api"
    "github.com/Enigma56/pokedex/internal/cache"
) 

func main() {
    cfg := api.Config{
        ApiClient: api.NewClient(),
        CurrLocationAreaURL: "https://pokeapi.co/api/v2/location-area",
    }

    cache := cache.NewCache()
    //fmt.Println(cache)


    commandMap := api.CommandMap
        
    for {
        userScanner := bufio.NewScanner(os.Stdin)
        
        fmt.Print("Pokedex > ")
        userScanner.Scan()
        userIn := userScanner.Text()
        userIn = strings.TrimRight(userIn, "\n")
        
        cmd, ok := commandMap[userIn]
        if !ok {
            log.Print("Command not recognized")
        }

        err := cmd(&cfg, &cache)
        if err != nil {
            fmt.Printf("Command failed with error: %v\n", err)   
        }
    }
}

