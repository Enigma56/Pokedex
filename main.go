package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    type cmd func()
    commandMap := map[string]cmd{
        "help": cmdHelp,
        "exit": cmdExit,
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

        //fmt.Println(userIn)
    }
}

func cmdHelp() {
    fmt.Println("\nWlcome to the Pokedex!")
    fmt.Println("Usage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
}

func cmdExit() {
    os.Exit(0)
}
