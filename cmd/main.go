package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "log"

    "github.com/Enigma56/pokedex/api"
) 

func main() {
    cfg := api.Config{
        ApiClient: api.NewClient(),
    }

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

        err := cmd(&cfg)
        if err != nil {
            log.Fatalf("Command failed with err: %s", err)
        }
    }
}

