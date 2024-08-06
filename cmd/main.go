package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"

    "github.com/Enigma56/pokedex/api"
)  

func main() {
    //restructure to have functions implement an interface that contains only a description method
    type cmd func() error
    commandMap := map[string]cmd{
        "help": api.cmdHelp,
        "exit": api.cmdExit,
        "map": api.pokeLocations
    }
    for {
        userScanner := bufio.NewScanner(os.Stdin)
        
        //Take input from user
        fmt.Print("Pokedex > ")
        userScanner.Scan()
        userIn := userScanner.Text()
        userIn = strings.TrimRight(userIn, "\n")
        
        cmd, ok := commandMap[userIn]
        if ok {
            cmd()
        }
    }
}

