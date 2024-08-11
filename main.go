package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "log"
    //"time"

    "github.com/Enigma56/pokedex/api"
    //"github.com/Enigma56/pokedex/internal/cache"
) 

func main() {
    cfg := api.Config{
        ApiClient: api.NewClient(),
        CurrLocationAreaURL: "https://pokeapi.co/api/v2/location-area",
    }
    commandMap := api.CommandMap
        
    for {
        userScanner := bufio.NewScanner(os.Stdin)
        
        fmt.Print("Pokedex > ")
        userScanner.Scan()
        userIn := userScanner.Text()
        userIn = strings.TrimRight(userIn, "\n")

        userArgs := strings.Split(userIn, " ")
        
        cmd, ok := commandMap[userArgs[0]]
        if !ok {
            log.Print("Command not recognized")
        }

        //cmdArgs := userArgs[1:]
        err := cmd(&cfg, userArgs[1:]...)
        if err != nil {
            fmt.Printf("Command failed with error: %v\n", err)   
        }
    }
}

