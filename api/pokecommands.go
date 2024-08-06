package api

type LocationData struct {
    Count int `json:"count"`
    Next string `json:"next"`
    Previous *string `json:"previous"`
    Results []Location `json:"results"`
}

type Location struct {
    Name string `json:"name"`
    Url string `json:"url"`
}

func printLocations(location LocationData) {
    maps := location.Results

    for _, loc := range maps {
        fmt.Println(loc.Name)
    }
} 

func pokeLocations() error {
    res, err := http.Get("https://pokeapi.co/api/v2/location-area")
    if err != nil {
        log.Fatal(err)
    }

    body, err := io.ReadAll(res.Body)
    res.Body.Close()

    if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

    //fmt.Printf("%s\n", body)

    //byteBody := []byte(body)
    //fmt.Println(byteBody)
    //fmt.Println(string(body))
    var location LocationData
    err1 := json.Unmarshal([]byte(body), &location)
    if err1 != nil {
        fmt.Println(err1)
    }

    printLocations(location)

    return nil
}
func cmdHelp() error {
    fmt.Println("\nWlcome to the Pokedex!")
    fmt.Println("Usage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")

    return nil
}

func cmdExit() error {
    os.Exit(0)

    return nil
}
